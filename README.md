[![GoDoc](https://godoc.org/github.com/3DSIM/simulation-goclient?status.svg)](https://godoc.org/github.com/3DSIM/simulation-goclient)

# simulation-goclient
Go client for interacting with the 3DSIM simulation api.  See http://docs.simulationapi.apiary.io for full API documentation.
Issues and PRs are welcome.  If you are a 3DSIM customer and have questions, feel free to visit https://community.3dsim.com to start a discussion.

## Background Info
We use https://goswagger.io to generate our Go APIs and clients.  This allows
us to build our APIs in a "design first" manner.

First we create a `swagger.yaml` file that defines the API.  Then we run a command
to generate the server code.

Additionally, this allows us to automatically generate client code.  The code in this
directory was all generated using the `go-swagger` tools.


## Organization
* `auth0` - package for getting auth0 tokens
* `client` - the client package that adds convenience methods for common operations
* `genclient` - the generated client code
   * `client/examples` - examples of how to use the `genclient` package
* `helpers` - a few simple helper utilities
* `models` - the generated models
* `sugar` - syntactic sugar for easily working with pointers

## Regenerating code
First install the swagger generator.  Currently we are using commit 5da5a44596c2c9e69568ea9b9f06ef4c48c6c5fe of https://github.com/go-swagger/go-swagger.

For mac users:
* brew tap go-swagger/go-swagger
* brew install go-swagger

For windows users:
* See https://github.com/go-swagger/go-swagger for options

The code generator needs a specification file.  The specification for the simulation API is stored in `github.com/3dsim/simulation-api-specification/swagger.yaml`.  Assuming that project
is cloned as a sibling project, the command to run to generate new client code is:
```
swagger generate client -A SimulationAPI -f ../simulation-api-specification/swagger.yaml --client-package genclient
```

## Using the client
TODO

## Client to API version compatibility

| Simulation API | Simulation Client |
| ------------- | ------------- |
| 0.14.x  | 0.3.x |

