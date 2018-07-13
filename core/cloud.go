package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/viper"
)

const (
	SERVER_URL_DEV = "https://api.dev.cbox.dplabs.io"
	SETTINGS_USER  = "cloud.auth.user"
	SETTINGS_JWT   = "cloud.auth.jwt"
)

type Cloud struct {
	User       string
	token      string
	baseURL    *url.URL
	httpClient *http.Client
}

func CloudLogin(jwt string) (string, error) {
	user, err := tools.VerifyJWT(jwt)
	if err != nil {
		return "", err
	}
	viper.Set(SETTINGS_USER, user)
	viper.Set(SETTINGS_JWT, jwt)
	return user, nil
}

func CloudClient() (*Cloud, error) {
	if !viper.IsSet(SETTINGS_USER) || !viper.IsSet(SETTINGS_JWT) {
		return nil, fmt.Errorf("cloud: user not authenticated")
	}

	url, err := url.Parse(SERVER_URL_DEV)
	if err != nil {
		return nil, fmt.Errorf("cloud: could not parse server's URL: %v", err)
	}

	cloud := Cloud{User: viper.GetString(SETTINGS_USER), token: viper.GetString(SETTINGS_JWT), baseURL: url, httpClient: http.DefaultClient}

	return &cloud, nil
}

func (cloud *Cloud) doRequest(method string, path string, body string) (string, error) {
	rel := &url.URL{Path: path}
	u := cloud.baseURL.ResolveReference(rel)

	var jsonStr = []byte(body)

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+cloud.token)

	resp, err := cloud.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("rest: could not read response body: %v", err)
	}
	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		return bodyString, nil
	}

	return "", fmt.Errorf("rest: request failed with '%s' (code: %d): %s", resp.Status, resp.StatusCode, bodyString)
}

func (cloud *Cloud) PublishSpace(space *models.Space) error {

	jsonSpace, err := json.Marshal(space)
	if err != nil {
		return fmt.Errorf("cloud: publish space: could not stringify object: %v", err)
	}

	_, err = cloud.doRequest("POST", "/v1/spaces", string(jsonSpace))

	return err
}
