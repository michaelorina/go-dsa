#!/bin/bash

TODAY=$(date +'%Y-%m-%d')
YESTERDAY=${YESTERDAY:-$(date -d "yesterday" +'%Y-%m-%d')}
NOW=$(date +'%H-%M-%S')
LOG_DIR="logs"
LOG_FILE="$LOG_DIR/$TODAY-$NOW.log"

mkdir -p "$LOG_DIR"

log() {
  echo "$1" >>"$LOG_FILE"
}

declare -A paths_to_copy=(
  ["$YESTERDAY/array/array.go"]="$TODAY/array/array.go"
  ["$YESTERDAY/slice/slice.go"]="$TODAY/slice/slice.go"
  ["$YESTERDAY/search/search.go"]="$TODAY/search/search.go"
  ["$YESTERDAY/sort/sort.go"]="$TODAY/sort/sort.go"
  ["$YESTERDAY/sort/linkedlist.go"]="$TODAY/sort/linkedlist.go"
  ["$YESTERDAY/sort/queue.go"]="$TODAY/sort/queue.go"
  ["$YESTERDAY/sort/stack.go"]="$TODAY/sort/stack.go"
  ["$YESTERDAY/quicksort/quicksort.go"]="$TODAY/quicksort/quicksort.go"
)

for src in "${!paths_to_copy[@]}"; do
  dest=${paths_to_copy[$src]}
  if [ -f "$src" ]; then
    cp "$src" "$dest"
    echo "📄 Copied $src to $dest"
    log "✅ Copied $src to $dest"
  else
    echo "⚠️ File not found: $src (skipped)"
    log "⚠️ Missing source: $src"
  fi
done

echo "✅ Copying complete for $TODAY"
log "🚀 Copy script complete"
