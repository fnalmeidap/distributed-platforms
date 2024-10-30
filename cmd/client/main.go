package main

import (
	calculatorproxy "distributed-platforms/internal/distribution/proxy"
	shared "distributed-platforms/internal/shared"
)

func main() {
	ior := shared.IOR{Host: "localhost", Port: 8080}
	c := calculatorproxy.New(ior)

	c.Sum(1, 2)
}
