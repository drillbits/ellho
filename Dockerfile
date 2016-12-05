FROM       golang:1.7.4-alpine
MAINTAINER drillbits <neji@drillbits.jp>

RUN apk --update add git

RUN go get github.com/drillbits/ellho

WORKDIR $GOPATH/src/github.com/drillbits/ellho

RUN go install

EXPOSE 5000

CMD ["ellho", "-p", "5000"]