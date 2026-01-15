package config

import (
	"github.com/bencoronard/demo-go-common-libs/vault"
	"go.uber.org/fx"
)

func NewVaultClient(lc fx.Lifecycle, e *envCfg) (vault.Client, error) {
	return vault.NewTokenClient(lc, e.Vault.URI, e.Vault.Token)
}
