# FROM golang:1.15.2
FROM alpine:latest

WORKDIR /opt/

COPY src/ .

RUN chmod +x cani

CMD ["/bin/sh", "-c", "./cani"]
