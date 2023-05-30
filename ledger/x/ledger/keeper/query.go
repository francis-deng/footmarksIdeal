package keeper

import (
	"ledger/x/ledger/types"
)

var _ types.QueryServer = Keeper{}
