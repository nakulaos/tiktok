FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/feed_api /app/feed_api
COPY  feed/api/etc /app/etc
COPY script /app/script

CMD ["./feed_api", "-f", "etc/feed_dev.yaml"]