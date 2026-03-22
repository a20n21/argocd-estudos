FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o server .

EXPOSE 8080

CMD ["./server"]