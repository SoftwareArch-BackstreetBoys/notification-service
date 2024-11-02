package notification

import (
	"encoding/json"
	"log"
	"time"

	"github.com/SoftwareArch-BackstreetBoys/notification-service/email"
	"github.com/SoftwareArch-BackstreetBoys/notification-service/logger"
	"github.com/SoftwareArch-BackstreetBoys/notification-service/models"
	"github.com/SoftwareArch-BackstreetBoys/notification-service/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NotifyEventUpdate(d *amqp.Delivery) {
	// Handle event update logic
	log.Println("Sending test email...")

	var notificationMsg models.NotificationMessage
	err := json.Unmarshal(d.Body, &notificationMsg)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	log.Printf("Received notification: %+v", notificationMsg)

	notificationLog := models.MongoNotificationLog{
		NotificationType: notificationMsg.NotificationType,
		Sender:           notificationMsg.Sender,
		Receiver:         notificationMsg.Receiver,
		Subject:          notificationMsg.Subject,
		BodyMessage:      notificationMsg.BodyMessage,
		Status:           notificationMsg.Status,
		Timestamp:        time.Now(), // Add current time as the timestamp
	}

	receiverEmail := util.GetUserEmailById(notificationMsg.Receiver)

	err = email.SendEmail(
		notificationLog.Sender,
		receiverEmail,
		notificationLog.Subject,
		notificationLog.BodyMessage,
	)

	if err != nil {
		notificationLog.Status = "failed"
		notificationLog.ErrorMessage = err.Error()
	} else {
		notificationLog.Status = "success"
	}

	logger.LogNotification(&notificationLog)

	log.Println("Processing event update")
	// Log the event update
}
