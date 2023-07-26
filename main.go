package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/text/message"
)

var (
	workerProgressUpdateDuration  = 5 * time.Second
	handlerProgressUpdateDuration = 15 * time.Second
)

func main() {
	concurrency := runtime.NumCPU()
	prefixRegex := regexp.MustCompile("^0x[0-9a-fA-F]{1,39}$")
	suffixRegex := regexp.MustCompile("^[0-9a-fA-F]{1,39}$")

	prefix := flag.String("prefix", "", "address prefix")
	suffix := flag.String("suffix", "", "address suffix")
	ignoreCase := flag.Bool("ignore-case", false, "case insensitive")
	password := flag.String("password", "", "when provided saves key into password protected key file")
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

	introMessage := fmt.Sprintf("Generating address with %d workers, prefix=%s, suffix=%s\n", concurrency, *prefix, *suffix)
	log.Println((introMessage))

	var searchPrefix string
	var searchSuffix string
	if *ignoreCase {
		searchPrefix = strings.ToLower(*prefix)
		searchSuffix = strings.ToLower(*suffix)
	} else {
		searchPrefix = *prefix
		searchSuffix = *suffix
	}

	progressChn := make(chan int)
	go func() {
		handleProgressUpdate(progressChn)
	}()

	keyFoundChn := make(chan *ecdsa.PrivateKey)
	trimPasswd := strings.TrimRight(*password, "\r\n")
	go func() {
		handleKeyFound(keyFoundChn, trimPasswd)
	}()

	// https://www.developer.com/languages/os-signals-go/
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)

	go func() {
		for {
			s := <-sigchnl
			handleStopSignal(s)
		}
	}()

	var wg sync.WaitGroup

	for i := 1; i <= concurrency; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			findAddressWorker(i, searchPrefix, searchSuffix, *ignoreCase, progressChn, keyFoundChn)
		}()
	}

	wg.Wait()
}

func findAddressWorker(
	id int,
	prefix string,
	suffix string,
	ignoreCase bool,
	progressChn chan int,
	keyFoundChn chan *ecdsa.PrivateKey,
) {
	start := time.Now()
	keysChecked := 0

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

			log.Println(fmt.Sprintf("Worker %d found address: %s", id, address))
			keyFoundChn <- privateKey

			// no break, find me more addresses
			// break
		}
		keysChecked++

		if time.Since(start) > workerProgressUpdateDuration {
			progressChn <- keysChecked
			keysChecked, start = 0, time.Now()
		}
	}

	// Exit when process is interrupted `Ctrl+C` or terminated other way
}

func handleStopSignal(signal os.Signal) {
	if signal == syscall.SIGTERM {
		log.Println("Replit is killing me! Got kill signal.")
		fmt.Println("Program will terminate now.")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		log.Println("Received interrupt signal. BYE!")
		fmt.Println("Closing.")
		os.Exit(0)
	}
	// any other signal - just ignore
}

func handleProgressUpdate(updateChn chan int) {
	p := message.NewPrinter(message.MatchLanguage("en"))
	keysChecked, start := 0, time.Now()
	for keysFromWorker := range updateChn {
		keysChecked += keysFromWorker
		if time.Since(start) > handlerProgressUpdateDuration {
			log.Println(p.Sprintf("Total keys checked: %d", keysChecked))
			start = time.Now()
		}
	}
}

func handleKeyFound(keyChn chan *ecdsa.PrivateKey, password string) {
	for privateKey := range keyChn {
		publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		privateKeyBytes := crypto.FromECDSA(privateKey)

		log.Println(strings.Join([]string{
			"Found key",
			fmt.Sprintf("Address    : %s", address),
			fmt.Sprintf("Public key : %s", hexutil.Encode(publicKeyBytes)[4:]),
			fmt.Sprintf("Private key: %s\n", hexutil.Encode(privateKeyBytes)[2:]),
		}, "\n"))
	}
}
