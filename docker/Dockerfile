FROM golang:1.9

MAINTAINER Hassan Youssef "hasusuf@gmail.com"

ADD build build

RUN ./build && rm build

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["dnsotls"]