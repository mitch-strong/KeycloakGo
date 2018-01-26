package keycloak

//Client Values from JSON file
type Client struct {
	Realm       string `json:"realm"`
	ID          string `json:"resource"`
	Credentials Creds  `json:"credentials"`
}

//Creds is a substruct of Keycloak
type Creds struct {
	Secret string `json:"secret"`
}

type action int

const (
	actionLogin action = iota
	actionLogout
	actionPageAccess
	actionInvalid
)

type Action string

var (
	ActionLogin      Action = "Login"
	ActionLogout     Action = "Logout"
	ActionPageAccess Action = "Access"
)
