FROM alpine:latest

WORKDIR /app/
COPY ./bin/crudsample .
ENTRYPOINT ["/app/crudsample"]

