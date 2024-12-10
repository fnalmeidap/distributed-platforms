package main

import (
	naminginvoker "distributed-platforms/internal/services/naming/invoker"

	"distributed-platforms/internal/shared"
	"fmt"
)

func main() {

	go namingServer()

	fmt.Println("'Servidor de Nomes' em execução...")
	fmt.Scanln()
}

func namingServer() {
	// Start messagingservice invoker
	i := naminginvoker.New(shared.LocalHost, shared.NamingPort)
	go i.Invoke()
}
