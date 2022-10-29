package position_manager

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/protohedge/protohedge.api/internal/contracts/position_manager_contract"
	"github.com/protohedge/protohedge.api/internal/mapping"
	"github.com/protohedge/protohedge.api/internal/models"
)

type Service interface {
	GetPositionManager(address common.Address) (*models.PositionManager, error)
}

type positionManagerService struct {
	ethClient *ethclient.Client
}

func NewPositionManagerService(ethClient *ethclient.Client) Service {
	return &positionManagerService{
		ethClient: ethClient,
	}
}

func (s *positionManagerService) GetPositionManager(address common.Address) (*models.PositionManager, error) {
	ctr, err := position_manager_contract.NewPositionManagerContract(address, s.ethClient)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	result, err := ctr.Stats(callOpts)

	if err != nil {
		return nil, err
	}

	positionManager := mapping.ToPositionManagerModel(result, address)
	return &positionManager, nil
}
