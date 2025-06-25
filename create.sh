#!/bin/bash

TODAY=$(date +'%Y-%m-%d')
YESTERDAY=$(date -d "yesterday" +'%Y-%m-%d')
NOW=$(date +'%H-%M-%S')
LOG_DIR="logs"
LOG_FILE="$LOG_DIR/$TODAY-$NOW.log"

mkdir -p "$LOG_DIR" # ✅ Create logs folder if missing

# Log helper function
log() {
  echo "$1" >>"$LOG_FILE"
}

# Create today's folder
if [ ! -d "$TODAY" ]; then
  echo "📂 Creating today's folder: $TODAY"
  mkdir -p "$TODAY/array"
  mkdir -p "$TODAY/slice"
  mkdir -p "$TODAY/search"
  log "✅ Created today's directory structure: $TODAY"
fi
# Copy yesterday's array.go if it exists
if [ -f "$YESTERDAY/array/array.go" ]; then
  cp "$YESTERDAY/array/array.go" "$TODAY/array/array.go"
  echo "📄 Copied yesterday's array.go to today's folder"
  log "✅ Copied $YESTERDAY/array/array.go to $TODAY/array/array.go"
fi

# Declare files
declare -a function_files=(
  "$TODAY/array/array.go"
  "$TODAY/slice/slice.go"
  "$TODAY/search/search.go"
)

declare -a test_files=(
  "tests/array_test.go"
  "tests/slice_test.go"
  "tests/search_test.go"
)

# Create function files
for file in "${function_files[@]}"; do
  if [ ! -f "$file" ]; then
    echo "📝 Creating function file: $file"
    touch "$file"
    log "✅ Created function file: $file"
  else
   log "ℹ️ File already exists: $file (skipped creation)"
  fi
done

# Create test files
for file in "${test_files[@]}"; do
  if [ ! -f "$file" ]; then
    echo "📝 Creating test file: $file"
    mkdir -p "$(dirname "$file")"
    touch "$file"
    log "✅ Created test file: $file"
  else
    log "ℹ️ Test File already exists: $file (skipped creation)"
  fi
done

rm -rf latest
ln -sfn "$TODAY" latest

echo "✅ Updated symlink: latest -> $TODAY"
echo "🚀 Setup complete for $TODAY"
log "✅ Symlink created to $TODAY"
log "🚀 Setup complete"
