# Uncomment code to build application
# FROM golang:1.16-alpine as builder

# WORKDIR /project/go-docker/

# # COPY go.mod, go.sum and download the dependencies
# COPY go.* ./
# RUN go mod download

# # COPY All things inside the project and build
# COPY . .
# RUN go build -o /project/go-docker/build/myapp .

FROM scratch
# COPY --from=builder /project/go-docker/build/myapp /project/go-docker/build/myapp
COPY goapp/todoapp .

EXPOSE 8080
ENTRYPOINT [ "/todoapp" ]