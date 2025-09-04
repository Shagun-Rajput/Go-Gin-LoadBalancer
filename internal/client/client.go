package client

import (
	"net/http"
	"net/url"
)

func ProxyToServer(selectedServer string, originalReq *http.Request) (*http.Response, error) {
	targetURL := url.URL{
		Scheme: "http",
		Host:   selectedServer,
		Path:   originalReq.URL.Path,
	}
	if originalReq.URL.RawQuery != "" {
		targetURL.RawQuery = originalReq.URL.RawQuery
	}

	req, err := http.NewRequest(originalReq.Method, targetURL.String(), originalReq.Body)
	if err != nil {
		return nil, err
	}

	// Copy headers
	for k, v := range originalReq.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
