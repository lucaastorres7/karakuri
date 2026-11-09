package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
}

func InitConfig() Config {
	godotenv.Load()

	return Config{
		Port: getEnv("PORT", 8000),
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