#!/usr/bin/env bash
# Build and ship. Run from the repository root.
#
#   ./deploy/deploy.sh                 # binary + every course's content
#   ./deploy/deploy.sh valuechain      # binary + one course
#
# One binary serves every course, so a fix deploys to all of them at once. Each course
# is a directory of content that rsyncs separately and restarts only its own instance.
set -euo pipefail

HOST=${HOST:-wilhelm@172.232.147.147}
KEY=${KEY:-$HOME/.ssh/id_ed25519_nopass}
DIR=${DIR:-/var/www/chulastudy}

cd "$(dirname "$0")/.."

COURSES=("$@")
if [ ${#COURSES[@]} -eq 0 ]; then
  COURSES=()
  for d in courses/*/; do COURSES+=("$(basename "$d")"); done
fi

echo "==> checking the content"
go test ./... >/dev/null

echo "==> building the frontend"
(cd web && npm ci --silent && npm run build)

echo "==> cross-compiling for linux/amd64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags='-s -w' -o dist/chulastudy ./cmd/server

echo "==> uploading the binary"
ssh -i "$KEY" "$HOST" "sudo install -d -o wilhelm -g wilhelm $DIR $DIR/courses"
# A running executable cannot be overwritten, but it can be replaced.
scp -q -i "$KEY" dist/chulastudy "$HOST:$DIR/chulastudy.new"
ssh -i "$KEY" "$HOST" "mv $DIR/chulastudy.new $DIR/chulastudy && chmod 755 $DIR/chulastudy && sudo restorecon -v $DIR/chulastudy"

for course in "${COURSES[@]}"; do
  echo "==> $course: syncing content"
  ssh -i "$KEY" "$HOST" "sudo install -d -o wilhelm -g wilhelm $DIR/courses/$course"
  # The audiobook lives outside the repo and changes far less often than the text, so
  # it is excluded here and shipped by sync-audio.sh.
  rsync -a --delete --exclude 'audio/' -e "ssh -i $KEY" \
    "courses/$course/" "$HOST:$DIR/courses/$course/"
  ssh -i "$KEY" "$HOST" "sudo restorecon -R $DIR/courses/$course"
done

for course in "${COURSES[@]}"; do
  echo "==> $course: restarting"
  ssh -i "$KEY" "$HOST" "sudo systemctl restart chulastudy@$course && sleep 1 && systemctl is-active chulastudy@$course"
done

echo "==> done"
