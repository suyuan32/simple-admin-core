# Backend zh-TW 功能部署指南

## ⚠️ 重要說明

**本文檔僅涵蓋 Backend (simple-admin-core) 的部署**

- **Backend**: `simple-admin-core` (本專案) ✅
- **Frontend**: `simple-admin-vben5-ui` (獨立專案，不在本文檔範圍) ❌

---

## 📦 Backend 交付成果

| 項目 | 檔案 | 狀態 | Commit |
|------|------|------|--------|
| zh-TW 語言檔案 | `api/internal/i18n/locale/zh-TW.json` | ✅ | efd2d8d |
| i18n Translator | `api/internal/i18n/translator.go` | ✅ | e94e1d1 |
| 單元測試 | `api/internal/i18n/translator_test.go` (81.8% 覆蓋率) | ✅ | e94e1d1 |
| User schema | `rpc/ent/schema/user.go` (locale 欄位) | ✅ | 451ed06 |
| Proto 定義 | `rpc/desc/user.proto` | ✅ | be16dc9 |
| API 定義 | `api/desc/core/user.api` | ✅ | be16dc9 |

---

## 🚀 本地運行測試

### 1. 啟動基礎設施

```bash
cd deploy/docker-compose/postgresql_redis
docker-compose up -d
```

### 2. 啟動 Backend 服務

**終端 1 - RPC**：
```bash
go run rpc/core.go -f rpc/etc/core-postgres.yaml
```

**終端 2 - API**：
```bash
go run api/core.go -f api/etc/core.yaml
```

（註：需要修改 `api/etc/core.yaml` 的 `CoreRpc.Target` 為 `127.0.0.1:9101`）

### 3. 測試 zh-TW 功能

```bash
# 測試 zh-TW
curl -H "Accept-Language: zh-TW" \
  -X POST http://localhost:9100/api/v1/core/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test"}'

# 測試 zh-Hant 正規化
curl -H "Accept-Language: zh-Hant" \
  -X POST http://localhost:9100/api/v1/core/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test"}'
```

**預期結果**：錯誤訊息顯示繁體中文（台灣）

---

## 🐳 Docker 部署

### 構建映像

```bash
# Backend RPC
docker build -f Dockerfile-rpc -t your-username/core-rpc:zh-tw .

# Backend API
docker build -f Dockerfile-api -t your-username/core-api:zh-tw .
```

### 使用 Docker Compose

修改 `deploy/docker-compose/all_in_one/postgresql/docker-compose.yaml`：

```yaml
services:
  core-rpc:
    image: your-username/core-rpc:zh-tw

  core-api:
    image: your-username/core-api:zh-tw
```

啟動：

```bash
cd deploy/docker-compose/all_in_one/postgresql
docker-compose up -d
```

---

## ✅ 驗證

```bash
# 健康檢查
curl http://localhost:9100/health

# zh-TW 測試
curl -H "Accept-Language: zh-TW" http://localhost:9100/api/v1/...

# 單元測試
go test ./api/internal/i18n/... -v
```

---

## 📝 Backend 任務完成狀態

### ✅ 本專案已完成
- [ZH-TW-001] Backend zh-TW.json 語言檔案
- [ZH-TW-002] i18n Translator 實作
- [ZH-TW-003] Backend 單元測試
- [ZH-TW-007] Ent Schema locale 欄位
- [ZH-TW-008] Proto & API 定義

### ❌ 不在本專案範圍（Frontend）
- [ZH-TW-004] 前端語言檔案
- [ZH-TW-005] Ant Design 整合
- [ZH-TW-006] 語言選擇器 UI

---

## 🔗 相關資源

- [E2E 測試計劃](./zh-TW-E2E-TEST-PLAN.md)
- [QA 檢查清單](./zh-TW-MANUAL-QA-CHECKLIST.md)
- [Technical Plan](../specs/001-traditional-chinese-i18n/plan.md)
