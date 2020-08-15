FROM golang:1.14.6-alpine as builder

ENV GOPATH /go

RUN apk update && apk add --no-cache git & mkdir $GOPATH/src/go-hpa/

WORKDIR $GOPATH/src/go-hpa/

COPY ./go-hpa.go ./go-hpa_test.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOPATH/src/go-hpa/go-hpa .

FROM scratch

COPY --from=builder /go/src/go-hpa/go-hpa /app/
WORKDIR /app
USER 1000
EXPOSE 8000

ENTRYPOINT ["./go-hpa"]
