package core

import (
	"log"
	"net/http"
	"net/url"

	"github.com/dplabs/cbox/src/models"
)

var (
	Env = "dev"
)

func CloudClient(cbox *models.CBox) *models.Cloud {

	serverURL, serverKey, userID, login, name, token := repo.LoadCloudSettings(Env)

	baseUrl, err := url.Parse(serverURL)
	if err != nil {
		log.Fatalf("cloud: could not parse server's URL: %v", err)
	}

	cloud := models.Cloud{
		Environment: Env,
		ServerKey:   serverKey,
		UserID:      userID,
		Login:       login,
		Name:        name,
		Token:       token,
		URL:         serverURL,
		BaseURL:     baseUrl,
		HttpClient:  http.DefaultClient,
		Cbox:        cbox,
	}

	return &cloud
}

func StoreCloudSettings(cloud *models.Cloud) {
	repo.StoreCloudSettings(cloud)
}

func DeleteCloudSettings() {
	repo.DeleteCloudSettings()
}
