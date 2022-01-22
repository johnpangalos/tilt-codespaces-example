# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.209.6/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

ENV CTLPTL_VERSION="0.7.2"
RUN curl -fsSL https://github.com/tilt-dev/ctlptl/releases/download/v$CTLPTL_VERSION/ctlptl.$CTLPTL_VERSION.linux.x86_64.tar.gz | tar -xzv -C /usr/local/bin ctlptl
RUN curl -fsSL https://github.com/tilt-dev/tilt/releases/download/v0.23.6/tilt.0.23.6.linux.x86_64.tar.gz | tar -xzv tilt && mv tilt /usr/local/bin/tilt
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && unzip awscliv2.zip && sudo ./aws/install
RUN curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.11.1/kind-linux-amd64 && chmod +x ./kind && mv ./kind /usr/local/bin/kind

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get -y install --no-install-recommends xdg-utils

USER vscode
RUN git config --global pull.rebase false
COPY ./post-create.sh /home/vscode/post-create.sh