# Generate API and Swagger

An example of how i configured a test-api and swagger using oapi-codegen

## API

The file oapi_codegen.yml includes both api's and components. It defines a health endpoint and a user object including cruds endpoints for it.
The generate.go and config.yml-files enables the creation of api.gen.go from oapi_codegen.yml, using the command go generate ./...
The api.gen.go file contains a User-struct, endpoints and a ServerInterface

## Handlers

There are three handlers in the handlers package, servers.go, check.go and user.go

- servers contains implementations of the ServerInterface created in api.gen.go. It also defines the requirements (interfaces) towards my check and user-handlers and calls them
- check implements the CheckHandler-interface defined in servers.go and performs the actual check.
- user implements the UserHandler-interface defined in servers.go, and defines a UserService-interface. It makes the syntactic-checks of the incoming calls and then calls on the UserService-interface

## Domain

Here a domain.User is defined, to make the app-logic independent of future changes in the api-spec. There is also a list of predefined error-messages

## Service

The user.go implements the UserService-interface and defines the DbInterface. It calls the DbInterface-functions, and performs the logical checks (business rules etc.)
Results from the Db-interaction and logical checks are sent back to the handlers
The service converts the api.User to domain.User on incoming from the handler- and back again when it returns the result to the handler

## Database

The implementation of tha DbInterface is a map[string]domain.User.
To be replaced by postgres or any other database

## Swagger

For the swagger part, the index.html and a copy of the yaml-file that was used to create the api is needed. the yaml-file is renamed to openapi.yaml
The only difference from the original-file is that the portnumber is added to the server config
