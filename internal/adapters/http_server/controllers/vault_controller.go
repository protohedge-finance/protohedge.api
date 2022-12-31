package http_controllers

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	http_server_mappings "github.com/protohedge/protohedge.api/internal/adapters/http_server/mappings"
	http_server_utils "github.com/protohedge/protohedge.api/internal/adapters/http_server/utils"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type vaultController struct {
	vaultRetriever            use_cases.VaultRetriever
	pnlRetriever              use_cases.PnlRetriever
	rebalanceHistoryRetriever use_cases.RebalanceHistoryRetriever
}

func NewVaultController(vaultRetriever use_cases.VaultRetriever, pnlRetriever use_cases.PnlRetriever, rebalanceHistoryRetriever use_cases.RebalanceHistoryRetriever) vaultController {
	return vaultController{
		vaultRetriever:            vaultRetriever,
		pnlRetriever:              pnlRetriever,
		rebalanceHistoryRetriever: rebalanceHistoryRetriever,
	}
}

func (s *vaultController) GetVault(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := common.HexToAddress(vars["address"])
	vault, err := s.vaultRetriever.RetrieveVault(r.Context(), address)

	if err != nil {
		fmt.Println(err)
		http_server_utils.ErrorRes(w)
		return
	}

	positionManagerDto := http_server_mappings.ToVaultResponseDto(vault)
	http_server_utils.SuccessResWithData(w, positionManagerDto)
}

func (s *vaultController) GetHistoricVaultPnl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := common.HexToAddress(vars["address"])
	historicPnl, err := s.pnlRetriever.RetrieveHistoricPnl(r.Context(), address)

	if err != nil {
		http_server_utils.ErrorRes(w)
		return
	}

	historicPnlDto := http_server_mappings.ToHistoricPnlDto(historicPnl)
	http_server_utils.SuccessResWithData(w, historicPnlDto)
}

func (s *vaultController) GetRebalanceHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := common.HexToAddress(vars["address"])
	rebalanceHistory, err := s.rebalanceHistoryRetriever.RetrieveRebalanceHistory(r.Context(), address)

	if err != nil {
		http_server_utils.ErrorRes(w)
		return
	}

	rebalanceHistoryDto := http_server_mappings.ToRebalanceHistoryDto(rebalanceHistory)
	http_server_utils.SuccessResWithData(w, rebalanceHistoryDto)
}
