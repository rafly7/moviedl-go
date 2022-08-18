#!/bin/bash
trap break INT
while true; do
  ls -d **/*.go **/**/*.go | entr -r go run main.go dev
done
echo "Goodbye"
trap - INT