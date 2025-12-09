package config

import (
	"log"
	"os"
	"strconv"
)

type envStorage struct {
	LLMToken    string
	Model       string
	LLMUrl      string
	ServerName  string
	ServerPort  string
	Debug       bool
	Temperature float64
	TopP        float64
	MaxTokens   int
}

var Env *envStorage

func EnvLoad() {
	Env = &envStorage{}

	Env.LLMToken = os.Getenv("LLM_TOKEN")
	Env.Model = os.Getenv("MODEL")
	Env.LLMUrl = os.Getenv("LLM_URL")
	Env.ServerName = os.Getenv("SERVER_NAME")
	Env.ServerPort = os.Getenv("SERVER_PORT")

	debugStr := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		log.Fatalf("'DEBUG' type is not boolean: %f\n", err)
		return
	}

	temperatureStr := os.Getenv("TEMPERATURE")
	temperature, err := strconv.ParseFloat(temperatureStr, 32)
	if err != nil {
		log.Fatalf("'TEMPERATURE' type invalid: %f\n", err)
	}

	topPStr := os.Getenv("TOP_P")
	topP, err := strconv.ParseFloat(topPStr, 32)
	if err != nil {
		log.Fatalf("'TOP_P' type invalid: %f\n", err)
	}

	maxTokensStr := os.Getenv("MAX_TOKENS")
	maxTokens, err := strconv.Atoi(maxTokensStr)
	if err != nil {
		log.Fatalf("'MAX_TOKENS' type invalid: %f\n", err)
	}

	Env.Debug = debug
	Env.Temperature = temperature
	Env.TopP = topP
	Env.MaxTokens = maxTokens
}
