#!/usr/bin/env bash
# Push a course's rendered audiobook without rebuilding or redeploying the app.
#
#   ./deploy/sync-audio.sh valuechain ~/Desktop/valuechain-audiobook
#
# Episodes are rendered one lecture at a time, so this ships whatever exists. The
# manifest is rebuilt first so it lists exactly what is on disk, and the service is
# restarted because it reads the manifest once at startup.
set -euo pipefail

COURSE=${1:?usage: sync-audio.sh <course> [audio-dir]}
AUDIO=${2:-$HOME/Desktop/$COURSE-audiobook}
HOST=${HOST:-wilhelm@172.232.147.147}
KEY=${KEY:-$HOME/.ssh/id_ed25519_nopass}
DIR=${DIR:-/var/www/chulastudy}

cd "$(dirname "$0")/.."

echo "==> rebuilding the manifest from what has rendered"
python3.12 tools/manifest.py "$AUDIO" /tmp/script.json

echo "==> syncing"
ssh -i "$KEY" "$HOST" "sudo install -d -o wilhelm -g wilhelm $DIR/courses/$COURSE/audio"
rsync -a --delete -e "ssh -i $KEY" "$AUDIO/" "$HOST:$DIR/courses/$COURSE/audio/"
ssh -i "$KEY" "$HOST" "sudo restorecon -R $DIR/courses/$COURSE/audio"

echo "==> restarting so it re-reads the manifest"
ssh -i "$KEY" "$HOST" "sudo systemctl restart chulastudy@$COURSE && sleep 1 && systemctl is-active chulastudy@$COURSE"
