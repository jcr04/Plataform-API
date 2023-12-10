package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // Nome do arquivo de configuração (sem a extensão)
	viper.SetConfigType("yaml")   // Extensão do arquivo de configuração
	viper.AddConfigPath(".")      // Procura o arquivo de configuração no diretório atual

	viper.SetDefault("DatabaseURL", "postgres://postgres:12345@localhost:5432/Plataform")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
