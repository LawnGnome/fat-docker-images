#!/bin/bash

set -e
set -x

MANIFEST="${MANIFEST:-manifest}"

cd "$(dirname "$0")"
find manifests -name '*.yml' -exec "$MANIFEST" pushml '{}' ';'
