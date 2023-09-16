FROM golang:1.13.5

ENV GODEBUG=netdns=cgo
ENV GOMAXPROCS=2

WORKDIR /app

COPY . .

RUN go build -o go-interpreter-rinha

CMD ["./go-interpreter-rinha"]