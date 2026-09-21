#!/usr/bin/env bash
# Build for the Linode and ship it. Run from the repository root.
#
#   ./deploy/deploy.sh
#
# The server has no Go toolchain, so the binary is cross-compiled here. The frontend is
# built first because it is embedded in the binary, and so is the study material — a
# deploy ships the questions along with the code.
set -euo pipefail

HOST=${HOST:-wilhelm@172.232.147.147}
KEY=${KEY:-$HOME/.ssh/id_ed25519_nopass}
DIR=${DIR:-/var/www/chulavaluechain}
SERVICE=chulavaluechain

cd "$(dirname "$0")/.."

echo "==> checking the question bank"
go test ./... >/dev/null

echo "==> building the frontend"
(cd web && npm ci --silent && npm run build)

echo "==> cross-compiling for linux/amd64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags='-s -w' -o dist/chulavaluechain ./cmd/server

echo "==> uploading"
ssh -i "$KEY" "$HOST" "sudo install -d -o wilhelm -g wilhelm $DIR"
# Upload beside the live binary, then move it into place: a running executable cannot
# be overwritten, but it can be replaced.
scp -i "$KEY" dist/chulavaluechain "$HOST:$DIR/chulavaluechain.new"
ssh -i "$KEY" "$HOST" "mv $DIR/chulavaluechain.new $DIR/chulavaluechain && chmod 755 $DIR/chulavaluechain"

echo "==> restarting $SERVICE"
ssh -i "$KEY" "$HOST" "sudo systemctl restart $SERVICE && sleep 1 && systemctl is-active $SERVICE"

echo "==> done: https://valuechain.wilhelm.my"
