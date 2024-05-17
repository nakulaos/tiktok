FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/user_rpc /app/user_rpc
COPY  user/rpc/etc /app/etc
COPY script /app/script

CMD ["./user_rpc", "-f", "etc/user_dev.yaml"]