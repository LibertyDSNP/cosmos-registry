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

func (k Keeper) RegistrationAll(c context.Context, req *types.QueryAllRegistrationRequest) (*types.QueryAllRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registrations []types.Registration
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	registrationStore := prefix.NewStore(store, types.KeyPrefix(types.RegistrationKeyPrefix))

	pageRes, err := query.Paginate(registrationStore, req.Pagination, func(key []byte, value []byte) error {
		var registration types.Registration
		if err := k.cdc.Unmarshal(value, &registration); err != nil {
			return err
		}

		registrations = append(registrations, registration)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegistrationResponse{Registration: registrations, Pagination: pageRes}, nil
}

func (k Keeper) Registration(c context.Context, req *types.QueryGetRegistrationRequest) (*types.QueryGetRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRegistration(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRegistrationResponse{Registration: val}, nil
}
