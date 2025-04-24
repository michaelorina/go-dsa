#!/bin/bash

# 💡 Get today's date in YYYY-MM-DD format for commit messages
TODAY=$(date +'%F')

# 🧼 Stage changes
echo "🔍 Staging changes..."
git add .

# ✍️ Commit with a standard message
echo "📝 Committing changes..."
git commit -m "📅 Daily DSA commit: $TODAY"

# ⏫ Push to origin main
echo "🚀 Pushing to GitHub..."
git push origin main

# ✅ Done
echo "✅ Push complete!"
