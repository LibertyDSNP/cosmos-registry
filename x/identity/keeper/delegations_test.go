package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/amparks100/registry/testutil/keeper"
	"github.com/amparks100/registry/x/identity/keeper"
	"github.com/amparks100/registry/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDelegation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Delegation {
	items := make([]types.Delegation, n)
	for i := range items {
		items[i].Index = fmt.Sprintf("%d", i)

		keeper.SetDelegation(ctx, items[i])
	}
	return items
}

func TestDelegationGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDelegation(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDelegation(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestDelegationRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDelegation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDelegation(ctx,
			item.Index,
		)
		_, found := keeper.GetDelegation(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestDelegationGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDelegation(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllDelegation(ctx))
}
