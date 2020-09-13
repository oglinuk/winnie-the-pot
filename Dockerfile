FROM golang:1.15.0
ADD . /go/src/winnie-the-pot
WORKDIR /go/src/winnie-the-pot
RUN go get
RUN go build
EXPOSE 22
CMD ["./winnie-the-pot"]