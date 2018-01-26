# Keycloak GOLANG Plugin Usage
This is a GOLANG application plugin deigned to work with keycloak for authentication

## Adding a Client
To create a connection with your keycloak client you will need add the 'Keycloak OIDC JSON' file to <$AppHomeDir>/json/ folder. </br>
This file can be found in the Keycloak Console in **Clients &rarr; <$ClientName> &rarr; Installation**

## Methods 
### Init
Init will create the connection between the keycloak client and your application </br>
**Params:** 
```go
Init(keycloakserver string, server string)
```

**Calling:**
```go
server = "http://" + <$AppHost> + ":" + <$AppPort>
keycloakserver = "http://" + <$KeycloakHost> + ":" + <$KeycloakPort>

keycloak.Init(keycloakserver, server)
```

### Login
Login and LoginCallback will direct the user the login screen and verify their login token </br>
**Params:** 
```go
HandleLogin(w http.ResponseWriter, r *http.Request)
HandleLoginCallback(w http.ResponseWriter, r *http.Request)
```

**Calling:**
```go
//Login User
//Unauthenticated
Route{
	"handleLogin",
	"GET",
	"/login",
	keycloak.HandleLogin,
},
//Login helper
//Authenticated
Route{
	"handleLoginCallback",
	"GET",
	"/loginCallback",
	keycloak.HandleLoginCallback,
},
```


### MiddleWare
Middleware provides a means to have users authenticated before accessing a page </br>
**Params:** 
```go
AuthMiddleware(next http.HandlerFunc) http.HandlerFunc 
```

**Calling:**
```go
//Example Function Redirect
//Authenticated
//indexHandler is an http.HandlerFunc
Route{
	"Index",
	"GET",
	"/",
	keycloak.AuthMiddleware(indexHandler),
},
```

### Logout
Middleware provides a means to have users authenticated before accessing a page </br>
**Params:** 
```go
Logout(w http.ResponseWriter, r *http.Request)
```

**Calling:**
```go
//Logout, redirects to login
///Unauthenticatec
Route{
	"logout",
	"GET",
	"/logout",
	keycloak.Logout,
},
```