FROM golang:1.22.0-alpine3.19 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . /app

RUN go get -d -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build the binary.
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o rinha .
RUN chmod +x ./rinha

FROM alpine:3.19

ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

RUN apk add --no-cache ca-certificates

COPY --from=builder --chown=nobody:root /app/rinha /usr/bin/rinha

CMD ["/usr/bin/rinha"]
