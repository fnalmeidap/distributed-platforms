package main

import (
	calculatorinvoker "distributed-platforms/internal/distribution/invoker"
	namingproxy "distributed-platforms/internal/services/naming/proxy"
	shared "distributed-platforms/internal/shared"
	"fmt"
)

func StartServer(inv calculatorinvoker.Invoker) {
	inv.Invoke()
}

func main() {
	fmt.Println("Server running...")

	// Proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Instantiate application logic, invoker, and request handler
	ior := shared.IOR{Host: shared.LocalHost, Port: shared.DefaultPort}
	invoker := calculatorinvoker.NewInvoker(ior.Host, ior.Port)

	naming.Bind("calculator", shared.NewIOR(invoker.Ior.Host, invoker.Ior.Port))

	StartServer(invoker)
}
