FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY *.go ./
COPY index.html ./
COPY images/ ./images/

RUN go build -o monkeylines .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/monkeylines .

EXPOSE 8080 8023

CMD ["./monkeylines"]
