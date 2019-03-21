FROM golang:alpine as builder

RUN apk add --no-cache make

ARG PACKAGE_NAME=spy-api
ARG BIN_DIR=/usr/local/bin

WORKDIR ./src/github.com/devchallenge/${PACKAGE_NAME}

COPY main.go Makefile ./
COPY ./vendor ./vendor
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN make build_static && cp ./bin/${PACKAGE_NAME} ${BIN_DIR} && rm -rf /go/src/github.com

FROM scratch
COPY --from=builder ${BIN_DIR}/${PACKAGE_NAME} ${BIN_DIR}/${PACKAGE_NAME}

ENV BIND 0.0.0.0:80
EXPOSE 80

ENTRYPOINT ["spy-api"]
