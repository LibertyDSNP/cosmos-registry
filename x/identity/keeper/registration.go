package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/amparks100/registry/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRegistration set a specific registration in the store from its index
func (k Keeper) SetRegistration(ctx sdk.Context, registration types.Registration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrationKeyPrefix))
	b := k.cdc.MustMarshal(&registration)
	store.Set(types.RegistrationKey(
		registration.Index,
	), b)
}

// GetRegistration returns a registration from its index
func (k Keeper) GetRegistration(
	ctx sdk.Context,
	index string,

) (val types.Registration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrationKeyPrefix))

	b := store.Get(types.RegistrationKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegistration removes a registration from the store
func (k Keeper) RemoveRegistration(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrationKeyPrefix))
	store.Delete(types.RegistrationKey(
		index,
	))
}

// GetAllRegistration returns all registration
func (k Keeper) GetAllRegistration(ctx sdk.Context) (list []types.Registration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Registration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) DoesHandleExist(ctx sdk.Context, handle string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Registration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Handle == handle {
			return true
		}
	}
	return false
}

// GetCurrentDsnpId gets the current value of the dsnpid iterator
func (k Keeper) GetCurrentDsnpId(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DsnpIdKey))
	byteKey := types.KeyPrefix(types.DsnpIdKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	// Parse bytes
	id, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the id should be always formattable to iint64
		panic("cannot decode current dsnpid")
	}

	return id
}

// SetDsnpId sets the current value of the dsnpid iterator
func (k Keeper) SetDsnpId(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DsnpIdKey))
	byteKey := types.KeyPrefix(types.DsnpIdKey)
	bz := []byte(strconv.FormatUint(id, 10))
	store.Set(byteKey, bz)
}

// GetDsnpIdBytes returns the byte representation of the ID
func GetDsnpIdBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDsnpIdFromBytes returns ID in uint64 format from a byte array
func GetDsnpIdFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
