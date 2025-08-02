package main

import (
	"fmt"
	"math/big"
)

func Arctan(x *big.Float, prec uint) *big.Float {
	f := func() *big.Float {
		return new(big.Float).SetPrec(prec).SetMode(big.ToNearestAway)
	}

	result := f().SetFloat64(0.0)
	term := f().Set(x)
	xSquared := f().Mul(x, x)
	sign := 1

	for i := 1; ; i += 2 {
		t := f().Quo(term, f().SetFloat64(float64(i)))

		if sign > 0 {
			result.Add(result, t)
		} else {
			result.Sub(result, t)
		}

		term.Mul(term, xSquared)
		sign *= -1

		if t.Cmp(f().SetFloat64(1e-300)) < 0 {
			break
		}
	}

	return result
}

func main() {
	prec := uint(100000)
	multiplier := new(big.Float).SetPrec(prec).SetFloat64(4.0)

	x1 := new(big.Float).SetPrec(prec).SetFloat64(1.0 / 2.0)
	atan1 := Arctan(x1, prec)

	x2 := new(big.Float).SetPrec(prec).SetFloat64(1.0 / 3.0)
	atan2 := Arctan(x2, prec)
	atan1 = new(big.Float).Add(atan1, atan2)
	atan1 = new(big.Float).Mul(atan1, multiplier)

	fmt.Println(atan1.Text('f', 50))
}