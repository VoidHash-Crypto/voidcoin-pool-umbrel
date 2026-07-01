#!/bin/bash
set -e

VERSION="${1:-1.0.0}"
REGISTRY="${REGISTRY:-void}"

echo "Building Void Pool Umbrel App v${VERSION}"
echo "=========================================="

cd "$(dirname "$0")/.."

# Build Node image
echo "Building VoidCoin node image..."
docker build -t ${REGISTRY}/void-pool-node:${VERSION} \
    -t ${REGISTRY}/void-pool-node:latest \
    -f umbrel-app/docker/node/Dockerfile \
    umbrel-app/docker/node/

# Build API image
echo "Building API image..."
docker build -t ${REGISTRY}/void-pool-api:${VERSION} \
    -t ${REGISTRY}/void-pool-api:latest \
    -f umbrel-app/docker/api/Dockerfile \
    .

# Build Stratum image
echo "Building Stratum image..."
docker build -t ${REGISTRY}/void-pool-stratum:${VERSION} \
    -t ${REGISTRY}/void-pool-stratum:latest \
    -f umbrel-app/docker/stratum/Dockerfile \
    .

# Build Web image
echo "Building Web image..."
docker build -t ${REGISTRY}/void-pool-web:${VERSION} \
    -t ${REGISTRY}/void-pool-web:latest \
    -f umbrel-app/docker/web/Dockerfile \
    .

echo ""
echo "Build complete!"
echo ""
echo "To push to registry:"
echo "  docker push ${REGISTRY}/void-pool-node:${VERSION}"
echo "  docker push ${REGISTRY}/void-pool-api:${VERSION}"
echo "  docker push ${REGISTRY}/void-pool-stratum:${VERSION}"
echo "  docker push ${REGISTRY}/void-pool-web:${VERSION}"
echo ""
echo "To test locally:"
echo "  cd umbrel-app && docker-compose up -d"
