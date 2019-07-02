FROM depscloud/base:latest

ARG VERSION=0.0.6

RUN install-depscloud-binary discovery ${VERSION}

RUN useradd -ms /bin/sh discovery
WORKDIR /home/discovery
USER discovery

ENTRYPOINT [ "discovery" ]
