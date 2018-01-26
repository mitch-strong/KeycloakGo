package keycloak

import (
	"context"
	"fmt"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

//Constants
const keycloakJSONFileName = "keycloak.json"

//Global Variables
var client Client                  //Client Object
var realm string                   //realm string from json file
var clientID string                //Client ID string from json file
var clientSecret string            //Client ID sectret from json file
var oauth2Config oauth2.Config     //oath2Config
var provider *oidc.Provider        //oidc provider
var err error                      //generic error object
var keycloakserver string          //keycloak server string passed from app
var server string                  //app server string passed from app
var verifier *oidc.IDTokenVerifier //verifier

//Init begins keycloak server
func Init(keycloakServer, Server string) {
	userLog = GetInstance()
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
