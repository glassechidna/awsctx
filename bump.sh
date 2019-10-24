#!/bin/sh
set -ux

git diff | wc -l

AWS_VERSION=$(cat go.mod | grep aws-sdk-go | awk '{print $2}')

git config --local user.email "aidan.steele+bot@glassechidna.com.au"
git config --local user.name "Aidan bot"
git commit -m "Upstream ${AWS_VERSION}" -a || exit 0
git tag -m "${AWS_VERSION}" -a "${AWS_VERSION}"

remote_repo="https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
git push "${remote_repo}" HEAD:master --follow-tags
