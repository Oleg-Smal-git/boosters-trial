# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /boosters-trial
ENV PROJECT_PATH='/boosters-trial'
ENV ENV='docker'

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN apk add --update make
RUN make build_all

EXPOSE 8080

CMD [ "/boosters-trial/bin/run-api" ]