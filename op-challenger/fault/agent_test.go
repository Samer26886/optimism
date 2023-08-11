package fault

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-challenger/fault/types"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
)

var (
	statusFetcherError = errors.New("status fetcher error")
)

// TestAgent_ShouldResolve tests the [Agent] resolution logic.
func TestAgent_ShouldResolve(t *testing.T) {
	log := testlog.Logger(t, log.LvlCrit)

	t.Run("AgreeWithProposedOutput", func(t *testing.T) {
		agreeWithProposedOutput := true

		agent := NewAgent(nil, 0, nil, newMockStatusFetcher(types.GameStatusDefenderWon, false), nil, nil, agreeWithProposedOutput, log)
		require.False(t, agent.ShouldResolve(context.Background()))

		agent = NewAgent(nil, 0, nil, newMockStatusFetcher(types.GameStatusChallengerWon, false), nil, nil, agreeWithProposedOutput, log)
		require.True(t, agent.ShouldResolve(context.Background()))
	})

	t.Run("DisagreeWithProposedOutput", func(t *testing.T) {
		agreeWithProposedOutput := false

		agent := NewAgent(nil, 0, nil, newMockStatusFetcher(types.GameStatusDefenderWon, false), nil, nil, agreeWithProposedOutput, log)
		require.True(t, agent.ShouldResolve(context.Background()))

		agent = NewAgent(nil, 0, nil, newMockStatusFetcher(types.GameStatusChallengerWon, false), nil, nil, agreeWithProposedOutput, log)
		require.False(t, agent.ShouldResolve(context.Background()))
	})

	t.Run("StatusFetchReturnsFalse", func(t *testing.T) {
		agreeWithProposedOutput := true
		agent := NewAgent(nil, 0, nil, newMockStatusFetcher(types.GameStatusChallengerWon, true), nil, nil, agreeWithProposedOutput, log)
		require.False(t, agent.ShouldResolve(context.Background()))
	})
}

type mockStatusFetcher struct {
	shouldError bool
	status      types.GameStatus
}

func newMockStatusFetcher(status types.GameStatus, shouldError bool) *mockStatusFetcher {
	return &mockStatusFetcher{
		shouldError: shouldError,
		status:      status,
	}
}

func (m *mockStatusFetcher) GetGameStatus(context.Context) (types.GameStatus, error) {
	if m.shouldError {
		return 0, statusFetcherError
	}
	return m.status, nil
}

// type stubResponder struct {}
// func (s *stubResponder) CanResolve(ctx context.Context) bool { return true }
// func (s *stubResponder) Resolve(ctx context.Context) error { panic("not implemented") }
// func (s *stubResponder) Respond(ctx context.Context, response types.Claim) error { panic("not implemented") }
// func (s *stubResponder) Step(ctx context.Context, stepData types.StepCallData) error { panic("not implemented") }
//
