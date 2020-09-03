package nap

import (
	"net/http"
	"testing"
)

func TestAPI_Call(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		return nil
	})
	res := NewResource("/", "GET", router)
	api.AddResource("get", res)
	if err := api.Call("get", nil, nil); err != nil {
		t.Fail()
	}

	resources := api.ResourceNames()
	if len(resources) != 1 || resources[0] != "get" {
		t.Fail()
	}
}

func TestAPI_Auth(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		return nil
	})
	res := NewResource("/basic-auth/{{.user}}/{{.password}}", "GET", router)
	api.AddResource("basicauth", res)
	api.SetAuth(&AuthBasic{
		Username: "user",
		Password: "passw0rd",
	})
	if err := api.Call("basicauth", map[string]string{
		"user":     "user",
		"password": "passw0rd",
	}); err != nil {
		t.Fail()
	}

}
