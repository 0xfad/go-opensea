package goopensea

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseURL = "https://api.opensea.io/api/v1"
)

type Opensea struct {
	APIKey string
	c      *http.Client
	ctx    context.Context
}

func New(apiKey string) (*Opensea, error) {
	o := &Opensea{
		APIKey: apiKey,
		c:      http.DefaultClient,
		ctx:    context.Background(),
	}

	return o, nil
}

func buildPath(path string) string {
	return fmt.Sprintf(baseURL + path)
}

func (o *Opensea) buildRequest(method, path string, params map[string]string) (*http.Request, error) {
	r, err := http.NewRequest(method, buildPath(path), nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("X-API-KEY", o.APIKey)

	if params != nil {
		q := url.Values{}
		for k, v := range params {
			q.Set(k, string(v))
		}

		r.URL.RawQuery = q.Encode()
	}

	return r, nil
}

func readBody(b io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(b)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func (o *Opensea) GetSingleCollection(name string) (*SingleCollectionResponse, error) {
	req, err := o.buildRequest(http.MethodGet, "/collection/"+name, nil)
	if err != nil {
		return nil, err
	}

	r, err := o.c.Do(req)
	if err != nil {
		return nil, err
	}

	var scr SingleCollectionResponse
	data, err := readBody(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &scr); err != nil {
		return nil, err
	}

	return &scr, nil
}

func (o Opensea) GetAssets(collectionName string, limit, offset int) (*AssetListResponse, error) {
	req, err := o.buildRequest(http.MethodGet, "/assets", map[string]string{
		"collection":      collectionName,
		"limit":           strconv.Itoa(limit),
		"offset":          strconv.Itoa(offset),
		"order_by":        "pk",
		"order_direction": "asc",
	})
	if err != nil {
		return nil, err
	}

	r, err := o.c.Do(req)
	if err != nil {
		return nil, err
	}

	var alr AssetListResponse
	data, err := readBody(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &alr); err != nil {
		return nil, err
	}

	return &alr, nil
}
