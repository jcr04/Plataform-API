package config

import (
	"github.com/spf13/viper"
)

// Config representa as configurações da aplicação.
type Config struct {
	DatabaseURL string
}

// LoadConfig carrega as configurações da aplicação.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // Nome do arquivo de configuração (sem a extensão)
	viper.SetConfigType("yaml")   // Extensão do arquivo de configuração
	viper.AddConfigPath(".")      // Procura o arquivo de configuração no diretório atual

	// Configuração padrão
	viper.SetDefault("DatabaseURL", "postgres://postgres:12345@localhost:5432/Plataform")

	// Sobrescreve com variáveis de ambiente, se existirem
	viper.AutomaticEnv()

	// Lê o arquivo de configuração
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
