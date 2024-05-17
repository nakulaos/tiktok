FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/comment_api /app/comment_api
COPY  comment/api/etc /app/etc
COPY script /app/script

CMD ["./comment_api", "-f", "etc/comment_dev.yaml"]