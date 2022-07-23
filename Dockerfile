FROM golang:1.18-alpine as build
LABEL stage=Build
ENV GITHUB_SHA=$GITHUB_SHA
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get . && go install .
RUN go build -o /FOSH-Proxy

FROM alpine:3
WORKDIR /go/src/app
COPY --from=build /FOSH-Proxy ./FOSH-Proxy
EXPOSE 8080
VOLUME ['/go/src/app/config.yml']

ENTRYPOINT  ["./FOSH-Proxy"]
