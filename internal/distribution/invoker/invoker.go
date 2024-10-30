package messaginginvoker

import (
	marshaller "distributed-platforms/internal/distribution/marshaller"
	miop "distributed-platforms/internal/distribution/miop"

	srh "distributed-platforms/internal/infra/srh"
	shared "distributed-platforms/internal/shared"

	//"fmt"
	calculator "distributed-platforms/internal/app/calculator"
	lease "distributed-platforms/internal/lease"
	"log"
	"time"
)

type Invoker struct {
	Ior          shared.IOR
	app          *calculator.Calculator
	leaseManager *lease.LeaseManager
}

func NewInvoker(h string, p int, app *calculator.Calculator, leaseMan *lease.LeaseManager) Invoker {
	i := shared.IOR{Host: h, Port: p}
	r := Invoker{Ior: i, app: app, leaseManager: leaseMan}
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

	miopPacket := miop.Packet{}

	var reply interface{}
	duration := time.Duration(20 * float64(time.Second)) // 20s
	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		_p1 := int(r.Params[0].(float64))
		_p2 := int(r.Params[1].(float64))

		exists := OperationLeaseExists(r.Operation, inv.leaseManager)
		if exists {
			inv.leaseManager.UpdateLease(r.Operation, duration)
		} else {
			inv.leaseManager.NewLease(r.Operation, duration)
		}

		switch r.Operation {
		case "Sum":

			reply = inv.app.Sum(_p1, _p2)
		default:
			log.Fatal("Invoker:: Operation '" + r.Operation + "' is unknown:: ")
		}

		// Prepare reply
		var params []interface{}
		params = append(params, reply)

		// Create miop reply packet
		miop := miop.CreateReplyMIOP(params)

		// Marshall miop packet
		b = m.Marshall(miop)

		// Send marshalled packet
		s.Send(b)

		go inv.leaseManager.CleanupExpiredLeases()
	}
}
