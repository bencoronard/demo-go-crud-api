package config

import (
	"github.com/bencoronard/demo-go-common-libs/vault"
	"go.uber.org/fx"
)

func NewVaultClient(lc fx.Lifecycle, p *Properties) (vault.Client, error) {
	return vault.NewTokenClient(lc, p.Env.Vault.URI, p.Env.Vault.Token)
}
