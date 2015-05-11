FROM google/golang

WORKDIR /gopath/src/github.com/etcinit/cerulean/
ADD . /gopath/src/github.com/etcinit/cerulean/
RUN go get ./... && go build ./... && go install ./...

CMD []
ENTRYPOINT ["/gopath/bin/cerulean"]
