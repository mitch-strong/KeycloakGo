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

var logs *logger
var once sync.Once

// start loggeando
func GetInstance() *logger {
	once.Do(func() {
		logs = createLogger("UserLogs.log")
	})
	return logs
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "User Action: ", log.Ldate|log.Ltime),
	}
}

func logAction(a action) {
	event := getAction(a)
	userLog.Println(event)
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
