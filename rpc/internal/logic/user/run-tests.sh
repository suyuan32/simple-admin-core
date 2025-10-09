#!/bin/bash
# RPC Unit Test Runner
# Runs all User logic tests and generates coverage report

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}User Module RPC Unit Tests${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Run tests with coverage
echo -e "${YELLOW}Running tests...${NC}"
go test -v -coverprofile=coverage.out -covermode=atomic ./...

echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${YELLOW}Coverage Summary${NC}"
echo -e "${BLUE}========================================${NC}"

# Display coverage summary
go tool cover -func=coverage.out | tail -1

echo ""
echo -e "${YELLOW}Generating HTML coverage report...${NC}"
go tool cover -html=coverage.out -o coverage.html

echo -e "${GREEN}✓ Coverage report generated: coverage.html${NC}"
echo ""
echo -e "${YELLOW}To view coverage report:${NC}"
echo -e "  open coverage.html"
echo ""
echo -e "${BLUE}========================================${NC}"
