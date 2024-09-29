package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// InitLogger инициализирует логгер с ротацией логов
func InitLogger() {
	log = logrus.New()
	os.MkdirAll("logs", os.ModePerm) // Создает директорию logs, если она не существует

	// Настройка ротации логов
	log.SetOutput(&lumberjack.Logger{
		Filename:   filepath.Join("logs", "server.log"),
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30,   // дни
		Compress:   true, // сжимать старые логи
	})

	// Установка формата логов
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Установка уровня логирования
	log.SetLevel(logrus.InfoLevel)
}

// LogError записывает ошибку в лог
func LogError(err error) {
	log.Errorf("Ошибка: %v", err)
}

// LogInfo записывает информационное сообщение в лог
func LogInfo(message string) {
	log.Info(message)
}
