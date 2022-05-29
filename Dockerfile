FROM golang:1.17-buster

ENV GO111MODULE=on

WORKDIR /go/src

COPY ./src /go/src

RUN apt-get update \
 && apt-get install -y vim \
 && go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
