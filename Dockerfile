FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /build/ ./...

CMD ["/build/app"]
