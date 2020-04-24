package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strings"
)

type Telemetry struct {
	Speed     int8
	Latitude  float32
	Longitude float32
	Direction float32
	Timestamp string
	Sat       int8
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	pieces := strings.Split(message.Topic(), "/")
	transportId := pieces[len(pieces)-1]

	var telemetry Telemetry
	if err := json.Unmarshal(message.Payload(), &telemetry); err != nil {
		panic(err)
	}
	fmt.Printf("Vehicle %s\tat lat: %f, lon: %f\n", transportId, telemetry.Latitude, telemetry.Longitude)

}

func main() {
	address := "opendata.dekart.com:1945"
	opts := mqtt.NewClientOptions()
	opts.AddBroker(address)
	opts.SetClientID("roataway-golang-sample")
	opts.SetCleanSession(true)

	qos := byte(0) // quality of service "maybe"
	topic := "telemetry/transport/+"

	fmt.Printf("Connecting to %s ...\n", address)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Printf("Subscribing to %s ...\n", topic)
	if token := c.Subscribe(topic, qos, onMessageReceived); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Printf("Receiving messages, press Ctrl+C to stop\n")
	for {
	}

}
