package client

import (
	"context"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"golang.org/x/xerrors"
	"ledger/x/ledger/types"
	"log"
)

var addressPrefix = "cosmos"

type TStmtC struct {
	localAccountName string

	parentCtx context.Context
	cc        cosmosclient.Client
}

type TStmt struct {
	Payor string `json:"payor"`
	Payee string `json:"payee"`
	Denom string `json:"denom"`
	Q     uint64 `json:"q"`
	Ts    uint64 `json:"ts"`
	Cid   string `json:"cid"`
}

func NewTStmt(ctx context.Context, creator string) (*TStmtC, error) {
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		return nil, err
	}

	return &TStmtC{
		localAccountName: creator,

		parentCtx: ctx,
		cc:        client,
	}, nil
}

func (ts *TStmtC) Get(cid string) (*TStmt, error) {
	msg := types.QueryGetTstmtRequest{Index: cid}

	qClient := types.NewQueryClient(ts.cc.Context())
	res, err := qClient.Tstmt(ts.parentCtx, &msg)

	if err != nil {
		return nil, err
	}
	if res.Tstmt.Index != res.Tstmt.Cid {
		return nil, xerrors.New("cid field was tampered maliciously by something")
	}

	return &TStmt{
		Payor: res.Tstmt.Payor,
		Payee: res.Tstmt.Payee,
		Denom: res.Tstmt.Denom,
		Q:     res.Tstmt.Q,
		Ts:    res.Tstmt.Ts,
		Cid:   res.Tstmt.Cid,
	}, nil
}

func (ts *TStmtC) Commit(tstmt TStmt) error {
	account, accountAddr, err := getLocalAccount(ts.parentCtx, ts.cc, ts.localAccountName)
	if err != nil {
		return err
	}

	// replace payor with the creator's address
	log.Printf("accountAddr: %s", accountAddr)
	msg := types.NewMsgCommitTstmt(accountAddr, accountAddr, tstmt.Payee, tstmt.Denom, tstmt.Q, tstmt.Ts, tstmt.Cid)

	txResp, err := ts.cc.BroadcastTx(ts.parentCtx, account, msg)
	if err != nil {
		return err
	}
	log.Println("txResp.code: ", txResp.Code)

	return nil
}

func getLocalAccount(ctx context.Context, client cosmosclient.Client, accountName string) (cosmosaccount.Account, string, error) {
	account := cosmosaccount.Account{}

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		return account, "", err
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		return account, "", err
	}

	return account, addr, nil
}
