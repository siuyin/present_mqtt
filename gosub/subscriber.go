package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eclipse/paho.mqtt.golang" // HL
	"github.com/siuyin/dflt"
)

// 010 OMIT
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// 020 OMIT
func main() {
	fmt.Println("Go MQTT subscriber")
	url := dflt.EnvString("URL", "tcp://localhost:1883")
	opts := mqtt.NewClientOptions().AddBroker(url) // HL
	cid := dflt.EnvString("ID", "go_subscriber")
	opts.SetClientID(cid)
	// 050 OMIT
	opts.SetDefaultPublishHandler(f)
	// 060 OMIT
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil { // HL
		log.Fatal(token.Error())
	}
	// M030 OMIT
	if token := c.Subscribe("my_topic", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	// M040 OMIT
	select {}
	// c.Disconnect(500)
}
