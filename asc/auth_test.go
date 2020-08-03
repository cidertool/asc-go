package asc

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenConfig(t *testing.T) {
	// This is a key that I generated solely for mocking purposes. This is not a
	// real secret, so don't get any funny ideas. If you need to regenerate it,
	// run this openssl command in a shell and copy the contents of key.pem to the string:
	//
	//   openssl ecparam -name prime256v1 -genkey -noout | openssl pkcs8 -topk8 -nocrypt -out key.pem
	//
	// This will generate the ASN.1 PKCS#8 representation of the private key needed
	// to create a valid token. If you are looking at this test to see how to make a key,
	// reference Apple's documentation on this subject instead.
	//
	// https://developer.apple.com/documentation/appstoreconnectapi/creating_api_keys_for_app_store_connect_api
	var privPEMData = []byte(`
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgHuRdbDHRCtzCr0RA
UM0BwX7QPb7lbZNLvXmeG/k9k2+hRANCAATd7nn03pbNquj7IwUMy5SrOFRm71Sb
PURJWPQa24fI+wNPDi4OzjkB2g6fa5BHqam1gRlZHe8BU3+IjuC3AUFz
-----END PRIVATE KEY-----		
`)

	token, err := NewTokenConfig("TEST", "TEST", time.Now().Add(20), privPEMData)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if components := strings.Split(token.token, "."); len(components) != 3 {
		t.Errorf("Expected token length to be 3, was %d", len(components))
	}
}

func TestNewTokenConfigBadPEM(t *testing.T) {
	_, err := NewTokenConfig("TEST", "TEST", time.Now().Add(20), []byte("TEST"))
	if err == nil {
		t.Error("Expected error for invalid PEM, got nil")
	}
}

func TestNewTokenConfigPrivateKeyNotPKCS8(t *testing.T) {
	var badKey = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIXpcxwADKgwJSwxz24ypAMDFUHPrirqhcx0vimrl9L2oAoGCCqGSM49
AwEHoUQDQgAE7Ee8TlNaDqWa6O/Yw/nqHVEiJwYS+wt5cd7DC85nhsDxaU8M2Uy5
oH1YGuY57H3BQ3zLPVPsN+A8xnInGDa8yQ==
-----END EC PRIVATE KEY-----
`)

	_, err := NewTokenConfig("TEST", "TEST", time.Now().Add(20), badKey)
	if err == nil {
		t.Error("Expected error for non-PKCS8 PEM, got nil")
	}
}

// func TestNewTokenConfigFailsSigning(t *testing.T) {
// 	var privPEMData = []byte(`
// -----BEGIN PRIVATE KEY-----
// MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgk865haucvLiNrbFE
// Pi4TSaSvdUE0KdR5RhZGXnP9r4ihRANCAATJfqeCoJ9UEsNvCSgEf1/Nqjby2/xm
// SZJ6C98MewyQAagP58pyy48r1jbnx7WLZOeQfaVZOoHyeV6Gbz01D/AJ
// -----END PRIVATE KEY-----
// `)

// 	_, err := NewTokenConfig("", "", time.Now(), privPEMData)
// 	if err == nil {
// 		t.Error("Expected error for bad signing string, got nil")
// 	}
// }

func TestAuthTransport(t *testing.T) {
	token := "TEST.TEST.TEST"
	transport := AuthTransport{token: token}
	client := transport.Client()

	req, _ := http.NewRequest("GET", "", nil)
	_, _ = client.Do(req)

	got, want := req.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", token)
	assert.Equal(t, got, want)
}

func TestAuthTransportCustomTransport(t *testing.T) {
	token := "TEST.TEST.TEST"
	transport := AuthTransport{token: token, Transport: http.DefaultTransport}
	client := transport.Client()

	req, _ := http.NewRequest("GET", "", nil)
	_, _ = client.Do(req)

	got, want := req.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", token)
	assert.Equal(t, got, want)
}
