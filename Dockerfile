# build stage
FROM golang:1.22-bullseye AS build-stage

WORKDIR /app

COPY go.mod go.sum main.go ./
COPY app ./app
COPY common ./common

RUN go mod download
RUN go build -o /app/blockaction-api .


# final stage
FROM debian:bullseye-slim AS execution-stage

ENV PORT=8080
ENV GIN_MODE=release
ENV LOG_LEVEL=info
ENV DATABASE_URL=""

ARG USERNAME=blockaction-service
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# install tools
RUN apt-get update \
    # * for ssl certificates
    && apt-get install -y ca-certificates \
    # * for non-root user
    && apt-get install -y sudo \
    # * for healthcheck
    && apt-get install -y curl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# configure non-root user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# configure healthcheck
HEALTHCHECK --start-period=5s --interval=30s --timeout=5s --retries=5 \
    CMD curl --silent --fail http://localhost:8080/health || exit 1

# [Optional] Set the default user. Omit if you want to keep the default as root.
USER $USERNAME

COPY --from=build-stage /app/blockaction-api /app/blockaction-api

EXPOSE 8080

CMD ["/app/blockaction-api"]
