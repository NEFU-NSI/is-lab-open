FROM golang:alpine
LABEL authors="sfc9982"

WORKDIR /app

COPY . .

RUN apk add --no-cache iputils

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /islabopen

EXPOSE 8080

CMD ["/islabopen"]