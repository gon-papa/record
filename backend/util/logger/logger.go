package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gon-papa/record/config"
	"golang.org/x/exp/slog"
)

// ログの初期化
func setUp() (*os.File, error) {
	cnf, err := config.GetConfig()
	if err != nil {
		fmt.Printf("config not read: %v\n", err)
		return nil, errors.New("not open file")
	}
	LogRotation(cnf)

	// ファイルに含める日付を生成
	t := time.Now()
	date := t.Format("2006-01-02")
	var fileName string = cnf.Log + date + ".log"

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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

// ログのローテーション設定
func LogRotation(cnf *config.Config) {
	// ディレクトリ内を全ファイルをスキャン
	files, err := filepath.Glob(filepath.Join(cnf.Log, "*"))
	if err != nil {
		fmt.Printf("not found directory%v\n", err)
	}

	if len(files) >= cnf.LogRotation {
		sort.Strings(files) // ソート(ファイル先頭が日付)
		for i := 0; i < len(files)-cnf.LogRotation; i++ {
			if err := os.Remove(files[i]); err != nil {
				fmt.Println(err)
			}
		}
	}
}

// configから文字列をとみとりslogのログレベルを返す
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
