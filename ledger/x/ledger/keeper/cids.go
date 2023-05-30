package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ledger/x/ledger/types"
)

// GetCidsCount get the total number of cids
func (k Keeper) GetCidsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CidsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCidsCount set the total number of cids
func (k Keeper) SetCidsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CidsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCids appends a cids in the store with a new id and update the count
func (k Keeper) AppendCids(
	ctx sdk.Context,
	cids types.Cids,
) uint64 {
	// Create the cids
	count := k.GetCidsCount(ctx)

	// Set the ID of the appended value
	cids.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CidsKey))
	appendedValue := k.cdc.MustMarshal(&cids)
	store.Set(GetCidsIDBytes(cids.Id), appendedValue)

	// Update cids count
	k.SetCidsCount(ctx, count+1)

	return count
}

// SetCids set a specific cids in the store
func (k Keeper) SetCids(ctx sdk.Context, cids types.Cids) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CidsKey))
	b := k.cdc.MustMarshal(&cids)
	store.Set(GetCidsIDBytes(cids.Id), b)
}

// GetCids returns a cids from its id
func (k Keeper) GetCids(ctx sdk.Context, id uint64) (val types.Cids, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CidsKey))
	b := store.Get(GetCidsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCids removes a cids from the store
func (k Keeper) RemoveCids(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CidsKey))
	store.Delete(GetCidsIDBytes(id))
}

// GetAllCids returns all cids
func (k Keeper) GetAllCids(ctx sdk.Context) (list []types.Cids) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CidsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Cids
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCidsIDBytes returns the byte representation of the ID
func GetCidsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCidsIDFromBytes returns ID in uint64 format from a byte array
func GetCidsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
