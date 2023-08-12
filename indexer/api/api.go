package api

import (
	"fmt"
	"net/http"

	"github.com/ethereum-optimism/optimism/indexer/api/handlers"
	"github.com/ethereum-optimism/optimism/indexer/api/middleware"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
)

const ethereumAddressRegex = `^0x[a-fA-F0-9]{40}$`

type Api struct {
	Router *chi.Mux
}

func NewApi(bv database.BridgeTransfersView, logger log.Logger) *Api {
	logger.Info("Initializing API...")

	r := chi.NewRouter()

	r.Use(middleware.ContextMiddleware(logger, bv))

	api := &Api{Router: r}

	r.Get("/healthz", handlers.HealthzHandler)
	r.Get(fmt.Sprintf("/api/v0/deposits/{address:%s}", ethereumAddressRegex), handlers.L1DepositsHandler)
	r.Get(fmt.Sprintf("/api/v0/withdrawals/{address:%s}", ethereumAddressRegex), handlers.L2WithdrawalsHandler)

	return api
}

func (a *Api) Listen(port string) error {
	return http.ListenAndServe(port, a.Router)
}
