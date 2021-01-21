package cmd

import (
	"encoding/json"
	"github.com/gojektech/heimdall/httpclient"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ExecuteRequest(url string, method string, request interface{}, authCreds map[string]string) ([]byte, error) {
	var client *httpclient.Client

	timeout := 100000 * time.Millisecond

	client = httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	reqJSON, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(method, url, strings.NewReader(string(reqJSON)))

	httpReq.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	if authCreds != nil {
		httpReq.SetBasicAuth(authCreds["key"], authCreds["secret"])
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respByte,  nil
}