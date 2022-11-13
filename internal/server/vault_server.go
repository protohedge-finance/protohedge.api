package server

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/protohedge/protohedge.api/internal/mapping"
	"github.com/protohedge/protohedge.api/internal/vault"
)

type vaultServer struct {
	vaultService vault.Service
}

func NewVaultServer(vaultService vault.Service) vaultServer {
	return vaultServer{
		vaultService: vaultService,
	}
}

func (s *vaultServer) GetVault(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := common.HexToAddress(vars["address"])
	vault, err := s.vaultService.GetVault(address)

	if err != nil {
		fmt.Println(err)
		ErrorRes(w)
		return
	}

	positionManagerDto := mapping.ToVaultResponseDto(vault)
	SuccessResWithData(w, positionManagerDto)
}
