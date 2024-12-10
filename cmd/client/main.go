package main

import (
	calculatorproxy "distributed-platforms/internal/distribution/proxy"
	namingproxy "distributed-platforms/internal/services/naming/proxy"
	shared "distributed-platforms/internal/shared"
	"fmt"
)

func main() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	ior := naming.Find("Calculadora")
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
