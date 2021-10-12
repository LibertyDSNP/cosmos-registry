package keeper_test

import (
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

func createNRegistration(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Registration {
	items := make([]types.Registration, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRegistration(ctx, items[i])
	}
	return items
}

func TestRegistrationGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegistration(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRegistration(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestRegistrationRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegistration(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegistration(ctx,
			item.Index,
		)
		_, found := keeper.GetRegistration(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRegistrationGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNRegistration(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllRegistration(ctx))
}
