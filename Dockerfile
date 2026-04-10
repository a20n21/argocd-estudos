FROM golang:1.25-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

# 👇 ISSO RESOLVE TUDO
RUN go mod init argocd-estudos || true
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o server .

EXPOSE 8080
CMD ["./server"]