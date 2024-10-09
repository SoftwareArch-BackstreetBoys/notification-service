package logger

import (
	"context"
	"log"
	"time"

	"github.com/SoftwareArch-BackstreetBoys/notification-service/configs"
	"github.com/SoftwareArch-BackstreetBoys/notification-service/models"
)

func LogNotification(notification *models.MongoNotificationLog) {
	collection := configs.GetCollection(configs.DB, "logs")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notification.Timestamp = time.Now()

	_, err := collection.InsertOne(ctx, notification)
	if err != nil {
			log.Fatalf("Failed to log notification: %s", err)
	}else{
		log.Println("Notification logged successfully")
	}

}
