package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetJson(url string, target interface{}) error {
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print("Error closing GetJson body")
		}
	}(r.Body)
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf(r.Status)
	}
	return json.NewDecoder(r.Body).Decode(target)
}
