package constants

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"strconv"
)

// ChainID network chain id
type ChainID int

const (
	Mainnet = ChainID(int(cfxaddress.NetowrkTypeMainnetID))
)

func (i ChainID) String() string {
	switch {
	case i == 1029:
		return "Mainnet"
	default:
		return "ChainID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
