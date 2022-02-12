package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/term"
)

/*
Generates a keystore file for a given password
*/
func main() {

	args := os.Args

	password, err := password()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex, publicKeyHex, err := keys(password)
	if err != nil {
		log.Fatal(err)
	}

	if len(args) > 1 && args[1] == "print" {
		log.Println("Password: " + password)
		log.Println("Private key: " + strings.ToLower(privateKeyHex))
		log.Println("Ethereum Address: 0x" + strings.ToLower(publicKeyHex))
	}

	clean(publicKeyHex)
}

/*
password requests a password off the user
*/
func password() (string, error) {

	log.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)

	return strings.TrimSpace(password), nil
}

/*
keys generates a new keystore with the given password
and returns the private and public keys in hex format
*/
func keys(password string) (string, string, error) {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:], account.Address.Hex()[2:], nil
}

/*
clean changes the keystore file name and stores it in './keystore'
*/
func clean(publicKey string) {

	pattern := "*" + strings.ToLower(publicKey)

	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	keystorePath := filepath.Join(".", "keystore")
	err = os.MkdirAll(keystorePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	filename := "keystore/0x" + publicKey + ".json"

	err = os.Rename(matches[0], filename)
	if err != nil {
		log.Fatal(err)
	}
}
