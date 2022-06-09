FROM golang:1.18-alpine as builder
LABEL maintainer="Simone Lazzaris <simone@vchain.us>"
WORKDIR /app
COPY go.* .
RUN go mod download -x
COPY *.go ./
RUN go build -o zeroservice

FROM scratch as runner
WORKDIR /app
COPY --from=builder /app/zeroservice .
EXPOSE 8080
CMD ["./zeroservice"]

FROM builder as debugger
WORKDIR /go
RUN apk add gcc build-base && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 8080
EXPOSE 1234
CMD ["dlv", "debug", "/app/zeroservice", "--headless", "--listen=:1234"]

