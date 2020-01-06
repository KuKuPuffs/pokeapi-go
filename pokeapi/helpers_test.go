package pokeapi_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dzDiscoveryZone/pokeapi-go/pokeapi"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *pokeapi.Client
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = pokeapi.NewClientWIthOpts(
		pokeapi.OptionBaseURL(server.URL),
		pokeapi.OptionHTTPClient(server.Client()))

	return func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// helper method to check the HTTP Method used for all API interactions
func assertHTTPMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if r.Method != want {
		t.Fatalf("incorrect HTTP Method used, got: %v, want: %v", r.Method, http.MethodPost)
	}
}

// helper method to check the URL path
func assertPath(t *testing.T, r *http.Request, wantPath string) {
	t.Helper()
	gotPath := r.URL.Path
	t.Log("path: " + gotPath)
	if gotPath != wantPath {
		t.Fatalf("incorrct path used, want: %v, got: %v", wantPath, gotPath)
	}
}
