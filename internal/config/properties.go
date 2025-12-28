package config

import "github.com/caarlos0/env/v11"

type Properties struct {
	Env EnvCfg
}

func NewProperties() (*Properties, error) {
	var props Properties
	if err := env.Parse(&props); err != nil {
		return nil, err
	}
	return &props, nil
}

type EnvCfg struct {
	App   AppCfg
	Vault VaultCfg
	OTEL  OTELCfg
	DB    DBCfg
}

type AppCfg struct {
	ListenPort    int    `env:"APP_LISTEN_PORT"`
	Environment   string `env:"APP_ENVIRONMENT"`
	PublicKeyPath string `env:"APP_PUBLIC_KEY_PATH"`
}

type VaultCfg struct {
	URI   string `env:"VAULT_URI"`
	Token string `env:"VAULT_TOKEN"`
}

type OTELCfg struct {
	MetricsEndpoint           string  `env:"OTEL_COL_METRICS_ENDPOINT"`
	TracesEndpoint            string  `env:"OTEL_COL_TRACES_ENDPOINT"`
	LogsEndpoint              string  `env:"OTEL_COL_LOGS_ENDPOINT"`
	MetricsSamplingFreqInMin  string  `env:"OTEL_METRICS_SAMPLING_FREQ_IN_MIN"`
	TracesSamplingProbability float64 `env:"OTEL_TRACES_SAMPLING_PROBABILITY"`
}

type DBCfg struct {
	ConnectionPoolCap         int `env:"CRUD_API_DB_CP_CAP"`
	ConnectionPoolIdleMin     int `env:"CRUD_API_DB_CP_IDLE_MIN"`
	ConnectionPoolIdleTimeout int `env:"CRUD_API_DB_CP_IDLE_TIMEOUT"`
	ConnectionTimeout         int `env:"CRUD_API_DB_CP_CONN_TIMEOUT"`
	ConnectionTTL             int `env:"CRUD_API_DB_CP_CONN_TTL"`
}
