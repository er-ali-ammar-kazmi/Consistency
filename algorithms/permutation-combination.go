package algorithms

func Permutation(n, r int) int {
	// nPr - n!/(n-r)!
	if r > n {
		return 0
	}
	return Fact(n) / Fact(n-r)
}

func Combination(n, r int) int {
	// nCr - n!/((n-r)!*r!)
	if r > n {
		return 0
	}
	return Fact(n) / (Fact(n-r) * Fact(r))
}
