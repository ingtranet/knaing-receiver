FROM golang:1.12.6-alpine

ENV TINI_VERSION v0.18.0

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

RUN apk add --no-cache git mercurial
WORKDIR /workspace/knaing-receiver
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk del git mercurial

COPY . .
RUN go build -o /app/knaing_receiver ./cmd/receiver
RUN rm -rf $GOPATH/pkg/mod

ENTRYPOINT ["/tini", "-s", "--"]
CMD ["/app/knaing_receiver"]