import paho.mqtt.client as mqtt


# Această funcție va fi apelată de fiecare dată când vine un mesaj
# prin MQTT, care conține date telemetrice
def on_message(client, userdata, message):
    print("Date de la `%s`: %s" % (message.topic, message.payload))
    # aici poți adăuga propriile instrucțiuni, pentru a face ceva
    # interesant cu aceste date


# Acest identificator se folosește de server pentru a diferenția
# mai multe aplicații care s-au conectat simultan.
client_id = "roataway-pydemo"
client = mqtt.Client(client_id)

# aici indicăm faptul că va fi apelată funcția ”on_message” (pe care
# am definit-o în linia 6) de fiecare dată când primim un mesaj de la
# server
client.on_message = on_message

# ne conectăm, indicând adresa și portu din documentație
client.connect("opendata.dekart.com", port=1945)

# indicăm un ”topic” la care ne abonăm, pentru a primi telemetrie actuală
client.subscribe("telemetry/transport/+")
client.loop_forever()
