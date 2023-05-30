package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCommitTstmt = "commit_tstmt"

var _ sdk.Msg = &MsgCommitTstmt{}

func NewMsgCommitTstmt(creator string, payor string, payee string, denom string, q uint64, ts uint64, cid string) *MsgCommitTstmt {
	return &MsgCommitTstmt{
		Creator: creator,
		Payor:   payor,
		Payee:   payee,
		Denom:   denom,
		Q:       q,
		Ts:      ts,
		Cid:     cid,
	}
}

func (msg *MsgCommitTstmt) Route() string {
	return RouterKey
}

func (msg *MsgCommitTstmt) Type() string {
	return TypeMsgCommitTstmt
}

func (msg *MsgCommitTstmt) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCommitTstmt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitTstmt) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
