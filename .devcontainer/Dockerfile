FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install air for hot reload
RUN go install github.com/cosmtrek/air@v1.49.0

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# Create Go module cache directory and set permissions
RUN mkdir -p /go/pkg && \
    mkdir -p /go/src && \
    mkdir -p /go/bin && \
    chmod -R 777 /go

# If you have a non-root user, you might want to change ownership
# RUN chown -R your_user:your_user /go
