FROM golang:1.21-alpine as pattern
WORKDIR /pattern

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o pattern ./cmd/main.go

EXPOSE 8000
CMD ["./pattern"]