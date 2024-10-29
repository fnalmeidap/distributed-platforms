package calculator

type Calculator struct{}

func (Calculator) Sum(p1, p2 int) int {
	return p1 + p2
}

func (Calculator) Sub(p1, p2 int) int {
	return p1 - p2
}

func (Calculator) Mul(p1, p2 int) int {
	return p1 * p2
}

func (Calculator) Div(p1, p2 int) int {
	if p2 == 0 {
		return 0
	}

	return p1 / p2
}