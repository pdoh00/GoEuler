package main

import (
	"math/big"
	"strconv"
)

func Problem16() interface{} {
	two := big.NewInt(2)
	oneThousand := big.NewInt(1000)
	value := two.Exp(two, oneThousand, nil).String()

	var sum int64

	for _, v := range value {
		x, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil{
			panic(err)
		}
		sum += x
	}

	return int(sum)
}
