package calculatorinvoker

import (
	calculator "distributed-platforms/internal/app/calculator"
	lifecyclemanager "distributed-platforms/internal/distribution/lifecycle_manager"
	marshaller "distributed-platforms/internal/distribution/marshaller"
	miop "distributed-platforms/internal/distribution/miop"
	srh "distributed-platforms/internal/infra/srh"
	shared "distributed-platforms/internal/shared"
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

func (inv Invoker) Invoke() {
	s := srh.NewSRH(inv.Ior.Host, inv.Ior.Port)
	lcm := lifecyclemanager.NewLifecycleManager()
	m := marshaller.Marshaller{}
	kDefaultLeaseDuration := time.Duration(shared.DefaultLeasingTimeSeconds * float64(time.Second))

	var ans int
	var c *calculator.Calculator

	iorToServer := shared.IOR{Host: shared.LocalHost, Port: shared.DefaultPortClientServer}
	go lcm.Lm.CleanupExpiredLeases(iorToServer)

	for {

		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket := m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		// Leasing remote pattern implementation
		lcm.Lease(kDefaultLeaseDuration, &c)

		if r.Operation == "LeaseType" {
			_p1 := r.Params[0].(string)
			var a int
			a = -1
			switch _p1 {
			case "lease_type_0":
				a = 0
			case "lease_type_1":
				a = 1
			case "lease_type_2":
				a = 2
			default:
				log.Fatal("Invoker:: Operation '" + r.Operation + "' is unknown:: ")
			}
			lcm.LeaseTypeSet(a) //_p2 is not used here.
			ans = 0
		} else {

			_p1 := int(r.Params[0].(float64))
			_p2 := int(r.Params[1].(float64))

			switch r.Operation {
			case "Sum":
				ans = c.Sum(_p1, _p2)
			case "Sub":
				ans = c.Sub(_p1, _p2)
			case "Mul":
				ans = c.Mul(_p1, _p2)
			case "Div":
				ans = c.Div(_p1, _p2)
			case "LeaseExtend":
				lcm.RenewLease_v2(_p1) //_p2 is not used here.
				ans = 0
			default:
				log.Fatal("Invoker:: Operation '" + r.Operation + "' is unknown:: ")
			}
		}

		// Prepare reply
		var params []interface{}
		params = append(params, ans)

		//TODO: make this variable depending if the lease is still valid or not
		if lcm.Lm.LeaseOkayFlag {
			params = append(params, "ok")
		} else {
			params = append(params, "no leasing resource available")
		}

		// Create miop reply packet
		miop := miop.CreateReplyMIOP(params)

		// Marshall miop packet
		b = m.Marshall(miop)

		// Send marshalled packet
		s.Send(b)

	}
}
