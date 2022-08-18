FROM python:3.9-slim

RUN pip install grpcio protobuf

COPY . .

COPY --from=cowait/gowait /usr/local/bin/cowait /usr/local/bin/cowait
ENTRYPOINT [ "/usr/local/bin/cowait" ]