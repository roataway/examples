Example of subscribing to the server and receiving a message, written in Go.

1. Install `go` itself: https://golang.org/doc/install
2. Make sure it is available in your environment `export PATH=$PATH:/usr/local/go/bin`
3. Install dependencies: `go get github.com/eclipse/paho.mqtt.golang` and `go get golang.org/x/net/proxy`
5. Build the binary: `go build main.go`
6. Run the program `./main`
