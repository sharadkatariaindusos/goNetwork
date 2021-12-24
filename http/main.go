// Contains all the essential functions related to HTTP protocol
package http

import (
	"bytes"
	"net/http"
)

/* Call creates an HTTP request based on the passed arguments

method is the type of request. Eg- GET, POST, PATCH, PUT, DELETE.

url is the request URL.

data is the request body (Optional, pass nil).

headers is a map of request headers that needs to be set for the request (Optional, pass empty map).

The function returns a response object and an error object.
*/
func Call(method HttpMethod, url string, data []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(string(method), url, bytes.NewBuffer(data))
	if len(headers) != 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if err != nil {
		return nil, err
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}
