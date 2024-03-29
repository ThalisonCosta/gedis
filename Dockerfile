# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gedis

EXPOSE 6379

CMD ["/gedis"]
