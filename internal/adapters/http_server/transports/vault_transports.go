package http_transports

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	http_server_mappings "github.com/protohedge/protohedge.api/internal/adapters/http_server/mappings"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	vault_endpoints "github.com/protohedge/protohedge.api/internal/core/endpoints/vault"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
	"go.uber.org/zap"
)

func NewVaultHTTPHandler(
	logger *zap.Logger,
	vaultRetriever use_cases.VaultRetriever,
	pnlRetriever use_cases.PnlRetriever,
	rebalanceInfoRetriever use_cases.RebalanceInfoRetriever,
	rebalanceHistoryRetriever use_cases.RebalanceHistoryRetriever) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	m := mux.NewRouter()

	getVaultEndpoint := vault_endpoints.CreateGetVaultEndpoint(vaultRetriever)
	getHistoricVaultPnlEndpoint := vault_endpoints.CreateGetHistoricVaultPnl(pnlRetriever)
	getRebalanceInfoEndpoint := vault_endpoints.CreateGetRebalanceInfo(rebalanceInfoRetriever)
	getRebalanceHistoryEndpoint := vault_endpoints.CreateGetRebalanceHistory(rebalanceHistoryRetriever)

	m.Handle("/vault/{address}", httptransport.NewServer(getVaultEndpoint, decodeGetVaultRequest, encodeMappedResponse(mapGetVault), options...))
	m.Handle("/vault/{address}/historicPnl", httptransport.NewServer(getHistoricVaultPnlEndpoint, decodeGetHistoricPnlRequest, encodeMappedResponse(mapGetHistoricPnl), options...))
	m.Handle("/vault/{address}/rebalanceHistory", httptransport.NewServer(getRebalanceHistoryEndpoint, decodeGetRebalanceHistoryRequest, encodeMappedResponse(mapGetRebalanceHistory), options...))
	m.Handle("/vault/{address}/rebalanceInfo", httptransport.NewServer(getRebalanceInfoEndpoint, decodeGetRebalanceInfoRequest, encodeMappedResponse(mapGetRebalanceInfo), options...))

	return m
}

func decodeGetVaultRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, errors.New("Address not found")
	}

	return vault_endpoints.GetVaultRequest{Address: address}, nil
}

func decodeGetHistoricPnlRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, errors.New("Address not found")
	}

	return vault_endpoints.GetHistoricVaultPnlRequest{Address: address}, nil
}

func decodeGetRebalanceHistoryRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, errors.New("Address not found")
	}

	return vault_endpoints.GetRebalanceHistoryRequest{Address: address}, nil
}

func decodeGetRebalanceInfoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, errors.New("Address not found")
	}

	return vault_endpoints.GetRebalanceInfoRequest{Address: address}, nil
}

type EncodeResponse = func(context.Context, http.ResponseWriter, interface{}) error

func mapGetVault(response interface{}) interface{} {
	return http_server_mappings.ToVaultResponseDto(response.(vault_endpoints.GetVaultResponse).Vault)
}

func mapGetHistoricPnl(response interface{}) interface{} {
	return http_server_mappings.ToHistoricPnlDto(response.(vault_endpoints.GetHistoricVaultPnlResponse).HistoricVaultPnl)
}

func mapGetRebalanceHistory(response interface{}) interface{} {
	return http_server_mappings.ToRebalanceHistoryDto(response.(vault_endpoints.GetRebalanceHistoryResponse).RebalanceHistory)
}

func mapGetRebalanceInfo(response interface{}) interface{} {
	return http_server_mappings.ToRebalanceInfoDto(response.(domain.RebalanceInfo))
}

func encodeMappedResponse(mapResponse func(interface{}) interface{}) EncodeResponse {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if e, ok := response.(errorer); ok && e.error() != nil {
			errorEncoder(ctx, e.error(), w)
			return nil
		}

		mappedResponse := mapResponse(response)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(mappedResponse)
	}
}

type errorer interface {
	error() error
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
