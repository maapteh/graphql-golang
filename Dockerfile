FROM golang:1.15.2-buster

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PORT=3000

WORKDIR /build
COPY . .
RUN go mod vendor
RUN go build

WORKDIR /dist
RUN cp /build/graphql-golang .

EXPOSE $PORT
CMD ["/dist/graphql-golang"]