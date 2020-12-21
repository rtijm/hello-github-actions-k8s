# builder image
FROM golang:1.15-alpine as builder
RUN mkdir /build
ADD *.go /build/
ADD go.* /build/
WORKDIR /build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o hello-world .

# generate clean, final image for end users
FROM alpine:3
COPY --from=builder /build/hello-world .

# executable
ENTRYPOINT [ "./hello-world" ]