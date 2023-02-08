package http_server

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v9"
	"github.com/protohedge/protohedge.api/internal/adapters"
	http_transports "github.com/protohedge/protohedge.api/internal/adapters/http_server/transports"
	"github.com/protohedge/protohedge.api/internal/adapters/repositories"
	cfg "github.com/protohedge/protohedge.api/internal/config"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
	"github.com/rs/cors"
)

var port = ":8080"

func CreateServer(config *cfg.Config) {
	logger := adapters.NewLogger(config)

	ethClient, err := ethclient.Dial(config.RpcUrl)

	redisOptions := &redis.Options{
		Addr:     config.Redis.Host,
		Username: config.Redis.Username,
		Password: config.Redis.Password,
	}

	if config.Redis.UseSsl {
		redisOptions.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	redisClient := redis.NewClient(redisOptions)

	if err != nil {
		panic(err)
	}

	awsConfig := adapters.NewClient(config)

	vaultRepository := repositories.NewVaultRepository(ethClient, redisClient, config, awsConfig)
	vaultRetriever := use_cases.NewVaultRetriever(vaultRepository)
	pnlRetriever := use_cases.NewPnlRetriever(vaultRepository)
	rebalanceHistoryRetriever := use_cases.NewRebalanceHistoryRetriever(vaultRepository)
	rebalanceInfoRetriever := use_cases.NewRebalanceInfoRetriever(vaultRepository)

	vaultHandler := http_transports.NewVaultHTTPHandler(logger, vaultRetriever, pnlRetriever, rebalanceInfoRetriever, rebalanceHistoryRetriever)
	healthHandler := http_transports.NewHealthHTTPHandler()

	mux := http.NewServeMux()
	mux.Handle("/vault/", vaultHandler)
	mux.Handle("/", healthHandler)

	log.Printf("Listening on port %s\n", port)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(port, handler))
}
