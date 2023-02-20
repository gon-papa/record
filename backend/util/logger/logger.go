package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/gon-papa/record/config"
	"golang.org/x/exp/slog"
)

func setUp() (*os.File, error) {
	cnf, err := config.GetConfig()
	if err != nil {
		fmt.Printf("config not read: %v\n", err)
		return nil, errors.New("not open file")
	}

	f, err := os.OpenFile(cnf.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("not open file: %v\n", err)
		return nil, errors.New("not open file")
	}

	l := slog.New(slog.HandlerOptions{
		AddSource: true,
		Level:     getLevel(cnf.LogLevel),
	}.NewTextHandler(f))

	logger = l
	return f, nil
}

// 日付でlogファイルを作成
// 古いログは削除していく

func getLevel(level string) slog.Leveler {
	switch level {
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

var (
	logger *slog.Logger = nil
)

func Debug(msg string, args ...any) {
	f, e := setUp()
	if e != nil {
		fmt.Printf("failed logger: %+v", e)
		os.Exit(1)
	}
	logger.Debug(msg, args...)
	defer f.Close()
}

func Info(msg string, args ...any) {
	f, e := setUp()
	if e != nil {
		fmt.Printf("failed logger: %+v", e)
		os.Exit(1)
	}
	logger.Info(msg, args...)
	defer f.Close()
}

func Warn(msg string, args ...any) {
	f, e := setUp()
	if e != nil {
		fmt.Printf("failed logger: %+v", e)
		os.Exit(1)
	}
	logger.Warn(msg, args...)
	defer f.Close()
}

func Error(msg string, err error, args ...any) {
	f, e := setUp()
	if e != nil {
		fmt.Printf("failed logger: %+v", e)
		os.Exit(1)
	}
	logger.Error(msg, err, args...)
	defer f.Close()
}
