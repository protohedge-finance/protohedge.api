package repositories

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v9"
	"github.com/protohedge/protohedge.api/internal/adapters"
	redis_mappings "github.com/protohedge/protohedge.api/internal/adapters/redis/mappings"
	"github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/abi/vault_contract"
	contract_mappings "github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/mappings"
	"github.com/protohedge/protohedge.api/internal/config"
	"github.com/protohedge/protohedge.api/internal/core/domain"
	"github.com/protohedge/protohedge.api/internal/core/ports"
)

type vaultRepository struct {
	ethClient   *ethclient.Client
	redisClient *redis.Client
	config      *config.Config
	awsClient   *adapters.AwsClient
}

func NewVaultRepository(ethClient *ethclient.Client, redisClient *redis.Client, config *config.Config, awsClient *adapters.AwsClient) ports.VaultRepository {
	return &vaultRepository{
		ethClient:   ethClient,
		redisClient: redisClient,
		config:      config,
		awsClient:   awsClient,
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
		fmt.Println("Error is: ")
		fmt.Println(err)
		return nil, err
	}

	vault := contract_mappings.ToVaultModel(result)
	return &vault, nil
}

func (v *vaultRepository) GetRebalanceHistory(ctx context.Context, address common.Address) ([]domain.RebalanceNote, error) {
	result, err := v.redisClient.ZRangeByLex(ctx, "rebalance_history", &redis.ZRangeBy{
		Min: fmt.Sprintf("[%s:", strings.ToLower(address.String())),
		Max: fmt.Sprintf("(%s;", strings.ToLower(address.String())),
	}).Result()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return redis_mappings.ToRebalanceHistoryModel(result)
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

func (v *vaultRepository) GetRebalanceInfo(ctx context.Context, address common.Address) (*domain.RebalanceInfo, error) {
	durationRemaining, rebalanceInterval, err := v.getRebalanceDurations(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &domain.RebalanceInfo{
		RebalanceInterval: rebalanceInterval,
		DurationRemaining: durationRemaining,
	}, nil
}

func (v *vaultRepository) getRebalanceDurations(ctx context.Context) (time.Duration, time.Duration, error) {
	output, err := v.awsClient.SchedulerClient.GetSchedule(ctx, &scheduler.GetScheduleInput{
		Name: &v.config.EventbridgeRuleName,
	})

	if err != nil {
		fmt.Println(err)
		return time.Duration(0), time.Duration(0), err
	}

	pattern := regexp.MustCompile(`\d+`)
	match := pattern.FindString(*output.ScheduleExpression)

	if match == "" {
		return time.Duration(0), time.Duration(0), errors.New("Could not extract rebalance rate")
	}

	unformattedDurationInHours, err := strconv.Atoi(match)

	if err != nil {
		fmt.Println(err)
		return time.Duration(0), time.Duration(0), err
	}

	rebalanceDuration := time.Duration(unformattedDurationInHours) * time.Hour

	timeDifference := time.Now().Sub(*output.StartDate)

	if timeDifference.Milliseconds() < 0 {
		return rebalanceDuration, output.StartDate.Sub(time.Now()), nil
	}

	timeRemainderRemaining := time.Duration(timeDifference.Milliseconds()%rebalanceDuration.Milliseconds()) * time.Millisecond
	timeRemaining := (time.Duration(1) * time.Hour) - timeRemainderRemaining

	return timeRemaining, rebalanceDuration, nil
}
