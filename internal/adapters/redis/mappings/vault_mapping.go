package redis_mappings

import (
	"strconv"
	"strings"
	"time"

	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToRebalanceHistoryModel(rebalanceHistoryDto []string) ([]domain.RebalanceNote, error) {
	rebalanceHistory := []domain.RebalanceNote{}

	for _, rebalance_note_record := range rebalanceHistoryDto {
		splitRebalanceNote := strings.Split(rebalance_note_record, ":")

		if len(splitRebalanceNote) != 3 {
			continue
		}

		timestamp, err := strconv.ParseInt(splitRebalanceNote[1], 10, 64)
		note := splitRebalanceNote[2]

		if err != nil {
			return nil, err
		}

		rebalanceHistory = append(rebalanceHistory, domain.RebalanceNote{
			Date: time.UnixMilli(timestamp).UTC(),
			Note: note,
		})
	}

	return rebalanceHistory, nil

}
