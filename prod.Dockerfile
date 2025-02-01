FROM golang:alpine3.20 AS builder

ENV GOOS=linux
ENV CGO_ENABLE=0
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

# go build app
# RUN go build -o api-server cmd/main.go
RUN go build -o api-server .

FROM alpine:3.14 AS production
RUN apk add --no-cach ca-certificates
RUN apk --no-cache add tzdata

ENV GIN_MODE=release
ENV APP_HOME=/app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
RUN pwd

COPY --from=builder /app/api-server app-server
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip

RUN ["chmod","+x","app-server"]
EXPOSE 8081
CMD ["./app-server"]
