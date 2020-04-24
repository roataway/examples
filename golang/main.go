package main

import (
	"fmt"
	"os"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func main() {
	address := "opendata.dekart.com:1945"
	opts := mqtt.NewClientOptions()
	opts.AddBroker(address)
	opts.SetClientID("roataway-golang-sample")
	opts.SetCleanSession(true)

	qos := byte(0)  // maybe once
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
	for ;; {}

}
