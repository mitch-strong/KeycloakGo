# Keycloak Oauth2 Plugin Adapter for GOLANG
Plugin to add [Keycloak](http://www.keycloak.org/) for authentication to GOLANG applications.

## Installation
```
go get github.com/mitch-strong/keycloakgo
go install
```

## Adding a Client
To create a connection with your keycloak client you will need add the 'Keycloak OIDC JSON' file to <$AppHomeDir>/json/ folder. </br>
This file can be found in the Keycloak Console in **Clients &rarr; <$ClientName> &rarr; Installation**

## Usage
For usage of this plugin see [Documentation](../master/USAGE.md)