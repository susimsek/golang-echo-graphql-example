FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# builder stagedeki binary dosyasını production imageye kopyalayın
COPY server /app/server

# Web servisi konteynar başlangıcında çalıştırın.
CMD ["/app/server"]