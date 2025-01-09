FROM golang:alpine

WORKDIR /Bird-Box-Go
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/Bird-Box-Go/bin/api"]
EXPOSE 8080