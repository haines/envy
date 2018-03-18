FROM circleci/golang:1.10

USER root

ENV BATS_VERSION=0.4.0
RUN set -x \
 && mkdir bats \
 && curl -fsSL "https://github.com/sstephenson/bats/archive/v${BATS_VERSION}.tar.gz" \
    | tar xz --strip-components=1 -C bats \
 && bats/install.sh /usr/local \
 && rm -Rf bats

USER circleci
