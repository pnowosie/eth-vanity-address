package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	prefix_regex := regexp.MustCompile("^0x[0-9a-fA-F]{1,39}$")
	suffix_regex := regexp.MustCompile("^[0-9a-fA-F]{1,39}$")

	prefix := flag.String("prefix", "", "address prefix")
	suffix := flag.String("suffix", "", "address suffix")
	flag.Parse()

	if *prefix == "" && *suffix == "" {
		log.Fatal("Must specify prefix or suffix")
	}

	if *prefix != "" && !prefix_regex.MatchString(*prefix) {
		log.Fatal("Prefix must begin with '0x' and contain only valid characters")
	}

	if *suffix != "" && !suffix_regex.MatchString(*suffix) {
		log.Fatal("Suffix must contain only valid characters")
	}

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("Address:", address)
	fmt.Println("Public key:", hexutil.Encode(publicKeyBytes)[4:])
	fmt.Println("Private key:", hexutil.Encode(privateKeyBytes)[2:])
}
