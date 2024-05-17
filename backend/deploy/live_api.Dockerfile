FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/live_api /app/live_api
COPY  live/api/etc /app/etc
COPY script /app/script

CMD ["./live_api", "-f", "etc/live_dev.yaml"]