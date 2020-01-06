package pokeapi_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	teardown := setup()
	defer teardown()

	pokemonURL := fmt.Sprintf("/pokemon/charizard")

	mux.HandleFunc(pokemonURL, func(w http.ResponseWriter, r *http.Request) {
		assertHTTPMethod(t, r, http.MethodGet)
		assertPath(t, r, "/pokemon/charizard")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprintf(w, fixture("get_pokemon.json"))
		if err != nil {
			t.Fatal("could not read json file with mocked data")
		}
	})

	clientURL := client.BaseURL.String()
	testServerURL := server.URL

	t.Logf("Client URL: %v, Server Endpoint URL: %v\n", clientURL, testServerURL)

	t.Run("server endpoint", func(t *testing.T) {
		if clientURL != testServerURL {
			t.Fatalf("incorrect URL endpoint for test, want: %v, got: %v", testServerURL, clientURL)
		}
	})

	t.Run("validate charizard", func(t *testing.T) {
		res, err := client.GetPokemon("charizard")
		if err != nil {
			t.Errorf("error retrieving pokemon: %v", err)
		}

		wantHeight := 17
		gotHeight := res.Height
		if gotHeight != wantHeight {
			t.Errorf("charizard height incorrect, want: %v, got: %v", wantHeight, res.Height)
		}
	})
}
