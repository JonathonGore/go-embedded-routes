# go-embedded-routes

This is an example program demonstrating how to use embedded structs to organize api endpoints for a `Go` web app.

## Prerequisites
Run `dep ensure` to gather the needed dependencies. 

## Running the program

Run `go run main.go` to start the server. This will start a server over http on port `3000`.

This creates two dummy routes that accept `GET` requests.

These can be reached by using curl:

`curl http://localhost:3000/user`
`curl http://localhost:3000/post`
