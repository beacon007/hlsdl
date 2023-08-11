package base

import (
	"crypto/tls"
	"net/http"
	"time"
)

var DefaultHttpClient *http.Client

func init() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 256
	t.MaxConnsPerHost = 128
	t.MaxIdleConnsPerHost = 16
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	DefaultHttpClient = &http.Client{
		Timeout:   60 * time.Second,
		Transport: t,
	}
}
