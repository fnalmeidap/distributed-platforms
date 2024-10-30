package main

import (
	"distributed-platforms/internal/infra/srh"
	shared "distributed-platforms/internal/shared"
)

func main() {
	// Instantiate application logic, invoker, and request handler
	app := &Application{}
	invoker := NewInvoker(app)
	handler := NewServiceRequestHandler(invoker)

	ior := shared.IOR{Host: "localhost", Port: 9876}
	SRH := srh.NewSRH(ior.Host, ior.Port)

}
