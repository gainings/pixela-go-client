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
