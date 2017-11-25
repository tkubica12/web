FROM ubuntu

COPY f5web /usr/local/bin/web

RUN chmod +x /usr/local/bin/f5web && apt-get update && apt-get install uuid-runtime

CMD ["/usr/local/bin/f5web"] 