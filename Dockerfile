# build stage
FROM golang:1.22-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum main.go ./
COPY app ./app
COPY common ./common

RUN go mod download
RUN go build -o /app/blockaction-api .


# final stage
FROM debian:bullseye-slim

ENV PORT=8080
ENV GIN_MODE=release
ENV LOG_LEVEL=info
ENV DATABASE_URL=""

ARG USERNAME=blockaction-service
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

# [Optional] Set the default user. Omit if you want to keep the default as root.
USER $USERNAME

COPY --from=builder /app/blockaction-api /app/blockaction-api

EXPOSE 8080

CMD ["/app/blockaction-api"]
