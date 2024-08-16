FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

RUN su vscode -c " \
    go install github.com/air-verse/air@latest && \
    go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest \
    "
