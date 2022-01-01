package aerror

import (
	"encoding/json"
	"net/http"
)

type AlpineError int

type ErrorResponse struct {
	ErrorName string `json:"error"`
}

const (
	// An Unknown Error Occurred
	Unknown AlpineError = iota
	// Verification of Transaction (ED25519) Failed
	FailedVerification
	// Incorrect PublicKey Size (32 Bytes)
	PublicKeyBadLength
	// Incorrect Signature Length (64 Bytes)
	SignatureBadLength
	// Hashing Error
	HashError
)

func AerrorToName(ae AlpineError) string {
	switch ae {
	case Unknown:
		return "Unknown"
	case FailedVerification:
		return "FailedVerification"
	case PublicKeyBadLength:
		return "PublicKeyBadLength"
	case SignatureBadLength:
		return "SignatureBadLength"
	case HashError:
		return "HashError"
	default:
		return "NotImplemented"
	}
}

func NewErrorResponse(w http.ResponseWriter, ae AlpineError) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ErrorResponse{ErrorName: AerrorToName(ae)})
}
