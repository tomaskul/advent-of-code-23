package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

// GetCachedRows sends a HTTP GET request using auth session cookie to
// specified (advent of code) URL. After successful request, data is saved to local file.
func GetCachedRows(url, fileName, fileExt, sessionCookie string) ([]string, error) {
	filePath := fmt.Sprintf("data/%s%s", fileName, fileExt)
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("getting data from server...\n")
		dataToSave, err := GetData(url, sessionCookie)
		if err != nil {
			return []string{}, err
		}

		if err = saveDataToFile(filePath, dataToSave); err != nil {
			return []string{}, err
		}

		data, err = os.ReadFile(filePath)
		if err != nil {
			return []string{}, err
		}
	} else {
		fmt.Printf("getting data from local cache...\n")
	}

	rows := strings.Split(string(data), "\n")
	return rows[:len(rows)-1], nil
}

// saveDataToFile saves data using specified name to solution data folder.
func saveDataToFile(filePath string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0750); err != nil {
		return fmt.Errorf("error creating directory for solution data files: %v", err)
	}
	if err := os.WriteFile(filePath, data, 0755); err != nil {
		return fmt.Errorf("error saving file %q: %v", filePath, err)
	}

	return nil
}
