FROM node:latest

COPY package.json package-lock.json ./
RUN npm install

COPY --from=cowait/gowait /bin/cowait /bin/cowait
ENTRYPOINT [ "cowait" ]

WORKDIR /task
COPY . .