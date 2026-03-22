FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder

ARG TARGETARCH
ARG LDFLAGS

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /usr/local/go/src/app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY ./ ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -ldflags="$LDFLAGS" -o /dist/app main.go \
    && cp -r etc /dist/etc \
    && mkdir -p /dist/desc \
    && if [ -d desc/swagger ]; then cp -r desc/swagger /dist/desc; fi


FROM --platform=$TARGETPLATFORM alpine:latest

WORKDIR /dist

COPY --from=builder /dist .

EXPOSE 8000 8001

CMD ["./app", "server"]