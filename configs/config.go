package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/ilyakaznacheev/cleanenv"
)

type conf struct {
	DBDriver      string `env:"DB_DRIVER"`
	DBHost        string `env:"DB_HOST"`
	DBPort        string `env:"DB_PORT"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBName        string `env:"DB_NAME"`
	WebServerPort string `env:"WEB_SERVER_PORT"`
	JWTSecret     string `env:"JWT_SECRET"`
	JWTExpiresIn  int    `env:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	cfg := &conf{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}
