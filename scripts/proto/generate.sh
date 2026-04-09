#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
PROTOC_BIN="${PROTOC_BIN:-/opt/homebrew/bin/protoc}"

if [[ ! -x "${PROTOC_BIN}" ]]; then
  echo "protoc not found at ${PROTOC_BIN}" >&2
  exit 1
fi

export PATH="/Users/adam/.local/bin:${PATH}"

cd "${REPO_ROOT}"

PROTO_FILES=()
while IFS= read -r file; do
  PROTO_FILES+=("${file}")
done < <(find "${REPO_ROOT}/proto" -name '*.proto' | sort)

if [[ ${#PROTO_FILES[@]} -eq 0 ]]; then
  echo "No proto files found under ${REPO_ROOT}/proto" >&2
  exit 1
fi

"${PROTOC_BIN}" \
  -I "${REPO_ROOT}/proto" \
  --go_out="${REPO_ROOT}/modules/exchange-core-go" \
  --go_opt=module=local.exchange-demo/exchange-core-go \
  --go-grpc_out="${REPO_ROOT}/modules/exchange-core-go" \
  --go-grpc_opt=module=local.exchange-demo/exchange-core-go \
  "${PROTO_FILES[@]}"
