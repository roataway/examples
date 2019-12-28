//
//  ViewController.swift
//  ios-swift-moscapsule
//
//  Created by Vlad Drahnea on 12/28/19.
//  Copyright © 2019 Vlad Drahnea. All rights reserved.
//

import UIKit
import Moscapsule

class ViewController: UIViewController {
    
    var mqttClient: MQTTClient?

    override func viewDidLoad() {
        super.viewDidLoad()
        self.mqttClientSetup()
    }
    
    func mqttClientSetup() {
//        configure mqtt connection params, clientId is an arbitrary string, can be treated somehow on broker probably
        let mqttConfig = MQTTConfig(clientId: "Moscapsule-iOS-client", host: "opendata.dekart.com", port: 1945, keepAlive: 60)
        
//        callback for new message event handling
        mqttConfig.onMessageCallback = { mqttMessage in
            let json = try? JSONSerialization.jsonObject(with: mqttMessage.payload!, options: [])
            print(json!)
        }

        // create new MQTT Connection, it connects automatically
        self.mqttClient = MQTT.newConnection(mqttConfig)

        // subscribe to transport topic
        mqttClient!.subscribe("telemetry/transport/+", qos: 2)
    }
}

