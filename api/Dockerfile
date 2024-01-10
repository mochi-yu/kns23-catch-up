FROM golang:1.21-alpine as build

WORKDIR /app

# Copy dependencies list
COPY go.mod go.sum ./

# Build
COPY main.go .
RUN go build -o main main.go

# Copy artifacts to a clean image
FROM alpine:3.19
COPY --from=build /app/main /main
ENTRYPOINT [ "/main" ]
