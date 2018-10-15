#!/bin/sh

set -e
latest_tag=$(git describe --abbrev=0 --tags)
goxc
ghr -u gainings -r go-comments $latest_tag dist/snapshot/
