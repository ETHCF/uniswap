package sqrtprice_math

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/numbergroup/uniswap/uint256"
)

func mustBigInt(value string) *big.Int {
	bi, ok := new(big.Int).SetString(value, 10)
	if !ok {
		panic("failed to create big.Int")
	}
	return bi
}

func Test_GetAmount0Delta(t *testing.T) {

	sqrtPricex96, _ := uint256.FromBig(mustBigInt("271628393057704077821457785659010"))
	sqrtAtTick, _ := uint256.FromBig(mustBigInt("271692128865584629052887551158270"))
	liquidityDelta, _ := uint256.FromBig(mustBigInt("44860547286813223"))

	result := GetAmount0Delta(sqrtPricex96, sqrtAtTick, liquidityDelta, true)
	require.Equal(t, result.ToBig().String(), "3069555042")

	result = GetAmount0Delta(sqrtPricex96, sqrtAtTick, liquidityDelta, false)
	require.Equal(t, result.ToBig().String(), "3069555041")

}

func Test_GetAmount0DeltaNumerator(t *testing.T) {
	sqrtPricex96, _ := uint256.FromBig(mustBigInt("271628393057704077821457785659010"))
	liquidityDelta, _ := uint256.FromBig(mustBigInt("44860547286813223"))
	sqrtAtTick, _ := uint256.FromBig(mustBigInt("271692128865584629052887551158270"))
	num1, num2 := GetAmount0DeltaNumerator(sqrtPricex96, sqrtAtTick, liquidityDelta)
	require.Equal(t, num1.ToBig().String(), "3554218730918478130122983419778198180020092928")
	require.Equal(t, num2.ToBig().String(), "63735807880551231429765499260")
}

func Test_GetAmount1Delta(t *testing.T) {
	sqrtPricex96, _ := uint256.FromBig(mustBigInt("271628393057704077821457785659010"))
	sqrtAtTick, _ := uint256.FromBig(mustBigInt("271692128865584629052887551158270"))
	liquidityDelta, _ := uint256.FromBig(mustBigInt("44860547286813223"))

	result := GetAmount1Delta(sqrtPricex96, sqrtAtTick, liquidityDelta, true)
	require.Equal(t, result.ToBig().String(), "36088470722439553")

	result = GetAmount1Delta(sqrtPricex96, sqrtAtTick, liquidityDelta, false)
	require.Equal(t, result.ToBig().String(), "36088470722439552")
}

func Test_GetNextSqrtPriceFromInput(t *testing.T) {
	sqrtPX96, _ := uint256.FromBig(mustBigInt("79228162514264337593543950336"))
	liquidity, _ := uint256.FromBig(mustBigInt("1000000000000000000"))
	amountIn, _ := uint256.FromBig(mustBigInt("1000000"))

	result := GetNextSqrtPriceFromInput(sqrtPX96, liquidity, amountIn, true)
	expected := "79228162514185109431029765227"
	require.Equal(t, result.ToBig().String(), expected)

	result = GetNextSqrtPriceFromInput(sqrtPX96, liquidity, amountIn, false)
	expected = "79228162514343565756058214673"
	require.Equal(t, result.ToBig().String(), expected)
}

func Test_GetNextSqrtPriceFromOutput(t *testing.T) {
	sqrtPX96, _ := uint256.FromBig(mustBigInt("79228162514264337593543950336"))
	liquidity, _ := uint256.FromBig(mustBigInt("1000000000000000000"))
	amountOut, _ := uint256.FromBig(mustBigInt("1000000"))

	result := GetNextSqrtPriceFromOutput(sqrtPX96, liquidity, amountOut, true)
	expected := "79228162514185109431029685998"
	require.Equal(t, result.ToBig().String(), expected)

	result = GetNextSqrtPriceFromOutput(sqrtPX96, liquidity, amountOut, false)
	expected = "79228162514343565756058293902"
	require.Equal(t, result.ToBig().String(), expected)
}

func Test_GetAmount0DeltaRounded(t *testing.T) {
	sqrtPricex96, _ := uint256.FromBig(mustBigInt("271628393057704077821457785659010"))
	sqrtAtTick, _ := uint256.FromBig(mustBigInt("271692128865584629052887551158270"))
	liquidityDelta, _ := uint256.FromBig(mustBigInt("44860547286813223"))

	result := GetAmount0DeltaRounded(sqrtPricex96, sqrtAtTick, liquidityDelta)
	require.Equal(t, result.ToBig().String(), "3069555042")

	negativeLiquidity, _ := uint256.FromBig(mustBigInt("-44860547286813223"))
	result = GetAmount0DeltaRounded(sqrtPricex96, sqrtAtTick, negativeLiquidity)
	require.Equal(t, result.ToBig().String(), "115792089237316195423570985008687907853269984665640564039457584007910060084895")
}

func Test_GetAmount1DeltaRounded(t *testing.T) {
	sqrtPricex96, _ := uint256.FromBig(mustBigInt("271628393057704077821457785659010"))
	sqrtAtTick, _ := uint256.FromBig(mustBigInt("271692128865584629052887551158270"))
	liquidityDelta, _ := uint256.FromBig(mustBigInt("44860547286813223"))

	result := GetAmount1DeltaRounded(sqrtPricex96, sqrtAtTick, liquidityDelta)
	require.Equal(t, result.ToBig().String(), "36088470722439553")

	negativeLiquidity, _ := uint256.FromBig(mustBigInt("-44860547286813223"))
	result = GetAmount1DeltaRounded(sqrtPricex96, sqrtAtTick, negativeLiquidity)
	require.Equal(t, result.ToBig().String(), "115792089237316195423570985008687907853269984665640564039457547919442407200384")
}
