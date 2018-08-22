FROM golang:rc-alpine3.7

ENV STATUSOK_VERSION 0.2

RUN wget https://github.com/lmorel3/statusok/releases/download/$STATUSOK_VERSION/statusok \
    && mv ./statusok /go/bin/StatusOk \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

VOLUME /config
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
ENTRYPOINT /docker-entrypoint.sh
