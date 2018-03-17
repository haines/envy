FROM golang:1.10 as build

WORKDIR /go/src/github.com/haines/envy
COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    make all

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /go/src/github.com/haines/envy/target/envy /usr/local/bin/

RUN adduser -D envy

USER envy
WORKDIR /home/envy

ENTRYPOINT ["envy"]
