package core

import (
	"log"

	"github.com/dplabs/cbox/src/repository"
	"github.com/dplabs/cbox/src/tools/console"

	"github.com/dplabs/cbox/src/models"
)

const (
	defaultSpaceID          = "default"
	defaultSpaceDescription = "Default space to store commands"
)

var (
	Version = "development"
	Build   = "-"

	cboxWorkDirectory string
	repo              *repository.Repository
)

func Load(path string) *models.CBox {

	repo = repository.InitRepository(path)

	Env = repo.LoadSettings(Env)
	spaces, isNewRepository := repo.LoadSpaces()

	cbox := &models.CBox{
		Spaces:  []*models.Space{},
		Version: Version,
		Build:   Build,
	}

	if isNewRepository {
		createDefaultSpace(cbox)
		console.PrintInfo("Initial setup: an empty space has been created\n")
	}

	for _, space := range spaces {
		err := cbox.SpaceCreate(space)
		if err != nil {
			log.Fatalf("load: could not create space: %v", err)
		}
	}
	return cbox
}

func createDefaultSpace(cbox *models.CBox) {
	defaultSpace := models.Space{
		Label:       defaultSpaceID,
		Description: defaultSpaceDescription,
	}
	defaultSpace.Selector = models.NewSelector(models.TypeNone, "", defaultSpace.Label, "")

	err := cbox.SpaceCreate(&defaultSpace)
	if err != nil {
		log.Fatalf("init: could not create space: %v", err)
	}

	Save(cbox)
}

func Save(cbox *models.CBox) {
	for _, space := range cbox.Spaces {
		repo.Persist(space)
	}
}

func DeleteSpaceFile(selector *models.Selector) {
	repo.Delete(selector)
}
