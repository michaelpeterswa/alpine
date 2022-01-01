package wallet

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Wallet struct {
	Name       string              `json:"name"`
	PublicHex  string              `json:"public-hex"`
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

	hex := hex.EncodeToString(*pubKey)
	hexRepr := fmt.Sprintf("0x%s", hex)

	return &Wallet{
		Name:       s,
		PublicHex:  hexRepr,
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}
}

func SaveWallet(w Wallet) error {
	name := fmt.Sprintf("%s.alp.json", strings.ReplaceAll(currentWallet.Name, " ", ""))
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")

	err = enc.Encode(w)
	if err != nil {
		return err
	}
	return nil
}

func LoadWallet(s string) {
	w := &Wallet{}

	f, err := os.Open(s)
	if err != nil {
		log.Fatal("couldn't open wallet file")
	}

	err = json.NewDecoder(f).Decode(w)
	if err != nil {
		log.Fatal("couldn't decode wallet file: ", err)
	}

	currentWallet = w
}
