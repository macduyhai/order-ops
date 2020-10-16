FROM golang:1.14 as builder

WORKDIR /app_order

COPY . /app_order

RUN go mod download

RUN GOOS=linux

RUN echo "Asia/Ho_Chi_Minh" > /etc/timezone

RUN dpkg-reconfigure -f noninteractive tzdata

RUN go build -o main

FROM ubuntu:16.04

WORKDIR /app_order

COPY --from=builder /app_order/main .

EXPOSE 80

CMD ["/app_order/main"]