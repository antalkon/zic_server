package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// LoadEnvConfig загружает конфигурацию из системных переменных окружения и .env файла.
func LoadEnvConfig() error {
	viper.SetConfigName(".env") // Имя конфигурационного файла
	viper.SetConfigType("env")  // Формат файла
	viper.AddConfigPath(".")    // Поиск в текущей директории
	viper.AddConfigPath("/app") // Поиск в директории приложения
	viper.AutomaticEnv()        // Подключение системных переменных окружения

	// Читаем файл конфигурации, если он существует
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Warning: .env file not found, using environment variables only.")
		} else {
			log.Printf("Error reading .env file: %v", err)
			return err
		}
	} else {
		log.Printf("Loaded configuration from .env file: %s", viper.ConfigFileUsed())
	}

	return nil
}

// GetEnvValue возвращает значение переменной из конфигурации.
// Если переменная не найдена, вызывает панику.
func GetEnvValue(key string) string {
	// Сначала пробуем получить значение из окружения
	value := viper.GetString(key)

	// Если значение пустое, логируем и вызываем панику
	if value == "" {
		log.Printf("Error: Key '%s' is not set or empty", key)
		panic("Missing required environment variable: " + key)
	}

	log.Printf("Fetching key '%s': value='%s'", key, value)
	return value
}

// DebugEnvFile проверяет наличие файла .env в текущей директории.
func DebugEnvFile() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		log.Println(".env file does not exist in the current directory")
	} else if err != nil {
		log.Printf("Error checking .env file: %v", err)
	} else {
		log.Println(".env file exists in the current directory")
	}
}

// IsEnvVariableSet проверяет, установлена ли переменная окружения.
func IsEnvVariableSet(key string) bool {
	value, exists := os.LookupEnv(key)
	if exists {
		log.Printf("Environment variable '%s' is set: value='%s'", key, value)
	} else {
		log.Printf("Environment variable '%s' is not set", key)
	}
	return exists
}
