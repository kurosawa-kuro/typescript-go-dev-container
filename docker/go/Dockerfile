FROM golang:1.22

# Install essential tools
RUN apt-get update && apt-get install -y \
    build-essential \
    make \
    git

# Install Go tools
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@v1.49.0

# Set environment variables
ENV GO111MODULE=on
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Create and set work directory
WORKDIR /app/src

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

CMD ["air", "-c", "../.air.toml"]