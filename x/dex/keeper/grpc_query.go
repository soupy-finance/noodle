package keeper

import (
	"github.com/soupy-finance/noodle/x/dex/types"
)

var _ types.QueryServer = Keeper{}
