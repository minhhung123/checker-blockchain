package keeper

import (
	"context"
	"strings"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rules "github.com/minhhung123/checkers/x/checkers/rules"

	"github.com/minhhung123/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanPlayMove(goCtx context.Context, req *types.QueryCanPlayMoveRequest) (*types.QueryCanPlayMoveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	storedGame, found := k.GetStoredGame(ctx, req.IdValue)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, types.ErrGameNotFound.Error(), req.IdValue)
	}

	// Is the game already won ? What happen when it is forfeited
	if storedGame.Winner != rules.NO_PLAYER.Color {
		return &types.QueryCanPlayMoveResponse {
			Possible: false,
			Reason: types.ErrGameFinished.Error(),
		}, nil
	}

	// Is the player given a valid player
	var player rules.Player
	if strings.Compare(rules.RED_PLAYER.Color, req.Player) == 0 {
		player = rules.RED_PLAYER
	} else if strings.Compare(rules.BLACK_PLAYER.Color, req.Player) == 0 {
		player = rules.BLACK_PLAYER
	} else {
		return &types.QueryCanPlayMoveResponse {
			Possible : false,
			Reason   : types.ErrCreatorNotPlayer.Error(),  
		}, nil
	}

	// Is it the player's turn
	game, err := storedGame.ParseGame()
	if err != nil {
		return nil, err
	}
	if !game.TurnIs(player) {
		return &types.QueryCanPlayMoveResponse {
			Possible : false,
			Reason   : types.ErrNotPlayerTurn.Error(),
		}, nil
	}

	// Attempt the move in memory and report back
	_, moveErr := game.Move(
		rules.Pos{
			X: int(req.FromX),
			Y: int(req.FromY),
		},
		rules.Pos{
			X: int(req.ToX),
			Y: int(req.ToY),
		},
	)

	if moveErr != nil {
		return &types.QueryCanPlayMoveResponse{
			Possible : false,
			Reason   : fmt.Sprintf(types.ErrWrongMove.Error(), moveErr.Error()), 
		}, nil
	}

	return &types.QueryCanPlayMoveResponse{
		Possible : true,
		Reason   : "ok",
	}, nil
	
}
