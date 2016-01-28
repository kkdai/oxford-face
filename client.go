package face

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	StreamContent string = "application/octet-stream"
	JsonContent   string = "application/json"
)

type Client struct {
	key string
}

// New oxford client based on key
func NewClient(key string) *Client {
	c := new(Client)
	c.key = key
	return c
}

// Connect with API url and data, return response byte or error if http.Status is not OK
func (c *Client) Connect(mode string, url string, data *bytes.Buffer, useJson bool) ([]byte, error) {
	client := &http.Client{}
	r, _ := http.NewRequest(mode, url, data)

	if useJson {
		r.Header.Add("Content-Type", JsonContent)
	} else {
		r.Header.Add("Content-Type", StreamContent)
	}

	r.Header.Add("Ocp-Apim-Subscription-Key", c.key)

	resp, err := client.Do(r)
	if err != nil {
		log.Println("er:", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("er:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error happen! body:", string(body))
		return body, errors.New("Error on:" + string(body))
	}
	return body, nil
}
