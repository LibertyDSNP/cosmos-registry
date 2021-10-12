package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/amparks100/registry/testutil/keeper"
	"github.com/amparks100/registry/x/identity/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDelegationQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDelegation(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDelegationRequest
		response *types.QueryGetDelegationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDelegationRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetDelegationResponse{Delegation: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDelegationRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetDelegationResponse{Delegation: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDelegationRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Delegation(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestDelegationQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDelegation(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDelegationRequest {
		return &types.QueryAllDelegationRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DelegationAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Delegation), step)
			require.Subset(t, msgs, resp.Delegation)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DelegationAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Delegation), step)
			require.Subset(t, msgs, resp.Delegation)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DelegationAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DelegationAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
