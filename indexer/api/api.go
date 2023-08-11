package api

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
)

type Extensions struct {
	OptimismBridgeAddress string `json:"OptimismBridgeAddress"`
}

type TokenInfo struct {
	// TODO lazily typing ints go through them all with fine tooth comb once api is up
	ChainId    int        `json:"chainId"`
	Address    string     `json:"address"`
	Name       string     `json:"name"`
	Symbol     string     `json:"symbol"`
	Decimals   int        `json:"decimals"`
	LogoURI    string     `json:"logoURI"`
	Extensions Extensions `json:"extensions"`
}

// lazily typing numbers fixme
type Transaction struct {
	Timestamp       uint64 `json:"timestamp"`
	BlockNumber     int64  `json:"number"` // viem treats me as a BigInt maybe consider just stringing me too. Not 100% necessary though
	BlockHash       string `json:"hash"`
	TransactionHash string `json:"transactionHash"`
	// TODO maybe include me
	// ParentBlockHash   string `json:"parentHash"`
}

type DepositItem struct {
	Guid string `json:"guid"`
	From string `json:"from"`
	To   string `json:"to"`
	// TODO could consider OriginTx to be more generic to handling L2 to L2 deposits
	// this seems more clear today though
	Tx      Transaction `json:"Block"`
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
func newDepositResponse(deposits []*database.L1BridgeDepositWithTransactionHashes) DepositResponse {
	var items []DepositItem
	for _, deposit := range deposits {
		item := DepositItem{
			Guid: deposit.L1BridgeDeposit.TransactionSourceHash.String(),
			Tx: Transaction{
				BlockNumber:     420420,  // TODO
				BlockHash:       "0x420", // TODO
				TransactionHash: "0x420", // TODO
				Timestamp:       deposit.L1BridgeDeposit.Tx.Timestamp,
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
		Cursor:      "42042042-4204-4204-4204-420420420420", // TODO
		HasNextPage: false,                                  // TODO
		Items:       items,
	}
}

func (a *Api) L1DepositsHandler(w http.ResponseWriter, r *http.Request) {
	bv := a.BridgeTransfersView
	address := common.HexToAddress(chi.URLParam(r, "address"))

	deposits, err := bv.L1BridgeDepositsByAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := newDepositResponse(deposits)

	jsonResponse(w, response, http.StatusOK)
}

type Proof struct {
	TransactionHash string `json:"transactionHash"`
	BlockTimestamp  uint64 `json:"blockTimestamp"`
	BlockNumber     int    `json:"blockNumber"`
}

type Claim struct {
	TransactionHash string `json:"transactionHash"`
	BlockTimestamp  uint64 `json:"blockTimestamp"`
	BlockNumber     int    `json:"blockNumber"`
}

type WithdrawalItem struct {
	Guid            string    `json:"guid"`
	BlockTimestamp  uint64    `json:"blockTimestamp"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	TransactionHash string    `json:"transactionHash"`
	Amount          string    `json:"amount"`
	BlockNumber     int       `json:"blockNumber"`
	Proof           Proof     `json:"proof"`
	Claim           Claim     `json:"claim"`
	WithdrawalState string    `json:"withdrawalState"`
	L1Token         TokenInfo `json:"l1Token"`
	L2Token         TokenInfo `json:"l2Token"`
}

type WithdrawalResponse struct {
	Cursor      string           `json:"cursor"`
	HasNextPage bool             `json:"hasNextPage"`
	Items       []WithdrawalItem `json:"items"`
}

func newWithdrawalResponse(withdrawals []*database.L2BridgeWithdrawalWithTransactionHashes) WithdrawalResponse {
	var items []WithdrawalItem
	for _, withdrawal := range withdrawals {
		item := WithdrawalItem{
			Guid:            withdrawal.L2BridgeWithdrawal.TransactionWithdrawalHash.String(),
			BlockTimestamp:  withdrawal.L2BridgeWithdrawal.Tx.Timestamp,
			From:            withdrawal.L2BridgeWithdrawal.Tx.FromAddress.String(),
			To:              withdrawal.L2BridgeWithdrawal.Tx.ToAddress.String(),
			TransactionHash: withdrawal.L2TransactionHash.String(),
			Amount:          withdrawal.L2BridgeWithdrawal.Tx.Amount.Int.String(),
			BlockNumber:     420, // TODO
			Proof: Proof{
				TransactionHash: withdrawal.ProvenL1TransactionHash.String(),
				BlockTimestamp:  withdrawal.L2BridgeWithdrawal.Tx.Timestamp,
				BlockNumber:     420, // TODO Block struct instead
			},
			Claim: Claim{
				TransactionHash: withdrawal.FinalizedL1TransactionHash.String(),
				BlockTimestamp:  withdrawal.L2BridgeWithdrawal.Tx.Timestamp, // Using L2 timestamp for now, might need adjustment
				BlockNumber:     420,                                        // TODO block struct
			},
			WithdrawalState: "COMPLETE", // TODO
			L1Token: TokenInfo{
				ChainId:  1,
				Address:  withdrawal.L2BridgeWithdrawal.TokenPair.L1TokenAddress.String(),
				Name:     "Example",                                              // TODO
				Symbol:   "EXAMPLE",                                              // TODO
				Decimals: 18,                                                     // TODO
				LogoURI:  "https://ethereum-optimism.github.io/data/OP/logo.svg", // TODO
				Extensions: Extensions{
					OptimismBridgeAddress: "0x636Af16bf2f682dD3109e60102b8E1A089FedAa8",
				},
			},
			L2Token: TokenInfo{
				ChainId:  10,
				Address:  withdrawal.L2BridgeWithdrawal.TokenPair.L2TokenAddress.String(),
				Name:     "Example",                                              // TODO
				Symbol:   "EXAMPLE",                                              // TODO
				Decimals: 18,                                                     // TODO
				LogoURI:  "https://ethereum-optimism.github.io/data/OP/logo.svg", // TODO
				Extensions: Extensions{
					OptimismBridgeAddress: "0x36Af16bf2f682dD3109e60102b8E1A089FedAa86",
				},
			},
		}
		items = append(items, item)
	}

	return WithdrawalResponse{
		Cursor:      "42042042-0420-4204-2042-420420420420", // TODO
		HasNextPage: true,                                   // TODO
		Items:       items,
	}
}

func (a *Api) L2WithdrawalsHandler(w http.ResponseWriter, r *http.Request) {
	bv := a.BridgeTransfersView
	address := common.HexToAddress(chi.URLParam(r, "address"))

	withdrawals, err := bv.L2BridgeWithdrawalsByAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := newWithdrawalResponse(withdrawals)

	jsonResponse(w, response, http.StatusOK)
}

func (a *Api) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, "ok", http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Api struct {
	Router              *chi.Mux
	BridgeTransfersView database.BridgeTransfersView
}

func NewApi(bv database.BridgeTransfersView) *Api {
	r := chi.NewRouter()

	api := &Api{Router: r, BridgeTransfersView: bv}

	// these regex are .+ because I wasn't sure what they should be
	// don't want a regex for addresses because would prefer to validate the address
	// with go-ethereum and throw a friendly error message
	r.Get("/api/v0/deposits/{address:.+}", api.L1DepositsHandler)
	r.Get("/api/v0/withdrawals/{address:.+}", api.L2WithdrawalsHandler)
	r.Get("/healthz", api.HealthzHandler)

	return api

}

func (a *Api) Listen(port string) error {
	return http.ListenAndServe(port, a.Router)
}
