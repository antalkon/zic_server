package config

import (
	"github.com/spf13/viper"
)

// LoadDataConfig загружает конфигурацию из файла /configs/config.yaml
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

// SaveDataConfig сохраняет изменения конфигурации в файл
func SaveDataConfig() error {
	// Сохраняем изменения в файл
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
