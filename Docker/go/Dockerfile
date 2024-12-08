FROM golang:1.23.3-alpine

RUN mkdir /app

WORKDIR /app

RUN apk add build-base
RUN go install github.com/air-verse/air@latest

ADD ../../go.mod .

RUN go mod download
ADD ../.. .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]