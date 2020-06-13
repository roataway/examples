package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var loc *time.Location = loadLocation("Europe/Chisinau")

type telemetry struct {
	Direction float32   `json:"direction"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Sat       int8      `json:"sat"`
	Speed     int8      `json:"speed"`
	Timestamp time.Time `json:"timestamp"`
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	pieces := strings.Split(message.Topic(), "/")
	transportID := pieces[len(pieces)-1]

	var t telemetry
	if err := json.Unmarshal(message.Payload(), &t); err != nil {
		panic(err)
	}

	fmt.Printf(
		"[%s] Vehicle %-10s at (%f, %f), %2d km/h\n",
		t.Timestamp.In(loc).Format("2006-01-02 15:04:05 -07:00"),
		transportID,
		t.Latitude,
		t.Longitude,
		t.Speed,
	)
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

func loadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}
