package pnl

import (
	"context"
	"errors"
	"math/big"

	"github.com/gmx-delta-neutral/gmx-neutral.api/internal/client"
	"github.com/gmx-delta-neutral/gmx-neutral.api/internal/models"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

type Service interface {
	GetPnl() (*models.Pnl, error)
}

type pnlService struct {
	queryClient *client.QueryClient
}

func NewPnlService(queryClient *client.QueryClient) Service {
	return &pnlService{
		queryClient: queryClient,
	}
}

func (s *pnlService) GetPnl() (*models.Pnl, error) {
	tokenPositionsRequest := api.GetTokenPositionsRequest{}
	response, err := s.queryClient.Position.GetTokenPositions(context.Background(), &tokenPositionsRequest)
	if err != nil {
		return nil, err
	}

	fullPnl := new(big.Int)
	fullPnlPercentage := new(big.Float)

	for _, tokenPosition := range response.TokenPositions {
		pnl := new(big.Int).SetBytes(tokenPosition.Pnl)
		fullPnl = new(big.Int).Add(fullPnl, pnl)
		pnlPercentage, ok := new(big.Float).SetString(tokenPosition.PnlPercentage)

		if !ok {
			return nil, errors.New("Could not parse PNL percentage")
		}

		fullPnlPercentage = new(big.Float).Add(fullPnlPercentage, pnlPercentage)
	}

	return &models.Pnl{
		Amount:     fullPnl,
		Percentage: fullPnlPercentage,
	}, nil
}
