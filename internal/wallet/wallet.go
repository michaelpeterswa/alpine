package wallet

import (
	"crypto/ed25519"
	"log"
)

type Wallet struct {
	Name       string              `json:"name"`
	PublicKey  *ed25519.PublicKey  `json:"public-key"`
	PrivateKey *ed25519.PrivateKey `json:"private-key"`
}

func GenerateEd25519KeyPair() (*ed25519.PublicKey, *ed25519.PrivateKey) {
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal("keygen failed")
	}
	return &pubKey, &privKey
}

func CreateNewWallet(s string) *Wallet {
	pubKey, privKey := GenerateEd25519KeyPair()

	return &Wallet{
		Name:       s,
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}
}

func SaveWallet(w Wallet) {
	// nop
}

func LoadWallet(s string) {
	// nop
}
