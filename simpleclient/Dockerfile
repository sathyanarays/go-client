FROM golang:1.17.0-alpine3.13

WORKDIR /src

COPY go.mod .

COPY go.sum .

COPY *.go .

RUN go build main.go

RUN chmod +x /src/main

CMD ["/src/main"]
