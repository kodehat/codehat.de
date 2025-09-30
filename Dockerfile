FROM golang:1.25.1-alpine3.21 AS backend

ARG VERSION=dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY static static/
COPY templates templates/
COPY .git .git/
COPY *.go build.sh ./
COPY pkg pkg/

RUN apk add --no-cache bash curl git && \
  ./build.sh -v "$VERSION"

FROM alpine:3.22.1

LABEL org.opencontainers.image.authors='dev@codehat.de' \
      org.opencontainers.image.url='https://www.codehat.de' \
      org.opencontainers.image.documentation='https://github.com/kodehat/codehat.de' \
      org.opencontainers.image.source='https://github.com/kodehat/codehat.de' \
      org.opencontainers.image.vendor='kodehat' \
      org.opencontainers.image.licenses='AGPL-3.0'

WORKDIR /opt

COPY --from=backend /app/codehat.de ./codehat.de

# "curl" is added only for Docker healthchecks!
RUN apk add --no-cache tzdata curl && \
  adduser -D -H nonroot && \
  chmod +x ./codehat.de

EXPOSE 1313/tcp

USER nonroot:nonroot

ENTRYPOINT [ "/opt/codehat.de" ]
