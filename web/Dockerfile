#FROM ubuntu
FROM scratch

MAINTAINER Ton Schoots <ton@maiastra.com>

EXPOSE 8080

WORKDIR /
COPY hvd_web_server /
COPY html /html
#COPY templates /templates

ENTRYPOINT ["./hvd_web_server"]
