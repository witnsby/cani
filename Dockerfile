# FROM golang:1.15.2
FROM alpine:latest

WORKDIR /opt/

COPY cani .
RUN chmod +x cani

CMD ["/bin/sh", "-c", "./cani"]
