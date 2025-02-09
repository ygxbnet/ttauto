FROM golang AS builder

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0

WORKDIR /build

COPY ./ .

RUN set -ex \
  && cd /build \
  && go build -o ttauto


FROM alpine:latest

COPY --from=builder /build/ttauto /usr/bin/ttauto

RUN apk update \
  && apk add tzdata --no-cache \
  && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo "Asia/Shanghai" > /etc/timezone \
  && chmod +x /usr/bin/ttauto

WORKDIR /data

ENTRYPOINT [ "/usr/bin/ttauto" ]
