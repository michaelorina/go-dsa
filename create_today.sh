#!/bin/bash

TODAY=$(date +'%Y-%m-%d')
NOW=$(date +'%H-%M-%S')
LOG_DIR="logs"
LOG_FILE="$LOG_DIR/$TODAY-$NOW.log"

mkdir -p "$LOG_DIR"

log() {
  echo "$1" >>"$LOG_FILE"
}

# Create today's folders
if [ ! -d "$TODAY" ]; then
  echo "📂 Creating today's folder structure: $TODAY"
  mkdir -p "$TODAY/array"
  mkdir -p "$TODAY/slice"
  mkdir -p "$TODAY/search"
  mkdir -p "$TODAY/sort"
  log "✅ Created directory structure for $TODAY"
fi

# Define function files to create
declare -a function_files=(
  "$TODAY/array/array.go"
  "$TODAY/slice/slice.go"
  "$TODAY/search/search.go"
  "$TODAY/sort/sort.go"
  "$TODAY/sort/linkedlist.go"
  "$TODAY/sort/queue.go"
  "$TODAY/sort/stack.go"
)

# Define test files to create
declare -a test_files=(
  "tests/array_test.go"
  "tests/slice_test.go"
  "tests/search_test.go"
  "tests/sort_test.go"
  "tests/linkedlist_test.go"
  "tests/queue_test.go"
  "tests/stack_test.go"
)

# Create function files if missing
for file in "${function_files[@]}"; do
  if [ ! -f "$file" ]; then
    echo "📝 Creating function file: $file"
    touch "$file"
    log "✅ Created function file: $file"
  else
    log "ℹ️ File exists: $file (skipped)"
  fi
done

# Create test files if missing
for file in "${test_files[@]}"; do
  if [ ! -f "$file" ]; then
    echo "📝 Creating test file: $file"
    mkdir -p "$(dirname "$file")"
    touch "$file"
    log "✅ Created test file: $file"
  else
    log "ℹ️ Test file exists: $file (skipped)"
  fi
done

# Update symlink
rm -rf latest
ln -sfn "$TODAY" latest

echo "✅ Symlink updated: latest -> $TODAY"
echo "🚀 Setup complete for $TODAY"
log "✅ Symlink set"
log "🚀 Creation script complete"
