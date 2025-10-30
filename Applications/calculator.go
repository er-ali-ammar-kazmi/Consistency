package app

type Calculator interface {
	Addition(a, b float64) float64
	Subtraction(a, b float64) float64
	Multiplication(a, b float64) float64
	Division(a, b float64) float64
}

type Calci struct {
}

type RowCalci struct {
	Arr  []float64
	Cacl Calci
}

func NewCalculator() Calci {
	return Calci{}
}

func (c Calci) Addition(a, b float64) float64 {
	return a + b
}
func (c Calci) Subtraction(a, b float64) float64 {
	return a - b
}
func (c Calci) Multiplication(a, b float64) float64 {
	return a * b
}
func (c Calci) Division(a, b float64) float64 {
	return a / b
}

func (c RowCalci) RowAddition() float64 {
	if len(c.Arr) < 2 {
		return helper(c)
	}
	output := float64(0)
	for _, num := range c.Arr {
		output = c.Cacl.Addition(output, num)
	}
	return output
}

func (c RowCalci) RowMultiplication() float64 {
	if len(c.Arr) < 2 {
		return helper(c)
	}

	output := float64(1)
	for _, num := range c.Arr {
		output = c.Cacl.Multiplication(output, num)
	}
	return output
}

func helper(c RowCalci) float64 {
	if len(c.Arr) == 1 {
		return c.Arr[0]
	}
	return 0
}
