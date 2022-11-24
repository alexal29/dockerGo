FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY ./cmd/*.go ./

COPY ./public ./public

RUN go build -o /main

CMD [ "/main" ]