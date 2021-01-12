# Simple Chat Backend in Go

### Instructions

They are located in the *docs/index.html* file

### Prerequisites

Installed Go version >= 1.12 since it uses Go Modules.

### Installation
`
go mod tidy
`

### How to run it
`
go run cmd/server.go
`

##### Next Steps
Since this project has been made as part of a challenge that was not supposed to last more than some hours,
there are some unfinished features that would be very nice to have.

To name a few:

1. Add integration tests for all controllers.
1. Use api errors to handle different statuses for errors and provide a json error response.
1. Support video type messages & extend functionality/metadata for image type messages.
1. Add tests for the sqlite implementation of the db interface.
1. Add tests for the jwt module.
1. Provide an abstraction from the hash function used for passwords for easier testing and modifications.
1. Add Dockerfile and k8s for deployment.

