package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// create a new DelegatorStartingInfo
func NewDelegatorStartingInfo(previousPeriod uint64, stake sdk.Dec, height uint64, cumulativeRewards sdk.Dec) DelegatorStartingInfo {
	return DelegatorStartingInfo{
		PreviousPeriod:    previousPeriod,
		Stake:             stake,
		Height:            height,
		CumulativeRewards: cumulativeRewards,
	}
}
