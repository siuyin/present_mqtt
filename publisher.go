package main

// M010 OMIT
import (
	"fmt"
	"log"
	"time"

	"github.com/eclipse/paho.mqtt.golang" // HL
	"github.com/siuyin/dflt"
)

// M020 OMIT
func main() {
	fmt.Println("Go MQTT publisher")
	url := dflt.EnvString("URL", "tcp://localhost:1883")
	opts := mqtt.NewClientOptions().AddBroker(url) // HL
	opts.SetClientID("go_client")
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil { // HL
		log.Fatal(token.Error())
	}
	// M030 OMIT
	i := 0
	for {
		msg := fmt.Sprintf("go gopher: %d", i)
		c.Publish("my_topic", 0, false, msg) // HL
		i++
		time.Sleep(1 * time.Second)
	}
	// c.Disconnect(500)
}

// M040 OMIT
