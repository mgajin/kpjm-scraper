package client

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// secret is key that is used for generating hashed x-kp-signature
	secret = "eMhrpdzwDgWf7zsLsAeS"
	// baseURL is base URL for kupujemprodajem
	baseURL = "https://www.kupujemprodajem.com/mapi.php?"
)

// KpjmClient is wrapper around http client.
// It is used for sending requests to kupujemprodajem
type KpjmClient struct {
	apiClient *http.Client
}

// New returns new KpjmClient.
func New() *KpjmClient {
	return &KpjmClient{
		apiClient: &http.Client{},
	}
}

// NewRequest returns creates GET Request for given route.
// Returns created request and error.
func NewRequest(route string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseURL+route, nil)

	if err != nil {
		return nil, err
	}

	signature := hashRoute(route)
	req.Header.Set("x-kp-signature", signature)

	return req, nil
}

// SendRequest sends Request to kupujemprodajem.
// Returns response body and error.
func (kc *KpjmClient) SendRequest(request *http.Request) ([]byte, error) {
	res, err := kc.apiClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// hashRoute returns hashed string that will be used as x-kp-signature.
func hashRoute(route string) string {
	h := sha1.New()
	_, _ = io.WriteString(h, route+secret)
	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}
