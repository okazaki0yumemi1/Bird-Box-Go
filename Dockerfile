FROM golang:alpine

WORKDIR /bird-box-go
COPY . .

RUN apk add alsa-utils
RUN go build -o ./bin/api ./cmd/api && go build -o ./bin/migrate ./cmd/migrate

CMD ["/bird-box-go/bin/api"]
EXPOSE 8080