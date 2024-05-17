FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/mq_api  /app/mq_api
COPY  mq/etc  /app/etc
COPY script  /app/script

CMD ["./mq_api", "-f", "etc/mq_dev.yaml"]