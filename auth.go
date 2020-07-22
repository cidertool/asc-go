package asc

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// AuthTransport is an http.RoundTripper implementation that stores the JWT created.
// If the token expires, the Rotate function should be called to update the stored token.
type AuthTransport struct {
	Transport http.RoundTripper
	token     string
}

// RoundTrip implements the http.RoundTripper interface to set the Authorization header
func (t AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return t.transport().RoundTrip(req)
}

// Client returns a new http.Client instance for use with asc.Client.
func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// NewTokenConfig returns a new AuthTransport instance that customizes the Authentication header of the request during transport.
// It can be customized further by supplying a custom http.RoundTripper instance to the Transport field.
func NewTokenConfig(keyID string, issuerID string, expiry time.Time, privateKey []byte) (*AuthTransport, error) {
	key, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Audience:  []string{"appstoreconnect-v1"},
		Issuer:    issuerID,
		ExpiresAt: &jwt.Time{Time: expiry},
	})
	t.Header["kid"] = keyID
	token, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &AuthTransport{token: token}, nil
}

func parsePrivateKey(blob []byte) (key interface{}, err error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, fmt.Errorf("no PEM blob found")
	}
	key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
