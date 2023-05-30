package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ledger/x/ledger/types"
)

func (k Keeper) CidsAll(goCtx context.Context, req *types.QueryAllCidsRequest) (*types.QueryAllCidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cidss []types.Cids
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	cidsStore := prefix.NewStore(store, types.KeyPrefix(types.CidsKey))

	pageRes, err := query.Paginate(cidsStore, req.Pagination, func(key []byte, value []byte) error {
		var cids types.Cids
		if err := k.cdc.Unmarshal(value, &cids); err != nil {
			return err
		}

		cidss = append(cidss, cids)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCidsResponse{Cids: cidss, Pagination: pageRes}, nil
}

func (k Keeper) Cids(goCtx context.Context, req *types.QueryGetCidsRequest) (*types.QueryGetCidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	cids, found := k.GetCids(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCidsResponse{Cids: cids}, nil
}
