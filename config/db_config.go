package config

type DatabaseCfg struct {
	Host       string
	Port       int
	Username   string
	Password   string
	DBName     string
	SSLEnabled bool
}
