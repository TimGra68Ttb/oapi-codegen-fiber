# Generate API and Swagger
An example of how i configured a test-api and swagger using oapi-codegen

## API
In the api-folder i have put an example of a yaml-specification, including both api's and components. It has a health-endpoint and crud operations for the User-stuct that is specified in the components-section

There is also a generate.go-file that enables me to create the api.gen.go file using go generate ./...

## Database
I'm using a map[string]api.User as a database, accessible through the DbInterface.
That way i can easily change to any other type of database.


## Swagger
For the swagger part i need the index.html and a copy of the yaml-file that i used to create the api. I renamed it to openapi.yaml
The only difference from the original-file is that i added portnumber to the server config on row 12
