FROM golang:1.15.6-buster

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

CMD ["/build/main"]