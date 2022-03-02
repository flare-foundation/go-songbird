package validators

import (
	"github.com/flare-foundation/flare/ids"
)

type Source interface {
	LastAccepted() (ids.ID, error)
	LoadValidators(blockID ids.ID) (map[ids.ShortID]uint64, error)
}
