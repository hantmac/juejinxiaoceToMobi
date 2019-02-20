package httpClient

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// reqType is one of HTTP request strings (GET, POST, PUT, DELETE, etc.)
func DoRequest(reqType string, url string, headers map[string]string, data []byte, timeoutSeconds int) ([]byte, map[string][]string, error) {
	var reader io.Reader
	if data != nil && len(data) > 0 {
		reader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(reqType, url, reader)
	if err != nil {
		log.WithError(err).Fatal("http.Newrequest failed")
		return nil, nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	var (
		statusCode int
		body       []byte
		timeout    time.Duration
		header     map[string][]string
	)
	timeout = time.Duration(time.Duration(timeoutSeconds) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.WithError(err).Fatal("client.Do(req) failed")
	}

	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	statusCode = resp.StatusCode
	header = resp.Header
	if statusCode >= 400 {
		log.Info(resp)
		log.WithError(err).Fatal("req failed, status code:" + string(statusCode))
	}
	return body, header, nil

}
