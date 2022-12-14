FROM golang:1.18-alpine

WORKDIR /app

EXPOSE 3070

COPY . /app/

RUN apk add alpine-sdk
RUN go mod tidy
RUN go build -a -o ./main ./cmd/service

CMD ["./main", "--config-file", "./build/app.yaml"]