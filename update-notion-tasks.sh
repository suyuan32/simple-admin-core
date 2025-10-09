#!/bin/bash
# Interactive Notion Tasks Update Script
# Guides user through setup and executes notion-auto-update.sh

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m'

function print_header() {
    echo -e "${BOLD}${CYAN}"
    echo "╔════════════════════════════════════════════════════════════════╗"
    echo "║        Notion Tasks Auto-Update - Interactive Setup           ║"
    echo "╚════════════════════════════════════════════════════════════════╝"
    echo -e "${NC}"
}

function print_step() {
    echo -e "${BOLD}${BLUE}➜ $1${NC}"
}

function print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

function print_error() {
    echo -e "${RED}✗ $1${NC}"
}

function print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

function print_info() {
    echo -e "${CYAN}ℹ $1${NC}"
}

function pause() {
    echo ""
    read -p "按 Enter 繼續..."
    echo ""
}

# ============= Main Script =============

clear
print_header

echo "這個腳本會幫助您更新 8 個 Spec-004 Notion 任務到 'Done' 狀態。"
echo ""

# Step 1: Check prerequisites
print_step "步驟 1/4: 檢查系統依賴"

if ! command -v jq &> /dev/null; then
    print_error "jq 未安裝"
    echo "請安裝 jq："
    echo "  macOS: brew install jq"
    echo "  Linux: sudo apt-get install jq"
    exit 1
fi
print_success "jq 已安裝"

if ! command -v curl &> /dev/null; then
    print_error "curl 未安裝"
    exit 1
fi
print_success "curl 已安裝"

echo ""
pause

# Step 2: Get Notion API Key
clear
print_header
print_step "步驟 2/4: 獲取 Notion API Key"

echo "請按照以下步驟獲取 Notion API Key："
echo ""
echo "1. 訪問：https://www.notion.so/my-integrations"
echo "2. 點擊 '+ New integration'"
echo "3. 填寫信息："
echo "   - Name: Simple Admin Task Updater"
echo "   - Associated workspace: 選擇您的工作空間"
echo "   - Capabilities: 勾選 Read, Update, Insert"
echo "4. 點擊 'Submit'"
echo "5. 複製顯示的 Internal Integration Token"
echo "   (格式: secret_xxxxxx...)"
echo ""

read -p "請輸入您的 Notion API Key: " NOTION_API_KEY

if [ -z "$NOTION_API_KEY" ]; then
    print_error "API Key 不能為空"
    exit 1
fi

