package main

import (
	"fmt"
	"net/url"
	"testing"
)

func addQueryParams(baseURL string, params map[string]string) (string, error) {
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := parsedUrl.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	encodedUrl := q.Encode()
	parsedUrl.RawQuery = encodedUrl
	URLString := parsedUrl.String()

	return URLString, nil
}

func Test(t *testing.T) {
	type testCase struct {
		baseURL  string
		params   map[string]string
		expected string
	}

	testCases := []testCase{
		{
			"https://api.boot.dev/search",
			map[string]string{"page": "2", "limit": "10"},
			"https://api.boot.dev/search?limit=10&page=2",
		},
		{
			"https://api.boot.dev/items?sort=asc",
			map[string]string{"page": "3"},
			"https://api.boot.dev/items?page=3&sort=asc",
		},
		{
			"https://api.boot.dev/users",
			map[string]string{},
			"https://api.boot.dev/users",
		},
		{
			"https://api.boot.dev/search",
			map[string]string{"q": "go lang", "filter": "new"},
			"https://api.boot.dev/search?filter=new&q=go+lang",
		},
		{
			"https://api.boot.dev/items?page=1&sort=asc",
			map[string]string{"page": "5"},
			"https://api.boot.dev/items?page=5&sort=asc",
		},
		{
			"https://api.boot.dev/report",
			map[string]string{"format": "json", "verbose": "true", "region": "us"},
			"https://api.boot.dev/report?format=json&region=us&verbose=true",
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result, err := addQueryParams(test.baseURL, test.params)
		if err != nil {
			failCount++
			t.Errorf("Input: %s %v\nExpected: %s\nActual Error: %v\n",
				test.baseURL, test.params, test.expected, err)
			continue
		}

		if result != test.expected {
			failCount++
			t.Errorf("Input: %s %v\nExpected: %s\nActual: %s\n",
				test.baseURL, test.params, test.expected, result)
		} else {
			passCount++
			fmt.Printf("Input: %s %v\nExpected: %s\nActual: %s\nPass\n",
				test.baseURL, test.params, test.expected, result)
		}
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
