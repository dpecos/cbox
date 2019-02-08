package repository

import (
	"log"
	"path"

	"github.com/dplabs/cbox/src/tools"
	homedir "github.com/mitchellh/go-homedir"
)

type Repository struct {
	Path string
}

const (
	cboxDir = ".cbox"
)

func InitRepository(repoPath string) *Repository {

	if repoPath == "" {
		var err error
		repoPath, err = homedir.Dir()
		if err != nil {
			log.Fatalf("init: could not get HOME: %v", err)
		}
	}

	repoPath = path.Join(repoPath, cboxDir)
	tools.CreateDirectoryIfNotExists(repoPath)

	repo := Repository{
		Path: repoPath,
	}

	return &repo
}

func (repo *Repository) resolve(paths ...string) string {
	return path.Join(repo.Path, path.Join(paths...))
}
