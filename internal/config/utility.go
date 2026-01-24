package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	xhttp "github.com/bencoronard/demo-go-common-libs/http"
	xjwt "github.com/bencoronard/demo-go-common-libs/jwt"
)

func NewJwtVerifier(p *Properties) (xjwt.Verifier, error) {

	data, err := os.ReadFile(p.Env.App.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to parse public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	v, err := xjwt.NewAsymmVerifier(rsaPub)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func NewAuthHeaderResolver(v xjwt.Verifier) xhttp.AuthHeaderResolver {
	return xhttp.NewHttpAuthHeaderResolver(v)
}
