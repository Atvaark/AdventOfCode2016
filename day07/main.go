package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	addresses, error := openInput("input.txt")
	if error != nil {
		fmt.Printf("unable to open input: %v", error)
		os.Exit(1)
	}

	tlsCount, sslCount := getFeatureCount(addresses)
	fmt.Println("result1: ", tlsCount)
	fmt.Println("result2: ", sslCount)
}

func getFeatureCount(addresses []string) (TLScount int, SSKcount int) {
	for _, address := range addresses {
		tls, ssl := getFeatures(address)
		if tls {
			TLScount++
		}

		if ssl {
			SSKcount++
		}
	}

	return
}

func openInput(name string) ([]string, error) {
	file, error := os.Open(name)
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func getFeatures(address string) (supportsTLS bool, supportsSSL bool) {
	l := len(address)

	abba := false
	hypernetAbba := false
	ssl := false

	inBracket := false

	var abas []string
	var babs []string

	for i := 0; i < l-2; i++ {
		if address[i] == '[' {
			inBracket = true
			continue
		}

		if address[i] == ']' {
			inBracket = false
			continue
		}

		if i+3 < l {
			abbaRange := address[i : i+4]
			if abbaRange[0] != abbaRange[1] && abbaRange[0] == abbaRange[3] && abbaRange[1] == abbaRange[2] {
				if inBracket {
					hypernetAbba = true
				} else {
					abba = true
				}
			}
		}

		sslRange := address[i : i+3]
		if sslRange[0] != sslRange[1] && sslRange[0] == sslRange[2] {
			if inBracket {
				babs = append(babs, sslRange)
			} else {
				abas = append(abas, sslRange)
			}
		}
	}

CheckAba:
	for _, aba := range abas {
		var bab [3]byte
		bab[0] = aba[1]
		bab[1] = aba[0]
		bab[2] = aba[1]
		abaAsBab := string(bab[:])

		for _, bab := range babs {
			if abaAsBab == bab {
				ssl = true
				break CheckAba
			}
		}
	}

	return abba && !hypernetAbba, ssl
}
