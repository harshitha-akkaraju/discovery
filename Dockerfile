FROM depscloud/base:latest

ARG VERSION=0.0.6

RUN install-depscloud-binary rds ${VERSION}

RUN useradd -ms /bin/sh rds
WORKDIR /home/rds
USER rds

ENTRYPOINT [ "rds" ]
