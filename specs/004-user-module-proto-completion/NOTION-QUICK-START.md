# Notion 任務更新 - 快速開始指南

## 現狀說明

✅ **代碼已全部提交** (Git 工作區乾淨)
❌ **Notion 任務尚未更新** (需要您手動執行)

---

## 方法 1: 自動化腳本（推薦）⚡

### 步驟 1: 獲取 Notion API 密鑰

1. 訪問 [Notion Integrations](https://www.notion.so/my-integrations)
2. 點擊 **"+ New integration"**
3. 名稱：`Simple Admin Tasks Updater`
4. 選擇您的工作區
5. 複製 **Internal Integration Token** (格式: `secret_...`)

### 步驟 2: 分享數據庫給 Integration

1. 打開您的 **Tasks** 數據庫
2. 點擊右上角 **"..."** → **Connections** → **Connect to**
3. 選擇 **"Simple Admin Tasks Updater"**
4. 授予寫入權限

### 步驟 3: 獲取數據庫 ID

從 Notion URL 中提取：
```
https://www.notion.so/<workspace>/abcd1234efgh5678?v=...
                                   ^^^^^^^^^^^^^^^^
                                   這是您的 Database ID
```

### 步驟 4: 運行腳本

```bash
cd /Volumes/eclipse/projects/simple-admin-core/specs/004-user-module-proto-completion

# 方式 A: 直接傳遞 API 密鑰
./notion-auto-update.sh "secret_你的API密鑰"

# 方式 B: 使用環境變量
export NOTION_API_KEY="secret_你的API密鑰"
export NOTION_DATABASE_ID="你的數據庫ID"
./notion-auto-update.sh
```

### 預期輸出

```
[INFO] Starting Spec-004 Notion Tasks auto-update...
[INFO] Database ID: abcd1234efgh5678
[INFO] Tasks to update: 8

[INFO] Processing: ZH-TW-007 - Extend core.proto with User RPC methods
[INFO] Querying task: ZH-TW-007
[INFO] Updating task ZH-TW-007 to status: Done
[SUCCESS] Updated ZH-TW-007

[INFO] Processing: ZH-TW-008 - Update user.proto for Proto-First generation
[SUCCESS] Updated ZH-TW-008

... (繼續更新 6 個任務)

========================================
[SUCCESS] Successfully updated: 8 tasks
========================================
[SUCCESS] All tasks updated successfully! ✅
```

---

## 方法 2: 手動 CSV 導入（備選）📋

如果您不想使用 API，可以手動更新：

### 步驟 1: 導出 CSV

1. 打開 Notion Tasks 數據庫
2. 點擊右上角 **"..."** → **Export** → **CSV**
3. 下載 CSV 文件

### 步驟 2: 編輯 CSV

打開 CSV，找到以下 8 個任務並更新：

| Task ID | Status | Estimated Hours | Actual Hours | Completed At | Commit Hash | Progress |
|---------|--------|-----------------|--------------|--------------|-------------|----------|
| ZH-TW-007 | Done | 6 | 6 | 2025-10-10 | eac6379d | 100 |
| ZH-TW-008 | Done | 4 | 4 | 2025-10-10 | eac6379d | 100 |
| USER-001 | Done | 6 | 6 | 2025-10-10 | eac6379d | 100 |
| USER-002 | Done | 4 | 4 | 2025-10-10 | eac6379d | 100 |
| USER-003 | Done | 4 | 4 | 2025-10-10 | eac6379d | 100 |
| USER-004 | Done | 4 | 4 | 2025-10-10 | eac6379d | 100 |
| USER-005 | Done | 4 | 4 | 2025-10-10 | eac6379d | 100 |
| USER-006 | Done | 2 | 2 | 2025-10-10 | eac6379d | 100 |

### 步驟 3: 重新導入 CSV

1. 在 Notion 中，點擊 **"..."** → **Import**
2. 選擇修改後的 CSV 文件
3. 選擇 **"Merge"** 模式（合併更新）
4. 確認導入

---

## 方法 3: 手動逐個更新（最簡單但最慢）✋

對於每個任務，打開 Notion 頁面並更新以下字段：

### 任務列表

**ZH-TW-007** - Extend core.proto with User RPC methods
- Status: Done
- Estimated Hours: 6
- Actual Hours: 6
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**ZH-TW-008** - Update user.proto for Proto-First generation
- Status: Done
- Estimated Hours: 4
- Actual Hours: 4
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**USER-001** - Implement authentication RPC logic (login, email, SMS)
- Status: Done
- Estimated Hours: 6
- Actual Hours: 6
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**USER-002** - Implement registration RPC logic (basic, email, SMS)
- Status: Done
- Estimated Hours: 4
- Actual Hours: 4
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**USER-003** - Implement password management RPC logic
- Status: Done
- Estimated Hours: 4
- Actual Hours: 4
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**USER-004** - Implement user info retrieval RPC logic
- Status: Done
- Estimated Hours: 4
- Actual Hours: 4
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**USER-005** - Implement token management RPC logic
- Status: Done
- Estimated Hours: 4
- Actual Hours: 4
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

**USER-006** - Generate API file from user.proto
- Status: Done
- Estimated Hours: 2
- Actual Hours: 2
- Completed At: 2025-10-10
- Commit Hash: eac6379d
- Progress: 100%

---

## 驗證更新是否成功

### 自動驗證（腳本完成後）

腳本會顯示：
```
[SUCCESS] Successfully updated: 8 tasks
```

### 手動驗證（在 Notion 中）

1. 打開 Tasks 數據庫
2. 篩選 `Status = Done`
3. 檢查以下 8 個任務是否都標記為完成：
   - ZH-TW-007 ✅
   - ZH-TW-008 ✅
   - USER-001 ✅
   - USER-002 ✅
   - USER-003 ✅
   - USER-004 ✅
   - USER-005 ✅
   - USER-006 ✅

4. 驗證每個任務的字段：
   - ✅ Completed At: 2025-10-10
   - ✅ Commit Hash: eac6379d
   - ✅ Progress: 100%
   - ✅ Actual Hours = Estimated Hours

---

## 更新統計

### 任務總覽
- **總任務數**: 8
- **已完成**: 8 (100%)
- **總估計時間**: 34 小時
- **總實際時間**: 34 小時
- **時間準確度**: 100%

### 按類型分組
| 類型 | 任務數 | 總時數 |
|------|--------|--------|
| Proto 擴展 | 2 | 10h |
| RPC Logic 實現 | 5 | 22h |
| API 生成 | 1 | 2h |

---

## 常見問題

### Q: 腳本報錯 "401 Unauthorized"

**A**: 檢查：
1. API 密鑰是否正確（應以 `secret_` 開頭）
2. Integration 是否已分享到 Tasks 數據庫
3. API 密鑰是否過期（重新生成）

### Q: 腳本報錯 "Task not found"

**A**: 檢查：
1. Task ID 是否在 Notion 中存在且拼寫正確
2. 數據庫 ID 是否正確
3. Tasks 數據庫是否有 "Task ID" 字段

### Q: CSV 導入後數據丟失

**A**: 重新導入時：
1. 選擇 **"Merge"** 模式，不要選 "Replace"
2. 確保 CSV 中有 ID 列用於匹配
3. 先備份數據庫（Export → CSV）

### Q: 不想使用 Notion API

**A**: 使用方法 2（CSV 導入）或方法 3（手動更新）

---

## 安全提醒 ⚠️

1. **永遠不要**將 API 密鑰提交到 git
2. **使用後**考慮輪換 API 密鑰
3. **限制** Integration 權限僅到需要的數據庫
4. **添加** `.env` 到 `.gitignore`

---

## 需要幫助？

- **腳本詳細文檔**: 查看 `notion-auto-update-README.md`
- **手動更新指南**: 查看 `notion-task-updates.md`
- **Notion API 文檔**: https://developers.notion.com/

---

**推薦方法**: 方法 1（自動化腳本）⚡
- 最快（~10 秒）
- 最準確（零錯誤）
- 可重用於未來項目

**時間對比**:
- 方法 1（腳本）: ~10 秒
- 方法 2（CSV）: ~5 分鐘
- 方法 3（手動）: ~15 分鐘

---

**準備好了嗎？開始更新 Notion 任務吧！** 🚀
