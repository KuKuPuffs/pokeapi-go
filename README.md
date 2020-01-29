# pokeapi-go
![pokemon](pokemon.png)
## This is a Go API Client for the pokeapi website


You can instantiate the pokeapi client in 2 ways:

1\. Supplying a default http.Client to the NewClient method which will create the proper timeout of 2 minutes as well as setting the transport layer to insecure in order to react with the REST API.
```go
    HTTPClient := &http.Client{}

    client := pokeapi.NewClient(HTTPClient)
```

<br />

2\. Using the functional options approach and providing an endpoint URL as well as   an http.CLient
```go
    HTTPClient := &http.Client{Timeout: time.Millisecond * 100,
    			Transport: &http.Transport{TLSClientConfig: &tls.Config{
    				InsecureSkipVerify: true,
    			}}}

    client := pokeapi.NewClientWIthOpts(
    		pokeapi.OptionBaseURL("https://pokemonurl.com/api"),
    		pokeapi.OptionHTTPClient(HTTPClient))
```