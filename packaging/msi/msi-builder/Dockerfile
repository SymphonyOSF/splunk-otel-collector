# Based on https://github.com/fleetdm/fleet/blob/main/tools/wix-docker/Dockerfile
# https://github.com/fleetdm/fleet/blob/main/LICENSE (MIT Expat)

ARG DOCKER_REPO=docker.io
FROM ${DOCKER_REPO}/debian:bullseye-slim as wix

RUN true \
    && dpkg --add-architecture i386 \
    && apt update \
    && apt install -y --no-install-recommends \
    ca-certificates \
    wine \
    wine32 \
    wget \
    unzip \
    osslsigncode \
    && mkdir /wix \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /home/wine
ENV HOME=/home/wine WINEPREFIX=/home/wine/.wine WINEARCH=win32 PATH="/home/wine/bin:$PATH" WINEDEBUG=-all

COPY packaging/msi/msi-builder/make-aliases.sh /home/wine/make-aliases.sh

# Install .NET framework and WiX Toolset binaries
RUN wine wineboot && \
    wget https://dl.winehq.org/wine/wine-mono/6.4.0/wine-mono-6.4.0-x86.msi -nv -O mono.msi \
    && wine msiexec /i mono.msi \
    && rm -f mono.msi \
    && wget https://github.com/wixtoolset/wix3/releases/download/wix314rtm/wix314-binaries.zip -nv -O wix.zip \
    && mkdir wix \
    && unzip wix.zip -d wix \
    && rm -f wix.zip \
    && /home/wine/make-aliases.sh \
    && rm -f /home/wine/make-aliases.sh \
    && mkdir $WINEPREFIX/drive_c/temp


FROM wix as msi-builder

ARG JMX_METRIC_GATHERER_RELEASE

USER root
RUN apt-get update -y
RUN apt-get install -y curl unzip

COPY bin/otelcol_windows_amd64.exe /project/bin/otelcol_windows_amd64.exe
COPY cmd/ /project/cmd
COPY packaging/fpm/ /project/internal/buildscripts/packaging/fpm
COPY packaging/msi/ /project/internal/buildscripts/packaging/msi
COPY dist/agent-bundle_windows_amd64.zip /project/dist/agent-bundle_windows_amd64.zip

RUN cp /project/internal/buildscripts/packaging/msi/msi-builder/docker-entrypoint.sh /docker-entrypoint.sh && \
    chmod a+x /docker-entrypoint.sh

ENV OUTPUT_DIR=/project/dist
ENV JMX_METRIC_GATHERER_RELEASE=${JMX_METRIC_GATHERER_RELEASE}
ENV VERSION=""

WORKDIR /work

ENTRYPOINT ["/docker-entrypoint.sh"]
