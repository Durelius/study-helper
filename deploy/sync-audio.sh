#!/usr/bin/env bash
# Push newly rendered episodes without rebuilding or redeploying the app.
#
#   ./deploy/sync-audio.sh
#
# Episodes are rendered one lecture at a time, so this exists to ship them as they
# appear. The manifest is rebuilt first so it lists exactly what is on disk, and the
# service is restarted because it reads the manifest once at startup.
set -euo pipefail

HOST=${HOST:-wilhelm@172.232.147.147}
KEY=${KEY:-$HOME/.ssh/id_ed25519_nopass}
DIR=${DIR:-/var/www/chulavaluechain}
AUDIO=${AUDIO:-$HOME/Desktop/valuechain-audiobook}
SERVICE=chulavaluechain

cd "$(dirname "$0")/.."

echo "==> rebuilding the manifest from what has rendered"
python3.12 tools/manifest.py "$AUDIO" /tmp/script.json

echo "==> syncing"
ssh -i "$KEY" "$HOST" "sudo install -d -o wilhelm -g wilhelm $DIR/audio"
rsync -a --delete -e "ssh -i $KEY" "$AUDIO/" "$HOST:$DIR/audio/"
ssh -i "$KEY" "$HOST" "sudo restorecon -R $DIR/audio"

echo "==> restarting $SERVICE so it re-reads the manifest"
ssh -i "$KEY" "$HOST" "sudo systemctl restart $SERVICE && sleep 1 && systemctl is-active $SERVICE"

curl -s "https://valuechain.wilhelm.my/api/episodes?player=-" |
  python3 -c "import json,sys; d=json.load(sys.stdin); print(f'==> live: {len(d[\"episodes\"])} episodes, {d[\"minutes\"]:.0f} minutes')"
