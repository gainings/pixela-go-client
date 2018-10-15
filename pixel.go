package pixela

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//DrawPixel draw specific pixel
func (c Client) DrawPixel(username, token string, id, date, quantity string) error {
	type RequestBody struct {
		Date     string `json:"date"`
		Quantity string `json:"quantity"`
	}
	rb := RequestBody{
		Date:     date,
		Quantity: quantity,
	}
	rbJSON, err := json.Marshal(rb)
	if err != nil {
		return err
	}

	u := fmt.Sprintf("%s/users/%s/graphs/%s", c.URL, username, id)
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(rbJSON))
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
