FROM debian:stretch-slim

LABEL maintainer="wooos <819844849@qq.com>"

RUN apt-get update \
    && apt-get install -y ca-certificates

WORKDIR /app

ADD bin/alerter /app/

ENTRYPOINT ["/app/alerter"]