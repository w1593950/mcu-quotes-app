# Build stage
FROM golang:latest AS builder

ENV GOOS=linux
ENV GOARCH=arm64
ENV GO111MODULE='on'

WORKDIR /app
COPY . .
RUN go build -o mcu-quotes-app .

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/mcu-quotes-app /app

EXPOSE 8080
CMD ["./mcu-quotes-app"]

