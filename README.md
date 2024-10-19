# Notification Service

This service listens for events from a RabbitMQ message broker and processes notifications. Follow the steps below to set up the environment and run the service.

## Prerequisites

Before running the notification service, ensure you have the following installed on your system:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Steps to Run

### 1. Create a Shared Docker Network

Both RabbitMQ and the Notification Service need to communicate over a shared Docker network. Before starting the containers, create this network:

```bash
docker network create shared-network
```
