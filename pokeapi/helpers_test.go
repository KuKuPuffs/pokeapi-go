package pokeapi_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dzdiscoveryzone/pokeapi-go/pokeapi"
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

	got := r.URL.Path
	if got != wantPath {
		t.Errorf("incorrct URL path used, got %v, want: %v", r.URL.Path, wantPath)
	}
}

// helper method to validate the query parameter passed into the URL path
//func assertQueryParam(t *testing.T, r *http.Request, query string, want string) {
//	t.Helper()
//
//	param := r.URL.Query().Get(query)
//	if param != want {
//		t.Error("Url Param 'name' is missing")
//	}
//}
