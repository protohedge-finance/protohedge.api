package http_transports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	http_server_mappings "github.com/protohedge/protohedge.api/internal/adapters/http_server/mappings"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	vault_endpoints "github.com/protohedge/protohedge.api/internal/core/endpoints/vault"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
	"go.uber.org/zap"
)

var addressNotFound = fmt.Errorf("%w: Empty address", domain.ErrBadRequest)

func NewVaultHTTPHandler(
	logger *zap.Logger,
	vaultRetriever use_cases.VaultRetriever,
	pnlRetriever use_cases.PnlRetriever,
	rebalanceInfoRetriever use_cases.RebalanceInfoRetriever,
	rebalanceNotesRetriever use_cases.RebalanceNotesRetriever) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	m := mux.NewRouter()

	getVaultEndpoint := vault_endpoints.CreateGetVaultEndpoint(vaultRetriever)
	getHistoricVaultPnlEndpoint := vault_endpoints.CreateGetHistoricVaultPnl(pnlRetriever)
	getRebalanceInfoEndpoint := vault_endpoints.CreateGetRebalanceInfo(rebalanceInfoRetriever)
	getRebalanceNotesEndpoint := vault_endpoints.CreateGetRebalanceNotes(rebalanceNotesRetriever)

	m.Handle("/vault/{address}", httptransport.NewServer(getVaultEndpoint, decodeGetVaultRequest, encodeMappedResponse(mapGetVault), options...))
	m.Handle("/vault/{address}/historicPnl", httptransport.NewServer(getHistoricVaultPnlEndpoint, decodeGetHistoricPnlRequest, encodeMappedResponse(mapGetHistoricPnl), options...))
	m.Handle("/vault/{address}/rebalanceNotes", httptransport.NewServer(getRebalanceNotesEndpoint, decodeGetRebalanceNotesRequest, encodeMappedResponse(mapGetRebalanceNotes), options...))
	m.Handle("/vault/{address}/rebalanceInfo", httptransport.NewServer(getRebalanceInfoEndpoint, decodeGetRebalanceInfoRequest, encodeMappedResponse(mapGetRebalanceInfo), options...))

	return m
}

func decodeGetVaultRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, addressNotFound
	}

	return vault_endpoints.GetVaultRequest{Address: address}, nil
}

func decodeGetHistoricPnlRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, addressNotFound
	}

	return vault_endpoints.GetHistoricVaultPnlRequest{Address: address}, nil
}

func decodeGetRebalanceNotesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, addressNotFound
	}

	return vault_endpoints.GetRebalanceNotesRequest{Address: address}, nil
}

func decodeGetRebalanceInfoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	address, ok := vars["address"]
	if !ok {
		return nil, addressNotFound
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

func mapGetRebalanceNotes(response interface{}) interface{} {
	return http_server_mappings.ToRebalanceNotesDto(response.(vault_endpoints.GetRebalanceNotesResponse).RebalanceNotes)
}

func mapGetRebalanceInfo(response interface{}) interface{} {
	return http_server_mappings.ToRebalanceInfoDto(response.(vault_endpoints.GetRebalanceInfoResponse).RebalanceInfo)
}

func encodeMappedResponse(mapResponse func(interface{}) interface{}) EncodeResponse {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		mappedResponse := mapResponse(response)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(mappedResponse)
	}
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var message string
	var errorCode string
	if errors.Is(err, domain.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		message = err.Error()
		errorCode = "notFound"
	} else if errors.Is(err, domain.ErrBadRequest) {
		w.WriteHeader(http.StatusBadRequest)
		message = err.Error()
		errorCode = "badRequest"
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		errorCode = "unknownError"
		message = domain.ErrUnknownError.Error()
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":     message,
		"errorCode": errorCode,
	})
}
