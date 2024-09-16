// The web service consumes from final-topic and displays the final message on the webpage.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/apache/pulsar-client-go/pulsar"
)

var finalMessage string

func main() {
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar client: %v", err)
    }
    defer client.Close()

    // Create a consumer for final-topic
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "final-topic",
        SubscriptionName: "web-subscription",
    })
    if err != nil {
        log.Fatalf("Could not create Pulsar consumer: %v", err)
    }
    defer consumer.Close()

    // Start an HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if finalMessage == "" {
            msg, err := consumer.Receive(context.Background())
            if err != nil {
                log.Fatalf("Could not receive message: %v", err)
            }
            finalMessage = string(msg.Payload())
            consumer.Ack(msg)
        }

        fmt.Fprintf(w, "Final Message: %s", finalMessage)
    })

    fmt.Println("Web server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
