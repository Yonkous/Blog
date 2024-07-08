# Build aşaması
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /blog

# Run aşaması
FROM alpine:latest

WORKDIR /

COPY --from=builder /blog /blog

EXPOSE 8081

CMD ["/blog"]
