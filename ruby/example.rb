require 'rubygems'
require 'mqtt'

# Conectarea la server
MQTT::Client.connect('opendata.dekart.com', 1945) do |c|
  # Abonarea la MQTT topic-ul prin care se distribuie datele telemetrice
  c.get('telemetry/transport/+') do |topic,message|
    # Acest bloc va fi executat de fiecare data cind a fost receptionat
    # un mesaj
    puts "Date noi de la #{topic}: #{message}"
  end
end
