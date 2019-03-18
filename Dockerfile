FROM golang:alpine as builder
ARG PACKAGE_NAME=spy-api

WORKDIR ./src/github.com/alexandear/${PACKAGE_NAME}

COPY ./main.go ./
COPY ./vendor ./vendor
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ${PACKAGE_NAME} . && \
    cp ${PACKAGE_NAME} /usr/local/bin/ && \
    rm -rf /go/src/github.com

FROM scratch
COPY --from=builder /usr/local/bin/${PACKAGE_NAME} /usr/local/bin/${PACKAGE_NAME}

ENV BIND 0.0.0.0:80
EXPOSE 80

ENTRYPOINT ["spy-api"]
