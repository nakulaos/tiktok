FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/relation_rpc /app/relation_rpc
COPY  relation/rpc/etc /app/etc
COPY script /app/script

CMD ["./relation_rpc", "-f", "etc/relation_dev.yaml"]