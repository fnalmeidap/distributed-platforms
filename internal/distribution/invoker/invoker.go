package messaginginvoker

import (
	marshaller "distributed-platforms/internal/distribution/marshaller"
	"distributed-platforms/internal/distribution/miop"

	//miop "distributed-platforms/internal/distribution/miop"
	srh "distributed-platforms/internal/infra/srh"
	shared "distributed-platforms/internal/shared"

	//"fmt"
	//"distributed-platforms/test/mymom/services/messagingservice"
	//"distributed-platforms/test/mymom/services/messagingservice/event"
	calculator "distributed-platforms/internal/app/calculator"
	"log"
)

type Invoker struct {
	Ior shared.IOR
	app *calculator.Calculator
}

func NewInvoker(h string, p int, app *calculator.Calculator) Invoker {
	i := shared.IOR{Host: h, Port: p}
	r := Invoker{Ior: i, app: app}
	return r
}

func (inv Invoker) Invoke() {
	s := srh.NewSRH(inv.Ior.Host, inv.Ior.Port)
	m := marshaller.Marshaller{}

	miopPacket := miop.Packet{}

	var reply interface{}

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		switch r.Operation {
		case "sum":
			//_p1 := r.Params[0].(string)
			_p2 := r.Params[1].(int)
			_p3 := r.Params[2].(int)
			reply = inv.app.Sum(_p2, _p3)
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
	}
}
