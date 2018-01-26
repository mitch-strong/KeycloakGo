package keycloak

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var userLog *logger
var logs *logger
var once sync.Once

//GetInstance returns a new logger to a file
func GetInstance() *logger {
	once.Do(func() {
		logs = createLogger("UserLogs.log")
	})
	return logs
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "", log.Ldate|log.Ltime),
	}
}

func logAction(username string, a action) {
	event := getAction(a)
	userLog.Println(username+": ", event)
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
