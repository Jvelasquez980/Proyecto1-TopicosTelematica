FROM golang:1.22

WORKDIR /app/mom-service

COPY ./mom-service /app/mom-service
COPY ./orders-service /app/orders-service

RUN go mod tidy
RUN go build -o service .

CMD ["./service"]
