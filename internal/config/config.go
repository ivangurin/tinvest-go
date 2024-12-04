package config

import (
	"context"
	"os"
	"strconv"
	"tinvest-go/internal/pkg/logger"

	"github.com/joho/godotenv"
)

const (
	AppName   = "ivangurin.tinvest-go"
	DbDsn     = "./database/database.sqlite"
	DbDsnTest = "../../../database/database_test.sqlite"
	JaegerUrl = "http://jaeger:14268/api/traces"

	// nolint:gosec
	envParamBotToken = "TINVEST_BOT_TOKEN"
	envParamBotDebug = "TINVEST_BOT_DEBUG"
)

var (
	BotToken string
	BotDebug bool
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf(context.Background(), "failed to load .env file: %s", err.Error())
	}

	BotToken = os.Getenv(envParamBotToken)
	BotDebug, err = strconv.ParseBool(os.Getenv(envParamBotDebug))
	if err != nil {
		logger.Errorf(context.Background(), "failed to parse %s: %s", envParamBotDebug, err.Error())
	}
}
