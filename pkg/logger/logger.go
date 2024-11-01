package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// Logger – интерфейс для логирования, позволяющий легко заменить логгер в тестах или разных средах
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// CustomLogger включает Logrus и поддержку ротации логов через Lumberjack
type CustomLogger struct {
	log *logrus.Logger
}

// Config структура для параметров конфигурации логгера
type Config struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Level      logrus.Level
}

// InitLogger инициализирует логгер с ротацией логов и возвращает его
func InitLogger(config Config) (*CustomLogger, error) {
	logger := logrus.New()

	// Создаем директорию для логов, если она не существует
	logDir := filepath.Dir(config.Filename)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	// Настройка ротации логов
	logger.SetOutput(&lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	})

	// Устанавливаем формат и уровень логирования
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetLevel(config.Level)

	return &CustomLogger{log: logger}, nil
}

// Метод Info для логгирования информационных сообщений
func (c *CustomLogger) Info(args ...interface{}) {
	c.log.Info(args...)
}

// Метод Warn для логгирования предупреждений
func (c *CustomLogger) Warn(args ...interface{}) {
	c.log.Warn(args...)
}

// Метод Error для логгирования ошибок
func (c *CustomLogger) Error(args ...interface{}) {
	c.log.Error(args...)
}

// Метод Fatal для логгирования критических ошибок и завершения программы
func (c *CustomLogger) Fatal(args ...interface{}) {
	c.log.Fatal(args...)
}
