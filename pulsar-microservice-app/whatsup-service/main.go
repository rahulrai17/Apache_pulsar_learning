// This service consumes from world-topic, appends "Whatsup", and sends the final message to final-topic.

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

    // Create a consumer for world-topic
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "world-topic",
        SubscriptionName: "whatsup-subscription",
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

    // Append the word
    finalMessage := string(msg.Payload()) + " Whatsup"

    // Publish final message to final-topic
    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: "final-topic",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar producer: %v", err)
    }
    defer producer.Close()

    _, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
        Payload: []byte(finalMessage),
    })
    if err != nil {
        log.Fatalf("Could not send message: %v", err)
    }
    fmt.Println("Published:", finalMessage)

    consumer.Ack(msg)
}
