package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func DoHTTPRequest(client *http.Client, req *http.Request, out interface{}, options ...Option) (*http.Response, error) {
	if client == nil {
		return nil, fmt.Errorf("client must not be nil")
	}
	if req == nil {
		return nil, fmt.Errorf("req must not be nil")
	}
	option := defaultOption()
	option.apply(options...)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read respsonse error %w", err)
	}
	if resp.StatusCode != option.expectedStatusCode {
		return nil, fmt.Errorf("unexpected code %d: %s", resp.StatusCode, data)
	}
	if out != nil {
		if err = json.Unmarshal(data, out); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w - %s", err, data)
		}
	}
	return resp, nil
}

func NewRequest(method, baseURL, path string, query Query, body io.Reader) (*http.Request, error) {
	return NewRequestWithContext(context.Background(), method, baseURL, path, query, body)
}

func NewRequestWithContext(
	ctx context.Context, method, baseURL, path string, query Query, body io.Reader,
) (*http.Request, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}

	if path != "" {
		u = u.JoinPath(path)
	}
	if query != nil {
		u.RawQuery = query.String()
	}

	return http.NewRequestWithContext(ctx, method, u.String(), body)
}

func NewGet(baseURL, path string, query Query) (*http.Request, error) {
	return NewRequest(http.MethodGet, baseURL, path, query, nil)
}

func NewGetWithContext(ctx context.Context, baseURL, path string, query Query) (*http.Request, error) {
	return NewRequestWithContext(ctx, http.MethodGet, baseURL, path, query, nil)
}

func NewPost(baseURL, path string, query Query, body io.Reader) (*http.Request, error) {
	return NewRequest(http.MethodPost, baseURL, path, query, body)
}

func NewPostWithContext(
	ctx context.Context, baseURL, path string, query Query, body io.Reader,
) (*http.Request, error) {
	return NewRequestWithContext(ctx, http.MethodPost, baseURL, path, query, body)
}

func NewPostJSON(baseURL, path string, query Query, body interface{}) (*http.Request, error) {
	return NewPostJSONWithContext(context.Background(), baseURL, path, query, body)
}

func NewPostJSONWithContext(
	ctx context.Context, baseURL, path string, query Query, body interface{},
) (*http.Request, error) {
	var buff io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body failed: %w", err)
		}
		buff = bytes.NewBuffer(data)
	}
	out, err := NewPostWithContext(ctx, baseURL, path, query, buff)
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}
