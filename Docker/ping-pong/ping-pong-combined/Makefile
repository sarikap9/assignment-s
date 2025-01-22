# Define Docker image names
PYTHON_IMAGE=ping-pong-python
GO_IMAGE=ping-pong-go

# Build Python Docker image
build-python:
	docker build -f Dockerfile -t $(PYTHON_IMAGE) --target python-build .

# Build Go Docker image
build-go:
	docker build -f Dockerfile -t $(GO_IMAGE) --target go-build .

# Run Python Docker container
run-python:
	docker run -p 3001:3001 $(PYTHON_IMAGE)

# Run Go Docker container
run-go:
	docker run -p 3002:3002 $(GO_IMAGE)

# Clean up Docker containers and images
clean:
	docker rm -f $(PYTHON_IMAGE) $(GO_IMAGE) || true
	docker rmi $(PYTHON_IMAGE) $(GO_IMAGE) || true
