package main

import (
	"fmt"
	"net/http"

	"github.com/SoftwareArch-BackstreetBoys/notification-service/notification"
	"github.com/SoftwareArch-BackstreetBoys/notification-service/queue"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Healthy")
}

func main() {
    go queue.ConsumeMessages(notification.NotifyEventUpdate)

    // Start HTTP server for health checks
    http.HandleFunc("/health", healthHandler)

    // Run health check server in a separate goroutine
    go func() {
        if err := http.ListenAndServe(":8082", nil); err != nil {
            fmt.Println("Error starting HTTP server:", err)
        }
    }()

    // Keep the service running
    select {}
}
