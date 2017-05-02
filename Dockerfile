FROM daocloud.io/golang:1.8-alpine

LABEL maintainer="Haichuang.Li" \
      io.daocloud.dce.plugin.name="AppUpdate" \
      io.daocloud.dce.plugin.description="AppUpdate " \
      io.daocloud.dce.plugin.categories="app-tool" \
      io.daocloud.dce.plugin.vendor="DaoCloud" \
      io.daocloud.dce.plugin.required-dce-version=">=2.6.0" \
      io.daocloud.dce.plugin.nano-cpus-limit="500000000" \
      io.daocloud.dce.plugin.memory-bytes-limit="524288000"

RUN apk add --update nginx supervisor build-base libffi-dev openssl-dev gcc \
    linux-headers musl-dev \
  && rm -rf /var/cache/apk/* \
  && ln -sf /dev/stdout /var/log/nginx/access.log \
  && ln -sf /dev/stderr /var/log/nginx/error.log

COPY nginx /etc/nginx/
COPY dist /usr/share/nginx/html/
COPY plugin.json /usr/share/nginx/html/plugin.json

COPY app_update /go/src/github.com/sctlee/app_update
RUN go install github.com/sctlee/app_update


COPY supervisord.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/supervisord.sh

CMD ["sh", "/usr/local/bin/supervisord.sh"]
