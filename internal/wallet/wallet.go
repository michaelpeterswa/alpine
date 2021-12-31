package wallet

import (
	"crypto/ed25519"
	"log"
)

func GenerateEd25519KeyPair() (*ed25519.PublicKey, *ed25519.PrivateKey) {
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal("keygen failed")
	}
	return &pubKey, &privKey
}
