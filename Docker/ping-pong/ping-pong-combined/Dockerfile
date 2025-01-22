# Python build stage
FROM python:3.9-slim as python-build
WORKDIR /app
COPY app-python.py .
RUN pip install flask
CMD ["python", "app-python.py"]

# Go build stage
FROM golang:1.19 as go-build
WORKDIR /app
COPY app-go.go .
RUN GO111MODULE=off go build -o app-go .
CMD ["./app-go"]
