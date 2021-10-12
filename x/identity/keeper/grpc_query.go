package keeper

import (
	"github.com/amparks100/registry/x/identity/types"
)

var _ types.QueryServer = Keeper{}
