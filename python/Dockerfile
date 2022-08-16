FROM python:3.9-slim

COPY . .

COPY --from=cowait/gowait /usr/local/bin/cowait /usr/local/bin/cowait
ENTRYPOINT [ "/usr/local/bin/cowait" ]