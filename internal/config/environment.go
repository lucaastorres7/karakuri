package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        					 int
	DBUser      					 string
	DBPassword  					 string
	DBName      					 string
	DBHost      					 string
	DBPort      					 int
	TokenSecret 					 string
	TokenExpirationSeconds int
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Port:       						getEnv("PORT", 8000),
		DBUser:     						getEnv("DB_USER", "user"),
		DBPassword: 						getEnv("DB_PASSWORD", "password"),
		DBName:     						getEnv("DB_NAME", "database"),
		DBHost:     						getEnv("DB_HOST", "localhost"),
		DBPort:     						getEnv("DB_PORT", 5432),
		TokenSecret: 						getEnv("TOKEN_SECRET", "secret"),
		TokenExpirationSeconds: getEnv("TOKEN_EXPIRATION_SECONDS", 60*60*24), // 24 hours
	}
}

func getEnv[T int | string | float64 | bool](key string, fallback T) T {
	raw, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	var v any
	var err error

	switch any(fallback).(type){
	case int:
		v, err = strconv.Atoi(raw)
	case string:
		v = raw
	case float64:
		v, err = strconv.ParseFloat(raw, 64)
	case bool:
		v, err = strconv.ParseBool(raw)
	}
	if err != nil {
		return fallback
	}
	
	return v.(T)
}