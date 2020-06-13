A Go example of subscribing to the server and receiving messages.

1. Install Go itself:  
https://golang.org/doc/install

2. Make sure `GOPATH` is set up:  
https://golang.org/doc/gopath_code.html#GOPATH

3. Run the program: `go run .`

Example output:

``` plain
$ go run .
Connecting to opendata.dekart.com:1945 ...
Subscribing to telemetry/transport/+ ...
Receiving messages, press Ctrl+C to stop
[2020-06-25 20:41:03 +03:00] Vehicle 0000119    at (47.032436, 28.820768),  8 km/h
[2020-06-25 20:41:03 +03:00] Vehicle 0000263    at (46.988174, 28.885153), 13 km/h
[2020-06-25 20:41:03 +03:00] Vehicle 0000074    at (47.023495, 28.833403),  0 km/h
[2020-06-25 20:41:02 +03:00] Vehicle 0000237    at (46.987362, 28.885031),  0 km/h
[2020-06-25 20:41:02 +03:00] Vehicle 0000135    at (46.990456, 28.786341), 12 km/h
[2020-06-25 20:41:03 +03:00] Vehicle 000026267  at (47.027409, 28.771643),  0 km/h
...
```
