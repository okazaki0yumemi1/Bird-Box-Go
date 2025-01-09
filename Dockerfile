FROM golang:alpine

WORKDIR /Bird-Box-Go
COPY . .

RUN go build cmd/api/main.go

CMD ["/Bird-Box-Go/main"]
EXPOSE 8080