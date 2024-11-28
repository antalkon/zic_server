package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// LoadEnvConfig загружает конфигурацию из файла .env
func LoadEnvConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Читаем файл конфигурации
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading .env file: %v", err)
		return err
	}

	// Логируем содержимое
	log.Printf("Loaded config from .env file: %s", viper.ConfigFileUsed())

	// НЕ используем viper.AutomaticEnv(), чтобы избежать конфликта
	return nil
}

// GetEnvValue возвращает значение переменной из .env
func GetEnvValue(key string) string {
	value := viper.GetString(key)
	log.Printf("Fetching key '%s': value='%s'", key, value)
	return value
}

// DebugEnvFile проверяет существование файла .env
func DebugEnvFile() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		log.Println(".env file does not exist in the current directory")
	} else if err != nil {
		log.Printf("Error checking .env file: %v", err)
	} else {
		log.Println(".env file exists")
	}
}
