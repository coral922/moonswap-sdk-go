package entities

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/coral922/moonswap-sdk-go/constants"
	"github.com/coral922/moonswap-sdk-go/utils"
)

var (
	ErrDiffChainID = fmt.Errorf("diff chain id")
	ErrDiffToken   = fmt.Errorf("diff token")
	ErrSameAddrss  = fmt.Errorf("same address")

	_WCFXCurrency, _ = newCurrency(constants.Decimals18, "WCFX", "Wrapped Conflux")

	WCFX = map[constants.ChainID]*Token{
		constants.Mainnet: {
			Currency: _WCFXCurrency,
			ChainID:  constants.Mainnet,
			Address:  utils.ValidateAndParseMainnetAddress("cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx"),
		},
	}
)

/**
 * Represents an ERC20 token with a unique address and some metadata.
 */
type Token struct {
	*Currency

	constants.ChainID
	common.Address
}

func NewToken(chainID constants.ChainID, address common.Address, decimals int, symbol, name string) (*Token, error) {
	currency, err := newCurrency(decimals, symbol, name)
	if err != nil {
		return nil, err
	}

	return &Token{
		Currency: currency,
		ChainID:  chainID,
		Address:  address,
	}, nil
}

/**
 * Returns true if the two tokens are equivalent, i.e. have the same chainId and address.
 * @param other other token to compare
 */
func (t *Token) Equals(other *Token) bool {
	if t == other {
		return true
	}

	return t.ChainID == other.ChainID && t.Address == other.Address
}

/**
 * Returns true if the address of this token sorts before the address of the other token
 * @param other other token to compare
 * @throws if the tokens have the same address
 * @throws if the tokens are on different chains
 */
func (t *Token) SortsBefore(other *Token) (bool, error) {
	if t.ChainID != other.ChainID {
		return false, ErrDiffChainID
	}
	if t.Address == other.Address {
		return false, ErrSameAddrss
	}

	return strings.ToLower(t.Address.String()) < strings.ToLower(other.Address.String()), nil
}

// NewCFXToken creates a token that currency is CFX
func NewCFXToken(chainID constants.ChainID, address common.Address) *Token {
	return &Token{
		Currency: CONFLUX,
		ChainID:  chainID,
		Address:  address,
	}
}
