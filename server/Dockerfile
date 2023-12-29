FROM golang:latest

COPY . /app
WORKDIR /app/cmd/sushi-script

RUN go mod download

RUN go build -o ./server

EXPOSE 8080

CMD ["./server"]