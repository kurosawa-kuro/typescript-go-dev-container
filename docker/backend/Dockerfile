FROM golang:1.21.5-bookworm

# Install air for Go hot reload
RUN go install github.com/cosmtrek/air@v1.49.0

# Create Go module cache directory and set permissions
RUN mkdir -p /go/{pkg,src,bin} && \
    chmod -R 777 /go

# Set working directory
WORKDIR /workspaces/typescript-go-dev-container/backend/src

# Copy go.mod and go.sum first
COPY backend/src/go.mod backend/src/go.sum ./

# Download dependencies
RUN go mod download

# Create non-root user
RUN useradd -m -s /bin/bash -u 1000 vscode && \
    chown -R vscode:vscode /workspaces /go

# Switch to non-root user
USER vscode