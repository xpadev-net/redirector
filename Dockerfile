FROM golang:1.22.0 as builder
WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine:3.12.0
WORKDIR /app
LABEL org.opencontainers.image.source=https://github.com/xpadev-net/redirector
COPY --from=builder /build/main /app/main

CMD ["./main"]