FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/favorite_api /app/favorite_api
COPY  favorite/api/etc /app/etc
COPY script /app/script

CMD ["./favorite_api", "-f", "etc/favorite_dev.yaml"]