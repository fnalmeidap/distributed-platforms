package calculatorproxy

import (
	"distributed-platforms/internal/distribution/marshaller"
	"distributed-platforms/internal/distribution/miop"
	"distributed-platforms/internal/distribution/requestor"
	"distributed-platforms/internal/infra/srh"
	"distributed-platforms/internal/shared"
	"fmt"
)

type CalculatorProxy struct {
	Ior shared.IOR
}

func New(i shared.IOR) CalculatorProxy {
	r := CalculatorProxy{Ior: i}
	return r
}

func (p *CalculatorProxy) Sum(p1, p2 int) (int, string) {
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	req := shared.Request{Operation: "Sum", Params: params}
	inv := shared.Invocation{Ior: p.Ior, Request: req}
	fmt.Println("Request params...")
	fmt.Printf("Operation: %s\n", req.Operation)
	fmt.Printf("Parameters: [%d,%d]\n", req.Params...)

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64)), r.Rep.Result[1].(string)
}

func (p *CalculatorProxy) Sub(p1, p2 int) (int, string) {
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	req := shared.Request{Operation: "Sub", Params: params}
	inv := shared.Invocation{Ior: p.Ior, Request: req}
	fmt.Println("Request params...")
	fmt.Printf("Operation: %s\n", req.Operation)
	fmt.Printf("Parameters: [%d,%d]\n", req.Params...)

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64)), r.Rep.Result[1].(string)
}

func (p *CalculatorProxy) Div(p1, p2 int) (int, string) {
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	req := shared.Request{Operation: "Div", Params: params}
	inv := shared.Invocation{Ior: p.Ior, Request: req}
	fmt.Println("Request params...")
	fmt.Printf("Operation: %s\n", req.Operation)
	fmt.Printf("Parameters: [%d,%d]\n", req.Params...)

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64)), r.Rep.Result[1].(string)
}

func (p *CalculatorProxy) Mul(p1, p2 int) (int, string) {
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	req := shared.Request{Operation: "Mul", Params: params}
	inv := shared.Invocation{Ior: p.Ior, Request: req}
	fmt.Println("Request params...")
	fmt.Printf("Operation: %s\n", req.Operation)
	fmt.Printf("Parameters: [%d,%d]\n", req.Params...)

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64)), r.Rep.Result[1].(string)
}

func (p *CalculatorProxy) LeaseExtend(T int) (int, string) {
	params := make([]interface{}, 2)
	params[0] = T
	params[1] = 0

	req := shared.Request{Operation: "LeaseExtend", Params: params}
	inv := shared.Invocation{Ior: p.Ior, Request: req}
	fmt.Println("Extending lease by:", T, "seconds")

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64)), r.Rep.Result[1].(string)
}

func (p *CalculatorProxy) LeaseTypeSet(lease string) (int, string) {
	params := make([]interface{}, 2)
	params[0] = lease
	params[1] = 0

	req := shared.Request{Operation: "LeaseType", Params: params}
	inv := shared.Invocation{Ior: p.Ior, Request: req}

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64)), r.Rep.Result[1].(string)
}

func (p *CalculatorProxy) AliveCheck(ior shared.IOR) {
	s := srh.NewSRH(ior.Host, ior.Port)

	m := marshaller.Marshaller{}

	for {
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket := m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		if r.Operation == "ReleaseWarn" {
			_p1 := r.Params[0].(string)
			fmt.Print("lease of id", _p1, "will expire\n> ")
		}

		s.Send(b) // i dont care about this answer
	}
}
