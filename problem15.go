package main

import (
	"math/big"
)

func Problem15() interface{} {
	//The number of paths are the central binomial coefficients
	// (2n)! / n!^2 where nxn is the dimension of the graph

	//for graph of 20x20 => n=20 = 2n = 40
	n := big.NewInt(20)
	var twoN = new(big.Int)
	twoN.Mul(n, big.NewInt(2))

	numerator := factorial(twoN)

	nBang := factorial(n)
	denominator := nBang.Exp(nBang, big.NewInt(2), nil)

	answer := numerator.Div(numerator, denominator)

	return answer.Int64()
}

func factorial(n *big.Int) *big.Int {
	var result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}

	return result
}
