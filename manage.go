package jenny

import (
	"net/http"
)

func Get() (int, error) {
	resp, err := http.Get("http://example.com/")
	return resp.StatusCode, err
}
