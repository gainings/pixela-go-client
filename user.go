package pixela

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RegisterUser is register new user
func (c Client) RegisterUser(agree, notMinor string) error {
	if c.UserName == "" || c.Token == "" {
		return errors.New("Plz set user information in Client")
	}

	type RequestBody struct {
		Username string `json:"username"`
		Token    string `json:"token"`
		Agree    string `json:"agreeTermsOfService"`
		NotMinor string `json:"notMinor"`
	}
	rb := RequestBody{
		Username: c.UserName,
		Token:    c.Token,
		Agree:    agree,
		NotMinor: notMinor,
	}
	rbJSON, err := json.Marshal(rb)
	if err != nil {
		return err
	}

	u := fmt.Sprintf("%s/users", c.URL)
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(rbJSON))
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
func (c *Client) UpdateToken(newToken string) error {
	if c.UserName == "" || c.Token == "" {
		return errors.New("Plz set user information in Client")
	}

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

	u := fmt.Sprintf("%s/users/%s", c.URL, c.UserName)
	req, err := http.NewRequest("PUT", u, bytes.NewBuffer(rbJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-USER-TOKEN", c.Token)
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
	c.Token = newToken
	return nil
}

//DeleteUser delete registered user
func (c Client) DeleteUser() error {
	if c.UserName == "" || c.Token == "" {
		return errors.New("Plz set user information in Client")
	}

	u := fmt.Sprintf("%s/users/%s", c.URL, c.UserName)
	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-USER-TOKEN", c.Token)
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
