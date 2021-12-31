package blockchain

import (
	"crypto/ed25519"
	"encoding/json"
)

func VerifyTransaction(st SignedTransaction) (bool, error) {
	message, err := json.Marshal(st.Tx)
	if err != nil {
		return false, err
	}
	result := ed25519.Verify(*st.PublicKey, message, st.Signature)

	return result, nil
}
