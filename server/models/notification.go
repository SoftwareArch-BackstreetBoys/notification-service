package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoNotificationLog struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`     // MongoDB ObjectID
	NotificationType string            `bson:"notification_type"` // Type of notification (e.g., "email", "event_update", "event_deletion", "event_join")
	Sender          string             `bson:"sender"`            // Sender's email or ID
	Receiver        string             `bson:"receiver"`          // Receiver's email or ID
	Subject         string             `bson:"subject,omitempty"` // Subject for email notifications (optional)
	BodyMessage     string             `bson:"body_message,omitempty"`// Message body (optional)
	EventId         string             `bson:"event_id,omitempty"`// ID of the event related to the notification (optional)
	Status          string             `bson:"status"`            // Delivery status (e.g., "sent", "failed")
	Timestamp       time.Time          `bson:"timestamp"`         // Time of notification
	ErrorMessage    string             `bson:"error_message,omitempty"` // Error message if notification fails (optional)
}

type NotificationMessage struct {
	NotificationType string 					`json:"notification_type"`
	Sender           string 					`json:"sender"`
	Receiver         string 					`json:"receiver"`
	Subject          string 					`json:"subject"`
	BodyMessage      string 					`json:"body_message"`
	Status					 string 					`json:"status"`
}

