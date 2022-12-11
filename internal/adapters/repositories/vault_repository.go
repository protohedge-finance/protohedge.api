package repositories

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v9"
	"github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/abi/vault_contract"
	contract_mappings "github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/mappings"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type vaultRepository struct {
	ethClient   *ethclient.Client
	redisClient *redis.Client
}

func NewVaultRepository(ethClient *ethclient.Client, redisClient *redis.Client) ports.VaultRepository {
	return &vaultRepository{
		ethClient:   ethClient,
		redisClient: redisClient,
	}
}

func (v *vaultRepository) GetVault(ctx context.Context, address common.Address) (*domain.Vault, error) {
	ctr, err := vault_contract.NewVaultContract(address, v.ethClient)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	result, err := ctr.Stats(callOpts)

	if err != nil {
		return nil, err
	}

	vault := contract_mappings.ToVaultModel(result)
	return &vault, nil
}

func (v *vaultRepository) GetHistoricVaultPnl(ctx context.Context, address common.Address) ([]domain.TimePoint, error) {
	historicPnl := []domain.TimePoint{}

	result, err := v.redisClient.ZRangeByLex(ctx, "pnl", &redis.ZRangeBy{
		Min: "-",
		Max: "+",
	}).Result()

	if err != nil {
		return nil, err
	}

	for _, pnlTimestamp := range result {
		splitPnlTimestamp := strings.Split(pnlTimestamp, ":")
		fmt.Println(len(splitPnlTimestamp))

		if len(splitPnlTimestamp) != 2 {
			continue
		}

		timestamp, err := strconv.ParseInt(splitPnlTimestamp[0], 10, 64)

		if err != nil {
			return nil, err
		}

		pnl, success := new(big.Int).SetString(splitPnlTimestamp[1], 10)
		if !success {
			return nil, fmt.Errorf("unable to convert pnl %s to big int", splitPnlTimestamp[1])
		}

		historicPnl = append(historicPnl, domain.TimePoint{
			Timestamp: time.UnixMilli(timestamp).UTC(),
			Point:     pnl,
		})

	}

	return historicPnl, nil
}
