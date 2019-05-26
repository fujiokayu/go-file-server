FROM golang:latest

RUN mkdir /contents-server
COPY src/main.go /go/contents-server

RUN mkdir /static
COPY static/* /go/static/

CMD ["go"、"run", "/go/content-server/main.go"]
