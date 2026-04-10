#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

echo "Starting notification-service-java in local demo mode..."
EXCHANGE_NOTIFICATION_KAFKA_ENABLED="${EXCHANGE_NOTIFICATION_KAFKA_ENABLED:-false}" \
SERVER_PORT="${NOTIFICATION_SERVER_PORT:-8094}" \
EXCHANGE_NOTIFICATION_GRPC_PORT="${NOTIFICATION_GRPC_PORT:-10084}" \
"${SCRIPT_DIR}/run_notification_service.sh"

echo
echo "Starting market-data-service-java in local demo mode..."
EXCHANGE_MARKET_DATA_REDIS_ENABLED="${EXCHANGE_MARKET_DATA_REDIS_ENABLED:-false}" \
SERVER_PORT="${MARKET_DATA_SERVER_PORT:-8095}" \
"${SCRIPT_DIR}/run_market_data_service.sh"

echo
echo "Portfolio demo is ready:"
echo "  Notification demo: http://127.0.0.1:${NOTIFICATION_SERVER_PORT:-8094}/demo/private-feed.html"
echo "  Market-data demo:  http://127.0.0.1:${MARKET_DATA_SERVER_PORT:-8095}/demo/market-data.html"
echo
echo "Seed helpers:"
echo "  /Users/adam/trading_system/scripts/seed/send_private_event_demo.sh"
echo "  /Users/adam/trading_system/scripts/seed/send_market_data_demo.sh"
