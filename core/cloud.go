package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dpecos/cbox/tools/console"
	uuid "github.com/satori/go.uuid"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/viper"
)

const (
	SETTINGS_USER_ID    = "cloud.auth.user.id"
	SETTINGS_USER_LOGIN = "cloud.auth.user.login"
	SETTINGS_USER_NAME  = "cloud.auth.user.name"
	SETTINGS_JWT        = "cloud.auth.jwt"
)

type Cloud struct {
	UserID     string
	Login      string
	Name       string
	token      string
	baseURL    *url.URL
	httpClient *http.Client
}

func CloudLogin(jwt string) (string, string, string, error) {
	userID, login, name, err := tools.VerifyJWT(jwt)
	if err != nil {
		return "", "", "", err
	}
	viper.Set(SETTINGS_USER_ID, userID)
	viper.Set(SETTINGS_USER_LOGIN, login)
	viper.Set(SETTINGS_USER_NAME, name)
	viper.Set(SETTINGS_JWT, jwt)
	return userID, login, name, nil
}

func CloudLogout() {
	viper.Set(SETTINGS_USER_ID, "")
	viper.Set(SETTINGS_USER_LOGIN, "")
	viper.Set(SETTINGS_USER_NAME, "")
	viper.Set(SETTINGS_JWT, "")
}

func CloudClient() (*Cloud, error) {
	if !viper.IsSet(SETTINGS_USER_ID) || !viper.IsSet(SETTINGS_USER_LOGIN) || !viper.IsSet(SETTINGS_USER_NAME) || !viper.IsSet(SETTINGS_JWT) {
		return nil, fmt.Errorf("cloud: user not authenticated")
	}

	url, err := url.Parse(CloudURL())
	if err != nil {
		return nil, fmt.Errorf("cloud: could not parse server's URL: %v", err)
	}

	cloud := Cloud{
		UserID:     viper.GetString(SETTINGS_USER_ID),
		Login:      viper.GetString(SETTINGS_USER_LOGIN),
		Name:       viper.GetString(SETTINGS_USER_NAME),
		token:      viper.GetString(SETTINGS_JWT),
		baseURL:    url,
		httpClient: http.DefaultClient,
	}

	return &cloud, nil
}

func (cloud *Cloud) doRequest(method string, path string, query map[string]string, body string) (string, error) {
	rel := &url.URL{Path: path}
	url := cloud.baseURL.ResolveReference(rel)

	var jsonStr = []byte(body)

	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+cloud.token)

	if len(query) != 0 {
		q := req.URL.Query()
		for param, value := range query {
			q.Add(param, value)
		}
		req.URL.RawQuery = q.Encode()
	}

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

	return "", fmt.Errorf("rest: request failed with '%s' (code: %d):\n%s", resp.Status, resp.StatusCode, console.ColorRed(bodyString))
}

func (cloud *Cloud) SpacePublish(space *models.Space) error {

	jsonSpace, err := json.Marshal(space)
	if err != nil {
		return fmt.Errorf("could not stringify object: %v", err)
	}

	_, err = cloud.doRequest("POST", "/v1/spaces", nil, string(jsonSpace))

	return err
}

func (cloud *Cloud) SpaceRetrieve(selector *models.Selector, id *uuid.UUID) (*models.Space, error) {

	query := make(map[string]string)

	if selector != nil {
		query["selector"] = selector.String()
	} else {
		query["selector"] = id.String()
	}

	response, err := cloud.doRequest("GET", "/v1/spaces", query, "")
	if err != nil {
		return nil, err
	}

	var space models.Space
	err = json.Unmarshal([]byte(response), &space)
	if err != nil {
		return nil, fmt.Errorf("could not parse response: %v", err)
	}

	return &space, err
}

func (cloud *Cloud) CommandList(selector *models.Selector) ([]models.Command, error) {

	query := make(map[string]string)
	query["selector"] = selector.String()

	response, err := cloud.doRequest("GET", "/v1/commands", query, "")
	if err != nil {
		return nil, err
	}

	var commands []models.Command
	err = json.Unmarshal([]byte(response), &commands)
	if err != nil {
		return nil, fmt.Errorf("could not parse response: %v", err)
	}

	return commands, nil
}
