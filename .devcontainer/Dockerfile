FROM docker.io/library/alpine:3.21.2

ENV TZ Asia/Tokyo
ENV EDITOR vi
ENV STAGE dev
ENV PS1 "[\A]\u@hooka:\W\\$ "
ENV GOPATH /usr/local/webhook
ENV PATH ${GOPATH}/bin:${PATH}

RUN apk update && \
    apk add doas vim tzdata go && \
    rm -f /var/cache/apk/* && \
    adduser -u 2000 -D -h /home/dev -G users dev && \
    addgroup dev wheel && \
    echo -e "permit nopass :root¥npermit nopass :wheel¥n" > /etc/doas.d/doas.conf && \
    mkdir -p "${GOPATH}/pkg" "${GOPATH}/bin" && \
    chown -R dev:users "${GOPATH}/pkg" "${GOPATH}/bin"

USER dev
RUN go install golang.org/x/tools/cmd/goimports@latest
WORKDIR ${GOPATH}/src

