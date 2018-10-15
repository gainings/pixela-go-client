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
	type RequestBody struct {
		Username string `json:"username"`
		Token    string `json:"token"`
		Agree    string `json:"agreeTermsOfService"`
		NotMinor string `json:"notMinor"`
	}
	rb := RequestBody{
		Username: username,
		Token:    token,
		Agree:    agree,
		NotMinor: notMinor,
	}
	rbJSON, err := json.Marshal(rb)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL+"/users", bytes.NewBuffer(rbJSON))
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
	bodyJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type ResponseBody struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	body := ResponseBody{}
	err = json.Unmarshal(bodyJSON, &body)
	if err != nil {
		return err
	}
	if !body.IsSuccess {
		return errors.New(body.Message)
	}
	return nil
}

//UpdateToken update user's token
func (c Client) UpdateToken(username, oldToken, newToken string) error {
	type RequestBody struct {
		Token string `json:"newToken"`
	}
	rb := RequestBody{
		Token: newToken,
	}
	rbJSON, err := json.Marshal(rb)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.URL+"/users/"+username, bytes.NewBuffer(rbJSON))
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
	bodyJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type ResponseBody struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	body := ResponseBody{}
	err = json.Unmarshal(bodyJSON, &body)
	if err != nil {
		return err
	}
	if !body.IsSuccess {
		return errors.New(body.Message)
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
