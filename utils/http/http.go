package utils_http

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

type HttpHeaderType = string

const (
	ContentType   HttpHeaderType = "Content-Type"
	Authorization HttpHeaderType = "Authorization"
)

type HttpHeader struct {
	Type  HttpHeaderType
	Value string
}

func RequestWithTimeout(method string, headers []HttpHeader, url string, body []byte, timeout *time.Duration) (*http.Request, *http.Client, context.CancelFunc, error) {
	defaultTimeout := 2 * time.Second
	requestTimeout := timeout
	if requestTimeout == nil {
		requestTimeout = &defaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), *requestTimeout)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	for _, header := range headers {
		req.Header.Add(header.Type, header.Value)
	}
	return req, client, cancel, err
}

func PerformRequest(context string, path string, method string, headers []HttpHeader, body []byte, timeout *time.Duration) (*http.Response, error) {
	log.Debugf("%s: Performing request with path: %s and method: %s", context, path, method)
	req, client, cancel, err := RequestWithTimeout(
		method,
		headers,
		path,
		body,
		timeout,
	)
	defer cancel()
	res, err := client.Do(req)
	if err != nil {
		log.Error("%s: Could nor perform request", context)
		return nil, err
	}
	if res.StatusCode != 200 {
		log.Error("%s: Request failed with status code %d", context, res.StatusCode)
		return nil, err
	}
	return res, nil
}
