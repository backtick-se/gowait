FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY ./vendor /app/vendor
RUN go build ./vendor/...

COPY . /app

FROM builder as build-executor
RUN go build -o /bin/cowait ./cmd/executor

FROM builder as build-daemon
RUN go build -o /bin/cowaitd ./cmd/daemon

FROM debian:stable-slim as executor

COPY --from=build-executor /bin/cowait /usr/local/bin/cowait
ENTRYPOINT ["/usr/local/bin/cowait"]

FROM debian:stable-slim as daemon

COPY --from=build-daemon /bin/cowaitd /bin/cowaitd
ENTRYPOINT ["cowaitd"]