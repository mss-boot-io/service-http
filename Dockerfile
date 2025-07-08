FROM alpine

WORKDIR /app

COPY ./application /app/application

ENTRYPOINT ["/app/application"]