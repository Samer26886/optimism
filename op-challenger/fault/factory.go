package fault

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// MinimalDisputeGameFactoryCaller is a minimal interface around [bindings.DisputeGameFactoryCaller].
// This needs to be updated if the [bindings.DisputeGameFactoryCaller] interface changes.
type MinimalDisputeGameFactoryCaller interface {
	GameCount(opts *bind.CallOpts) (*big.Int, error)
	GameAtIndex(opts *bind.CallOpts, _index *big.Int) (struct {
		Proxy     common.Address
		Timestamp *big.Int
	}, error)
}

type FaultDisputeGame struct {
	Proxy     common.Address
	Timestamp *big.Int
}

// GameLoader is a minimal interface for fetching on chain dispute games.
type GameLoader interface {
	FetchAllGames(ctx context.Context) ([]FaultDisputeGame, error)
}

type gameLoader struct {
	caller MinimalDisputeGameFactoryCaller
}

// NewGameLoader creates a new services that can be used to fetch on chain dispute games.
func NewGameLoader(caller MinimalDisputeGameFactoryCaller) *gameLoader {
	return &gameLoader{
		caller: caller,
	}
}

// FetchAllGames fetches all dispute games from the factory.
// todo(refcell): batch call this to give mr. rpc a break.
func (l *gameLoader) FetchAllGames(ctx context.Context) ([]FaultDisputeGame, error) {
	gameCount, err := l.caller.GameCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	games := make([]FaultDisputeGame, gameCount.Uint64())
	for i := uint64(0); i < gameCount.Uint64(); i++ {
		game, err := l.caller.GameAtIndex(&bind.CallOpts{Context: ctx}, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		games[i] = FaultDisputeGame{
			Proxy:     game.Proxy,
			Timestamp: game.Timestamp,
		}
	}

	return games, nil
}
