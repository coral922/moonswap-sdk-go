package utils

import (
	"fmt"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/coral922/moonswap-sdk-go/constants"
)

// ValidateSolidityTypeInstance determines if a value is a legal SolidityType
func ValidateSolidityTypeInstance(value *big.Int, t constants.SolidityType) error {
	if value.Cmp(constants.Zero) < 0 || value.Cmp(constants.SolidityTypeMaxima[t]) > 0 {
		return fmt.Errorf(`%v is not a %s`, value, t)
	}
	return nil
}

// ValidateAndParseMainnetAddress warns if addresses are not checksummed
func ValidateAndParseMainnetAddress(address string) common.Address {
	cAddr := cfxaddress.MustNew(address, uint32(constants.Mainnet))
	return cAddr.MustGetCommonAddress()
}
