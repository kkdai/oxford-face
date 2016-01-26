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

func NewClient(key string) *Client {
	c := new(Client)
	c.key = key
	return c
}

func (c *Client) Connect(url string, data *bytes.Buffer, useJson bool) ([]byte, error) {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", url, data)

	if useJson {
		r.Header.Add("Content-Type", JsonContent)
	} else {
		r.Header.Add("Content-Type", StreamContent)
	}

	r.Header.Add("Ocp-Apim-Subscription-Key", c.key)

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println("Error happen! body:", string(body))
		return body, errors.New("Error on:" + string(body))
	}
	return body, nil
}
