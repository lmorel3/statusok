FROM golang:alpine

ENV STATUSOK_VERSION 0.2

RUN wget https://github.com/lmorel3/statusok/releases/download/$STATUSOK_VERSION/statusok \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
    && chmod 755 statusok && mv statusok /bin

VOLUME /config

CMD ["/bin/statusok", "--config", "/config/config.json"]
