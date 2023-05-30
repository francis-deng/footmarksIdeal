package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ledger/testutil/keeper"
	"ledger/testutil/nullify"
	"ledger/x/ledger/keeper"
	"ledger/x/ledger/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTstmt(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Tstmt {
	items := make([]types.Tstmt, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTstmt(ctx, items[i])
	}
	return items
}

func TestTstmtGet(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNTstmt(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTstmt(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTstmtRemove(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNTstmt(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTstmt(ctx,
			item.Index,
		)
		_, found := keeper.GetTstmt(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTstmtGetAll(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNTstmt(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTstmt(ctx)),
	)
}
