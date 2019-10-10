#!/bin/sh
set -ux

git diff

git config --local user.email "action@github.com"
git config --local user.name "GitHub Action"
git commit -m "Update to match upstream" -a || true

remote_repo="https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
git push "${remote_repo}" HEAD:master
