FROM golang:latest

WORKDIR /introduction-bot
COPY . /introduction-bot
RUN go build .

EXPOSE 3000
ENTRYPOINT ["/introduction-bot/introduction-bot"]
