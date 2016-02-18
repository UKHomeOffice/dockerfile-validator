FROM scratch
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

ADD dockerfile-validator /
COPY rules.yaml /rules.yaml
COPY upload.html /upload.html

CMD ["/dockerfile-validator"]
