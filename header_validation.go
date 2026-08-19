package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func validateRequiredHeaders(headers http.Header, required []string) error {
	missing := []string{}
	blank := []string{}

	for _, headerName := range required {
		found := false
		hasValue := false
		for actualName, values := range headers {
			if strings.EqualFold(actualName, headerName) {
				found = true
				for _, value := range values {
					if strings.TrimSpace(value) != "" {
						hasValue = true
					}
				}
			}
		}

		if !found {
			missing = append(missing, headerName)
		} else if !hasValue {
			blank = append(blank, headerName)
		}
	}

	messages := []string{}
	if len(missing) > 0 {
		messages = append(messages, "missing required headers: "+strings.Join(missing, ", "))
	}
	if len(blank) > 0 {
		messages = append(messages, "blank required headers: "+strings.Join(blank, ", "))
	}
	if len(messages) == 0 {
		return nil
	}
	return errors.New(strings.Join(messages, "; "))
}

func main() {
	headers := http.Header{
		"Authorization": {"Bearer token123"},
		"X-Request-ID":  {""},
	}
	err := validateRequiredHeaders(headers, []string{"Authorization", "X-Request-ID", "Content-Type"})
	if err != nil {
		fmt.Println("Result:", err)
	} else {
		fmt.Println("Headers are valid!")
	}
}

/* output: challenges git:(main) ✗ go run header_validation.go
Result: missing required headers: Content-Type; blank required headers: X-Request-ID
➜  challenges git:(main) ✗ */
