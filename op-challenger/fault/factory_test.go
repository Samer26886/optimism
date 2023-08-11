package fault

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

var (
	gameCountErr = errors.New("game count error")
	gameIndexErr = errors.New("game index error")
)

// TestGameLoader_FetchAllGames tests that the game loader correctly fetches all games.
func TestGameLoader_FetchAllGames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		caller      *mockMinimalDisputeGameFactoryCaller
		expectedErr error
		expectedLen int
	}{
		{
			name:        "success",
			caller:      newMockMinimalDisputeGameFactoryCaller(10, false, false),
			expectedErr: nil,
			expectedLen: 10,
		},
		{
			name:        "game count error",
			caller:      newMockMinimalDisputeGameFactoryCaller(10, true, false),
			expectedErr: gameCountErr,
			expectedLen: 0,
		},
		{
			name:        "game index error",
			caller:      newMockMinimalDisputeGameFactoryCaller(10, false, true),
			expectedErr: gameIndexErr,
			expectedLen: 0,
		},
		{
			name:        "no games",
			caller:      newMockMinimalDisputeGameFactoryCaller(0, false, false),
			expectedErr: nil,
			expectedLen: 0,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			loader := NewGameLoader(test.caller)

			games, err := loader.FetchAllGames(context.Background())
			require.Equal(t, test.expectedErr, err)
			require.Len(t, games, test.expectedLen)
			expectedGames := test.caller.games
			if test.expectedErr != nil {
				expectedGames = make([]struct {
					proxy     common.Address
					timestamp *big.Int
				}, 0)
			}
			require.ElementsMatch(t, expectedGames, translateGames(games))
		})
	}
}

func generateMockGames(count uint64) []struct {
	proxy     common.Address
	timestamp *big.Int
} {
	games := make([]struct {
		proxy     common.Address
		timestamp *big.Int
	}, count)

	for i := uint64(0); i < count; i++ {
		games[i] = struct {
			proxy     common.Address
			timestamp *big.Int
		}{
			proxy:     common.BigToAddress(big.NewInt(int64(i))),
			timestamp: big.NewInt(int64(i)),
		}
	}

	return games
}

func translateGames(games []FaultDisputeGame) []struct {
	proxy     common.Address
	timestamp *big.Int
} {
	translated := make([]struct {
		proxy     common.Address
		timestamp *big.Int
	}, len(games))

	for i, game := range games {
		translated[i] = translateFaultDisputeGame(game)
	}

	return translated
}

func translateFaultDisputeGame(game FaultDisputeGame) struct {
	proxy     common.Address
	timestamp *big.Int
} {
	return struct {
		proxy     common.Address
		timestamp *big.Int
	}{
		proxy:     game.Proxy,
		timestamp: game.Timestamp,
	}
}

func generateMockGameErrors(count uint64, injectErrors bool) []bool {
	errors := make([]bool, count)

	if injectErrors {
		for i := uint64(0); i < count; i++ {
			errors[i] = true
		}
	}

	return errors
}

type mockMinimalDisputeGameFactoryCaller struct {
	gameCountErr bool
	indexErrors  []bool
	gameCount    uint64
	games        []struct {
		proxy     common.Address
		timestamp *big.Int
	}
}

func newMockMinimalDisputeGameFactoryCaller(count uint64, gameCountErr bool, indexErrors bool) *mockMinimalDisputeGameFactoryCaller {
	return &mockMinimalDisputeGameFactoryCaller{
		indexErrors:  generateMockGameErrors(count, indexErrors),
		gameCountErr: gameCountErr,
		gameCount:    count,
		games:        generateMockGames(count),
	}
}

func (m *mockMinimalDisputeGameFactoryCaller) GameCount(opts *bind.CallOpts) (*big.Int, error) {
	if m.gameCountErr {
		return nil, gameCountErr
	}

	return big.NewInt(int64(m.gameCount)), nil
}

func (m *mockMinimalDisputeGameFactoryCaller) GameAtIndex(opts *bind.CallOpts, _index *big.Int) (struct {
	Proxy     common.Address
	Timestamp *big.Int
}, error) {
	index := _index.Uint64()
	if m.indexErrors[index] {
		return struct {
			Proxy     common.Address
			Timestamp *big.Int
		}{}, gameIndexErr
	}

	return struct {
		Proxy     common.Address
		Timestamp *big.Int
	}{
		Proxy:     m.games[index].proxy,
		Timestamp: m.games[index].timestamp,
	}, nil
}
