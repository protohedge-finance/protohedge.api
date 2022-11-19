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
	vaultRetriever use_cases.VaultRetriever
}

func NewVaultController(vaultRetriever use_cases.VaultRetriever) vaultController {
	return vaultController{
		vaultRetriever: vaultRetriever,
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
