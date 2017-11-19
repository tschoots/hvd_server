#FROM ubuntu
FROM scratch

MAINTAINER Ton Schoots <ton@maiastra.com>

EXPOSE 8080

WORKDIR /
COPY a-frame_web_server /
COPY a-frame /a-frame
COPY templates /templates

ENTRYPOINT ["./a-frame_web_server"]
