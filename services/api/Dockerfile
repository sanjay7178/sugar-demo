FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api /app/
EXPOSE 8081

CMD ["/app/api"]
