FROM golang:latest AS build

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN cd /build && git clone https://github.com/josephyiwang/JOSEPH_Challenge.git
RUN cd /build/JOSEPH_Challenge && go mod tidy && go build -o /build/main

FROM ubuntu:latest

WORKDIR /app
COPY --from=build /build/main .
COPY cert.pem .
COPY key.pem .


RUN mkdir /static
WORKDIR /static
COPY static/index.html .

EXPOSE 443

CMD ["./main"]