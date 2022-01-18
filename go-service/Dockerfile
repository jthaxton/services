# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

# ### get migrate
# RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
# RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
# RUN apt-get update
# RUN apt-get install -y migrate
# ###

COPY . ./
RUN go mod download
RUN go install
RUN file="$(ls)" && echo $file
RUN echo $(ls)
RUN apk update
RUN apk add make
# RUN apk add golang-migrate
# RUN make migrateup


RUN go build -o /docker-gs-ping


EXPOSE 8080

CMD [ "/docker-gs-ping" ]