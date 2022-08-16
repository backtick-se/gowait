FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app
RUN go build -o /bin/cowait ./cmd/client

FROM alpine:latest as client

COPY --from=builder /bin/cowait /usr/local/bin/cowait
ENTRYPOINT ["/usr/local/bin/cowait"]