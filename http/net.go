package gola

import (
	"bytes"
	"io"
	"net/http"
)

func HTTPRequest(method string, headers http.Header, url string, load []byte) ([]byte, error) {
	errHandle := func(err error) ([]byte, error) {
		return nil, err
	}

	var request *http.Request
	var err error

	request, err = http.NewRequest(method, url, nil)
	if err != nil {
		return errHandle(err)
	}
	if load != nil {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(load))
		if err != nil {
			return errHandle(err)
		}
	}

	request.Header = headers
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return errHandle(err)
	}
	defer resp.Body.Close()

	//unmarshal load
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errHandle(err)
	}

	return body, nil
}
