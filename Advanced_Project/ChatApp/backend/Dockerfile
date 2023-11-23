# Stage 1: Build stage
FROM golang:alpine3.16 AS builder

RUN mkdir /app
ADD ./backend /app
WORKDIR /app

# Build the application
RUN go build -o main .

# Stage 2: Final stage
FROM alpine:3.16

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

EXPOSE 9000

CMD ["/app/main"]


