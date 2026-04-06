# Generate API and Swagger
An example of how i configured a test-api and swagger using oapi-codegen

## API
In the api-folder i have put an example of a yaml-specification, including both api's and components. In my yml i have defined a healt endpoint and a user object including cruds for it.
The generate.go and config.yml-files enables me to create the api.gen.go file using go generate ./...
The api.gen.go file contains a User-struct, endpoints and a ServerInterface

## Handlers
I have three handlers servers.go, check.go and user.go
* servers.go -> contains implementations of the ServerInterface created in api.gen.go. It also defines the requirements (interfaces) towards my check and user-handlers and calls them
* check.go -> implements the CheckHandler-interface defined in servers.go and performs the actual check.
* user.go -> implements the UserHandler-interface defined in servers.go, and defines a UserService-interface. It makes the syntactic-checks of the incoming calls and then calls on the UserService-interface

## Domain
Here a domain.User is defined, to make the app-logic independent of future changes in the api-spec. There is also a list of predefined error-messages

## Service
In user.go i implement the UserService-interface and define the DbInterface.I do the calls to the DbInterface, and perform the logical checks (business rules etc.) 
Results from the Db-interaction and logical checks are sent back to the handlers
The service converts the api.User to domain.User on incoming from the handler- and back again when it returns the result to the handler

## Database
I'm using a map[string]domain.User as a database, accessible through the DbInterface.
That way i can easily change to any other type of database.

## Swagger
For the swagger part i need the index.html and a copy of the yaml-file that i used to create the api. I renamed it to openapi.yaml
The only difference from the original-file is that i added portnumber to the server config on row 12
