package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ledger/testutil/keeper"
	"ledger/x/ledger/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LedgerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
