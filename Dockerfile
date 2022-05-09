# syntax=docker/dockerfile:1
############## Builder Image ################
FROM golang:1.18-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /app/dummy-gin-api

############## Target Image ################
FROM golang:1.18-alpine
WORKDIR /app
COPY --from=builder /app/dummy-gin-api /app/dummy-gin-api

### just for testing here, not needed in Dockerfile, would be defined in kubernetes deployment
# EXPOSE 16489
# CMD [ "/app/dummy-gin-api" ]