FROM alpine:3.3

RUN apk add --update bash && rm -rf /var/cache/apk/*

EXPOSE 8080
CMD ./user-ws

COPY bin/user-ws.linux user-ws
