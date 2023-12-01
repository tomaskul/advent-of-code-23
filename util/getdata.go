package util

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// GetData sends a HTTP GET request using auth session cookie to
// specified (advent of code) URL.
func GetData(url, sessionCookie string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	client := &http.Client{
		Timeout: time.Duration(time.Second * 3),
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response: %d", response.StatusCode)
	}

	return io.ReadAll(response.Body)
}

// GetData sends a HTTP GET request using auth session cookie to
// specified (advent of code) URL and splits data into rows (\n delimitated).
func GetRows(url, sessionCookie string) []string {
	bytes, err := GetData(url, sessionCookie)
	if err != nil {
		return []string{}
	}

	rows := strings.Split(string(bytes), "\n")
	return rows[:len(rows)-1]
}
