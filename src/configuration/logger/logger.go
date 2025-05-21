package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {

	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()}, // Corrigido: getOutotLogs para getOutputLogs
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder, // Mantém a cor para o nível de log
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error // Declarar err para evitar erro de shadowing com o log, _ abaixo
	log, err = logConfig.Build()
	if err != nil {
		panic(err) // Se a configuração do logger falhar, é um erro crítico
	}
}

// Info logs a message at InfoLevel.
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

// Warn logs a message at WarnLevel.
// Adicionada a função Warn.
// O parâmetro 'err' foi removido pois um aviso geralmente não está associado a um 'error' Go.
// Se um erro Go existir, logger.Error deve ser usado.
func Warn(message string, tags ...zap.Field) {
	log.Warn(message, tags...) // Usa o método Warn do zap.Logger
	log.Sync()
}

// Error logs a message at ErrorLevel.
func Error(message string, err error, tags ...zap.Field) {
	if err != nil { // Adicionada verificação para não adicionar "error": null se err for nil
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(message, tags...) // Corrigido de log.Info para log.Error
	log.Sync()
}

func getOutputLogs() string { // Corrigido: getOutotLogs para getOutputLogs
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "warn": // Adicionando o nível "warn"
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
