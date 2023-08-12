package handlers

import (
	"net/http"
	"strconv"

	"github.com/ethereum-optimism/optimism/indexer/api/middleware"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
)

type DepositItem struct {
	Guid string `json:"guid"`
	From string `json:"from"`
	To   string `json:"to"`
	// TODO could consider OriginTx to be more generic to handling L2 to L2 deposits
	// this seems more clear today though
	Tx      Transaction `json:"Tx"`
	RelayTx Transaction `json:"RelayTx"`
	Amount  string      `json:"amount"`
	L1Token TokenInfo   `json:"l1Token"`
	L2Token TokenInfo   `json:"l2Token"`
}

type DepositResponse struct {
	Cursor      string        `json:"cursor"`
	HasNextPage bool          `json:"hasNextPage"`
	Items       []DepositItem `json:"items"`
}

// TODO this is original spec but maybe include the l2 block info too for the relayed tx
func newDepositResponse(deposits *database.L1BridgeDepositsResponse) DepositResponse {
	items := make([]DepositItem, 0, len(deposits.Deposits))
	for _, deposit := range deposits.Deposits {
		item := DepositItem{
			Guid: deposit.L1BridgeDeposit.TransactionSourceHash.String(),
			Tx: Transaction{
				// BlockNumber:     420420,  // TODO
				// BlockHash:       "0x420", // TODO
				TransactionHash: deposit.L1TransactionHash.String(), // TODO
				Timestamp:       deposit.L1BridgeDeposit.Tx.Timestamp,
			},
			RelayTx: Transaction{
				// BlockNumber:     420420,  // TODO
				// BlockHash:       "0x420", // TODO
				TransactionHash: deposit.L2TransactionHash.String(), // TODO
				// Timestamp:       deposit.L1BridgeDeposit.Tx.Timestamp, // TODO
			},
			From:   deposit.L1BridgeDeposit.Tx.FromAddress.String(),
			To:     deposit.L1BridgeDeposit.Tx.ToAddress.String(),
			Amount: deposit.L1BridgeDeposit.Tx.Amount.Int.String(),
			L1Token: TokenInfo{
				ChainId:  1,
				Address:  deposit.L1BridgeDeposit.TokenPair.L1TokenAddress.String(),
				Name:     "TODO",
				Symbol:   "TODO",
				Decimals: 420,
				LogoURI:  "TODO",
				Extensions: Extensions{
					OptimismBridgeAddress: "0x420", // TODO
				},
			},
			L2Token: TokenInfo{
				ChainId:  10,
				Address:  deposit.L1BridgeDeposit.TokenPair.L2TokenAddress.String(),
				Name:     "TODO",
				Symbol:   "TODO",
				Decimals: 420,
				LogoURI:  "TODO",
				Extensions: Extensions{
					OptimismBridgeAddress: "0x420", // TODO
				},
			},
		}
		items = append(items, item)
	}

	return DepositResponse{
		Cursor:      deposits.Cursor,
		HasNextPage: deposits.HasNextPage,
		Items:       items,
	}
}

func L1DepositsHandler(w http.ResponseWriter, r *http.Request) {
	btv := middleware.GetBridgeTransfersView(r.Context())
	logger := middleware.GetLogger(r.Context())

	address := common.HexToAddress(chi.URLParam(r, "address"))
	cursor := r.URL.Query().Get("cursor")
	limitQuery := r.URL.Query().Get("limit")

	defaultLimit := 100
	limit := defaultLimit
	if limitQuery != "" {
		parsedLimit, err := strconv.Atoi(limitQuery)
		if err != nil {
			http.Error(w, "Limit could not be parsed into a number", http.StatusBadRequest)
			logger.Error("Invalid limit")
			logger.Error(err.Error())
		}
		limit = parsedLimit
	}

	deposits, err := btv.L1BridgeDepositsByAddress(address, cursor, limit)
	if err != nil {
		http.Error(w, "Internal server error reading deposits", http.StatusInternalServerError)
		logger.Error("Unable to read deposits from DB")
		logger.Error(err.Error())
	}

	response := newDepositResponse(deposits)

	jsonResponse(w, logger, response, http.StatusOK)
}
