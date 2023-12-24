FROM golang:1.21.5-alpine3.19

WORKDIR /app

COPY . .

RUN go build -o bin/server cmd/server/main.go

EXPOSE 8080

CMD ["./bin/server"]