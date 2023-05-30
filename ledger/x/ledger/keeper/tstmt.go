package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ledger/x/ledger/types"
)

// SetTstmt set a specific tstmt in the store from its index
func (k Keeper) SetTstmt(ctx sdk.Context, tstmt types.Tstmt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TstmtKeyPrefix))
	b := k.cdc.MustMarshal(&tstmt)
	store.Set(types.TstmtKey(
		tstmt.Index,
	), b)
}

// GetTstmt returns a tstmt from its index
func (k Keeper) GetTstmt(
	ctx sdk.Context,
	index string,

) (val types.Tstmt, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TstmtKeyPrefix))

	b := store.Get(types.TstmtKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTstmt removes a tstmt from the store
func (k Keeper) RemoveTstmt(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TstmtKeyPrefix))
	store.Delete(types.TstmtKey(
		index,
	))
}

// GetAllTstmt returns all tstmt
func (k Keeper) GetAllTstmt(ctx sdk.Context) (list []types.Tstmt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TstmtKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tstmt
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
