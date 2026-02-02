#!/bin/bash
set -euo pipefail

URL="${1:-http://127.0.0.1:8080/roll}"
N="${2:-50}"

for i in $(seq 1 "$N"); do
  curl -s "$URL" >/dev/null
  sleep 0.1
done

echo "ok sent=$N url=$URL"