package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1102, "black address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1103, "game is not parseable: %s")
	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1104, "game by id not found: %s")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1105, "message creator is not a player: %s")
	ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1106, "player tried to play out of turn: %s")
	ErrWrongMove        = sdkerrors.Register(ModuleName, 1107, "wrong move")
	ErrRedAlreadyPlayed = sdkerrors.Register(ModuleName, 1108, "red palyer has already played")
	ErrBlackAlreadyPlayed = sdkerrors.Register(ModuleName, 1109, "black player has already played")
	ErrInvalidDeadline = sdkerrors.Register(ModuleName, 1110, "deadline cannot be parsed: %s")
	ErrGameFinished = sdkerrors.Register(ModuleName, 1111, "game is already finished")
	ErrRedCannotPay = sdkerrors.Register(ModuleName, 1112, "red cannot pay the wager")
	ErrBlackCannotPay = sdkerrors.Register(ModuleName, 1113, "black cannot pay the wager")
	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1114, "cannot find winner by color: %s")
	ErrNothingToPay = sdkerrors.Register(ModuleName, 1115, "there is nothing to pay, should not have been called")
	ErrCannotRefundWager = sdkerrors.Register(ModuleName, 1116, "cannot refund the wager to: %s")
	ErrCannotPayWinnings = sdkerrors.Register(ModuleName, 1117, "cannot pay winnings to winners")
	ErrNotInRefundState = sdkerrors.Register(ModuleName, 1118, "game is not in a state to refund, move count: %d")

)
