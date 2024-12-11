package main

import (
	naminginvoker "distributed-platforms/internal/services/naming/invoker"

	"distributed-platforms/internal/shared"
	"fmt"
)

func main() {
	StartNamingServer()
}

func StartNamingServer() {
	fmt.Println("Naming Service running...")

	i := naminginvoker.New(shared.LocalHost, shared.NamingServicePort)
	i.Invoke()
}
