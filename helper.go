package keycloak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//randSeq generates a random string of letters of the given length (Helper function)
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//gets the keycloak JSON informtion from the json file in the json directrory of the users app
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
