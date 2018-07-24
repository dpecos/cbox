package core

import (
	"log"
	"path"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	homedir "github.com/mitchellh/go-homedir"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

const CBOX_PATH = ".cbox"

var (
	basePath string
)

func resolveInCboxDir(content string) string {
	cboxBasePath := basePath
	if cboxBasePath == "" {
		var err error
		cboxBasePath, err = homedir.Dir()
		if err != nil {
			log.Fatalf("respository: could not get HOME: %v", err)
		}
	}
	return path.Join(cboxBasePath, CBOX_PATH, content)
}

func CheckCboxDir(path string) string {
	basePath = path

	cboxPath := resolveInCboxDir("")
	tools.CreateDirectoryIfNotExists(cboxPath)

	configFile := resolveInCboxDir("config.yml")
	tools.CreateFileIfNotExists(configFile)

	spacesPath := resolveInCboxDir("spaces")
	if tools.CreateDirectoryIfNotExists(spacesPath) {
		id, _ := uuid.NewV4()
		defaultSpace := models.Space{
			Label:       DEFAULT_SPACE_ID,
			Description: DEFAULT_SPACE_DESCRIPTION,
		}
		defaultSpace.ID = id

		cbox := LoadCbox(path)
		cbox.SpaceAdd(&defaultSpace)
		PersistCbox(cbox)
	}

	return cboxPath
}

func LoadCbox(path string) *models.CBox {
	basePath = path

	cbox := models.CBox{
		Spaces: []models.Space{},
	}

	spaces := SpaceListFiles()

	for _, space := range spaces {
		cbox.SpaceAdd(space)
	}
	return &cbox
}

func PersistCbox(cbox *models.CBox) {
	for _, space := range cbox.Spaces {
		SpaceStoreFile(&space)
	}
}

func InitCBox(path string) {
	cboxPath := CheckCboxDir(path)

	viper.AddConfigPath(cboxPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	defaultSettings()
}

func defaultSettings() {
	viper.SetDefault("cbox.default-space", "default")
}
