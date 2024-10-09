package main

import (
	"github.com/SoftwareArch-BackstreetBoys/notification-service/notification"
	"github.com/SoftwareArch-BackstreetBoys/notification-service/queue"
)

func main() {
    go queue.ConsumeMessages(notification.NotifyEventUpdate)

    // Keep the service running
    select {}
}
