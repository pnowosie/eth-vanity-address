package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	prefixRegex := regexp.MustCompile("^0x[0-9a-fA-F]{1,39}$")
	suffixRegex := regexp.MustCompile("^[0-9a-fA-F]{1,39}$")

	prefix := flag.String("prefix", "", "address prefix")
	suffix := flag.String("suffix", "", "address suffix")
	concurrency := flag.Int("concurrency", 4, "concurrent goroutines")
	ignoreCase := flag.Bool("ignore-case", false, "case insensitive")
	flag.Parse()

	if *prefix == "" && *suffix == "" {
		log.Fatal("Must specify prefix or suffix")
	}

	if *prefix != "" && !prefixRegex.MatchString(*prefix) {
		log.Fatal("Prefix must begin with '0x' and contain only valid characters")
	}

	if *suffix != "" && !suffixRegex.MatchString(*suffix) {
		log.Fatal("Suffix must contain only valid characters")
	}

	introMessage := fmt.Sprintf("Generating address with prefix=%s , suffix=%s\n", *prefix, *suffix)
	fmt.Println((introMessage))

	var searchPrefix string
	var searchSuffix string
	if *ignoreCase {
		searchPrefix = strings.ToLower(*prefix)
		searchSuffix = strings.ToLower(*suffix)
	} else {
		searchPrefix = *prefix
		searchSuffix = *suffix
	}

	var wg sync.WaitGroup

	for i := 1; i <= *concurrency; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			findAddressWorker(i, searchPrefix, searchSuffix, *ignoreCase)
		}()
	}

	wg.Wait()
}

func findAddressWorker(id int, prefix string, suffix string, ignoreCase bool) {
	start := time.Now()
	fmt.Printf("Worker %d starting...\n", id)

	for {
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
		var processedAddress string
		if ignoreCase {
			processedAddress = strings.ToLower(address)
		} else {
			processedAddress = address
		}

		if strings.HasPrefix(processedAddress, prefix) && strings.HasSuffix(processedAddress, suffix) {
			publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
			privateKeyBytes := crypto.FromECDSA(privateKey)

			fmt.Printf("\nWorker %d found address:\n", id)
			fmt.Println("Address:", address)
			fmt.Println("Public key:", hexutil.Encode(publicKeyBytes)[4:])
			fmt.Println("Private key:", hexutil.Encode(privateKeyBytes)[2:])

			elapsed := time.Since(start)
			fmt.Printf("Total time: %s\n", elapsed)
			break
		}
	}
	// Exit as soon as any worker finds the address
	os.Exit(0)
}
