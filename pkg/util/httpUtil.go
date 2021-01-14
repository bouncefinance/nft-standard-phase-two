package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"nft_standard/config"
	"os"
)

func PostUrlRetry(url string, params map[string]string, body interface{}, headers map[string]string) ([]byte, error) {
	i := 0
	data, err := PostUrl(url, params, body, headers)
	for err != nil && i < config.RETRY_TIME {
		data, err = PostUrl(url, params, body, headers)
		i++
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func PostUrl(url string, params map[string]string, body interface{}, headers map[string]string) ([]byte, error) {
	var (
		bodyJson []byte
		req      *http.Request
		err      error
	)

	if body != nil {
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return nil, errors.Wrap(err, "json marshal request body error")
		}
	}

	req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest error")
	}

	contentType := "Content-type"
	req.Header.Set(contentType, headers[contentType])
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do error")
	}
	defer response.Body.Close()
	d, err_ := ioutil.ReadAll(response.Body)
	if err_ != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll error")
	}

	return d, nil
}

func Request(method string ,url string, params map[string]string, body interface{}, headers map[string]string) ([]byte, error) {
	var (
		bodyJson []byte
		req      *http.Request
		err      error
	)

	if body != nil {
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return nil, errors.Wrap(err, "json marshal request body error")
		}
	}

	fmt.Println(string(bodyJson))
	req, err = http.NewRequest(method, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest error")
	}

	contentType := "Content-type"
	req.Header.Set(contentType, headers[contentType])
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do error")
	}
	defer response.Body.Close()
	d, err_ := ioutil.ReadAll(response.Body)
	if err_ != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll error")
	}

	return d, nil
}

func GetURLRetry(url string, params map[string]string) ([]byte, error) {
	i := 0
	data, err := GetUrl(url, params)
	for err != nil && i < config.RETRY_TIME {
		data, err = GetUrl(url, params)
		i++
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUrl(url string, data map[string]string) ([]byte, error) {
	params := u.Values{}
	ur, err := u.Parse(url)
	if err != nil {
		return nil, errors.Wrap(err, "url.Parse error")
	}

	for key, value := range data {
		params.Set(key, value)
	}

	ur.RawQuery = params.Encode()
	urlPath := ur.String()

	res, err := http.Get(urlPath)
	if err != nil {
		return nil, errors.Wrap(err, "http.Get error")
	}
	defer res.Body.Close()

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll error")
	}
	return d, nil
}

func Request_(method string,url string , body io.Reader) {
	client := &http.Client{}
	req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := u.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
	req.URL.RawQuery = q.Encode()


	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))


}





