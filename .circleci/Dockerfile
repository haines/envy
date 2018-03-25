FROM circleci/golang:1.10

USER root

ENV DEP_VERSION=0.4.1 \
    DEP_SHA256SUM=31144e465e52ffbc0035248a10ddea61a09bf28b00784fd3fdd9882c8cbb2315

RUN set -x \
 && dep_binary=/usr/local/bin/dep \
 && curl -fsSL "https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64" -o $dep_binary \
 && echo "${DEP_SHA256SUM} ${dep_binary}" | sha256sum -c - \
 && chmod +x $dep_binary \
 && dep version

USER circleci

RUN set -x \
 && go get github.com/jstemmer/go-junit-report
