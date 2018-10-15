package pixela

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GraphInfo is need info to create graph
type GraphInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	UnitType string `json:"type"`
	Color    string `json:"color"`
}

//Validate check info need create Graph
func (gi GraphInfo) Validate() error {
	if gi.ID == "" {
		return errors.New("id is empty. plz input")
	}
	if gi.Name == "" {
		return errors.New("name is empty. plz input")
	}
	if gi.Unit == "" {
		return errors.New("unit name is empty. plz input")
	}

	switch gi.UnitType {
	case "int", "float":
	default:
		return errors.New("invalid unit type. expected int or float")
	}

	switch gi.Color {
	case "shibafu", "momiji", "sora", "ichou", "ajisai", "kuro":
	default:
		return errors.New("invalid color")
	}
	return nil
}

//CreateGraph create new graph
func (c Client) CreateGraph(username, token string, gi GraphInfo) error {
	giJSON, err := json.Marshal(gi)
	if err != nil {
		return err
	}

	u := fmt.Sprintf("%s/users/%s/graphs", c.URL, username)
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(giJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-USER-TOKEN", token)
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

//ListGraph return User's graph info list
func (c Client) ListGraph(username, token string) ([]GraphInfo, error) {
	u := fmt.Sprintf("%s/users/%s/graphs", c.URL, username)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-USER-TOKEN", token)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("return status code: " + res.Status)
	}
	bodyJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type ResponseBody struct {
		Graphs []GraphInfo `json:"graphs"`
	}
	body := ResponseBody{}
	err = json.Unmarshal(bodyJSON, &body)
	if err != nil {
		return nil, err
	}

	return body.Graphs, nil
}

//GetGraph get specific graphs's svg
func (c Client) GetGraph(username, token, id, date string) (string, error) {
	u := fmt.Sprintf("%s/users/%s/graphs/%s", c.URL, username, id)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-USER-TOKEN", token)
	if date != "" {
		q := req.URL.Query()
		q.Add("date", date)
		req.URL.RawQuery = q.Encode()
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("return status code: " + res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//UpdateGraph update specific graphs's information
func (c Client) UpdateGraph(username, token string, gi GraphInfo) error {
	giJSON, err := json.Marshal(gi)
	if err != nil {
		return err
	}

	u := fmt.Sprintf("%s/users/%s/graphs/%s", c.URL, username, gi.ID)
	req, err := http.NewRequest("PUT", u, bytes.NewBuffer(giJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-USER-TOKEN", token)
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

//DeleteGraph delete specific graphs
func (c Client) DeleteGraph(username, token, id string) error {
	u := fmt.Sprintf("%s/users/%s/graphs/%s", c.URL, username, id)
	req, err := http.NewRequest("DELETE", u, nil)
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
