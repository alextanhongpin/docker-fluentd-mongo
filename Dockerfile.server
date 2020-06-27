FROM golang:1.14-stretch

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY main.go main.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app .
CMD ["./app"]
