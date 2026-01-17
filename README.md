# Context

This directory holds code relating to sending golang functions over a network
call. I've always been interested in seeing if something like this is possible
though I haven't found anything online saying it is. My initial idea is to have
a client send a Go file to the server. The server receives the file, then runs
os.Exec to build/run the Go program contained in the file. While this isn't
technically passing a function, it does accomplish what I was interested in
exploring. After implementing that, I can see if I can trim down what's passed
in the Go file to only contain a function that the server somehow calls, but
we'll see.

# Instructions

- Create two terminal shells
- Navigate into `client` and `server`
- In `server` run `go run main.go`
- In `client` run `go run main.go`
- To modify source that's executed, update the `client/source/source.go` file

# How it works

The client sends a file to the server to be executed. The client must send a Go
file that is in the `source` package and contains a function called `Function`.
This is the interface that allows the server to execute the file.
