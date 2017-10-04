package jenny

import (
	"net/http"
)

func connect() (int, error) {
	resp, err := http.Get("http://example.com/")
	return resp.StatusCode, err
}
