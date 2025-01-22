# **Docker and Kubernetes mini-project**

## **Project Overview**

This project demonstrates the use of **Docker** and **Kubernetes** to deploy services for two different languages: **Python** and **Go**. The project includes:
1. A simple **Hello World Flask app** running inside a Docker container.
2. A **Ping-Pong** service where:
   - **Python** and **Go** both respond to the `/ping` endpoint with a "pong" message.
3. A **combined Docker setup** running both Python and Go services on different ports.
4. A **separate setup** with individual deployments for Python and Go services.

## **Hello World App (Python)**

The **Hello World** Flask app demonstrates a simple web application running inside a Docker container. The app listens for HTTP requests on port `5000` and returns a message `"Hello, Docker!"`. This example showcases how to build a basic Python web app and containerize it using Docker, followed by deploying it to a Kubernetes cluster.

1. **Flask Framework**: Flask is a lightweight web framework for Python, ideal for building simple web applications.
2. **Dockerization**: The app is containerized using a `Dockerfile`. The Dockerfile sets up a Python environment, installs dependencies, and runs the Flask app inside a container.
3. **Kubernetes Deployment**: The app is deployed on Kubernetes using the `deployment.yaml` and `service.yaml` files. The deployment ensures the app runs as a pod, while the service exposes it via a NodePort.

### **Setup and Usage**

#### **Prerequisites**
Before starting, ensure that you have Docker and Kubernetes enabled in Docker Desktop.

- **Docker Desktop** (with Kubernetes enabled) should be running on your machine.
  - [Install Docker Desktop](https://www.docker.com/products/docker-desktop)

#### **Step 1: Build the Docker Image**

In the `docker/hello-world-docker` directory, build the Docker image using the following command:

```bash
cd docker/hello-world-docker
docker build -t hello-world-docker .
```

This will create a Docker image with the tag `hello-world-docker`.

#### **Step 2: Run the Docker Container**

Run the container using the command:

```bash
docker run -p 5000:5000 hello-world-docker
```

This will expose the Flask app on port `5000` locally. You can access the application in your browser or using `curl`:

```bash
curl http://localhost:5000
```

You should receive a response:
```
Hello, Docker!
```
#### **Step 3: Kubernetes Deployment**

To deploy the app to a Kubernetes cluster, apply the Kubernetes deployment and service configuration:

```bash
kubectl apply -f docker/hello-world-docker/deployment.yaml
kubectl apply -f docker/hello-world-docker/service.yaml
```

This will deploy the Flask app to Kubernetes and expose it via a NodePort. You can access it using:

```bash
curl http://localhost:30001
```

---

## **Ping-Pong Service (Python & Go)**

### **Overview**

The **Ping-Pong service** provides two separate services:
1. A **Python Flask service** that responds with `"pong from Python"` to the `/ping` endpoint.
2. A **Go service** that responds with `"pong from Go"` to the `/ping` endpoint.

These services are built and containerized separately, using individual Dockerfiles for Python and Go, and then deployed to Kubernetes.

### **Separate Python & Go Deployment**

#### **Python (Ping-Pong)**

##### **Technical Explanation**

- **Flask Framework**: This Python app uses the Flask web framework to create a simple REST API.
- **Dockerization**: The Python service is containerized using a Dockerfile, which sets up a Python environment, installs the Flask dependency, and runs the app.
- **Kubernetes Deployment**: The app is deployed using Kubernetes configuration files (`deployment-python.yaml` and `service.yaml`) that ensure the service is available within the Kubernetes cluster.

##### **Step 1: Build the Docker Image**

In the `docker/ping-pong/ping-pong-python` directory, build the Docker image using:

```bash
cd docker/ping-pong/ping-pong-python
docker build -t ping-pong-python .
```

##### **Step 2: Run the Docker Container**

Run the Python container locally with the following command:

```bash
docker run -p 8081:8081 ping-pong-python
```

This exposes the Python service on port `8081`. You can test it using:

```bash
curl http://localhost:8081/ping
```

You should receive:
```
pong
```

##### **Step 3: Kubernetes Deployment**

Deploy the Python service to Kubernetes:

```bash
kubectl apply -f docker/ping-pong/ping-pong-python/deployment.yaml
kubectl apply -f docker/ping-pong/ping-pong-python/service.yaml
```

Access the service using:

```bash
curl http://<minikube-ip>:32183/ping
```

Where `<minikube-ip>` is your Minikube instance's IP.

#### **Go (Ping-Pong)**

##### **Technical Explanation**

- **Go HTTP Server**: This service is built using Go, and the program listens on port `8081` and responds with `"pong"` to the `/ping` endpoint.
- **Dockerization**: The Go service is containerized using a Dockerfile that sets up a Go environment, compiles the application, and runs it inside a container.
- **Kubernetes Deployment**: Similar to the Python service, the Go service is deployed using Kubernetes with separate deployment and service configurations.

##### **Step 1: Build the Docker Image**

In the `docker/ping-pong/ping-pong-go` directory, build the Docker image:

```bash
cd docker/ping-pong/ping-pong-go
docker build -t ping-pong-go .
```

##### **Step 2: Run the Docker Container**

Run the Go container locally using:

```bash
docker run -p 8081:8081 ping-pong-go
```

You can test it with:

```bash
curl http://localhost:8081/ping
```

You should receive:
```
pong
```

##### **Step 3: Kubernetes Deployment**

Deploy the Go service to Kubernetes:

```bash
kubectl apply -f docker/ping-pong/ping-pong-go/deployment.yaml
kubectl apply -f docker/ping-pong/ping-pong-go/service.yaml
```

Access the service using:

```bash
curl http://<minikube-ip>:3000/ping
```

---

## **Combined Service (Python & Go)**

### **Overview**

In the **combined service setup**, both the **Python** and **Go** services are containerized in a multi-stage Dockerfile and run on different ports (`3001` for Python and `3002` for Go`). This allows both services to run simultaneously in a single Docker container, which simplifies deployment.

### **Technical Explanation**

- **Multi-Stage Dockerfile**: The Dockerfile includes two stages: one for building the Python app and one for the Go app.
- **Separate Ports**: Each service runs on a different port inside the container, allowing them to coexist without conflict.
- **Kubernetes Deployment**: The service is deployed using Kubernetes configurations (`deployment.yaml` and `service.yaml`), ensuring both apps are accessible in the cluster.

### **Step 1: Build the Docker Image**

In the `docker/ping-pong/ping-pong-combined` directory, build the Docker images:

```bash
cd docker/ping-pong/ping-pong-combined
make build-python
make build-go
```

### **Step 2: Run the Docker Containers**

Run the combined service locally:

```bash
make run-python
make run-go
```

This will expose Python on port `3001` and Go on port `3002`.

### **Step 3: Kubernetes Deployment**

Deploy both services to Kubernetes:

```bash
kubectl apply -f docker/ping-pong/ping-pong-combined/deployment-python.yaml
kubectl apply -f docker/ping-pong/ping-pong-combined/service.yaml
```

### **Source Code Explanation**

- **Python Service**: Uses Flask to respond with `"pong from Python"` at the `/ping` endpoint.
- **Go Service**: Implements an HTTP server in Go that responds with `"pong from Go"` at the `/ping` endpoint.
