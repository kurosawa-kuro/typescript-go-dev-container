FROM node:20-bookworm-slim

# Set working directory
WORKDIR /workspaces/typescript-go-dev-container/frontend

# Create necessary directories and set permissions
RUN mkdir -p /workspaces/typescript-go-dev-container/frontend/node_modules && \
    chown -R node:node /workspaces

# Switch to non-root user
USER node