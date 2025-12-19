package config

type Config struct {
	DBUrl string
}

func Load() *Config {
	return &Config{
		DBUrl: ("postgres://postgres:<Password>@localhost:5432/userdb"),
	}
}
