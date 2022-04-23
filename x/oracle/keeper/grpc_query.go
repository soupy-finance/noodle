package keeper

import (
	"github.com/soupy-finance/noodle/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
