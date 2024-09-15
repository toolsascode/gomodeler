#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run ./cmd/gomodeler completion "$sh" > "completions/gomodeler.$sh"
done