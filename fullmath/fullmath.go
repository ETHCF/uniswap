package fullmath

import (
	cons "github.com/ethcf/uniswap/constants"
	ui "github.com/ethcf/uniswap/uint256"
)

func MulDivRoundingUp(a, b, denominator *ui.Int) *ui.Int {
	if a.IsZero() || b.IsZero() {
		return ui.NewInt(0)
	}
	product := ui.Umul(a, b)
	var q [5]uint64
	quotient := q[:]
	rem := ui.Udivrem(quotient, product[:], denominator)
	result := (*ui.Int)(quotient[0:4])
	if !rem.IsZero() {
		result.Add(result, cons.One)
	}
	return result
}

func MulDiv(a, b, denominator *ui.Int) *ui.Int {
	if a.IsZero() || b.IsZero() {
		return ui.NewInt(0)
	}
	product := ui.Umul(a, b)
	var q [5]uint64
	quotient := q[:]
	ui.Udivrem(quotient, product[:], denominator)
	result := (*ui.Int)(quotient[0:4])
	return result
}

func UnsafeMath_DivRoundingUp(x, y *ui.Int) *ui.Int {
	z := new(ui.Int).Div(x, y)
	if new(ui.Int).Mod(x, y).IsZero() {
		return z
	}
	z.Add(z, cons.One)
	return z
}
