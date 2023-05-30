package ledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ledger/x/ledger/keeper"
	"ledger/x/ledger/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the tstmt
	for _, elem := range genState.TstmtList {
		k.SetTstmt(ctx, elem)
	}
	// Set all the cids
	for _, elem := range genState.CidsList {
		k.SetCids(ctx, elem)
	}

	// Set cids count
	k.SetCidsCount(ctx, genState.CidsCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TstmtList = k.GetAllTstmt(ctx)
	genesis.CidsList = k.GetAllCids(ctx)
	genesis.CidsCount = k.GetCidsCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
