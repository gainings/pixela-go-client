package pixela

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://pixe.la/v1"

//Client is Client for access to pixe.la service
type Client struct {
	HTTPClient http.Client
	URL        string
}

//NewClient is return Cilent
func NewClient() *Client {
	return &Client{
		HTTPClient: http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
		URL: baseURL,
	}
}

//Register is register new user
func (c Client) Register(username, token string, agree, notMinor string) error {
	type registerInfo struct {
		Username string `json:"username"`
		Token    string `json:"token"`
		Agree    string `json:"agreeTermsOfService"`
		NotMinor string `json:"notMinor"`
	}
	ri := registerInfo{
		Username: username,
		Token:    token,
		Agree:    agree,
		NotMinor: notMinor,
	}
	riJSON, err := json.Marshal(ri)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL+"/users", bytes.NewBuffer(riJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("return status code: " + res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type RegisterResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	rres := RegisterResponse{}
	err = json.Unmarshal(body, &rres)
	if err != nil {
		return err
	}
	if !rres.IsSuccess {
		return errors.New(rres.Message)
	}
	return nil
}

//UpdateToken is update user's token
func (c Client) UpdateToken(username, oldToken, newToken string) error {
	type updateInfo struct {
		Token string `json:"newToken"`
	}
	ui := updateInfo{
		Token: newToken,
	}
	uiJSON, err := json.Marshal(ui)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.URL+"/users/"+username, bytes.NewBuffer(uiJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-USER-TOKEN", oldToken)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("return status code: " + res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type RegisterResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	rres := RegisterResponse{}
	err = json.Unmarshal(body, &rres)
	if err != nil {
		return err
	}
	if !rres.IsSuccess {
		return errors.New(rres.Message)
	}
	return nil
}