if [[ ! $NOTION_API_KEY == secret_* ]]; then
    print_warning "API Key 格式可能不正確（應該以 'secret_' 開頭）"
    read -p "是否繼續？(y/n): " confirm
    if [[ ! $confirm =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

print_success "API Key 已接收"
echo ""
pause

# Step 3: Get Database ID
clear
print_header
print_step "步驟 3/4: 獲取 Notion Database ID"

echo "請按照以下步驟獲取 Database ID："
echo ""
echo "1. 在 Notion 中打開您的 'Tasks' 數據庫頁面"
echo "2. 查看瀏覽器地址欄的 URL，格式如下："
echo ""
echo "   https://www.notion.so/workspace/DATABASE_ID?v=VIEW_ID"
echo "                                    ^^^^^^^^^^^^"
echo "                                    這就是 Database ID"
echo ""
echo "3. Database ID 是 32 個字符的十六進制字符串（不含破折號）"
echo ""
echo "示例："
echo "   URL: https://www.notion.so/myworkspace/1234567890abcdef1234567890abcdef?v=..."
echo "   Database ID: 1234567890abcdef1234567890abcdef"
echo ""

read -p "請輸入您的 Database ID: " NOTION_DATABASE_ID

if [ -z "$NOTION_DATABASE_ID" ]; then
    print_error "Database ID 不能為空"
    exit 1
fi

# Remove hyphens if present
NOTION_DATABASE_ID="${NOTION_DATABASE_ID//-/}"

if [ ${#NOTION_DATABASE_ID} -ne 32 ]; then
    print_warning "Database ID 長度不正確（應該是 32 個字符）"
    echo "您輸入的長度: ${#NOTION_DATABASE_ID}"
    read -p "是否繼續？(y/n): " confirm
    if [[ ! $confirm =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

print_success "Database ID 已接收"
echo ""
pause

# Step 4: Confirm connection
clear
print_header
print_step "步驟 4/4: 授權 Integration 訪問數據庫"

echo "在繼續之前，請確認您已完成以下步驟："
echo ""
echo "1. 在 Notion 中打開您的 'Tasks' 數據庫頁面"
echo "2. 點擊右上角 '...' (三個點)"
echo "3. 選擇 'Connections' → 'Connect to'"
echo "4. 找到並選擇 'Simple Admin Task Updater'"
echo "5. 點擊 'Confirm'"
echo ""

read -p "我已完成授權設置 (y/n): " authorized

if [[ ! $authorized =~ ^[Yy]$ ]]; then
    print_warning "請先完成授權設置，然後重新運行此腳本"
    exit 1
fi

print_success "授權已確認"
echo ""

# Step 5: Run the update script
clear
print_header
echo -e "${BOLD}準備更新 Notion 任務...${NC}"
echo ""
echo "即將更新以下 8 個任務到 'Done' 狀態："
echo ""
echo "  • ZH-TW-007: 擴展 core.proto 添加 User RPC 方法"
echo "  • ZH-TW-008: 更新 user.proto 支持 Proto-First 生成"
echo "  • USER-001: 實現認證 RPC 邏輯"
echo "  • USER-002: 實現註冊 RPC 邏輯"
echo "  • USER-003: 實現密碼管理 RPC 邏輯"
echo "  • USER-004: 實現用戶信息獲取 RPC 邏輯"
echo "  • USER-005: 實現令牌管理 RPC 邏輯"
echo "  • USER-006: 從 user.proto 生成 API 文件"
echo ""

read -p "確認執行更新？(y/n): " confirm_update

if [[ ! $confirm_update =~ ^[Yy]$ ]]; then
    print_warning "更新已取消"
    exit 0
fi

echo ""
print_step "正在執行更新..."
echo ""

# Export variables and run the actual update script
export NOTION_API_KEY
export NOTION_DATABASE_ID

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
UPDATE_SCRIPT="$SCRIPT_DIR/specs/004-user-module-proto-completion/notion-auto-update.sh"

if [ ! -f "$UPDATE_SCRIPT" ]; then
    print_error "找不到更新腳本: $UPDATE_SCRIPT"
    exit 1
fi

# Make script executable
chmod +x "$UPDATE_SCRIPT"

# Run the update script
if "$UPDATE_SCRIPT"; then
    echo ""
    print_success "所有任務已成功更新！ 🎉"
    echo ""
    echo "您可以在 Notion 中查看更新後的任務狀態。"
    echo ""

    # Ask if user wants to continue with tests
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    print_step "下一步：完成剩餘的單元測試"
    echo ""
    echo "當前測試進度："
    echo "  • 已完成：8/16 測試文件 (50%)"
    echo "  • 測試覆蓋率：~44% (目標 70%+)"
    echo "  • 剩餘工作：8 個測試文件，預估 2-3 小時"
    echo ""

    read -p "是否立即開始完成剩餘的單元測試？(y/n): " start_tests

    if [[ $start_tests =~ ^[Yy]$ ]]; then
        echo ""
        print_success "好的！我將開始實現剩餘的 8 個單元測試文件"
        echo ""
        echo "請告訴 Claude："
        echo "  \"請繼續完成剩餘的 8 個 RPC 單元測試文件\""
        echo ""
    else
        echo ""
        print_info "您可以稍後運行測試"
        echo "要開始測試，請告訴 Claude："
        echo "  \"請完成剩餘的 8 個 RPC 單元測試文件\""
        echo ""
    fi
else
    echo ""
    print_error "更新過程中出現錯誤"
    echo ""
    echo "請檢查以上錯誤信息，常見問題："
    echo "  1. API Key 不正確或已過期"
    echo "  2. Database ID 不正確"
    echo "  3. 未授權 Integration 訪問數據庫"
    echo "  4. Notion 任務 ID 不匹配"
    echo ""
    echo "詳細排錯指南請查看："
    echo "  $SCRIPT_DIR/NOTION-SETUP-GUIDE.md"
    echo ""
    exit 1
fi
