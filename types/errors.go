package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types/errors"
)

// Ethermint error codes
const (
	// DefaultCodespace reserves a Codespace for Ethermint.
	DefaultCodespace  = "ethermint"

	CodeInvalidValue   uint32 = 1
	CodeInvalidChainID uint32 = 2
	CodeInvalidSender  uint32 = 3
	CodeVMExecution    uint32 = 4
	CodeInvalidNonce   uint32 = 5
	CodeInternalError  uint32 = 6
)

// CodeToDefaultMsg takes the CodeType variable and returns the error string
//func CodeToDefaultMsg(code sdk.CodeType) string {
//	switch code {
//	case CodeInvalidValue:
//		return "invalid value"
//	case CodeInvalidChainID:
//		return "invalid chain ID"
//	case CodeInvalidSender:
//		return "could not derive sender from transaction"
//	case CodeVMExecution:
//		return "error while executing evm transaction"
//	case CodeInvalidNonce:
//		return "invalid nonce"
//	default:
//		return sdk.CodeToDefaultMsg(code)
//	}
//}

// ErrInvalidValue returns a standardized SDK error resulting from an invalid value.
func ErrInvalidValue(msg string) *sdk.Error {
	return sdk.Register(DefaultCodespace, CodeInvalidValue, msg)
}

// ErrInvalidChainID returns a standardized SDK error resulting from an invalid chain ID.
func ErrInvalidChainID(msg string) error {
	return sdk.Register(DefaultCodespace, CodeInvalidChainID, msg)
}

// ErrInvalidSender returns a standardized SDK error resulting from an invalid transaction sender.
func ErrInvalidSender(msg string) error {
	return sdk.Register(DefaultCodespace, CodeInvalidSender, msg)
}

// ErrVMExecution returns a standardized SDK error resulting from an error in EVM execution.
func ErrVMExecution(msg string) error {
	return sdk.Register(DefaultCodespace, CodeVMExecution, msg)
}

// ErrVMExecution returns a standardized SDK error resulting from an error in EVM execution.
func ErrInvalidNonce(msg string) error {
	return sdk.Register(DefaultCodespace, CodeInvalidNonce, msg)
}

func ErrInternalError(msg string) error {
	return sdk.Register(DefaultCodespace, CodeInternalError, msg)
}