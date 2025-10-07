# 專案範圍說明

## 🎯 本專案 (simple-admin-core)

**定位**: Backend 微服務系統

**包含**:
- ✅ API Service (REST API 閘道)
- ✅ RPC Service (gRPC 業務邏輯)
- ✅ Ent ORM Schema
- ✅ Proto 定義
- ✅ Backend i18n 支援

**Git Repository**: https://github.com/suyuan32/simple-admin-core

---

## 🚫 不在本專案範圍

### Frontend (simple-admin-vben5-ui)

**定位**: Vue 3 前端應用（獨立專案）

**包含**:
- ❌ Vue 3 組件
- ❌ 前端語言檔案 (zh-TW.ts, zh-CN.ts, en-US.ts)
- ❌ UI 元件
- ❌ 路由與狀態管理

**Git Repository**: https://github.com/suyuan32/simple-admin-vben5-ui

---

## 📋 zh-TW 功能任務分配

### Backend 任務（本專案）✅

| 任務 ID | 描述 | 狀態 | Commit |
|---------|------|------|--------|
| ZH-TW-001 | Backend zh-TW.json | ✅ 完成 | efd2d8d |
| ZH-TW-002 | i18n Translator | ✅ 完成 | e94e1d1 |
| ZH-TW-003 | Backend 測試 | ✅ 完成 | e94e1d1 |
| ZH-TW-007 | Ent Schema | ✅ 完成 | 451ed06 |
| ZH-TW-008 | Proto & API | ✅ 完成 | be16dc9 |

### Frontend 任務（獨立專案）❌

| 任務 ID | 描述 | 狀態 | 說明 |
|---------|------|------|------|
| ZH-TW-004 | 前端語言檔案 | ❌ Out of Scope | 需在 simple-admin-vben5-ui 處理 |
| ZH-TW-005 | Ant Design 整合 | ❌ Out of Scope | 需在 simple-admin-vben5-ui 處理 |
| ZH-TW-006 | 語言選擇器 UI | ❌ Out of Scope | 需在 simple-admin-vben5-ui 處理 |

---

## ⚠️ 重要提醒

1. **不要在 Backend 專案中修改 Frontend 程式碼**
2. **不要在 Backend 專案文檔中指示操作 Frontend**
3. **Frontend 和 Backend 有各自的 Git Repository**
4. **部署文檔應明確區分 Backend 和 Frontend**

---

## 🔗 正確的協作流程

### Backend 開發者
1. Clone `simple-admin-core`
2. 修改 Backend 程式碼
3. Commit 到 `simple-admin-core` Repository

### Frontend 開發者
1. Clone `simple-admin-vben5-ui`
2. 修改 Frontend 程式碼
3. Commit 到 `simple-admin-vben5-ui` Repository

### 整合測試
1. Backend 開發者發布 API 變更
2. Frontend 開發者依據 API 文檔更新前端
3. 兩個專案**獨立**進行版本發布

---

## 📝 文檔規範

### Backend 文檔 ✅
- 可以說明 Backend API 如何支援 zh-TW
- 可以說明 Accept-Language header 處理
- 可以說明 Backend 錯誤訊息格式

### Backend 文檔 ❌
- **不應該**指示如何修改 Frontend 程式碼
- **不應該**包含 Frontend 部署步驟
- **不應該**假設 Frontend 和 Backend 在同一專案

---

## ✅ 修正措施

### 已完成
1. ✅ 更新 Notion 任務狀態（標記 Frontend 任務為 Out of Scope）
2. ✅ 刪除包含 Frontend 指示的部署文檔
3. ✅ 創建僅涵蓋 Backend 的部署文檔

### 文檔變更
- ❌ 刪除: `docs/DEPLOYMENT-ZH-TW.md` (包含 Frontend 指示)
- ❌ 刪除: `docs/LOCAL-DEVELOPMENT-ZH-TW.md` (包含 Frontend 指示)
- ✅ 新增: `docs/BACKEND-DEPLOYMENT-ZH-TW.md` (僅 Backend)
- ✅ 新增: `docs/PROJECT-SCOPE-CLARIFICATION.md` (本文檔)

---

## 📊 正確的 Git Commits

**本專案的所有 commits 應該只涉及 Backend**:

```bash
git log --oneline --since="2025-10-07"
```

輸出應該只包含 Backend 相關變更：
- ✅ api/internal/i18n/
- ✅ rpc/ent/schema/
- ✅ rpc/desc/*.proto
- ✅ docs/
- ❌ 不應包含任何 Frontend 路徑

---

## 🎯 總結

- **simple-admin-core** = Backend 專案
- **simple-admin-vben5-ui** = Frontend 專案
- **兩者獨立開發、獨立部署、獨立版本管理**
- **本專案文檔應僅涵蓋 Backend**
