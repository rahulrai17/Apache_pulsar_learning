// This service consumes from hello-topic, appends "World", and publishes to world-topic.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar client: %v", err)
    }
    defer client.Close()

    // Create a consumer to listen to hello-topic
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "hello-topic",
        SubscriptionName: "world-subscription",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar consumer: %v", err)
    }
    defer consumer.Close()

    msg, err := consumer.Receive(context.Background())
    if err != nil {
        log.Fatalf("Could not receive message: %v", err)
    }
    fmt.Printf("Received message: %s\n", string(msg.Payload()))

    // Process and send next word
    word := string(msg.Payload()) + " World"
    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: "world-topic",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar producer: %v", err)
    }
    defer producer.Close()

    _, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
        Payload: []byte(word),
    })
    if err != nil {
        log.Fatalf("Could not send message: %v", err)
    }
    fmt.Println("Published:", word)

    // Acknowledge the message
    consumer.Ack(msg)
}
