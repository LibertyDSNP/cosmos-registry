package keeper

import (
	"github.com/amparks100/registry/x/registry/types"
)

var _ types.QueryServer = Keeper{}
