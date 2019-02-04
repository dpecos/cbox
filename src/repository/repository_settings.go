package repository

import (
	"log"

	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/viper"
)

const (
	configFileName = "config"
	configFilePath = configFileName + ".yml"
)

func (repo *Repository) LoadSettings(env string) string {
	configFile := repo.resolve(configFilePath)
	tools.CreateFileIfNotExists(configFile)

	viper.AddConfigPath(repo.Path)
	viper.SetConfigName(configFileName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	env = defaultSettings(env)

	return env
}

func defaultSettings(env string) string {
	viper.SetDefault("cbox.default-space", "default")
	viper.SetDefault("cbox.environment", env)

	if viper.IsSet("cbox.environment") {
		env = viper.GetString("cbox.environment")
	}

	return env
}
