package config

import (
	"fmt"
	"strconv"

	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

const envPath = "no-secrets.env"

var (
	Port                        string // PORT
	PprofPort                   string // PPROF_PORT
	RateLimit                   int    // RATE_LIMIT
	LogLevel                    string // LOG_LEVEL
	PersonDataFetchWorkersCount int    // PERSON_DATA_FETCH_WORKERS_COUNT
	MovieDataFetchWorkersCount  int    // MOVIE_DATA_FETCH_WORKERS_COUNT
	LogGoroutineCount           bool    // LOG_GOROUTINE_COUNT
)

func init() {
	LoadEnv()
}

func LoadEnv() {
	fmt.Println("Loading .env file...")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file. err", err)
	}

	Port = getEnvString("PORT", "3001")
	PprofPort = getEnvString("PPROF_PORT", "6060")
	LogLevel = getEnvString("LOG_LEVEL", "INFO")
	RateLimit = getEnvInt("RATE_LIMIT", 10)
	PersonDataFetchWorkersCount = getEnvInt("PERSON_DATA_FETCH_WORKERS_COUNT", 10)
	MovieDataFetchWorkersCount = getEnvInt("MOVIE_DATA_FETCH_WORKERS_COUNT", 10)
	LogGoroutineCount = getEnvBool("LOG_GOROUTINE_COUNT", false)
	fmt.Println("Load .env file completed")
}

func getEnvString(key, defaultValue string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Info("Environment variable not found. Using default value for ", key)
		return defaultValue
	}
	return str
}

func getEnvInt(key string, defaultValue int) int {
	str := os.Getenv(key)
	if str == "" {
		log.Info("Environment variable not found. Using default value for ", key)
		return defaultValue
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Info("Error parsing environment variable. Using default value for ", key)
		return defaultValue
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	str := os.Getenv(key)
	if str == "" {
		log.Info("Environment variable not found. Using default value for ", key)
		return defaultValue
	}
	value, err := strconv.ParseBool(str)
	if err != nil {
		log.Info("Error parsing environment variable. Using default value for ", key)
		return defaultValue
	}
	return value
}