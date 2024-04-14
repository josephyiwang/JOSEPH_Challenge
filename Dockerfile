FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN Export GO111MODULE=on
RUN go get github.com/josephyiwang/JOSEPH_Challenge
RUN cd /build && git clone https://github.com/josephyiwang/JOSEPH_Challenge.git

RUN cd /build/JOSEPH_Challenge && go build

EXPOSE 8080
EXPOSE 443

ENTRYPOINT [ "/build/JOSEPH_Challenge/main" ]