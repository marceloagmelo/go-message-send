FROM golang:1.13.6 AS builder

LABEL maintainer="Marcelo Melo marceloagmelo@gmail.com"


ENV APP_HOME /go/src/github.com/marceloagmelo/go-message-send

ADD . $APP_HOME

WORKDIR $APP_HOME

 RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go-message-send && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

###
# IMAGEM FINAL
###
FROM alpine:3.11

ENV GID 23550
ENV UID 23550
ENV USER golang

ENV APP_BUILDER /go/src/github.com/marceloagmelo/go-message-send/
ENV APP_HOME /opt/app

WORKDIR $APP_HOME

COPY --from=builder $APP_BUILDER/go-message-send $APP_HOME/go-message-send
COPY docker-container-start.sh $APP_HOME

RUN apk add --no-cache ca-certificates tzdata bash && \
    addgroup -g $GID -S $USER && \
    adduser -u $UID -S -h "$(pwd)" -G $USER $USER && \
    chown -fR $USER:0 $APP_HOME

ENV PATH $APP_HOME:$PATH

EXPOSE 8080

USER ${USER}

ENTRYPOINT [ "./docker-container-start.sh" ]
CMD [ "go-message-send" ]
