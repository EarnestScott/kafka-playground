FROM golang:1.19-bullseye

WORKDIR /app

COPY . .

RUN go build -o event-whisperer

CMD ["./event-whisperer"]
