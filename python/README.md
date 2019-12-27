# Exemplu de conexiune la MQTT server cu Python

1. Instalează librăria `paho-mqtt` prin comanda `pip install paho-mqtt`
2. Lansează `python demo.py`

Dacă totul e bine, vei vedea pe ecran cam așa ceva:

```
Date de la `telemetry/transport/0000313`: b'{"latitude":46.992603,"longitude":28.818197,"timestamp":"2019-11-10T15:11:23Z","speed":3,"direction":125.0,"sat":13}'
Date de la `telemetry/transport/0000138`: b'{"latitude":47.037434,"longitude":28.817371,"timestamp":"2019-12-27T22:35:15Z","speed":0,"direction":0.0,"sat":10}'
```

Vezi comentariile din `demo.py` pentru mai multe detalii.

# Articole recomandate
- https://cumulocity.com/guides/device-sdk/mqtt-examples/#hello-mqtt-python
- https://pypi.org/project/paho-mqtt
- În general, în lumea Python, ar fi bine să te familiarizezi cu `virtualenv`: https://docs.python-guide.org/dev/virtualenvs/