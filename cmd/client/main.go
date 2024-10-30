package main

import (
	calculatorproxy "distributed-platforms/internal/distribution/proxy"
	shared "distributed-platforms/internal/shared"
)

func main() {
	ior := shared.IOR{Host: "localhost", Port: 9876}
	c := calculatorproxy.New(ior)

	c.Sum(1, 2)
}
