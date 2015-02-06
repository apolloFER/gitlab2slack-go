FROM golang:1.3.3

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

ENTRYPOINT ["go-wrapper", "run"]