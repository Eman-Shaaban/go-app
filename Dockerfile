FROM alpine

RUN mkdir /app

ADD main /app

WORKDIR /app

ENTRYPOINT ["/app/main"]
