package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ledger/testutil/keeper"
	"ledger/testutil/nullify"
	"ledger/x/ledger/keeper"
	"ledger/x/ledger/types"
)

func createNCids(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Cids {
	items := make([]types.Cids, n)
	for i := range items {
		items[i].Id = keeper.AppendCids(ctx, items[i])
	}
	return items
}

func TestCidsGet(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNCids(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCids(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCidsRemove(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNCids(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCids(ctx, item.Id)
		_, found := keeper.GetCids(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCidsGetAll(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNCids(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCids(ctx)),
	)
}

func TestCidsCount(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNCids(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCidsCount(ctx))
}
