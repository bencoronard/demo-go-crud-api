package config

import (
	xhttp "github.com/bencoronard/demo-go-common-libs/http"
	xjwt "github.com/bencoronard/demo-go-common-libs/jwt"
)

func NewJwtVerifier() (xjwt.Verifier, error) {
	v, err := xjwt.NewAsymmVerifier(nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func NewAuthHeaderResolver() xhttp.AuthHeaderResolver {
	return xhttp.NewHttpAuthHeaderResolver(nil)
}
