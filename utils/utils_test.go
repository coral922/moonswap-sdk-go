package utils

import (
	"math/big"
	"testing"

	"github.com/coral922/moonswap-sdk-go/constants"
)

func TestValidateSolidityTypeInstance(t *testing.T) {
	output := ValidateSolidityTypeInstance(big.NewInt(-1), constants.Uint8)
	if output == nil {
		t.Errorf("should be an error")
	}

	output = ValidateSolidityTypeInstance(big.NewInt(255), constants.Uint8)
	if output != nil {
		t.Errorf("error should be nil")
	}
}
