package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/amparks100/registry/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDelegationId get the total number of delegations
func (k Keeper) GetDelegationId(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationIndex))
	byteKey := types.KeyPrefix(types.DelegationIndex)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetDelegationId set the total number of delegations
func (k Keeper) SetDelegationId(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationIndex))
	byteKey := types.KeyPrefix(types.DelegationIndex)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

//AppendDelegation appends a delegation in the store with new index
func (k Keeper) AppendDelegation(ctx sdk.Context, delegation types.Delegation) uint64 {
	index := k.GetDelegationId(ctx)

	delegation.Id = index

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&delegation)
	store.Set(GetDelegationIndexBytes(delegation.Id), appendedValue)

	// Update stored index
	k.SetDelegationId(ctx, index+1)

	return index + 1
}

// SetDelegations set a specific delegations in the store from its index
func (k Keeper) SetDelegation(ctx sdk.Context, delegation types.Delegation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationKeyPrefix))
	b := k.cdc.MustMarshal(&delegation)
	store.Set(types.DelegationKey(
		string(types.KeyPrefix(delegation.Index)),
	), b)
}

// GetDelegations returns a delegations from its index
func (k Keeper) GetDelegation(
	ctx sdk.Context,
	index string,

) (val types.Delegation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationKeyPrefix))

	b := store.Get(types.DelegationKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegations removes a delegations from the store
func (k Keeper) RemoveDelegation(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationKeyPrefix))
	store.Delete(types.DelegationKey(
		index,
	))
}

// GetAllDelegations returns all delegations
func (k Keeper) GetAllDelegation(ctx sdk.Context) (list []types.Delegation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllDelegationsForDsnpId returns all delegations
func (k Keeper) GetAllDelegationsForDsnpId(ctx sdk.Context, dsnpId uint64) (list []types.Delegation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Delegation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.DsnpId == dsnpId {
			list = append(list, val)
		}
	}

	return
}

// GetPostIDBytes returns the byte representation of the ID
func GetDelegationIndexBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
