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

func TestRegistrationQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistration(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRegistrationRequest
		response *types.QueryGetRegistrationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRegistrationRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRegistrationResponse{Registration: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRegistrationRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRegistrationResponse{Registration: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRegistrationRequest{
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
			response, err := keeper.Registration(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestRegistrationQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRegistration(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRegistrationRequest {
		return &types.QueryAllRegistrationRequest{
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
			resp, err := keeper.RegistrationAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Registration), step)
			require.Subset(t, msgs, resp.Registration)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RegistrationAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Registration), step)
			require.Subset(t, msgs, resp.Registration)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RegistrationAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RegistrationAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
