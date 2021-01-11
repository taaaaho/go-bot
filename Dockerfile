FROM golang:1.15.5-alpine3.12

RUN set -eux && \
  apk --update add --no-cache git && \
  go get -u github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv

RUN mkdir /go/src/work
WORKDIR /go/src/work
COPY . /go/src/work

COPY go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "/go/src/work/cmd/main.go"]
