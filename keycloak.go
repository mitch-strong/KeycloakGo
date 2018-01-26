package keycloak

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

//Constants
const keycloakJSONFileName = "keycloak.json"

//Global Variables
var client Client

//Realm is the keycloak realm
var Realm string

//ClientID is the keycloak ClientID
var ClientID string

//ClientSecret is the keycloak client secret
var ClientSecret string

//Oauth2Config is the oath2 config struct
var Oauth2Config oauth2.Config
var provider *oidc.Provider
var err error
var keycloakserver string
var server string

//Verifier is the oidc Verifier
var Verifier *oidc.IDTokenVerifier

//Init begins keycloak server
func Init(keycloakServer, Server string) {
	getKeycloakJSON()
	keycloakserver = keycloakServer
	server = Server
	ctx := context.Background()
	//Gets the provider for authentication (keycloak)
	provider, err = oidc.NewProvider(ctx, keycloakserver+"/auth/realms/"+Realm)
	if err != nil {
		fmt.Printf("This is an error with regard to the context: %v", err)
	}
	Verifier = provider.Verifier(&oidc.Config{ClientID: ClientID})

	// Configure an OpenID Connect aware OAuth2 client.
	Oauth2Config = oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  server + "/loginCallback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

}

func getKeycloakJSON() {
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(path.Dir(filename))

	jsonFile, err := os.Open(path + "/json/" + keycloakJSONFileName)
	if err != nil {
		fmt.Print(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()
	json.Unmarshal(byteValue, &client)
	if client.ID == "" || client.Realm == "" || client.Credentials.Secret == "" {
		fmt.Printf("Error reading keycloak file")
	}
	Realm = client.Realm
	ClientID = client.ID
	ClientSecret = client.Credentials.Secret
}
