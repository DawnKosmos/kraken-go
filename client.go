package kraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func (c *Client) POST(path string, queryParameters any, result any) (err error) {
	// Adds the path to the base Url
	reqLink, err := c.url.Parse(path)
	if err != nil {
		return err
	}

	//Adds QueryParameters
	values := make(url.Values)
	if queryParameters != nil {
		values, err = query.Values(queryParameters)
		if err != nil {
			return err
		}
	}
	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	fmt.Println(reqLink.String())

	// Prepare Get Request
	req, err := http.NewRequest("POST", reqLink.String(), strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	if c.a != nil { // Sign Request
		b64DecodedSecret, _ := base64.StdEncoding.DecodeString(c.a.SecretKey)
		signature := getSignature(path, values, b64DecodedSecret)
		req.Header.Set("API-Key", c.a.PublicKey)
		req.Header.Set("API-Sign", signature)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	}

	if c.isDebug {
		//Write down request
	}

	//Do Request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Read Json Body and Unmarshal
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(body, result)

	if c.isDebug {
		// log the return code
	}

	return err
}

func SignPOST(a *Account, req *http.Request, values url.Values) {
}

func getSignature(url_path string, values url.Values, secret []byte) string {
	sha := sha256.New()
	sha.Write([]byte(values.Get("nonce") + values.Encode()))

	mac := hmac.New(sha512.New, secret)
	mac.Write(append([]byte(url_path), sha.Sum(nil)...))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func getNonce() uint64 {
	return uint64(time.Now().UnixMilli())
}
