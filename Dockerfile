# BUILDER - Build binary for production
FROM golang:1.16 as builder

ENV GRPC_HEALTH_PROBE_VERSION=v0.3.6

ARG VERSION=unknown
ENV VERSION=$VERSION

WORKDIR /go/src/github.com/nicolaspearson/go.api/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -o api \
    -ldflags "-X $(go list -m)/pkg/version.VERSION=${VERSION}" \
    ./cmd/api/main.go

RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# RUNNER - Production image
FROM gcr.io/distroless/static-debian10

ARG RELEASE_VERSION=unknown
ENV RELEASE_VERSION=$RELEASE_VERSION

ARG VERSION=unknown
ENV VERSION=$VERSION

WORKDIR /bin/
COPY --from=builder /go/src/github.com/nicolaspearson/go.api/api .
COPY --from=builder /go/src/github.com/nicolaspearson/go.api/config/api.yaml ./config/
COPY --from=builder /bin/grpc_health_probe .
ENTRYPOINT [ "/bin/api" ]
