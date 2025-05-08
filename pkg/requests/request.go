package requests

import (
	"fmt"
	"io"
	"net/http"
)

func doRequest(method, path string, body io.Reader, auth bool) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:8000%s", path)

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if auth {

	}

	return http.DefaultClient.Do(req)
}
