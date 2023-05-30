package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ledger/x/ledger/types"
)

func (k Keeper) TstmtAll(goCtx context.Context, req *types.QueryAllTstmtRequest) (*types.QueryAllTstmtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tstmts []types.Tstmt
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tstmtStore := prefix.NewStore(store, types.KeyPrefix(types.TstmtKeyPrefix))

	pageRes, err := query.Paginate(tstmtStore, req.Pagination, func(key []byte, value []byte) error {
		var tstmt types.Tstmt
		if err := k.cdc.Unmarshal(value, &tstmt); err != nil {
			return err
		}

		tstmts = append(tstmts, tstmt)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTstmtResponse{Tstmt: tstmts, Pagination: pageRes}, nil
}

func (k Keeper) Tstmt(goCtx context.Context, req *types.QueryGetTstmtRequest) (*types.QueryGetTstmtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTstmt(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTstmtResponse{Tstmt: val}, nil
}
