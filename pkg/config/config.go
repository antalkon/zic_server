package config

import (
	"github.com/spf13/viper"
)

// LoadServerConfig загружает конфигурацию сервера из файла /configs/server.yaml
func LoadDataConfig() error {
	viper.SetConfigName("config")    // имя файла без расширения
	viper.SetConfigType("yaml")      // тип файла
	viper.AddConfigPath("./configs") // путь к файлу (директория configs)

	// Читаем файл конфигурации
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
