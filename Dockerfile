FROM golang:1.15.5

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY pkg ./pkg
COPY html ./html
COPY css ./css
COPY cmd/web ./cmd/web

WORKDIR cmd/web
RUN go build main.go
RUN echo "build successful"

ENTRYPOINT ["./main"]