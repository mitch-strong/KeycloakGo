package keycloak

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

//Constants
const keycloakJSONFileName = "keycloak.json"

//Global Variables
var client Client
var realm string
var clientID string
var clientSecret string
var oauth2Config oauth2.Config
var provider *oidc.Provider
var err error
var keycloakserver string
var server string

//Verifier is the oidc Verifier
var verifier *oidc.IDTokenVerifier

//Init begins keycloak server
func Init(keycloakServer, Server string) {
	getKeycloakJSON()
	keycloakserver = keycloakServer
	server = Server
	ctx := context.Background()
	//Gets the provider for authentication (keycloak)
	provider, err = oidc.NewProvider(ctx, keycloakserver+"/auth/realms/"+realm)
	if err != nil {
		fmt.Printf("This is an error with regard to the context: %v", err)
	}
	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  server + "/loginCallback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

}

func getKeycloakJSON() {

	jsonFile, err := os.Open("./json/" + keycloakJSONFileName)
	if err != nil {
		fmt.Print(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()
	json.Unmarshal(byteValue, &client)
	if client.ID == "" || client.Realm == "" || client.Credentials.Secret == "" {
		fmt.Printf("Error reading keycloak file")
	}
	realm = client.Realm
	clientID = client.ID
	clientSecret = client.Credentials.Secret
}
