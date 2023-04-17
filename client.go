package kraken

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
)

const URL = "https://api.kraken.com"

type Client struct {
	url     *url.URL
	a       *Account
	client  *http.Client
	isDebug bool
}

type Account struct {
	PublicKey string //Only needed for Private Api Calls
	SecretKey string //Only needed for Private Api Calls
}

func New(client *http.Client, a *Account, debug bool) (*Client, error) {
	var c *Client = new(Client)
	var err error
	if a != nil {
		c.a = a
	}
	c.url, err = url.Parse(URL)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	c.client = client

	return c, nil
}

func (c *Client) GET(path string, queryParameters any, result any) (err error) {
	// Adds the path to the base Url
	reqLink, err := c.url.Parse(path)
	if err != nil {
		return err
	}

	//Adds QueryParameters
	if queryParameters != nil {
		v, err := query.Values(queryParameters)
		if err != nil {
			return err
		}
		reqLink.RawQuery = v.Encode()
	}

	fmt.Println(reqLink.String())

	// Prepare Get Request
	req, err := http.NewRequest("GET", reqLink.String(), nil)
	if err != nil {
		return err
	}
	if c.a != nil { // Sign Request
		//SignGET(c.a, req, req.URL.RawQuery)
	}

	if c.isDebug {
		fmt.Println("Get Request", req.URL.String())
	}

	//Do Request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Read Json Body and Unmarshal
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)

	if c.isDebug {
		// log the return code
	}

	return err
}
