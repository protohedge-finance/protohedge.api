package http_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	http_controllers "github.com/protohedge/protohedge.api/internal/adapters/http_server/controllers"
	"github.com/protohedge/protohedge.api/internal/adapters/repositories"
	cfg "github.com/protohedge/protohedge.api/internal/config"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
	"github.com/rs/cors"
)

var port = ":8080"

func CreateServer(config *cfg.Config) {
	fmt.Println("rpc url is: ")
	fmt.Println(config.RpcUrl)
	ethClient, err := ethclient.Dial(config.RpcUrl)

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisConnectionString,
	})

	if err != nil {
		panic(err)
	}

	vaultRepository := repositories.NewVaultRepository(ethClient, redisClient)
	vaultRetriever := use_cases.NewVaultRetriever(vaultRepository)
	pnlRetriever := use_cases.NewPnlRetriever(vaultRepository)
	vaultController := http_controllers.NewVaultController(vaultRetriever, pnlRetriever)
	statusController := http_controllers.NewStatusController()

	router := mux.NewRouter()

	router.HandleFunc("/status", statusController.GetStatus)
	router.HandleFunc("/vault/{address}", vaultController.GetVault)
	router.HandleFunc("/vault/{address}/historicPnl", vaultController.GetHistoricVaultPnl)

	log.Printf("Listening on port %s\n", port)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}
