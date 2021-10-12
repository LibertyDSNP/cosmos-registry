package keeper

import (
	"context"

	"github.com/amparks100/registry/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DelegationAll(c context.Context, req *types.QueryAllDelegationRequest) (*types.QueryAllDelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var delegations []types.Delegation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	delegationsStore := prefix.NewStore(store, types.KeyPrefix(types.DelegationKeyPrefix))

	pageRes, err := query.Paginate(delegationsStore, req.Pagination, func(key []byte, value []byte) error {
		var delegation types.Delegation
		if err := k.cdc.Unmarshal(value, &delegation); err != nil {
			return err
		}

		delegations = append(delegations, delegation)
		return nil
	})

	if err != nil {
		ctx.Logger().Error("Could not find delegations", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDelegationResponse{Delegation: delegations, Pagination: pageRes}, nil
}

func (k Keeper) Delegation(c context.Context, req *types.QueryGetDelegationRequest) (*types.QueryGetDelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDelegation(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetDelegationResponse{Delegation: val}, nil
}
