package keycloak

import (
	"log"
	"os"
)

//Set up logging to LogFile
func setupLogger() {
	f, err := os.OpenFile("LogFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
}

func logAction(a action) {
	event := getAction(a)
	log.Println(event)
}

func getAction(a action) Action {
	switch a {
	case actionLogin:
		return ActionLogin

	case actionLogout:
		return ActionLogout

	case actionPageAccess:
		return ActionPageAccess
	}
	return ""
}
