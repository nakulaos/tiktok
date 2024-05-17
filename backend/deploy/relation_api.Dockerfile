FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY  output/relation_api /app/relation_api
COPY  relation/api/etc /app/etc
COPY script /app/script

CMD ["./relation_api", "-f", "etc/relation_dev.yaml"]