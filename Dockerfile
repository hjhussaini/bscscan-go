FROM golang:1.16.4-alpine

WORKDIR /app
COPY . /app

RUN go build -o bscscan-client

ENV PORT=8080
env API_KEY=VDFU1WV8A7M94KERKB8IW9E6GZY3X6BJEX

EXPOSE $PORT

CMD ["/app/bscscan-client"]
