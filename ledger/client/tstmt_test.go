package client

import (
	"context"
	"testing"
)

func TestTstmtRoundTrip(t *testing.T) {
	bName := "bob"
	c := "cosmos1gvx3a5w46dnl4nyfu3j5dn0wt3jst2f3l2j606"
	cid := "my cid"

	tstmt := TStmt{
		//Payor: b,
		Payee: c,
		Denom: "bto",
		Q:     1,
		Ts:    1999,
		Cid:   cid,
	}

	tsc, err := NewTStmt(context.Background(), bName)
	if err != nil {
		t.Fatal(err)
	}

	err = tsc.Commit(tstmt)
	if err != nil {
		t.Fatal(err)
	}

	gettedTstmt, err := tsc.Get(cid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("The following part means success")
	t.Logf("%+v", gettedTstmt)
}
