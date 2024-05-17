FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/feed_rpc /app/feed_rpc
COPY  feed/rpc/etc /app/etc
COPY script /app/script

CMD ["./feed_rpc", "-f", "etc/feed_dev.yaml"]