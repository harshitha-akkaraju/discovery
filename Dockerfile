FROM depscloud/base:latest

ARG VERSION=0.0.6

RUN install-depscloud-binary rds ${VERSION}

ENTRYPOINT [ "rds" ]
