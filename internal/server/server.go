package server

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	cfg "github.com/protohedge/protohedge.api/internal/config"
	"github.com/protohedge/protohedge.api/internal/position_manager"
	"github.com/rs/cors"
)

var port = ":8080"

func CreateServer(config *cfg.Config) {
	ethClient, err := ethclient.Dial(config.RpcUrl)

	if err != nil {
		panic(err)
	}

	positionManagerService := position_manager.NewPositionManagerService(ethClient)
	positionManagerServer := NewPositionManagerServer(positionManagerService)

	router := mux.NewRouter()
	router.HandleFunc("/position_manager/{address}", positionManagerServer.GetPositionManager)

	log.Printf("Listening on port %s\n", port)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}
