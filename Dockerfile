FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download  || true

COPY . .

RUN go build -o main

CMD ["./main"]