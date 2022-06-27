package utils

import "net/http"

type UtilsPort interface {
	HTTPRequest(url string, method string, payload []byte) (*http.Response, error)
}
