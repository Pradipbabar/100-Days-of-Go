FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY main.go ./

RUN go build -o /my-app

FROM alpine:3.16

COPY --from=build /my-app /my-app

ENTRYPOINT ["/my-app"]