VERSION 0.7
FROM golang:1.22.0-alpine3.19
WORKDIR /rinha

test:
	COPY go.mod ./
	COPY main.go ./
	RUN go test

build:
	COPY go.mod ./
	COPY main.go ./
	RUN go build

docker:
	FROM DOCKERFILE .
	ARG GITHUB_REPO=cciuenf/rinha-backend-2024-q1
	SAVE IMAGE --push ghcr.io/$GITHUB_REPO:latest
