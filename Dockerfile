FROM alpine:latest

MAINTAINER Raka Westu Mogandhi <rakamogandhi@hotmail.com>

WORKDIR "/opt"

ADD .docker_build/rakawm-snap /opt/bin/rakawm-snap
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/rakawm-snap"]
