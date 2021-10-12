package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/amparks100/registry/testutil/keeper"
	"github.com/amparks100/registry/x/identity/keeper"
	"github.com/amparks100/registry/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
