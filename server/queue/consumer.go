package queue

import (
	"log"

	"github.com/SoftwareArch-BackstreetBoys/notification-service/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessages(handleMessage func(*amqp.Delivery)) {
    // Connect to RabbitMQ
    ch, err := configs.ConnectRabbitMQ()

    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    // Declare the queue (same name as the publisher uses)
    q, err := ch.QueueDeclare(
        "event_notifications", // Queue name
        true,                  // Durable
        false,                 // Delete when unused
        false,                 // Exclusive
        false,                 // No-wait
        nil,                   // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    msgs, err := ch.Consume(
        q.Name,    // queue name
        "",        // consumer
        true,      // auto-ack
        false,     // exclusive
        false,     // no-local
        false,     // no-wait
        nil,       // args
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %s", err)
    }

    go func() {
        for d := range msgs {
            handleMessage(&d)
            // Acknowledge the message if processing is successful
        }
    }()
}
