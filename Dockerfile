# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN go build -o /mailer_api

FROM gcr.io/distroless/base-debian11
LABEL maintainer="Erick Amorim <github.com/ericklima-ca>"
COPY --from=builder /mailer_api /mailer_api
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/mailer_api"]
