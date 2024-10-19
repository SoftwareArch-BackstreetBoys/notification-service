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

### 2. Set up RabbitMQ Server

The Notification Service relies on a RabbitMQ instance for message brokering. You will need to create your own `docker-compose.yml` file for RabbitMQ. Below is an example configuration:

```yaml
version: "3.8"

networks:
  shared-network:
    external: true

services:
  rabbitmq:
    image: rabbitmq:3-management
    container_name: shared-rabbitmq
    ports:
      - "5672:5672" # AMQP port for producer/consumer
      - "15672:15672" # Management UI
    networks:
      - shared-network

Next, run the RabbitMQ container:
   ```bash
   docker-compose -f rabbitmq.yml up -d

You can verify that RabbitMQ is running by accessing the management UI in your browser at http://localhost:15672 (login with guest:guest).

### Step 3: Run Notification Service

```markdown
### 3. Run Notification Service

Once RabbitMQ is up and running, follow these steps to start the Notification Service:

1. Ensure that RabbitMQ is running and is accessible on the `shared-network`.

2. Run the Notification Service using Docker Compose:
   ```bash
   docker-compose -f notification-service.yml up -d


