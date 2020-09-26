FROM golang:1.14 as builder

ENV APP_USER app
ENV APP_HOME /go/src/imcargo

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY ./ .

RUN go mod download
RUN go mod verify
RUN go build -o imcargo

FROM debian:buster

ENV APP_USER app
ENV APP_HOME /go/src/imcargo

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

RUN mkdir static && cd static && touch index.html && mkdir avatar && cd avatar && touch index.html
COPY --chown=0:0 --from=builder $APP_HOME/imcargo $APP_HOME

EXPOSE 8080
CMD ["./imcargo"]