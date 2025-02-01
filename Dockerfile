FROM golang:alpine3.20 AS development
# Set the timezone to Bangkok
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" > /etc/timezone

# Set the TZ environment variable (optional but recommended)
ENV TZ=Asia/Bangkok

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN air init

COPY . .

EXPOSE 8081

CMD air