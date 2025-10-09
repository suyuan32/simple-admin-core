#!/bin/bash
# Notion Tasks Database Auto-Update Script
# Updates all Spec-004 tasks to "Done" status with completion metadata
#
# Usage: ./notion-auto-update.sh <NOTION_API_KEY>
#
# Prerequisites:
# - jq (JSON processor): brew install jq
# - curl (HTTP client): pre-installed on macOS/Linux
# - Notion API key with write access to Tasks database
#
# Environment Variables (optional):
# - NOTION_API_KEY: Your Notion integration token
# - NOTION_DATABASE_ID: Tasks database ID (default: from CLAUDE.md)

set -e  # Exit on error
set -u  # Exit on undefined variable

# ============= Configuration =============

NOTION_API_VERSION="2022-06-28"
NOTION_API_BASE="https://api.notion.com/v1"

# Default database ID (replace with actual ID from your Notion workspace)
# You can find this in the URL when viewing your database:
# https://www.notion.so/<workspace>/<DATABASE_ID>?v=...
NOTION_DATABASE_ID="${NOTION_DATABASE_ID:-YOUR_DATABASE_ID_HERE}"

# ============= Color Output =============

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

function log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

function log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

function log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

function log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# ============= Validation =============

# Check for required tools
if ! command -v jq &> /dev/null; then
    log_error "jq is not installed. Please install: brew install jq"
    exit 1
fi

if ! command -v curl &> /dev/null; then
    log_error "curl is not installed."
    exit 1
fi

# Get Notion API key
if [ $# -eq 1 ]; then
    NOTION_API_KEY="$1"
elif [ -n "${NOTION_API_KEY:-}" ]; then
    log_info "Using NOTION_API_KEY from environment"
else
    log_error "Usage: $0 <NOTION_API_KEY>"
    log_info "Or set environment variable: export NOTION_API_KEY='your-key'"
    exit 1
fi

# Validate database ID
if [ "$NOTION_DATABASE_ID" = "YOUR_DATABASE_ID_HERE" ]; then
    log_error "Please set NOTION_DATABASE_ID environment variable or update script"
    log_info "Find your database ID in Notion URL: https://www.notion.so/<workspace>/<DATABASE_ID>?v=..."
    exit 1
fi

# ============= Task Definitions =============

# Spec-004 tasks to update (matching notion-task-updates.md)
declare -A TASKS=(
    ["ZH-TW-007"]="Extend core.proto with User RPC methods|Done|6|6|eac6379d"
    ["ZH-TW-008"]="Update user.proto for Proto-First generation|Done|4|4|eac6379d"
    ["USER-001"]="Implement authentication RPC logic (login, email, SMS)|Done|6|6|eac6379d"
    ["USER-002"]="Implement registration RPC logic (basic, email, SMS)|Done|4|4|eac6379d"
    ["USER-003"]="Implement password management RPC logic|Done|4|4|eac6379d"
    ["USER-004"]="Implement user info retrieval RPC logic|Done|4|4|eac6379d"
    ["USER-005"]="Implement token management RPC logic|Done|4|4|eac6379d"
    ["USER-006"]="Generate API file from user.proto|Done|2|2|eac6379d"
)

# ============= Helper Functions =============

function query_task_by_name() {
    local task_name="$1"

    log_info "Querying task: $task_name"

    local response=$(curl -s -X POST "$NOTION_API_BASE/databases/$NOTION_DATABASE_ID/query" \
        -H "Authorization: Bearer $NOTION_API_KEY" \
        -H "Notion-Version: $NOTION_API_VERSION" \
        -H "Content-Type: application/json" \
        --data "{
            \"filter\": {
                \"property\": \"Task ID\",
                \"rich_text\": {
                    \"equals\": \"$task_name\"
                }
            }
        }")

    # Check for errors
    if echo "$response" | jq -e '.object == "error"' > /dev/null 2>&1; then
        local error_msg=$(echo "$response" | jq -r '.message')
        log_error "Notion API error: $error_msg"
        return 1
    fi

    # Extract page ID
    local page_id=$(echo "$response" | jq -r '.results[0].id // empty')

    if [ -z "$page_id" ]; then
        log_warning "Task not found: $task_name"
        return 1
    fi

    echo "$page_id"
}

function update_task() {
    local page_id="$1"
    local task_name="$2"
    local status="$3"
    local estimated_hours="$4"
    local actual_hours="$5"
    local commit_hash="$6"

    log_info "Updating task $task_name to status: $status"

    # Current timestamp
    local completed_at=$(date -u +"%Y-%m-%dT%H:%M:%S.000Z")

    # Build update payload
    local payload=$(cat <<EOF
{
    "properties": {
        "Status": {
            "status": {
                "name": "$status"
            }
        },
        "Estimated Hours": {
            "number": $estimated_hours
        },
        "Actual Hours": {
            "number": $actual_hours
        },
        "Completed At": {
            "date": {
                "start": "$completed_at"
            }
        },
        "Commit Hash": {
            "rich_text": [
                {
                    "text": {
                        "content": "$commit_hash"
                    }
                }
            ]
        },
        "Progress": {
            "number": 100
        }
    }
}
EOF
)

    local response=$(curl -s -X PATCH "$NOTION_API_BASE/pages/$page_id" \
        -H "Authorization: Bearer $NOTION_API_KEY" \
        -H "Notion-Version: $NOTION_API_VERSION" \
        -H "Content-Type: application/json" \
        --data "$payload")

    # Check for errors
    if echo "$response" | jq -e '.object == "error"' > /dev/null 2>&1; then
        local error_msg=$(echo "$response" | jq -r '.message')
        log_error "Failed to update $task_name: $error_msg"
        return 1
    fi

    log_success "Updated $task_name"
    return 0
}

# ============= Main Execution =============

log_info "Starting Spec-004 Notion Tasks auto-update..."
log_info "Database ID: $NOTION_DATABASE_ID"
log_info "Tasks to update: ${#TASKS[@]}"
echo ""

SUCCESS_COUNT=0
FAILED_COUNT=0

for task_id in "${!TASKS[@]}"; do
    IFS='|' read -r description status estimated actual commit <<< "${TASKS[$task_id]}"

    log_info "Processing: $task_id - $description"

    # Query task
    page_id=$(query_task_by_name "$task_id")

    if [ $? -ne 0 ] || [ -z "$page_id" ]; then
        log_warning "Skipping $task_id (not found in Notion)"
        ((FAILED_COUNT++))
        continue
    fi

    # Update task
    if update_task "$page_id" "$task_id" "$status" "$estimated" "$actual" "$commit"; then
        ((SUCCESS_COUNT++))
    else
        ((FAILED_COUNT++))
    fi

    echo ""
    sleep 0.5  # Rate limiting
done

# ============= Summary =============

echo "========================================"
log_info "Update Summary"
echo "========================================"
log_success "Successfully updated: $SUCCESS_COUNT tasks"

if [ $FAILED_COUNT -gt 0 ]; then
    log_warning "Failed to update: $FAILED_COUNT tasks"
fi

echo "========================================"

if [ $FAILED_COUNT -eq 0 ]; then
    log_success "All tasks updated successfully! ✅"
    exit 0
else
    log_warning "Some tasks failed to update. Please check manually."
    exit 1
fi
