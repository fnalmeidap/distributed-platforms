package calculatorproxy

import (
	"distributed-platforms/internal/distribution/requestor"
	"distributed-platforms/internal/shared"
)

type CalculatorProxy struct {
	Ior shared.IOR
}

func New(i shared.IOR) CalculatorProxy {
	r := CalculatorProxy{Ior: i}
	return r
}

func (p *CalculatorProxy) Sum(p1, p2 int) int {
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	req := shared.Request{Operation: "Sum", Params: params}

	inv := shared.Invocation{Ior: p.Ior, Request: req}

	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	return int(r.Rep.Result[0].(float64))
}
