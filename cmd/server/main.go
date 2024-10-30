package main

import (
	calculator "distributed-platforms/internal/app/calculator"
	messaginginvoker "distributed-platforms/internal/distribution/invoker"
	lease "distributed-platforms/internal/lease"
	shared "distributed-platforms/internal/shared"
)

func StartServer(inv messaginginvoker.Invoker) {
	inv.Invoke()
}

func main() {
	// Instantiate application logic, invoker, and request handler
	ior := shared.IOR{Host: shared.LocalHost, Port: shared.DefaultPort}
	leaseManager := &lease.LeaseManager{Leases: make(map[string]lease.Lease)}
	app := &calculator.Calculator{}
	invoker := messaginginvoker.NewInvoker(ior.Host, ior.Port, app, leaseManager)

	StartServer(invoker)
}
