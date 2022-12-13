# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . .

RUN go build -o /apiGateway

EXPOSE 2321

CMD [ "/apiGateway" ]