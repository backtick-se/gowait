FROM python:3.9-slim

RUN pip install grpcio protobuf

COPY --from=cowait/gowait /bin/cowait /bin/cowait
ENTRYPOINT [ "cowait" ]

WORKDIR /task
COPY . .
