FROM golang:latest
LABEL authors="sfc9982"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /islabopen

EXPOSE 8080

CMD [ "/islabopen" ]