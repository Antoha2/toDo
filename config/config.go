package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

func GetConfig() *Config {

	return &Config{
		DB: DBConfig{
			User:     "todoadmin",
			Password: "tododo",
			Host:     "localhost",
			Port:     5432,
			Dbname:   "tododb",
			Sslmode:  "",
		},
	}
}
