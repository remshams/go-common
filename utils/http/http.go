package utils_http

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
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

type QueryParam struct {
	Key   string
	Value string
}

func RequestWithTimeout(
	method string,
	headers []HttpHeader,
	url string,
	queryParams []QueryParam,
	body []byte,
	timeout *time.Duration,
) (*http.Request, *http.Client, context.CancelFunc, error) {
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
	query := req.URL.Query()
	for _, qp := range queryParams {
		query.Add(qp.Key, qp.Value)
	}
	log.Debugf("Request queries: %v", query)
	req.URL.RawQuery = query.Encode()
	return req, client, cancel, err
}

func PerformRequest(context string, path string, method string, headers []HttpHeader, queryParams []QueryParam, body []byte, timeout *time.Duration) (*http.Response, []byte, error) {
	log.Debugf("%s: Performing request with path: %s and method: %s", context, path, method)
	req, client, cancel, err := RequestWithTimeout(
		method,
		headers,
		path,
		queryParams,
		body,
		timeout,
	)
	defer cancel()
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("%s: Could not perform request", context)
		return nil, nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		logErrorResponse(context, res)
		return nil, nil, errors.New(fmt.Sprintf("Request failed with status code %d", res.StatusCode))
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("%s: Could not read response body", context)
		return nil, nil, err
	}
	return res, resBody, nil
}

func logErrorResponse(context string, res *http.Response) {
	log.Errorf("%s: Request failed with status code %d", context, res.StatusCode)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Could not convert response body")
	}
	log.Errorf("Response body: %s", string(body))
}

func CreateBasicAuthHeader(username string, password string) HttpHeader {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return HttpHeader{
		Type:  Authorization,
		Value: fmt.Sprintf("Basic %s", auth),
	}
}

func CreateBearerTokenHeader(token string) HttpHeader {
	return HttpHeader{
		Type:  Authorization,
		Value: fmt.Sprintf("Bearer %s", token),
	}
}
