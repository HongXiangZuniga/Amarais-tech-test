FROM golang:1.22.1 AS builder
WORKDIR /app
COPY . .
RUN make install && go build ./cmd/server/api.go
EXPOSE 8080
CMD ["./api"]

