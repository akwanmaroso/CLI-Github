package nap

import (
	"net/http"
	"testing"
)

func TestClient_ProcessRequest(t *testing.T) {
	client := NewClient()
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		return nil
	})
	res := NewResource("/get", "GET", router)
	if err := client.ProcessRequest("https://httpbin.org", res, nil, nil); err != nil {
		t.Fail()
	}
}
