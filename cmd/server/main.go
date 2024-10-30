package main

import (
	calculator "distributed-platforms/internal/app/calculator"
	messaginginvoker "distributed-platforms/internal/distribution/invoker"
	srh "distributed-platforms/internal/infra/srh"
	shared "distributed-platforms/internal/shared"
	"fmt"
)

func StartServer(srh *srh.SRH, inv messaginginvoker.Invoker) {

	fmt.Println("Server listening on", srh.Host)

	inv.Invoke()

}

func main() {
	// Instantiate application logic, invoker, and request handler
	ior := shared.IOR{Host: "localhost", Port: 9876}

	app := &calculator.Calculator{}
	invoker := messaginginvoker.NewInvoker(ior.Host, ior.Port, app)
	SRH := srh.NewSRH(ior.Host, ior.Port)
	//handler := NewServiceRequestHandler(invoker)
	StartServer(SRH, invoker)

}
