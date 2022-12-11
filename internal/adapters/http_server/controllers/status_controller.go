package http_controllers

import (
	"net/http"

	"github.com/protohedge/protohedge.api/internal/core/use_cases"
)

type statusController struct {
	vaultRetriever use_cases.VaultRetriever
	pnlRetriever   use_cases.PnlRetriever
}

func NewStatusController() statusController {
	return statusController{}
}

func (s *statusController) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("OK"))
}
