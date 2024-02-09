package graphql

import (
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
)

type authedTransport struct {
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(os.Getenv("API_KEY"), os.Getenv("API_KEY_VALUE"))
	return t.wrapped.RoundTrip(req)
}

func CreateClient() graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL_URL"), &http.Client{
		Transport: &authedTransport{wrapped: http.DefaultTransport},
	})
}
