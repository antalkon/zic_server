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
	// Создает директорию logs, если она не существует
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("Failed to create logs directory: %v", err) // Fatal: завершает программу
	}

	// Настройка ротации логов
	log.SetOutput(&lumberjack.Logger{
		Filename:   filepath.Join("logs", "server.log"),
		MaxSize:    10,   // Размер в MB до ротации
		MaxBackups: 5,    // Максимальное количество файлов бэкапа
		MaxAge:     30,   // Количество дней хранения логов
		Compress:   true, // Сжимать ли старые логи
	})

	// Устанавливаем формат логов
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Полный формат времени (с датой)
	})

	// Устанавливаем уровень логирования (можно изменять по необходимости)
	log.SetLevel(logrus.InfoLevel)
}

// LogError записывает сообщение об ошибке в лог
func LogError(err error) {
	log.Errorf("Ошибка: %v", err)
}

// LogInfo записывает информационное сообщение в лог
func LogInfo(message string) {
	log.Info(message)
}

// LogWarning записывает предупреждение в лог
func LogWarning(message string) {
	log.Warn(message)
}

// LogFatal записывает критическую ошибку и завершает выполнение программы
// LogFatal записывает критическую ошибку и завершает выполнение программы
func LogFatal(message string, err error) {
	log.Fatalf("%s: %v", message, err) // Изменяем формат вызова, чтобы обрабатывать строку и ошибку
}
