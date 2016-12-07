package main

import "testing"

func TestRun(t *testing.T) {

	addresses, error := openInput("input.txt")
	if error != nil {
		t.Errorf("unable to open input: %v", error)
		return
	}

	TLScount, SSLcount := getFeatureCount(addresses)

	const expectedTLScount = 115
	if TLScount != expectedTLScount {
		t.Errorf("invalid TLS count %d expected %d", TLScount, expectedTLScount)
		return
	}

	const expectedSSLcount = 231

	if SSLcount != expectedSSLcount {
		t.Errorf("invalid SSL count %d expected %d", SSLcount, expectedSSLcount)
	}
}
