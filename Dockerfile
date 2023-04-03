# Start from golang base image
FROM golang:alpine as builder

LABEL maintainer="dochien0204"

RUN apk update && apk add --no-cache git

WORKDIR /app/crud-company

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["./main"]