package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()

	time.Sleep(5 * time.Second)

	testCases := []struct {
		desc       string
		method     string
		endpoint   string
		StatusCode int
		body       []byte
	}{
		{"Get courses", http.MethodGet, "get", http.StatusOK, nil},
		{"Create courses", http.MethodPost, "create", http.StatusCreated, []byte(`{
			"id" : 11,
			"name": "ABC",
			"price": 1999,
			"author": "Ay09"
		}`)},
		{"Update courses", http.MethodPut, "update/11", http.StatusOK, []byte(`{
			"name": "Update AB",
			"price": 999,
			"author": "AJ"
		}`) },
		{"Delete courses", http.MethodDelete, "delete/11", http.StatusNoContent, nil},
	}

	for i, tc := range testCases {
		req, _ := request.NewMock(tc.method, "http://localhost:3000/"+tc.endpoint, bytes.NewBuffer(tc.body))
		client := http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue
		}
		if resp.StatusCode != tc.StatusCode{
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.StatusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()

	}
}
