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

// PerformOperation calls the correct application method based on the operation name.
/*func (app *Calculator) PerformOperation(req message.CalculationRequest) (message.CalculationResponse, error) {
	switch req.Operation {
	case "sum":
		result := app.Sum(req.Number1, req.Number2)
		return message.CalculationResponse{Result: result}, nil
	default:
		return message.CalculationResponse{}, fmt.Errorf("unsupported operation")
	}
}
*/
