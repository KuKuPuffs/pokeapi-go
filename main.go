package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dzDiscoveryZone/pokeapi-go/pokeapi"
)

func main() {
	fmt.Println("initializing the PokeAPI")
	httpClient := &http.Client{}
	poke := pokeapi.NewClient(httpClient)

	bs, err := poke.GetPokemon("charizard")
	if err != nil {
		log.Printf("error grabbing pokemon blastoise, err: %+v", err)
	}
	fmt.Println(bs)
	// fmt.Println(bs.Abilities)
	//for _, power := range bs.Abilities {
	//	fmt.Printf("power: %v, URL: %v\n", power.Ability.Name, power.Ability.URL)
	//}
}
