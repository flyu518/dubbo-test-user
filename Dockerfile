FROM golang:1.20.0-alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app

# Install dependencies
RUN apk --update --no-cache add ca-certificates make

# Download grpc_health_probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.11 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.io,direct/
RUN make init && go mod download
COPY . .


RUN go build -o /app/server .

# Deployment container
FROM scratch

WORKDIR /app

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /bin/grpc_health_probe /bin/
COPY --from=builder /app/server /app/

ENTRYPOINT ["/app/server"]