package rollup

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/superchain-registry/superchain"
)

const (
	opMainnet  = 10
	opGoerli   = 420
	opSepolia  = 11155420
	baseGoerli = 84531
	baseMainnet = 8453
	pgnMainnet = 424
	pgnSepolia = 58008
	zoraGoerli = 999
	zoraMainnet = 7777777
)


// SystemConfigProvider is an interface to retrieve SystemConfig variables, as registered in the onchain smart contract.
type SystemConfigProvider interface {
	GenesisSystemConfig() eth.SystemConfig
	DepositContractAddress() common.Address
}

// LoadOPStackRollupConfig loads the rollup configuration of the requested chain ID from the superchain-registry.
// Some chains may require a SystemConfigProvider to retrieve any values not part of the registry.
func LoadOPStackRollupConfig(sysConfig SystemConfigProvider, chainID uint64) (*Config, error) {
	chConfig, ok := superchain.OPChains[chainID]
	if !ok {
		return nil, fmt.Errorf("unknown chain ID: %d", chainID)
	}

	superChain, ok := superchain.Superchains[chConfig.Superchain]
	if !ok {
		return nil, fmt.Errorf("unknown superchain: %q", chConfig.Superchain)
	}

	var genesisSysConfig eth.SystemConfig
	if sysCfg, ok := superchain.GenesisSystemConfigs[chainID]; ok {
		genesisSysConfig = eth.SystemConfig{
			BatcherAddr: common.Address(sysCfg.BatcherAddr),
			Overhead:    eth.Bytes32(sysCfg.Overhead),
			Scalar:      eth.Bytes32(sysCfg.Scalar),
			GasLimit:    sysCfg.GasLimit,
		}
	} else if sysConfig != nil {
		genesisSysConfig = sysConfig.GenesisSystemConfig()
	} else {
		return nil, fmt.Errorf("unable to retrieve genesis SystemConfig")
	}

	var depositContractAddress common.Address
	if addrs, ok := superchain.Addresses[chainID]; ok {
		depositContractAddress = common.Address(addrs.OptimismPortalProxy)
	} else if sysConfig != nil {
		depositContractAddress = sysConfig.DepositContractAddress()
	} else {
		return nil, fmt.Errorf("unable to retrieve deposit contract address")
	}

	regolithTime := uint64(0)
	// two goerli testnets test-ran Bedrock and later upgraded to Regolith.
	// All other OP-Stack chains have Regolith enabled from the start.
	switch chainID {
	case baseGoerli:
		regolithTime = 1683219600
	case opGoerli:
		regolithTime = 1679079600
	}

	cfg := &Config{
		Genesis: Genesis{
			L1: eth.BlockID{
				Hash:   common.Hash(chConfig.Genesis.L1.Hash),
				Number: chConfig.Genesis.L1.Number,
			},
			L2: eth.BlockID{
				Hash:   common.Hash(chConfig.Genesis.L2.Hash),
				Number: chConfig.Genesis.L2.Number,
			},
			L2Time:       chConfig.Genesis.L2Time,
			SystemConfig: genesisSysConfig,
		},
		// The below chain parameters can be different per OP-Stack chain,
		// but since none of the superchain chains differ, it's not represented in the superchain-registry yet.
		// This restriction on superchain-chains may change in the future.
		// Test/Alt configurations can still load custom rollup-configs when necessary.
		BlockTime:              2,
		MaxSequencerDrift:      600,
		SeqWindowSize:          3600,
		ChannelTimeout:         300,
		L1ChainID:              new(big.Int).SetUint64(superChain.Config.L1.ChainID),
		L2ChainID:              new(big.Int).SetUint64(chConfig.ChainID),
		RegolithTime:           &regolithTime,
		BatchInboxAddress:      common.Address(chConfig.BatchInboxAddr),
		DepositContractAddress: depositContractAddress,
		L1SystemConfigAddress:  common.Address(chConfig.SystemConfigAddr),
	}
	return cfg, nil
}
