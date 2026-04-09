#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
SERVICE_DIR="${REPO_ROOT}/services/market-data-service-java"

JAVA_CANDIDATES=(
  "/Users/adam/.sdkman/candidates/java/21.0.3-tem"
  "/opt/homebrew/opt/openjdk"
)

MAVEN_CANDIDATES=(
  "/Users/adam/.sdkman/candidates/maven/current/bin/mvn"
  "/opt/homebrew/bin/mvn"
)

JAVA_HOME_RESOLVED="${JAVA_HOME:-}"
if [[ -z "${JAVA_HOME_RESOLVED}" ]]; then
  for candidate in "${JAVA_CANDIDATES[@]}"; do
    if [[ -x "${candidate}/bin/java" ]]; then
      JAVA_HOME_RESOLVED="${candidate}"
      break
    fi
  done
fi

if [[ -z "${JAVA_HOME_RESOLVED}" ]]; then
  echo "Unable to resolve JAVA_HOME for market-data-service-java tests." >&2
  exit 1
fi

MVN_BIN="${MVN_BIN:-}"
if [[ -z "${MVN_BIN}" ]]; then
  for candidate in "${MAVEN_CANDIDATES[@]}"; do
    if [[ -x "${candidate}" ]]; then
      MVN_BIN="${candidate}"
      break
    fi
  done
fi

if [[ -z "${MVN_BIN}" ]]; then
  if command -v mvn >/dev/null 2>&1; then
    MVN_BIN="$(command -v mvn)"
  else
    echo "Unable to find mvn for market-data-service-java tests." >&2
    exit 1
  fi
fi

export JAVA_HOME="${JAVA_HOME_RESOLVED}"
export PATH="${JAVA_HOME}/bin:$(dirname "${MVN_BIN}"):${PATH}"

cd "${SERVICE_DIR}"
exec "${MVN_BIN}" test "$@"
