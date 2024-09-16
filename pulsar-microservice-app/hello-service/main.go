// This service publishes "Hello" to the hello-topic.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
    // Create a Pulsar client
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar client: %v", err)
    }
    defer client.Close()

    // Create a producer
    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: "hello-topic",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar producer: %v", err)
    }
    defer producer.Close()

    // Publish a message
    msg := "Hello"
    _, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
        Payload: []byte(msg),
    })
    if err != nil {
        log.Fatalf("Could not send message: %v", err)
    }
    fmt.Println("Published:", msg)
}
