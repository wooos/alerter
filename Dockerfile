FROM debian:stretch-slim

LABEL maintainer="wooos <819844849@qq.com>"

WORKDIR /app

ADD alerter /app/

ENTRYPOINT ["/app/alerter"]