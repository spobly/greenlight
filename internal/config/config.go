package config

type Config struct {
	Port int
	Env  string
	Db   struct {
		DSN string
	}
}
