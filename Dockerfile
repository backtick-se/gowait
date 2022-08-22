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

FROM builder as build-cloud
RUN go build -o /bin/cloud ./cmd/cloud

#
# executor binary
#

FROM debian:stable-slim as executor
EXPOSE 1337

COPY --from=build-executor /bin/cowait /bin/cowait
ENTRYPOINT ["cowait"]

#
# daemon binary
#

FROM debian:stable-slim as daemon
EXPOSE 1337

COPY --from=build-daemon /bin/cowaitd /bin/cowaitd
ENTRYPOINT ["cowaitd"]

#
# cloud binary
#

FROM debian:stable-slim as cloud
EXPOSE 1338

COPY --from=build-cloud /bin/cloud /bin/cloud
ENTRYPOINT ["cloud"]
