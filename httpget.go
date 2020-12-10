package samplewasm

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		defer response.Body.Close()
		return nil, fmt.Errorf("%s: 404 not found", url)
	}
	return response.Body, nil
}
