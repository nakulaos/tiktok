FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/favorite_rpc /app/favorite_rpc
COPY  favorite/rpc/etc /app/etc
COPY script /app/script

CMD ["./favorite_rpc", "-f", "etc/favorite_dev.yaml"]