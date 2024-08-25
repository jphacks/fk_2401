FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

WORKDIR /backend

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1

RUN su vscode -c " \
    go install github.com/air-verse/air@latest && \
    go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest && \
    go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2 && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
    "
