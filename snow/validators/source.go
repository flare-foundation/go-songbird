package validators

import (
	"github.com/flare-foundation/flare/ids"
)

type Source interface {
	LoadValidators(blockID ids.ID) (map[ids.ID]uint64, error)
}
