package identity_test

import (
	"testing"

	keepertest "github.com/amparks100/registry/testutil/keeper"
	"github.com/amparks100/registry/x/identity"
	"github.com/amparks100/registry/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		PortId: types.PortID,
		RegistrationList: []types.Registration{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DelegationList: []types.Delegation{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, genesisState.PortId, got.PortId)
	require.Len(t, got.RegistrationList, len(genesisState.RegistrationList))
	require.Subset(t, genesisState.RegistrationList, got.RegistrationList)
	require.Len(t, got.DelegationList, len(genesisState.DelegationList))
	require.Subset(t, genesisState.DelegationList, got.DelegationList)
	// this line is used by starport scaffolding # genesis/test/assert
}
