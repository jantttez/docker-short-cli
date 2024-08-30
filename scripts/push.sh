#!/bin/bash

MESSAGE="1"
DATE=$(date)

PARENT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)

rm "${PARENT_DIR}/docker-short"

if [ -z "$MESSAGE" ]; then
    echo "Starting push to remote repo with auto message..."
    git add "${PARENT_DIR}/."
    git commit -m "$DATE"
    git push origin main
    echo "Pushing complete success"
else
    echo "Starting push to remote repo with custom message..."
    git add "${PARENT_DIR}/."
    git commit -m "$MESSAGE"
    git push origin main
    echo "Pushing complete success"
fi
