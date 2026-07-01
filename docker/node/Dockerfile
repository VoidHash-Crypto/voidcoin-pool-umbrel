FROM ubuntu:24.04

ARG VOIDCOIN_VERSION=1.0.2
ARG TARGETARCH

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    libminiupnpc17 \
    libnatpmp1 \
    libevent-pthreads-2.1-7 \
    libevent-2.1-7 \
    libzmq5 \
    libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*

# Download VoidCoin Core binaries from GitHub releases
RUN set -eux; \
    case "${TARGETARCH}" in \
        amd64) ARCH_SUFFIX="linux-cli" ;; \
        arm64) ARCH_SUFFIX="linux-cli" ;; \
        *) echo "Unsupported architecture: ${TARGETARCH}" && exit 1 ;; \
    esac; \
    curl -L "https://github.com/VoidHash-Crypto/VoidCoin-Core/releases/download/v${VOIDCOIN_VERSION}/voidcoin-${VOIDCOIN_VERSION}-${ARCH_SUFFIX}.tar.gz" | \
    tar -xz -C /tmp && \
    cp /tmp/voidcoin-${VOIDCOIN_VERSION}-${ARCH_SUFFIX}/bin/voidcoind /usr/local/bin/ && \
    cp /tmp/voidcoin-${VOIDCOIN_VERSION}-${ARCH_SUFFIX}/bin/voidcoin-cli /usr/local/bin/ && \
    rm -rf /tmp/voidcoin-* && \
    chmod +x /usr/local/bin/voidcoin*

RUN mkdir -p /root/.voidcoin

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# P2P port, RPC port, ZMQ port
EXPOSE 7777 7778 28332

ENTRYPOINT ["/entrypoint.sh"]
CMD ["-printtoconsole", "-prune=1000"]
