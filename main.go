package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dzdiscoveryzone/pokeapi-go/pokeapi"
)

func main() {
	fmt.Println("initializing the PokeAPI")
	httpClient := &http.Client{}
	poke := pokeapi.NewClient(httpClient)

	fmt.Println(poke.BaseURL.Path)

	bs, err := poke.GetPokemon("blastoise")
	if err != nil {
		log.Printf("error grabbing pokemon blastoise, err: %+v", err)
	}

	fmt.Println(bs.Name)

	for _, power := range bs.Abilities {
		fmt.Printf("power: %v, URL: %v\n", power.Ability.Name, power.Ability.URL)
	}
}
#I am a forked repository