package main

import (
	calculatorproxy "distributed-platforms/internal/distribution/proxy"
	shared "distributed-platforms/internal/shared"
	"fmt"
)

func main() {
	ior := shared.IOR{Host: shared.LocalHost, Port: shared.DefaultPort}
	c := calculatorproxy.New(ior)

	ans := c.Sum(1, 2)
	fmt.Println("Answer:", ans)

	ans = c.Sub(1, 2)
	fmt.Println("Answer:", ans)

	ans = c.Mul(1, 2)
	fmt.Println("Answer:", ans)

	ans = c.Div(1, 2)
	fmt.Println("Answer:", ans)
}
