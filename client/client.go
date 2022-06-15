package client

import (
	"encoding/json"
	"fmt"
	whm "github.com/aalkan/whm-go"
	"net/http"
	"reflect"
	"time"
)

type client struct {
	config     *whm.Config
	httpClient *http.Client
}

func NewClient(config *whm.Config) *client {
	return &client{
		config: config,
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (c *client) request(method, endpoint string, params, response interface{}) error {
	url := fmt.Sprintf("%s:%s/json-api%s%s", c.config.Host, c.config.Port, endpoint, "?api.version=1")
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", "whm "+c.config.User+":"+c.config.ApiToken)
	if params != nil {
		query := request.URL.Query()
		reflectVal := reflect.ValueOf(params).Elem()
		for i := 0; i < reflectVal.NumField(); i++ {
			value := reflectVal.Field(i).Interface().(string)
			if value != "" {
				query.Add(reflectVal.Type().Field(i).Tag.Get("json"), value)
			}
		}
		request.URL.RawQuery = query.Encode()
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
