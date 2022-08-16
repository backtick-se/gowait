FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app

FROM builder as daemonb
RUN go build -o /bin/daemon ./cmd/daemon

FROM builder as clientb
RUN go build -o /bin/cowait ./cmd/client

FROM alpine:latest as client

COPY --from=clientb /bin/cowait /usr/local/bin/cowait
ENTRYPOINT ["/usr/local/bin/cowait"]

FROM alpine:latest as daemon

COPY --from=daemonb /bin/daemon /usr/local/bin/daemon
ENTRYPOINT ["/usr/local/bin/daemon"]