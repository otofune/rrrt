package rrrt

import "net/http"

type MapperFunc func(*http.Request) (*http.Request, error)
type RequestMapper struct {
	ort    http.RoundTripper
	mapper MapperFunc
}

func (rm *RequestMapper) RoundTrip(req *http.Request) (*http.Response, error) {
	req, err := rm.mapper(req)
	if err != nil {
		return nil, err
	}
	return rm.ort.RoundTrip(req)
}

func NewRequestReplaceRoudtripper(origin http.RoundTripper, mapper MapperFunc) http.RoundTripper {
	return &RequestMapper{ort: origin, mapper: mapper}
}
