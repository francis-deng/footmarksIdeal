package ledger_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ledger/testutil/keeper"
	"ledger/testutil/nullify"
	"ledger/x/ledger"
	"ledger/x/ledger/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TstmtList: []types.Tstmt{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CidsList: []types.Cids{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CidsCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LedgerKeeper(t)
	ledger.InitGenesis(ctx, *k, genesisState)
	got := ledger.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TstmtList, got.TstmtList)
	require.ElementsMatch(t, genesisState.CidsList, got.CidsList)
	require.Equal(t, genesisState.CidsCount, got.CidsCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
