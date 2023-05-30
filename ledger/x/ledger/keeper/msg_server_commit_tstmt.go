package keeper

import (
	"context"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ledger/x/ledger/types"
)

var CreatorAndPayorMustBeenEqualError = errors.Register("ledger", 1, "creator and payor must be equal")
var CidNotExistedError = errors.Register("ledger", 2, "cid is not allowed to be empty")
var UsedCidInTheStoreError = errors.Register("ledger", 3, "used cid in the store")

func (k msgServer) CommitTstmt(goCtx context.Context, msg *types.MsgCommitTstmt) (*types.MsgCommitTstmtResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != msg.Payor {
		return nil, CreatorAndPayorMustBeenEqualError
	}
	if len(strings.TrimSpace(msg.Cid)) == 0 {
		return nil, CidNotExistedError
	}

	_, found := k.GetTstmt(ctx, msg.Cid)
	if found {
		return nil, UsedCidInTheStoreError
	}

	var tstmt = types.Tstmt{
		Payor: msg.Payor,
		Payee: msg.Payee,
		Denom: msg.Denom,
		Q:     msg.Q,
		Ts:    msg.Ts,
		Cid:   msg.Cid,
	}
	// cid as a key value
	tstmt.Index = msg.Cid

	payor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(err, "creator: %s", msg.Creator)
	}
	payee, err := sdk.AccAddressFromBech32(msg.Payee)
	if err != nil {
		return nil, errors.Wrapf(err, "payee: %s", msg.Payee)
	}
	c := sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.Q)))

	err = k.bankKeeper.SendCoins(ctx, payor, payee, []sdk.Coin{c})
	if err != nil {
		return nil, errors.Wrapf(err, "coin: %+v", c)
	}
	k.SetTstmt(ctx, tstmt)
	k.AppendCids(ctx, types.Cids{
		Cid: msg.Cid,
	})

	return &types.MsgCommitTstmtResponse{}, nil
}
