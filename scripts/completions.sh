#!/usr/bin/env bash
set -e

echo "Generating shell completions..."

# Remove existing completions directory
rm -rf completions
mkdir -p completions

# Generate completions for each shell
for sh in bash zsh fish; do
  echo "Generating completion for $sh..."
  go run main.go completion "$sh" > "completions/gofast.$sh"
done

echo "Shell completions generated successfully!"