package fullmath

import (
	"math/big"
	"testing"

	ui "github.com/ethcf/uniswap/uint256"
)

func TestMulDiv(t *testing.T) {
	tests := []struct {
		name        string
		a, b, denom string
		expected    string
	}{
		{
			name:     "zero multiplication",
			a:        "0",
			b:        "500",
			denom:    "1000000",
			expected: "0",
		},
		{
			name:     "basic division",
			a:        "1000",
			b:        "500",
			denom:    "100",
			expected: "5000",
		},
		{
			name:     "large numbers",
			a:        "1000000000000000000",
			b:        "2000000000000000000",
			denom:    "1000000000000000000",
			expected: "2000000000000000000",
		},
		{
			name:     "fractional result truncated",
			a:        "7",
			b:        "3",
			denom:    "2",
			expected: "10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := uintFromString(tt.a)
			b := uintFromString(tt.b)
			denom := uintFromString(tt.denom)
			expected := uintFromString(tt.expected)

			result := MulDiv(a, b, denom)
			if result.Cmp(expected) != 0 {
				t.Errorf("MulDiv(%s, %s, %s) = %s, want %s", tt.a, tt.b, tt.denom, result.String(), tt.expected)
			}
		})
	}
}

func TestMulDivRoundingUp(t *testing.T) {
	tests := []struct {
		name        string
		a, b, denom string
		expected    string
	}{
		{
			name:     "zero multiplication",
			a:        "0",
			b:        "500",
			denom:    "1000000",
			expected: "0",
		},
		{
			name:     "exact division no rounding",
			a:        "1000",
			b:        "500",
			denom:    "100",
			expected: "5000",
		},
		{
			name:     "fractional result rounded up",
			a:        "7",
			b:        "3",
			denom:    "2",
			expected: "11",
		},
		{
			name:     "large numbers with remainder",
			a:        "1000000000000000001",
			b:        "1000000000000000001",
			denom:    "1000000000000000000",
			expected: "1000000000000000003",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := uintFromString(tt.a)
			b := uintFromString(tt.b)
			denom := uintFromString(tt.denom)
			expected := uintFromString(tt.expected)

			result := MulDivRoundingUp(a, b, denom)
			if result.Cmp(expected) != 0 {
				t.Errorf("MulDivRoundingUp(%s, %s, %s) = %s, want %s", tt.a, tt.b, tt.denom, result.String(), tt.expected)
			}
		})
	}
}

func TestUnsafeMath_DivRoundingUp(t *testing.T) {
	tests := []struct {
		name     string
		x, y     string
		expected string
	}{
		{
			name:     "exact division",
			x:        "10",
			y:        "2",
			expected: "5",
		},
		{
			name:     "division with remainder rounds up",
			x:        "10",
			y:        "3",
			expected: "4",
		},
		{
			name:     "large numbers exact",
			x:        "1000000000000000000",
			y:        "1000000000000000000",
			expected: "1",
		},
		{
			name:     "large numbers with remainder",
			x:        "1000000000000000001",
			y:        "1000000000000000000",
			expected: "2",
		},
		{
			name:     "small remainder",
			x:        "101",
			y:        "100",
			expected: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := uintFromString(tt.x)
			y := uintFromString(tt.y)
			expected := uintFromString(tt.expected)

			result := UnsafeMath_DivRoundingUp(x, y)
			if result.Cmp(expected) != 0 {
				t.Errorf("UnsafeMath_DivRoundingUp(%s, %s) = %s, want %s", tt.x, tt.y, result.String(), tt.expected)
			}
		})
	}
}

func uintFromString(s string) *ui.Int {
	big_d, _ := new(big.Int).SetString(s, 10)
	result, _ := ui.FromBig(big_d)
	return result
}
