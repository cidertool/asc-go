package asc

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthTransport struct {
	Token     string
	Transport http.RoundTripper
}

func (t AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.Token))
	return t.transport().RoundTrip(req)
}

func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

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
	return &AuthTransport{Token: token}, nil
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
