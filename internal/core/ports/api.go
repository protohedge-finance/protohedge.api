package ports

import (
	cfg "github.com/protohedge/protohedge.api/internal/config"
)

type Api interface {
	Start(config *cfg.Config)
}
