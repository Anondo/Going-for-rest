

FROM golang:alpine

RUN apk update && apk add git 

WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV GOPATH=/go
# Add the source code:
ADD . $GOPATH/src/

RUN go get -u github.com/julienschmidt/httprouter
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-sql-driver/mysql

# Build it:
RUN cd $GOPATH/src; go build -o main; cp main /app/

ENTRYPOINT ["./main"]