#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
SERVICE_DIR="${REPO_ROOT}/services/market-data-service-java"
RUNTIME_DIR="${REPO_ROOT}/.demo-runtime"

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
  echo "Unable to resolve JAVA_HOME for market-data-service-java." >&2
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
    echo "Unable to find mvn for market-data-service-java." >&2
    exit 1
  fi
fi

export JAVA_HOME="${JAVA_HOME_RESOLVED}"
export PATH="${JAVA_HOME}/bin:$(dirname "${MVN_BIN}"):${PATH}"

mkdir -p "${RUNTIME_DIR}"

if [[ "${SKIP_BUILD:-0}" != "1" ]]; then
  (
    cd "${SERVICE_DIR}"
    "${MVN_BIN}" -DskipTests package >/dev/null
  )
fi

JAR_PATH="${SERVICE_DIR}/target/market-data-service-java-0.1.0-SNAPSHOT.jar"
if [[ ! -f "${JAR_PATH}" ]]; then
  echo "Market-data jar not found at ${JAR_PATH}" >&2
  exit 1
fi

LOG_PATH="${RUNTIME_DIR}/market-data-service.log"
PID_PATH="${RUNTIME_DIR}/market-data-service.pid"

if [[ -f "${PID_PATH}" ]]; then
  existing_pid="$(cat "${PID_PATH}")"
  if kill -0 "${existing_pid}" >/dev/null 2>&1; then
    echo "market-data-service-java is already running with pid ${existing_pid}" >&2
    exit 0
  fi
  rm -f "${PID_PATH}"
fi

SERVER_PORT="${SERVER_PORT:-8095}"
REDIS_ENABLED="${EXCHANGE_MARKET_DATA_REDIS_ENABLED:-false}"

(
  cd "${REPO_ROOT}"
  nohup env \
    SERVER_PORT="${SERVER_PORT}" \
    EXCHANGE_MARKET_DATA_REDIS_ENABLED="${REDIS_ENABLED}" \
    java -jar "${JAR_PATH}" \
    >"${LOG_PATH}" 2>&1 &
  echo $! > "${PID_PATH}"
)

for _ in {1..30}; do
  if curl -sS -o /dev/null "http://127.0.0.1:${SERVER_PORT}/demo/market-data.html" >/dev/null 2>&1; then
    echo "market-data-service-java started"
    echo "  pid: $(cat "${PID_PATH}")"
    echo "  http: http://127.0.0.1:${SERVER_PORT}"
    echo "  demo: http://127.0.0.1:${SERVER_PORT}/demo/market-data.html"
    echo "  log: ${LOG_PATH}"
    exit 0
  fi
  sleep 1
done

echo "market-data-service-java did not become healthy in time. Check ${LOG_PATH}" >&2
exit 1
