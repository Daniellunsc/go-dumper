FROM golang

ADD . /go/src/godumper

WORKDIR /go/src/godumper

RUN go get ./...

RUN go install godumper

ENTRYPOINT /go/bin/godumper

EXPOSE 8080