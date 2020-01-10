FROM golang

ADD . /go/src/godumper

WORKDIR /go/src/godumper

RUN go get ./...

RUN go install godumper

ENV DATABASE_DB "db_name"
ENV DATABASE_HOSTNAME "hostname"
ENV DATABASE_PASSWORD "password"
ENV DATABASE_PORT "3306"
ENV DATABASE_USER "database_user"

ENV TELEGRAM_BOTKEY "botkey"
ENV TELEGRAM_CHATID 0000

ENTRYPOINT /go/bin/godumper

EXPOSE 8080