package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dpecos/cbox/tools/console"
	"github.com/gofrs/uuid"

	"github.com/dpecos/cbox/pkg/models"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/viper"
)

var (
	cloudSettingsUserID    string
	cloudSettingsUserLogin string
	cloudSettingsUserName  string
	cloudSettingsJWT       string
)

func readCloudConfig() {
	var env = Env
	if env != "" && env != "prod" {
		env = "_" + env
	}
	if env == "prod" {
		env = ""
	}
	cloudSettingsUserID = fmt.Sprintf("cloud%s.auth.user.id", env)
	cloudSettingsUserLogin = fmt.Sprintf("cloud%s.auth.user.login", env)
	cloudSettingsUserName = fmt.Sprintf("cloud%s.auth.user.name", env)
	cloudSettingsJWT = fmt.Sprintf("cloud%s.auth.jwt", env)
}

type Cloud struct {
	UserID     string
	Login      string
	Name       string
	token      string
	baseURL    *url.URL
	httpClient *http.Client
}

func CloudLogin(jwt string) (string, string, string, error) {
	readCloudConfig()

	userID, login, name, err := tools.VerifyJWT(jwt)
	if err != nil {
		return "", "", "", err
	}
	viper.Set(cloudSettingsUserID, userID)
	viper.Set(cloudSettingsUserLogin, login)
	viper.Set(cloudSettingsUserName, name)
	viper.Set(cloudSettingsJWT, jwt)
	return userID, login, name, nil
}

func CloudLogout() {
	readCloudConfig()

	viper.Set(cloudSettingsUserID, "")
	viper.Set(cloudSettingsUserLogin, "")
	viper.Set(cloudSettingsUserName, "")
	viper.Set(cloudSettingsJWT, "")
}

func CloudClient() (*Cloud, error) {
	readCloudConfig()

	url, err := url.Parse(CloudURL())
	if err != nil {
		return nil, fmt.Errorf("cloud: could not parse server's URL: %v", err)
	}

	if !viper.IsSet(cloudSettingsUserID) || !viper.IsSet(cloudSettingsUserLogin) || !viper.IsSet(cloudSettingsUserName) || !viper.IsSet(cloudSettingsJWT) {
		return nil, fmt.Errorf("cloud: user not authenticated")
	}

	cloud := Cloud{
		UserID:     viper.GetString(cloudSettingsUserID),
		Login:      viper.GetString(cloudSettingsUserLogin),
		Name:       viper.GetString(cloudSettingsUserName),
		token:      viper.GetString(cloudSettingsJWT),
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