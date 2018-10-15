package pixela

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RegisterPixel Register specific pixel
func (c Client) RegisterPixel(username, token string, id, date, quantity string) error {
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

//GetPixelQuantity get quantity from specific pixel
func (c Client) GetPixelQuantity(username, token string, id, date string) (float64, error) {
	u := fmt.Sprintf("%s/users/%s/graphs/%s/%s", c.URL, username, id, date)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("X-USER-TOKEN", token)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, errors.New("return status code: " + res.Status)
	}
	bodyJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	type ResponseBody struct {
		Quantity float64
	}
	body := ResponseBody{}
	err = json.Unmarshal(bodyJSON, &body)
	if err != nil {
		return 0, err
	}
	return body.Quantity, nil
}

//UpdatePixelQuantity update quantity for already registered pixel
func (c Client) UpdatePixelQuantity(username, token string, id, date, quantity string) error {
	type RequestBody struct {
		Quantity string `json:"quantity"`
	}
	rb := RequestBody{
		Quantity: quantity,
	}
	rbJSON, err := json.Marshal(rb)
	if err != nil {
		return err
	}

	u := fmt.Sprintf("%s/users/%s/graphs/%s/%s", c.URL, username, id, date)
	req, err := http.NewRequest("PUT", u, bytes.NewBuffer(rbJSON))
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

//IncrementPixelQuantity increment quantity of the Day
func (c Client) IncrementPixelQuantity(username, token string, id string) error {
	u := fmt.Sprintf("%s/users/%s/graphs/%s/increment", c.URL, username, id)
	req, err := http.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-USER-TOKEN", token)
	req.Header.Set("Content-Length", "0")
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

//DecrementPixelQuantity decrement quantity of the Day
func (c Client) DecrementPixelQuantity(username, token string, id string) error {
	u := fmt.Sprintf("%s/users/%s/graphs/%s/decrement", c.URL, username, id)
	req, err := http.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-USER-TOKEN", token)
	req.Header.Set("Content-Length", "0")
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

//DeletePixelQuantity dlete quantity for already registered pixel
func (c Client) DeletePixelQuantity(username, token string, id, date, quantity string) error {
	u := fmt.Sprintf("%s/users/%s/graphs/%s/%s", c.URL, username, id, date)
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
