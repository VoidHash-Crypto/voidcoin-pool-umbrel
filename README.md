# Void Pool for Umbrel

Self-hosted VoidCoin mining pool for Umbrel.

## Installation

### Option 1: Community App Store (Recommended)

1. Open your Umbrel dashboard
2. Go to **App Store** → **Community App Stores** (three dots menu)
3. Add this URL: `https://github.com/VoidHash-Crypto/voidcoin-umbrel-store`
4. Find "Void Pool" in the app store and click **Install**

### Option 2: Standalone Docker Compose

Run Void Pool on any server (not just Umbrel):

```bash
# Clone the repo
git clone https://github.com/VoidHash-Crypto/void-pool-umbrel.git
cd void-pool-umbrel

# Copy and configure environment
cp .env.example .env

# Generate secure passwords
sed -i "s/NODE_RPC_PASSWORD=.*/NODE_RPC_PASSWORD=$(openssl rand -hex 32)/" .env
sed -i "s/DB_PASSWORD=.*/DB_PASSWORD=$(openssl rand -hex 32)/" .env

# Set your pool address (REQUIRED - edit this file)
nano .env  # Set POOL_ADDRESS to your VoidCoin address

# Start all services
docker compose up -d
```

## Quick Start

1. Install Void Pool using one of the methods above
2. Wait for the VoidCoin node to sync (first sync takes 15-30 minutes)
3. Open the web UI at `http://umbrel.local:3080`
4. Point your miners to `stratum+tcp://umbrel.local:3333`

## Miner Configuration

```
URL: stratum+tcp://YOUR_UMBREL_IP:3333
Username: YOUR_VOID_ADDRESS.WORKER_NAME
Password: x
```

### Mining Modes

| Worker Suffix | Mode | Description |
|---------------|------|-------------|
| `.solo` | Solo | Keep 100% of blocks you find |
| (none) | PPLNS | Shared rewards with other miners |

**Examples:**
```
void:qr08krt...7gnl732fwz.rig1       → PPLNS
void:qr08krt...7gnl732fwz.rig1.solo  → Solo
```

## Configuration

Edit `.env` or configure via Umbrel app settings:

| Setting | Default | Description |
|---------|---------|-------------|
| `POOL_ADDRESS` | (required) | Your VoidCoin address for pool fees |
| `POOL_NAME` | My Void Pool | Pool name shown in UI |
| `POOL_FEE` | 1.0 | PPLNS fee percentage |
| `SOLO_FEE` | 0.5 | Solo mining fee percentage |
| `MIN_PAYOUT` | 5 | Minimum VOID for auto-payout |
| `STRATUM_PORT` | 3333 | Stratum port for miners |
| `APP_PORT` | 3080 | Web UI port |

## Ports

| Port | Service | Description |
|------|---------|-------------|
| 3080 | Web UI | Pool dashboard |
| 3333 | Stratum | Miner connections (forward this on your router) |

## Data Storage

All data stored in the app data directory:
- `node/` - VOID blockchain (pruned, ~500MB)
- `postgres/` - Pool database
- `redis/` - Cache

## Troubleshooting

**Node not syncing?**
```bash
docker logs void-node
```

**Miners can't connect?**
- Ensure port 3333 is forwarded on your router
- Check firewall: `sudo ufw allow 3333`

**Pool shows 0 hashrate?**
- Hashrate appears after miners submit shares
- Stats update every 5 minutes

**Check all container status:**
```bash
docker ps | grep forge
```

## Building Images Locally

To build the Docker images yourself instead of using pre-built images:

```bash
./build.sh
```

## Support

- Issues: https://github.com/VoidHash-Crypto/void-pool-umbrel/issues
- Website: https://hashforge.void.org
- Main Pool: https://hashforge.void.org (use this if you just want to mine without running your own pool)
