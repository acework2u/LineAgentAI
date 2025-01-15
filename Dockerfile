FROM golang:alpine3.20 AS development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod downloa

RUN go install github.com/air-verse/air@latest
RUN air init

COPY . .

EXPOSE 8180

CMD air