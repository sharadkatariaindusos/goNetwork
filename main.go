// goNetwork package provides network essential functions that comes handy in day-to-day developmnent operations.
package goNetwork

import (
	HttpProtocol "goNetwork/http"
	"net/http"
)

/* HttpCall perfoms network operations based on the HTTP protocol.

method is the type of request. Eg- GET, POST, PATCH, PUT, DELETE.

url is the request URL.

data is the request body (Optional, pass nil).

headers is a map of request headers that needs to be set for the request (Optional, pass empty map).

The function returns a response object and an error object.
*/
func HttpCall(method string, url string, data []byte, headers map[string]string) (*http.Response, error) {
	return HttpProtocol.Call(HttpProtocol.HttpMethod(method), url, data, headers)
}
