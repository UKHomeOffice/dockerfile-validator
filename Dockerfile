FROM scratch
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

ADD dockerfile-validator /
COPY rules.yaml /rules.yaml

CMD ["/dockerfile-validator"]
