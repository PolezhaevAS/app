FROM golang:latest

RUN mkdir /app
RUN mkdir /etc/auth
COPY . /app
COPY ./auth/config /etc/auth

WORKDIR /app

RUN go build -o main auth/cmd/main.go

CMD ["/app/main"]