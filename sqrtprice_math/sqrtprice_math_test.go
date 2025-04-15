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

// func Test_GetAmount0DeltaNumerator(t *testing.T) {

// 	sqrtPricex96, _ := uint256.FromBig(mustBigInt("271628393057704077821457785659010"))
// 	liquidityDelta, _ := uint256.FromBig(mustBigInt("44860547286813223"))
// 	sqrtAtTick, _ := uint256.FromBig(mustBigInt("271692128865584629052887551158270"))
// 	num1, num2 := GetAmount0DeltaNumerator(sqrtPricex96, sqrtAtTick, liquidityDelta)
// 	require.Equal(t, num1.ToBig().String(), "3554218730918478130122983419778198180020092928")
// 	require.Equal(t, num2.ToBig().String(), "63735807880551231429765499260")

// }
