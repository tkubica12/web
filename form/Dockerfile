FROM ubuntu

COPY f5web /usr/local/bin/form

RUN chmod +x /usr/local/bin/form && apt-get update && apt-get install uuid-runtime

CMD ["/usr/local/bin/form"] 