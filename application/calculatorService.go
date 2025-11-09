package app

type Calculator interface {
	Addition(a, b float32) float32
	Subtraction(a, b float32) float32
	Multiplication(a, b float32) float32
	Division(a, b float32) float32
}

type Calci struct {
}

type RowCalci struct {
	Arr  []float32
	Cacl Calci
}

func NewCalculator() Calci {
	return Calci{}
}

func (c Calci) Addition(a, b float32) float32 {
	return a + b
}
func (c Calci) Subtraction(a, b float32) float32 {
	return a - b
}
func (c Calci) Multiplication(a, b float32) float32 {
	return a * b
}
func (c Calci) Division(a, b float32) float32 {
	return a / b
}

func (c RowCalci) RowAddition() float32 {
	if len(c.Arr) < 2 {
		return helper(c)
	}
	output := float32(0)
	for _, num := range c.Arr {
		output = c.Cacl.Addition(output, num)
	}
	return output
}

func (c RowCalci) RowMultiplication() float32 {
	if len(c.Arr) < 2 {
		return helper(c)
	}

	output := float32(1)
	for _, num := range c.Arr {
		output = c.Cacl.Multiplication(output, num)
	}
	return output
}

func helper(c RowCalci) float32 {
	if len(c.Arr) == 1 {
		return c.Arr[0]
	}
	return 0
}
