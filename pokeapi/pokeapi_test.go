package pokeapi_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	teardown := setup()
	defer teardown()

	pokemon := "charizard"
	pokemonURL := fmt.Sprintf("/pokemon/%v", pokemon)

	mux.HandleFunc(pokemonURL, func(w http.ResponseWriter, r *http.Request) {
		assertHTTPMethod(t, r, http.MethodGet)
		assertPath(t, r, pokemonURL)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprintf(w, fixture("get_pokemon.json"))
		if err != nil {
			t.Fatalf("fmt.Fprintf(w, fixture(get_pokemon.json)) = %v", err)
		}
	})

	clientURL := client.BaseURL.String()
	testServerURL := server.URL

	t.Logf("Client URL: %v, Server Endpoint URL: %v\n", clientURL, testServerURL)

	t.Run("validate test URL", func(t *testing.T) {
		if clientURL != testServerURL {
			t.Fatalf("clientURL = %v, want: %v\n", clientURL, testServerURL)
		}
	})

	t.Run(pokemon, func(t *testing.T) {
		res, err := client.GetPokemon(pokemon)
		if err != nil {
			t.Errorf("client.GetPokemon() = %v\n", err)
		}

		wantHeight := 17
		gotHeight := res.Height
		if gotHeight != wantHeight {
			t.Errorf("res.Height = %v, want: %v\n", res.Height, wantHeight)
		}
	})
}
