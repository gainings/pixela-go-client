package pixela

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//RegisterUser is register new user
func (c Client) RegisterUser(username, token string, agree, notMinor string) error {
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

//UpdateToken update user's token
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

	type UpdateTokenResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	utres := UpdateTokenResponse{}
	err = json.Unmarshal(body, &utres)
	if err != nil {
		return err
	}
	if !utres.IsSuccess {
		return errors.New(utres.Message)
	}
	return nil
}

//DeleteUser delete registered user
func (c Client) DeleteUser(username, token string) error {
	req, err := http.NewRequest("DELETE", c.URL+"/users/"+username, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-USER-TOKEN", token)
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

	type DeleteUserResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	dures := DeleteUserResponse{}
	err = json.Unmarshal(body, &dures)
	if err != nil {
		return err
	}
	if !dures.IsSuccess {
		return errors.New(dures.Message)
	}
	return nil
}
