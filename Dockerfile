FROM golang:1.11

RUN go get golang.org/x/lint/golint

ENV GO111MODULE on

WORKDIR /go/src/rds

COPY . .

RUN make deps && make test && make install

ENTRYPOINT [ "rds" ]