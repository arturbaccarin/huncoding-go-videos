package main

import (
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://www.youtube.com/watch?v=UrZeMk_AxLM&list=PLm-xZWCprwYRAsLvf43sg5ZWuIvmmnDp9&index=9

func main() {

	/*
		logTest := zap.Config{}
		logger, err := logTest.Build()
	*/

	logger, _ := zap.NewProduction() // forma simples de instanciar um logger e sair usando
	defer logger.Sync()

	_, err := strconv.Atoi("test")

	logger.Error("failed to convert string to int", zap.Error(err))

	logger.Info("failed to fecth URL",
		zap.String("url", "test"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	logConfiguration := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "mensagem",
			LevelKey:     "nivel",
			TimeKey:      "horario",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, _ = logConfiguration.Build()
	logger.Info("failed to fecth URL",
		zap.String("url", "test"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}
