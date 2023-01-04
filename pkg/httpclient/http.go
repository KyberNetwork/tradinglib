package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KyberNetwork/tradinglib/pkg/sb"
)

func DoHTTPRequest(client *http.Client, req *http.Request, out interface{}) error {
	if client == nil {
		return fmt.Errorf("client must not be nil")
	}
	if req == nil {
		return fmt.Errorf("req must not be nil")
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return fmt.Errorf("read respsonse error %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected code %d: %s", resp.StatusCode, data)
	}
	if out != nil {
		if err = json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("unmarshal error: %w - %s", err, data)
		}
	}
	return nil
}

func NewRequest(method, baseURL, path string, query Query, body io.Reader) (*http.Request, error) {
	url := baseURL + path
	if query != nil {
		url = sb.Concat(url, "?", query.String())
	}
	return http.NewRequest(method, url, body)
}

func NewPostJSONWithReader(baseURL, path string, query Query, body io.Reader) (*http.Request, error) {
	req, err := NewRequest(http.MethodPost, baseURL, path, query, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func NewPostJSON(baseURL, path string, query Query, body interface{}) (*http.Request, error) {
	var buff io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body failed: %w", err)
		}
		buff = bytes.NewBuffer(data)
	}
	return NewPostJSONWithReader(baseURL, path, query, buff)
}
