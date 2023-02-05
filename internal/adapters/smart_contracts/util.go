package smart_contracts

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func MapError(err error) error {
	if errors.Is(err, bind.ErrNoCode) {
		return fmt.Errorf("%w, %v", domain.ErrNotFound, err)
	}

	return err
}
