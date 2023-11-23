# Chat Application

## Overview

This is a web chat application built using React for the frontend, Golang for the backend, and WebSocket for real-time communication.

## Features

- Real-time chat functionality
- Frontend built with React
- Backend server in Golang using WebSocket for communication

## Prerequisites

Make sure you have the following tools installed:

- Docker
- Docker Compose
- Minikube (for Kubernetes deployment)
- kubectl (for Kubernetes deployment)
- Jenkins (for CI/CD)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/chat-application.git
cd chat-application
```

### 2. Run Locally with Docker Compose

```bash
docker-compose up
```

Visit `http://localhost:3000` in your web browser to access the chat application.

### 3. CI/CD with Jenkins

1. Create a Jenkins pipeline job.
2. Configure the pipeline to use the provided `Jenkinsfile` in the repository.
3. Set up Jenkins credentials for Docker Hub (used in the Jenkinsfile).

### 4. Deploy to Kubernetes

Ensure Minikube is running:

```bash
minikube start
```

Deploy to Kubernetes:

```bash
kubectl apply -f Kubernetes/deployment.yaml -f Kubernetes/service.yaml
```

Access the application using the Minikube IP and NodePort specified in the service.

## Directory Structure

- `frontend/`: React frontend code
- `backend/`: Golang WebSocket server code
- `docker-compose.yaml`: Docker Compose file for local development
- `Jenkinsfile`: Jenkins pipeline configuration
- `Kubernetes/`: Kubernetes deployment and service configurations
