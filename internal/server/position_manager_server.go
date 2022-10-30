package server

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/protohedge/protohedge.api/internal/mapping"
	"github.com/protohedge/protohedge.api/internal/position_manager"
)

type positionManagerServer struct {
	positionManagerService position_manager.Service
}

func NewPositionManagerServer(positionManagerService position_manager.Service) positionManagerServer {
	return positionManagerServer{
		positionManagerService: positionManagerService,
	}
}

func (s *positionManagerServer) GetPositionManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := common.HexToAddress(vars["address"])
	positionManager, err := s.positionManagerService.GetPositionManager(address)

	if err != nil {
		fmt.Println(err)
		ErrorRes(w)
		return
	}

	positionManagerDto := mapping.ToPositionManagerResponseDto(positionManager)
	SuccessResWithData(w, positionManagerDto)
}
