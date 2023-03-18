package redis_mappings

import (
	"strconv"
	"strings"
	"time"

	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToRebalanceNotesModel(rebalanceNotesDto []string) ([]domain.RebalanceNote, error) {
	rebalanceNotes := []domain.RebalanceNote{}

	for _, rebalance_note_record := range rebalanceNotesDto {
		splitRebalanceNote := strings.Split(rebalance_note_record, ":")

		if len(splitRebalanceNote) != 3 {
			continue
		}

		timestamp, err := strconv.ParseInt(splitRebalanceNote[1], 10, 64)
		note := splitRebalanceNote[2]

		if err != nil {
			return nil, err
		}

		rebalanceNotes = append(rebalanceNotes, domain.RebalanceNote{
			Date: time.UnixMilli(timestamp).UTC(),
			Note: note,
		})
	}

	return rebalanceNotes, nil
}
