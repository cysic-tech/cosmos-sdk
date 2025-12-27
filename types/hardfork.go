package types

import (
	"strings"
)

type HardforkConfig struct {
	LazyDistributionHeight int64
}

// GetHardforkConfig returns the hardfork configuration for a given chain-id.
func GetHardforkConfig(chainID string) HardforkConfig {
	if strings.HasPrefix(chainID, "cysic_4398-1") {
		// testnet config
		return HardforkConfig{
			LazyDistributionHeight: 345000,
		}
	}

	if strings.HasPrefix(chainID, "cysicmint_4399-1") {
		// mainnet config
		return HardforkConfig{
			LazyDistributionHeight: 235000,
		}
	}

	// default config (local development networks, etc.)
	return HardforkConfig{
		LazyDistributionHeight: 0,
	}
}

// IsLazyDistributionEnabled checks if the lazy distribution hardfork is enabled.
func IsLazyDistributionEnabled(ctx Context) bool {
	config := GetHardforkConfig(ctx.ChainID())
	return ctx.BlockHeight() >= config.LazyDistributionHeight
}
