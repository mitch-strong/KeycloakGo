package keycloak

import (
	"math/rand"
)

//randSeq generates a random string of letters of the given length (Helper function)
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
