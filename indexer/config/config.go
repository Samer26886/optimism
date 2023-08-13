package config

import (
	"os"
	"reflect"

	"github.com/BurntSushi/toml"

	"github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/common"
)

// Config represents the `indexer.toml` file used to configure the indexer
type Config struct {
	Chain   ChainConfig
	RPCs    RPCsConfig `toml:"rpcs"`
	DB      DBConfig
	API     APIConfig
	Metrics MetricsConfig
	Logger  log.CLIConfig
}

// fetch this via onchain config from RPCsConfig and remove from config in future
type L1Contracts struct {
	OptimismPortal         common.Address
	L2OutputOracle         common.Address
	L1CrossDomainMessenger common.Address
	L1StandardBridge       common.Address
	L1ERC721Bridge         common.Address

	// Some more contracts -- ProxyAdmin, SystemConfig, etcc
	// Ignore the auxiliary contracts?

	// Legacy contracts? We'll add this in to index the legacy chain.
	// Remove afterwards?
}

func (c L1Contracts) ToSlice() []common.Address {
	fields := reflect.VisibleFields(reflect.TypeOf(c))
	v := reflect.ValueOf(c)

	contracts := make([]common.Address, len(fields))
	for i, field := range fields {
		contracts[i] = (v.FieldByName(field.Name).Interface()).(common.Address)
	}

	return contracts
}

// ChainConfig configures of the chain being indexed
type ChainConfig struct {
	// Configure known chains with the l2 chain id
	Preset      int
	L1Contracts L1Contracts
}

// RPCsConfig configures the RPC urls
type RPCsConfig struct {
	L1RPC string `toml:"l1-rpc"`
	L2RPC string `toml:"l2-rpc"`
}

// DBConfig configures the postgres database
type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// APIConfig configures the API server
type APIConfig struct {
	Host string
	Port int
}

// MetricsConfig configures the metrics server
type MetricsConfig struct {
	Host string
	Port int
}

// LoadConfig loads the `indexer.toml` config file from a given path
func LoadConfig(path string) (Config, error) {
	var conf Config

	// Read the config file.
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	// Replace environment variables.
	data = []byte(os.ExpandEnv(string(data)))

	// Decode the TOML data.
	if _, err := toml.Decode(string(data), &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
