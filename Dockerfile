FROM alpine:latest

RUN apk add --no-cache tzdata

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY bin/app .

EXPOSE 8000

ENTRYPOINT ["./app"]