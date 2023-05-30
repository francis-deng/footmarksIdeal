package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"ledger/x/ledger/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated tstmt",
			genState: &types.GenesisState{
				TstmtList: []types.Tstmt{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated cids",
			genState: &types.GenesisState{
				CidsList: []types.Cids{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid cids count",
			genState: &types.GenesisState{
				CidsList: []types.Cids{
					{
						Id: 1,
					},
				},
				CidsCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
