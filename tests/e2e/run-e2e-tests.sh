#!/bin/bash
# E2E Test Execution Script
# Automates the complete E2E testing process

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
API_PORT=9100
RPC_PORT=9101
MAX_WAIT=30
REPORT_DIR="./reports"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Simple Admin Core - E2E Test Runner${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Function to check if port is in use
check_port() {
    local port=$1
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1 ; then
        return 0
    else
        return 1
    fi
}

# Function to wait for service
wait_for_service() {
    local url=$1
    local name=$2
    local max_attempts=$MAX_WAIT
    local attempt=1

    echo -e "${YELLOW}Waiting for $name to be ready...${NC}"

    while [ $attempt -le $max_attempts ]; do
        if curl -s -f "$url" > /dev/null 2>&1; then
            echo -e "${GREEN}✓ $name is ready${NC}"
            return 0
        fi
        echo -n "."
        sleep 1
        ((attempt++))
    done

    echo -e "\n${RED}✗ $name failed to start within ${max_attempts}s${NC}"
    return 1
}

# Check prerequisites
echo -e "${BLUE}1. Checking prerequisites...${NC}"

if ! command -v newman &> /dev/null; then
    echo -e "${RED}✗ Newman is not installed${NC}"
    echo -e "${YELLOW}Install with: npm install -g newman newman-reporter-htmlextra${NC}"
    exit 1
fi
echo -e "${GREEN}✓ Newman is installed${NC}"

if ! command -v lsof &> /dev/null; then
    echo -e "${YELLOW}⚠ lsof not available, cannot check ports${NC}"
fi

# Check if services are running
echo ""
echo -e "${BLUE}2. Checking services...${NC}"

if check_port $RPC_PORT; then
    echo -e "${GREEN}✓ RPC service is running on port $RPC_PORT${NC}"
else
    echo -e "${RED}✗ RPC service is not running on port $RPC_PORT${NC}"
    echo -e "${YELLOW}Start with: cd rpc && go run core.go -f etc/core.yaml${NC}"
    exit 1
fi

if check_port $API_PORT; then
    echo -e "${GREEN}✓ API service is running on port $API_PORT${NC}"
else
    echo -e "${RED}✗ API service is not running on port $API_PORT${NC}"
    echo -e "${YELLOW}Start with: cd api && go run core.go -f etc/core.yaml${NC}"
    exit 1
fi

# Wait for services to be fully ready
echo ""
echo -e "${BLUE}3. Verifying service health...${NC}"

if ! wait_for_service "http://localhost:$API_PORT/captcha" "API service"; then
    echo -e "${RED}API service health check failed${NC}"
    exit 1
fi

# Create reports directory
mkdir -p "$REPORT_DIR"

# Run E2E tests
echo ""
echo -e "${BLUE}4. Running E2E tests...${NC}"
echo ""

newman run user-module-e2e.postman_collection.json \
    --environment <(echo '{
        "name": "E2E Test Environment",
        "values": [
            {"key": "baseUrl", "value": "http://localhost:'$API_PORT'", "enabled": true},
            {"key": "testUsername", "value": "e2eTestUser_'$TIMESTAMP'", "enabled": true},
            {"key": "testEmail", "value": "e2etest_'$TIMESTAMP'@example.com", "enabled": true},
            {"key": "testPassword", "value": "TestPass123!", "enabled": true}
        ]
    }') \
    --reporters cli,htmlextra,json \
    --reporter-htmlextra-export "$REPORT_DIR/e2e-report-$TIMESTAMP.html" \
    --reporter-json-export "$REPORT_DIR/e2e-results-$TIMESTAMP.json" \
    --bail

TEST_EXIT_CODE=$?

echo ""
echo -e "${BLUE}========================================${NC}"

if [ $TEST_EXIT_CODE -eq 0 ]; then
    echo -e "${GREEN}✓ All E2E tests passed!${NC}"
    echo -e "${GREEN}Report: $REPORT_DIR/e2e-report-$TIMESTAMP.html${NC}"
else
    echo -e "${RED}✗ Some E2E tests failed${NC}"
    echo -e "${YELLOW}Check report: $REPORT_DIR/e2e-report-$TIMESTAMP.html${NC}"
fi

echo -e "${BLUE}========================================${NC}"

exit $TEST_EXIT_CODE
