FROM golang:alpine AS build

WORKDIR /app

ENV GONOPROXY=github.com/SnackLog/*
ENV GOSUMDB=off

RUN apk add git

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o /app/seed-postgres-db

FROM alpine:latest

WORKDIR /app
RUN mkdir /data
COPY LICENSES /licenses
COPY --from=build /app/seed-postgres-db .

ENTRYPOINT ["./seed-postgres-db"]

