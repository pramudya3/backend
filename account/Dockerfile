FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 2000

CMD ["./main"]
