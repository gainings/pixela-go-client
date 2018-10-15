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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type CreateGraphResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	cgres := CreateGraphResponse{}
	err = json.Unmarshal(body, &cgres)
	if err != nil {
		return err
	}
	if !cgres.IsSuccess {
		return errors.New(cgres.Message)
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type ListGraphResponse struct {
		Graphs []GraphInfo `json:"graphs"`
	}
	lgres := ListGraphResponse{}
	err = json.Unmarshal(body, &lgres)
	if err != nil {
		return nil, err
	}

	return lgres.Graphs, nil
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type UpdateGraphResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	ugres := UpdateGraphResponse{}
	err = json.Unmarshal(body, &ugres)
	if err != nil {
		return err
	}
	if !ugres.IsSuccess {
		return errors.New(ugres.Message)
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	type DeleteGraphResponse struct {
		Message   string `json:"message"`
		IsSuccess bool   `json:"isSuccess"`
	}
	dgres := DeleteGraphResponse{}
	err = json.Unmarshal(body, &dgres)
	if err != nil {
		return err
	}
	if !dgres.IsSuccess {
		return errors.New(dgres.Message)
	}
	return nil
}
