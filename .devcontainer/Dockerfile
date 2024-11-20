FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Install air for Go hot reload
RUN go install github.com/cosmtrek/air@v1.49.0

# Create Go module cache directory and set permissions
RUN mkdir -p /go/{pkg,src,bin} && \
    chmod -R 777 /go

# Set default user
USER vscode
