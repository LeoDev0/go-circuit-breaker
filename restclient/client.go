package restclient

import (
	"fmt"
	"net/http"
)

func GetEndpoint() error {
	response, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("bad response: %d", response.StatusCode)
	}

	return nil
}
