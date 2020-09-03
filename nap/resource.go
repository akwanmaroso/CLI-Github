package nap

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

type RestResource struct {
	Endpoint string
	Method   string
	Router   *CBRouter
}

func NewResource(endpoint, method string, router *CBRouter) *RestResource {
	return &RestResource{
		Endpoint: endpoint,
		Method:   method,
		Router:   router,
	}
}

func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}

	t, err := template.New("resource").Parse(r.Endpoint)
	if err != nil {
		log.Fatalf("Unable to parse endpoint: %v\n", err)
	}

	buffer := &bytes.Buffer{}
	t.Execute(buffer, params)

	endpoint, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalf("Unable to read endpoint: %v\n", err)
	}

	return string(endpoint)
}
