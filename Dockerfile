FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

EXPOSE 8000

RUN go mod download

RUN go build -o mock-serpro

CMD ["./mock-serpro", "-http.addr=0.0.0.0:8000"]