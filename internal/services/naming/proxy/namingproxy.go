package namingproxy

import (
	"distributed-platforms/internal/distribution/requestor"
	"distributed-platforms/internal/shared"
)

type NamingProxy struct {
	Ior shared.IOR
}

func New(h string, p int) NamingProxy {
	i := shared.IOR{Host: h, Port: p}
	r := NamingProxy{Ior: i}
	return r
}

func (h *NamingProxy) Bind(_p1 string, _p2 shared.IOR) bool {

	// 1. Configure input parameters
	params := make([]interface{}, 2)
	params[0] = _p1
	params[1] = _p2

	// Configure remote request
	req := shared.Request{Operation: "Bind", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return r.Rep.Result[0].(bool)
}

func (h *NamingProxy) Find(_p1 string) shared.IOR {

	// 1. Configure input parameters
	params := make([]interface{}, 1)
	params[0] = _p1

	// Configure remote request
	req := shared.Request{Operation: "Find", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// Invoke Requestor
	requestor := requestor.Requestor{}
	_r1 := requestor.Invoke(inv).Rep.Result
	_r2 := _r1[0].(map[string]interface{})

	//4. Return something to the publisher
	_ior := shared.IOR{}
	_ior.Host = _r2["Host"].(string)
	_ior.Port = int(_r2["Port"].(float64))
	_ior.Id = int(_r2["Id"].(float64))
	_ior.TypeName = _r2["TypeName"].(string)
	_ior.LeaseName = _r2["LeaseName"].(string)

	return _ior
}

func (h *NamingProxy) List() []shared.IOR {

	// 1. Configure input parameters
	params := make([]interface{}, 1)
	params[0] = "" // no parameters

	// Configure remote request
	req := shared.Request{Operation: "List", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// Invoke Requestor
	requestor := requestor.Requestor{}
	_r1 := requestor.Invoke(inv).Rep.Result
	_r2 := _r1[0].(map[string]interface{})

	_r3 := []shared.IOR{}
	for key, value := range _r2 {
		_r4 := value.(map[string]interface{})
		tempIor := shared.IOR{}
		tempIor.TypeName = key
		tempIor.Host = _r4["Host"].(string)
		tempIor.Port = int(_r4["Port"].(float64))
		tempIor.Id = int(_r4["Id"].(float64))
		_r3 = append(_r3, tempIor)
	}

	//4. Return something to the publisher
	return _r3
}

func (h *NamingProxy) Unbind(_p1 string) bool {

	// 1. Configure input parameters
	params := make([]interface{}, 1)
	params[0] = _p1

	// Configure remote request
	req := shared.Request{Operation: "Unbind", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return r.Rep.Result[0].(bool)
}
