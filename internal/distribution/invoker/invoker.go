package calculatorinvoker

import (
	marshaller "distributed-platforms/internal/distribution/marshaller"
	miop "distributed-platforms/internal/distribution/miop"

	srh "distributed-platforms/internal/infra/srh"
	shared "distributed-platforms/internal/shared"

	calculator "distributed-platforms/internal/app/calculator"
	lease "distributed-platforms/internal/lease"
	"log"
	"time"
)

type Invoker struct {
	Ior shared.IOR
}

func NewInvoker(h string, p int) Invoker {
	i := shared.IOR{Host: h, Port: p}
	r := Invoker{Ior: i}
	return r
}

func OperationLeaseExists(op string, leaseManager *lease.LeaseManager) bool {
	exists := false
	for key := range leaseManager.Leases {
		if op == key {
			exists = true
		} else {
			exists = false
		}
	}
	return exists
}

func (inv Invoker) Invoke() {
	s := srh.NewSRH(inv.Ior.Host, inv.Ior.Port)
	m := marshaller.Marshaller{}
	var ans int

	c := calculator.Calculator{}

	// Lease manager
	leaseManager := lease.LeaseManager{Leases: make(map[string]lease.Lease)}
	duration := time.Duration(shared.DefaultLeasingTimeSeconds * float64(time.Second))
	go leaseManager.CleanupExpiredLeases()

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket := m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		_p1 := int(r.Params[0].(float64))
		_p2 := int(r.Params[1].(float64))

		exists := OperationLeaseExists(r.Operation, &leaseManager)
		if exists {
			leaseManager.UpdateLease(r.Operation, duration)
		} else {
			leaseManager.NewLease(r.Operation, duration)
		}

		switch r.Operation {
		case "Sum":
			ans = c.Sum(_p1, _p2)
		case "Sub":
			ans = c.Sub(_p1, _p2)
		case "Mul":
			ans = c.Mul(_p1, _p2)
		case "Div":
			ans = c.Div(_p1, _p2)
		default:
			log.Fatal("Invoker:: Operation '" + r.Operation + "' is unknown:: ")
		}

		// Prepare reply
		var params []interface{}
		params = append(params, ans)

		// Create miop reply packet
		miop := miop.CreateReplyMIOP(params)

		// Marshall miop packet
		b = m.Marshall(miop)

		// Send marshalled packet
		s.Send(b)
	}
}
