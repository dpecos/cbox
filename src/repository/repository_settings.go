package repository

import (
	"log"
	"os"

	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/viper"
)

const (
	configFileName = "config"
	configFilePath = configFileName + ".yml"
)

func (repo *Repository) GetEnv() string {
	var env string

	if viper.IsSet("cbox.environment") {
		env = viper.GetString("cbox.environment")
	} else {
		env = "prod"
	}

	if os.Getenv("CBOX_ENV") != "" {
		env = os.Getenv("CBOX_ENV")
		if env != "test" && env != "prod" {
			log.Fatalf("unknown env value '%s' (from CBOX_ENV)", env)
		}
	}

	return env
}

func (repo *Repository) loadSettings() {

	env := repo.GetEnv()

	configFile := repo.resolve(configFilePath)
	tools.CreateFileIfNotExists(configFile)

	viper.AddConfigPath(repo.Path)
	viper.SetConfigName(configFileName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	defaultSettings(env)
}

func defaultSettings(env string) {
	viper.SetDefault("cbox.default-space", "default")
	viper.SetDefault("cbox.environment", env)
}
