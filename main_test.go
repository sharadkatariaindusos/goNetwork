package goNetwork

import (
	"net/http"
	"testing"
)

func TestHttpCall(t *testing.T) {
	scenarios := []struct {
		method   string
		url      string
		data     []byte
		headers  map[string]string
		expected int
	}{
		{method: "GET", url: "https://jsonplaceholder.typicode.com/posts/1", data: nil, headers: map[string]string{"content-type": "text/plain"}, expected: http.StatusOK},
		{method: "GET", url: "https://reqres.in/api/users/23", data: nil, headers: map[string]string{"content-type": "text/plain"}, expected: http.StatusNotFound},
		{method: "POST", url: "https://reqres.in/api/login", data: []byte(`{"email": "peter@klaven"}`), headers: map[string]string{"content-type": "application/json"}, expected: http.StatusBadRequest},
		{method: "POST", url: "https://jsonplaceholder.typicode.com/posts/", data: []byte(`{"title": "Test", "body": "Sharad", "userId": 12}`), headers: map[string]string{"content-type": "application/json"}, expected: http.StatusCreated},
		{method: "PUT", url: "https://jsonplaceholder.typicode.com/posts/1", data: []byte(`{"title": "Test", "body": "Sharad", "userId": 12}`), headers: map[string]string{"content-type": "application/json"}, expected: http.StatusOK},
		{method: "PATCH", url: "https://jsonplaceholder.typicode.com/posts/1", data: []byte(`{"title": "Test"}`), headers: map[string]string{"content-type": "application/json"}, expected: http.StatusOK},
		{method: "DELETE", url: "https://jsonplaceholder.typicode.com/posts/1", data: nil, headers: map[string]string{"content-type": "text/plain"}, expected: http.StatusOK},
		{method: "DELETE", url: "https://reqres.in/api/users/2", data: nil, headers: map[string]string{"content-type": "text/plain"}, expected: http.StatusNoContent},
	}

	for index, s := range scenarios {
		got, err := HttpCall(s.method, s.url, s.data, s.headers)
		if err != nil {
			t.Errorf("Did not get expected result for method %v at index %v. Expected status %v, Got error %v", s.method, index, s.expected, err)
		}

		defer got.Body.Close()

		if got.StatusCode != s.expected {
			t.Errorf("Did not get expected result for method %v at index %v. Expected status %v, Got status %v", s.method, index, s.expected, got.StatusCode)
		}
	}
}
