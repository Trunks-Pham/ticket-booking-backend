FROM golang:1.22.2-alpine3.19

WORKDIR /src/app

# Cài đặt air từ module mới
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy
