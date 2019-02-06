package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (cloud *Cloud) ServerLogin(jwt string) (string, error) {
	userID, login, name, err := tools.VerifyJWT(jwt, cloud.ServerKey)

	cloud.UserID = userID
	cloud.Login = login
	cloud.Name = name
	cloud.Token = jwt

	if err != nil {
		return "", err
	}

	return name, nil
}

func (cloud *Cloud) doRequest(method string, path string, query map[string]string, body string) (string, error) {
	rel := &url.URL{Path: path}
	url := cloud.BaseURL.ResolveReference(rel)

	var jsonStr = []byte(body)

	version := cloud.Cbox.Version
	if version == "development" {
		version = "0.0.0"
	}

	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+cloud.Token)
	req.Header.Set("cbox-version", version)

	if len(query) != 0 {
		q := req.URL.Query()
		for param, value := range query {
			q.Add(param, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	if cloud.Environment == "dev" {
		strReq, _ := httputil.DumpRequest(req, true)
		tty.Debug(fmt.Sprintf("---\n\n%s\n\n~~~\n", string(strReq)))
	}

	resp, err := cloud.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("rest: could not read response body: %v", err)
	}
	bodyString := string(bodyBytes)

	if cloud.Environment == "dev" {
		tty.Debug(fmt.Sprintf("%s\n\n---\n", bodyString))
	}

	if resp.StatusCode == http.StatusOK {
		return bodyString, nil
	} else if resp.StatusCode == http.StatusNotAcceptable {
		return "", fmt.Errorf("rest: client version not supported by server: %s\n%s", req.Header.Get("cbox-version"), tty.ColorRed(bodyString))
	}

	return "", fmt.Errorf("rest: request failed with '%s' (code: %d):\n%s", resp.Status, resp.StatusCode, tty.ColorRed(bodyString))
}

func (cloud *Cloud) SpacePublish(space *Space) error {

	jsonSpace, err := json.Marshal(space)
	if err != nil {
		return fmt.Errorf("cloud: could not stringify object: %v", err)
	}

	_, err = cloud.doRequest("POST", "/v1/spaces", nil, string(jsonSpace))

	return err
}

func (cloud *Cloud) SpaceUnpublish(selector *Selector) error {

	query := make(map[string]string)

	query["selector"] = selector.String()

	_, err := cloud.doRequest("DELETE", "/v1/spaces", query, "")

	return err
}

// This method retrieves details about an space from the cloud, but not its entries
func (cloud *Cloud) SpaceFind(selector *Selector) (*Space, error) {

	query := make(map[string]string)

	query["selector"] = selector.String()

	response, err := cloud.doRequest("GET", "/v1/spaces", query, "")
	if err != nil {
		return nil, err
	}

	var space Space
	err = json.Unmarshal([]byte(response), &space)
	if err != nil {
		return nil, fmt.Errorf("cloud: could not parse response: %v", err)
	}

	space.Selector, _ = ParseSelector(space.ID)

	return &space, err
}

func (cloud *Cloud) CommandList(selector *Selector) ([]*Command, error) {

	query := make(map[string]string)
	query["selector"] = selector.String()

	response, err := cloud.doRequest("GET", "/v1/commands", query, "")
	if err != nil {
		return nil, err
	}

	tty.Print("\n")
	var commands []*Command
	err = json.Unmarshal([]byte(response), &commands)
	if err != nil {
		return nil, fmt.Errorf("cloud: could not parse response: %v", err)
	}

	for _, command := range commands {
		selector, _ := ParseSelector(command.ID)
		command.Selector = selector
	}

	return commands, nil
}
