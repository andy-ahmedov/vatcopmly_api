package vatcomply

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetLatestRates() (*Response, error) {
	var response Response

	url := "https://api.vatcomply.com/rates"

	err := c.GetRequest(&response, url)

	return &response, err
}

func (c Client) GetBaseRates(base string) (*Response, error) {
	var response Response

	url := fmt.Sprintf("https://api.vatcomply.com/rates?base=%s", base)

	err := c.GetRequest(&response, url)

	return &response, err
}

func (c Client) GetDateeRates(date string) (*Response, error) {
	var response Response

	url := fmt.Sprintf("https://api.vatcomply.com/rates?date=%s", date)

	err := c.GetRequest(&response, url)

	return &response, err
}

func (c Client) GetGeolocation() (*Geolocate, error) {
	var geolocate Geolocate

	url := "https://api.vatcomply.com/geolocate"

	err := c.GetRequest(&geolocate, url)

	return &geolocate, err

}

func (c Client) GetRequest(typeStruct interface{}, url string) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, typeStruct)
	if err != nil {
		return err
	}

	return err
}
