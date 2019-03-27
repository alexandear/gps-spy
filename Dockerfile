FROM golang:alpine as builder

RUN apk add --no-cache make

ARG PACKAGE_NAME=spy-api
ARG BIN_DIR=/usr/local/bin

WORKDIR ./src/github.com/devchallenge/${PACKAGE_NAME}

COPY ./vendor ./vendor
COPY main.go Makefile ./
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN make test build_static && cp ./bin/${PACKAGE_NAME} ${BIN_DIR} && rm -rf /go/src/github.com

FROM scratch
COPY --from=builder ${BIN_DIR}/${PACKAGE_NAME} ${BIN_DIR}/${PACKAGE_NAME}

EXPOSE 80
ENV HOST 0.0.0.0
ENV PORT 80

ENTRYPOINT ["spy-api", "server"]
