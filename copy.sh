#!/bin/bash

TODAY=$(date +'%Y-%m-%d')
YESTERDAY=$(date -d "yesterday" +'%Y-%m-%d')
NOW=$(date +'%H-%M-%S')

# Copy yesterday's array.go if it exists
if [ -f "$YESTERDAY/array/array.go" ]; then
  cp "$YESTERDAY/array/array.go" "$TODAY/array/array.go"
  echo "📄 Copied yesterday's array.go to today's folder"
  log "✅ Copied $YESTERDAY/array/array.go to $TODAY/array/array.go"
fi

# Copy yesterday's slice.go if it exists
if [ -f "$YESTERDAY/slice/slice.go" ]; then
  cp "$YESTERDAY/slice/slice.go" "$TODAY/slice/slice.go"
  echo "📄 Copied yesterday's slice.go to today's folder"
  log "✅ Copied $YESTERDAY/slice/slice.go to $TODAY/slice/slice.go"
fi

# Copy yesterday's slice.go if it exists
if [ -f "$YESTERDAY/search/ssearch.go" ]; then
  cp "$YESTERDAY/search/search.go" "$TODAY/search/search.go"
  echo "📄 Copied yesterday's search.go to today's folder"
  log "✅ Copied $YESTERDAY/search/search.go to $TODAY/search/search.go"
fi
