package requestor

import (
	"distributed-platforms/distribution/marshaller"
	"distributed-platforms/distribution/miop"
	"distributed-platforms/infra/crh"
	"distributed-platforms/shared"
)

type Requestor struct {
}

func newRequestor() Requestor {
	return Requestor{}
}

func (Requestor) Invoke(i shared.Invocation) shared.Termination {
	reqPacket := miop.CreateRequestMIOP(i.Request.Op, i.Request.Params)

	// Serialization
	m := marshaller.Marshaller{}
	b := m.Marshall(reqPacket)

	c := crh.NewCRH(i.Ior.Host, i.Ior.Port)
	r := c.SendReceive(b)

	repPacket := m.Unmarshal(r)
	rt := miop.ExtractReply(repPacket)

	t := shared.Termination{Rep: rt}

	return t
}
