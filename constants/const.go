package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

type TradeType int

const (
	ExactInput TradeType = iota
	ExactOutput
)

type Rounding int

const (
	RoundDown Rounding = iota
	RoundHalfUp
	RoundUp
)

// Valid check this rounding mode is valid
func (r Rounding) Valid() bool {
	return r == RoundDown ||
		r == RoundHalfUp ||
		r == RoundUp
}

const (
	Decimals18     = 18
	MoonSwapSymbol = "MLP"
	MoonSwapName   = "MoonSwap LP Token"
)

var (
	MinimumLiquidity = big.NewInt(1000)

	Zero  = big.NewInt(0)
	One   = big.NewInt(1)
	Two   = big.NewInt(2)
	Three = big.NewInt(3)
	Five  = big.NewInt(5)
	Ten   = big.NewInt(10)

	B100  = big.NewInt(100)
	B997  = big.NewInt(997)
	B1000 = big.NewInt(1000)
)

type SolidityType string

const (
	Uint8   SolidityType = "uint8"
	Uint256 SolidityType = "uint256"
)

var (
	SolidityTypeMaxima = map[SolidityType]*big.Int{
		Uint8:   big.NewInt(0xff),
		Uint256: math.MaxBig256,
	}
)

var (
	FactoryAddress = common.HexToAddress("0x865f55a399bf9250ae781adfbed71e70c12bd2d8")
	InitCodeHash   = common.FromHex("a6330451e4d6d3fc19f31fc5ee71147d88812b0da79f64b03ed210fd594d84e9")
)
