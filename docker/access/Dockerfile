FROM golang:latest

RUN mkdir /app
RUN mkdir /etc/access
COPY . /app
COPY ./access/config /etc/access

WORKDIR /app

RUN go build -o main access/cmd/main.go

CMD ["/app/main", "-init"]