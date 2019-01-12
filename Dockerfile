FROM golang:alpine

RUN apk add --no-cache --update git

ENV GOPATH=/go
# Add the source code:
ADD . $GOPATH/src/gorest/

RUN go get -u github.com/golang/dep/cmd/dep

RUN cd $GOPATH/src/gorest && dep init && dep ensure -v

RUN go install -v $GOPATH/src/gorest

WORKDIR $GOPATH/src/gorest

ENTRYPOINT ["gorest"]
