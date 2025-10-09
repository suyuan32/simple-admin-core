# 專案範圍說明

## 🎯 本專案 (simple-admin-core) - Monorepo 架構

**定位**: Full-stack 微服務系統（Monorepo）

**專案結構**:
```
simple-admin-core/
├── api/                    # Backend API Service
├── rpc/                    # Backend RPC Service
└── web/                    # Frontend (simple-admin-vben5-ui)
    └── apps/
        └── simple-admin-core/
```

**包含**:
- ✅ Backend API Service (REST API 閘道)
- ✅ Backend RPC Service (gRPC 業務邏輯)
- ✅ Frontend (Vue 3 + Vben5)
- ✅ Ent ORM Schema
- ✅ Proto 定義
- ✅ 全端 i18n 支援

**Git Repository**: https://github.com/chimerakang/simple-admin-core

---

## 📦 Monorepo 目錄結構

### Backend (`api/`, `rpc/`)
- API Service (Port 9100)
- RPC Service (Port 9101)
- i18n 支援: `api/internal/i18n/locale/zh-TW.json`

### Frontend (`web/`)
- Vue 3 + Vben5 前端應用
- i18n 支援: `web/apps/simple-admin-core/src/locales/langs/zh-TW/`
- Port: 5555 (dev), 80 (production)

---

## 📋 zh-TW 功能任務分配

### Backend 任務 ✅

| 任務 ID | 描述 | 檔案 | 狀態 | Commit |
|---------|------|------|------|--------|
| ZH-TW-001 | Backend zh-TW.json | `api/internal/i18n/locale/zh-TW.json` | ✅ 完成 | efd2d8d |
| ZH-TW-002 | i18n Translator | `api/internal/i18n/translator.go` | ✅ 完成 | e94e1d1 |
| ZH-TW-003 | Backend 測試 | `api/internal/i18n/translator_test.go` | ✅ 完成 | e94e1d1 |
| ZH-TW-007 | Ent Schema | `rpc/ent/schema/user.go` | ✅ 完成 | 451ed06 |
| ZH-TW-008 | Proto & API | `rpc/desc/user.proto` | ✅ 完成 | be16dc9 |

### Frontend 任務 ✅

| 任務 ID | 描述 | 檔案路徑 | 狀態 |
|---------|------|----------|------|
| ZH-TW-004 | 前端語言檔案 | `web/apps/simple-admin-core/src/locales/langs/zh-TW/*.json` | ✅ 完成 |
| ZH-TW-005 | Ant Design 整合 | `web/apps/simple-admin-core/src/locales/index.ts` | ✅ 完成 |
| ZH-TW-006 | 語言選擇器 UI | `web/packages/constants/src/core.ts` | ✅ 完成 |

---

## 🚀 Monorepo 開發流程

### 1. Backend 開發
```bash
cd simple-admin-core

# 修改 Backend 程式碼
vim api/internal/...
vim rpc/internal/...

# 提交 Backend 變更
git add api/ rpc/
git commit -m "feat: backend feature"
```

### 2. Frontend 開發
```bash
cd simple-admin-core/web

# 修改 Frontend 程式碼
vim apps/simple-admin-core/src/...

# 提交 Frontend 變更
cd ..
git add web/
git commit -m "feat: frontend feature"
```

### 3. 全端功能開發
```bash
# 同時修改 Backend 和 Frontend
git add api/ rpc/ web/
git commit -m "feat: full-stack feature"
```

---

## 📝 文檔規範

### Backend 文檔
- 說明 Backend API 支援 zh-TW
- 說明 Accept-Language header 處理
- 說明 Backend 錯誤訊息格式

### Frontend 文檔
- 說明 Frontend 語言檔案結構
- 說明 i18n 整合方式
- 說明語言切換流程

### 全端文檔
- 說明 Frontend ↔ Backend i18n 協作
- 說明完整的部署流程
- 說明 Monorepo 開發規範

---

## 🎯 Monorepo 優勢

1. **統一版本管理**: Frontend 和 Backend 在同一 Repository
2. **原子性提交**: 全端功能可在單一 commit 完成
3. **簡化 CI/CD**: 單一 Repository 的持續整合
4. **共享配置**: 共享 ESLint, Prettier, Git hooks 等配置

---

## 📊 Git Commits 規範

**本專案 commits 可以包含**:
- ✅ Backend: `api/`, `rpc/`, `ent/`, `proto/`
- ✅ Frontend: `web/`
- ✅ Docs: `docs/`, `README.md`
- ✅ Config: `Makefile`, `.gitignore`, etc.

**Commit 範例**:
```bash
# Backend only
git commit -m "feat(api): add zh-TW i18n support"

# Frontend only
git commit -m "feat(web): add zh-TW language files"

# Full-stack
git commit -m "feat: add zh-TW i18n support for full-stack"
```

---

## ✅ Monorepo 架構優勢總結

- **統一管理**: Frontend + Backend 在同一專案
- **協作便利**: 跨端功能開發更容易
- **版本同步**: 避免前後端版本不一致
- **CI/CD 簡化**: 單一流水線完成全端部署
