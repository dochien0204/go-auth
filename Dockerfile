# Start from golang base image
FROM golang:alpine as builder

LABEL maintainer="dochien0204"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["./main"]