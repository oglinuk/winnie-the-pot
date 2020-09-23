FROM golang:1.15.0 as base
ADD . /go/src/winnie-the-pot
WORKDIR /go/src/winnie-the-pot
RUN go get
RUN make

FROM ubuntu:20.04
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -y openssh-server && apt upgrade -y
RUN mkdir ~/.ssh
RUN ssh-keygen -q -t rsa -N '' -f ~/.ssh/id_rsa
COPY --from=base /go/src/winnie-the-pot/winnie-the-pot ./
EXPOSE 22
CMD ["./winnie-the-pot"]