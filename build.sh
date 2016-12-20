#!/bin/bash

set -e
set -x

MANIFEST="${MANIFEST:-manifest}"

cd "$(dirname "$0")"
for f in manifests/*.yml; do
	"$MANIFEST" pushml "$f"
done
