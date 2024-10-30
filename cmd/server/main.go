package main

import (
	calculatorinvoker "distributed-platforms/internal/distribution/invoker"
	shared "distributed-platforms/internal/shared"
)

func StartServer(inv calculatorinvoker.Invoker) {
	inv.Invoke()
}

func main() {
	// Instantiate application logic, invoker, and request handler
	ior := shared.IOR{Host: shared.LocalHost, Port: shared.DefaultPort}
	invoker := calculatorinvoker.NewInvoker(ior.Host, ior.Port)

	StartServer(invoker)
}
