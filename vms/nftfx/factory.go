package nftfx

import (
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
)

// ID that this Fx uses when labeled
var (
	ID = ids.ID{'n', 'f', 't', 'f', 'x'}
)

type Factory struct{}

func (f *Factory) New(*snow.Context) (interface{}, error) { return &Fx{}, nil }
