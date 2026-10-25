FROM golang:1.26

WORKDIR /app

COPY . .

RUN go build -o server ./cmd/server

CMD ["./server"]