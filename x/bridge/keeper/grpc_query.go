package keeper

import (
	"github.com/soupy-finance/noodle/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
