package http_server

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	http_controllers "github.com/protohedge/protohedge.api/internal/adapters/http_server/controllers"
	contract_repositories "github.com/protohedge/protohedge.api/internal/adapters/smart_contracts/repositories"
	cfg "github.com/protohedge/protohedge.api/internal/config"
	"github.com/protohedge/protohedge.api/internal/core/use_cases"
	"github.com/rs/cors"
)

var port = ":8080"

func CreateServer(config *cfg.Config) {
	ethClient, err := ethclient.Dial(config.RpcUrl)

	if err != nil {
		panic(err)
	}

	vaultRepository := contract_repositories.NewVaultRepository(ethClient)
	vaultRetriever := use_cases.NewVaultRetriever(vaultRepository)
	positionManagerServer := http_controllers.NewVaultController(vaultRetriever)

	router := mux.NewRouter()
	router.HandleFunc("/vault/{address}", positionManagerServer.GetVault)

	log.Printf("Listening on port %s\n", port)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}
