# zh-TW 功能部署指南（Monorepo）

## 📋 專案架構

本專案採用 **Monorepo** 架構，Frontend 和 Backend 在同一 Repository：

```
simple-admin-core/ (Monorepo)
├── api/                              # Backend API Service
│   └── internal/i18n/locale/
│       └── zh-TW.json               # Backend 繁體中文
├── rpc/                              # Backend RPC Service
└── web/                              # Frontend
    └── apps/simple-admin-core/
        └── src/locales/langs/
            └── zh-TW/               # Frontend 繁體中文
                ├── common.json
                ├── sys.json
                ├── component.json
                ├── fms.json
                ├── mcms.json
                └── page.json
```

---

## 📦 zh-TW 功能交付成果

### Backend ✅
| 項目 | 檔案 | 狀態 | Commit |
|------|------|------|--------|
| zh-TW 語言檔案 | `api/internal/i18n/locale/zh-TW.json` | ✅ | efd2d8d |
| i18n Translator | `api/internal/i18n/translator.go` | ✅ | e94e1d1 |
| 單元測試 | `api/internal/i18n/translator_test.go` | ✅ | e94e1d1 |
| User schema | `rpc/ent/schema/user.go` | ✅ | 451ed06 |
| Proto & API | `rpc/desc/user.proto`, `api/desc/core/user.api` | ✅ | be16dc9 |

### Frontend ✅
| 項目 | 檔案路徑 | 狀態 |
|------|----------|------|
| 6 個語言檔案 | `web/apps/simple-admin-core/src/locales/langs/zh-TW/*.json` | ✅ |
| Ant Design 整合 | `web/apps/simple-admin-core/src/locales/index.ts` | ✅ |
| 語言選擇器 | `web/packages/constants/src/core.ts` | ✅ |

---

## 🚀 部署方法

### 方法一：本地開發（推薦）

#### 1. 啟動基礎設施

```bash
cd deploy/docker-compose/postgresql_redis
docker-compose up -d
```

#### 2. 啟動 Backend

**終端 1 - RPC**:
```bash
go run rpc/core.go -f rpc/etc/core-postgres.yaml
```

**終端 2 - API**:
```bash
go run api/core.go -f api/etc/core.yaml
```
（註：需修改 `CoreRpc.Target` 為 `127.0.0.1:9101`）

#### 3. 啟動 Frontend

**終端 3**:
```bash
cd web
pnpm install
pnpm run dev:core
```

#### 4. 訪問系統

- Frontend: http://localhost:5555
- Backend API: http://localhost:9100

#### 5. 測試 zh-TW 功能

1. 開啟瀏覽器訪問 http://localhost:5555
2. 點擊語言選擇器（右上角地球圖示）
3. 選擇「繁體中文（台灣）」
4. 驗證：
   - ✅ UI 顯示繁體中文（台灣用語）
   - ✅ API 錯誤訊息顯示繁體中文

---

### 方法二：Docker Compose（推薦用於開發/測試環境）

#### 選項 A：從 Monorepo 建置 Frontend（推薦）

使用 Monorepo Docker Compose 配置，會從源碼建置 Frontend：

```bash
cd deploy/docker-compose/monorepo
docker-compose up -d
```

**優點**：
- ✅ Frontend 從 `web/` 源碼建置，包含最新變更
- ✅ 自動包含 zh-TW 語言檔案
- ✅ 適合開發和測試

訪問：http://localhost

#### 選項 B：使用預建 Images

使用 all-in-one Docker Compose 配置：

```bash
cd deploy/docker-compose/all_in_one/postgresql
docker-compose up -d
```

此配置包含：
- PostgreSQL 資料庫
- Redis
- Core RPC Service
- Core API Service
- Backend UI (Frontend - 預建 image)

訪問：http://localhost

**注意**：此方法使用預建的 Docker images，可能不包含最新的 zh-TW 語言檔案變更。

#### 選項 C：自行建置所有映像

```bash
# Backend RPC
docker build -f Dockerfile-rpc -t simple-admin/core-rpc:zh-tw .

# Backend API
docker build -f Dockerfile-api -t simple-admin/core-api:zh-tw .

# Frontend (from Monorepo)
docker build -f Dockerfile.web -t simple-admin/backend-ui-vben5:zh-tw .
```

修改 docker-compose.yaml 使用自建映像後啟動。

---

### 方法三：Kubernetes 部署

#### 1. 更新 K8s 配置

```yaml
# deploy/k8s/core-rpc.yaml
spec:
  containers:
  - name: core-rpc
    image: simple-admin/core-rpc:zh-tw

# deploy/k8s/core-api.yaml
spec:
  containers:
  - name: core-api
    image: simple-admin/core-api:zh-tw

# deploy/k8s/backend-ui.yaml
spec:
  containers:
  - name: backend-ui
    image: simple-admin/frontend:zh-tw
```

#### 2. 部署

```bash
cd deploy/k8s
kubectl apply -f .
```

---

## 🔧 Makefile 快速指令

本專案提供 Makefile 簡化常見操作：

### Backend 指令

```bash
# 產生程式碼
make gen-api          # 生成 API 代碼
make gen-rpc          # 生成 RPC 代碼
make gen-ent          # 生成 Ent ORM 代碼

# 建置
make build-linux      # 建置 Linux 執行檔
make build-win        # 建置 Windows 執行檔
make build-mac        # 建置 macOS 執行檔

# Docker
make docker           # 建置 Backend Docker images
make publish-docker   # 發布 Docker images

# 測試
make test             # 執行測試
make lint             # 執行代碼檢查
```

### Frontend 指令（Monorepo）

```bash
# 安裝依賴
make install-web      # 安裝 Frontend 依賴

# 開發
make dev-web          # 啟動 Frontend 開發服務器 (port 5666)

# 建置
make build-web        # 建置 Frontend 生產版本

# Docker
make docker-web       # 建置 Frontend Docker image
make publish-docker-web  # 發布 Frontend Docker image
```

---

## ✅ 功能驗證

### 1. Backend 驗證

```bash
# 健康檢查
curl http://localhost:9100/health

# zh-TW API 測試
curl -H "Accept-Language: zh-TW" \
  -X POST http://localhost:9100/api/v1/core/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test"}'

# 預期：錯誤訊息顯示繁體中文
# {"code":10001,"msg":"使用者不存在"}
```

### 2. Frontend 驗證

開啟瀏覽器：http://localhost:5555 或 http://localhost

**檢查項目**：
- ✅ 語言選擇器顯示「繁體中文（台灣）」選項
- ✅ 切換後 UI 顯示台灣用語（使用者、選單、資料庫、檔案）
- ✅ 表單、按鈕、提示訊息顯示繁體中文
- ✅ Ant Design 組件（日期選擇器、分頁）顯示繁體中文

### 3. 整合測試

- ✅ 使用者登入後語言偏好儲存至資料庫
- ✅ 登出後重新登入保持語言設定
- ✅ Frontend 和 Backend 錯誤訊息一致使用繁體中文

---

## 🐛 常見問題

### 問題 1：Frontend 語言選項沒有 zh-TW

**原因**：`web/` 目錄未包含 Frontend 程式碼

**解決**：
```bash
# 確認 Frontend 程式碼存在
ls web/apps/simple-admin-core/src/locales/langs/zh-TW/

# 如果不存在，請從 simple-admin-vben5-ui 複製到 web/
```

### 問題 2：Backend API 返回簡體中文

**檢查**：
```bash
# 確認 zh-TW.json 存在
ls api/internal/i18n/locale/zh-TW.json

# 測試 Accept-Language header
curl -v -H "Accept-Language: zh-TW" http://localhost:9100/api/v1/...
```

### 問題 3：pnpm 命令找不到

**解決**：
```bash
# 安裝 pnpm
npm install -g pnpm

# 或使用 npx
npx pnpm install
npx pnpm run dev:core
```

---

## 📊 部署檢查清單

### 部署前
- [ ] Backend 程式碼包含 `api/internal/i18n/locale/zh-TW.json`
- [ ] Frontend 程式碼存在於 `web/` 目錄
- [ ] Frontend 包含 `web/apps/simple-admin-core/src/locales/langs/zh-TW/` 目錄
- [ ] 所有 zh-TW 語言檔案存在（6 個 JSON 檔案）

### 部署後
- [ ] Backend RPC 服務啟動成功
- [ ] Backend API 服務啟動成功
- [ ] Frontend 服務啟動成功
- [ ] Backend API 返回繁體中文錯誤訊息
- [ ] Frontend UI 顯示繁體中文
- [ ] 語言選擇器功能正常

---

## 🧪 測試

### 單元測試

```bash
# Backend 測試
go test ./api/internal/i18n/... -v

# 測試覆蓋率
go test ./api/internal/i18n/... -cover
```

**預期覆蓋率**: > 80%

### E2E 測試

參考：`docs/zh-TW-E2E-TEST-PLAN.md`

**主要測試案例**：
- TC-001: 語言選擇器顯示
- TC-002: 切換到繁體中文
- TC-003: 頁面文字檢查
- TC-004: API 錯誤訊息檢查
- TC-005: 使用者偏好設定儲存

### 人工 QA

參考：`docs/zh-TW-MANUAL-QA-CHECKLIST.md`

**檢查重點**：
- 台灣術語正確性（使用者 vs 用户）
- UI 文字自然度
- 錯誤訊息適當性

---

## 🔗 相關文檔

- [專案範圍說明 (Monorepo)](./PROJECT-SCOPE-CLARIFICATION.md)
- [E2E 測試計劃](./zh-TW-E2E-TEST-PLAN.md)
- [人工 QA 檢查清單](./zh-TW-MANUAL-QA-CHECKLIST.md)
- [Technical Plan](../specs/001-traditional-chinese-i18n/plan.md)

---

## 📞 支援

- GitHub Issues: https://github.com/chimerakang/simple-admin-core/issues
- Discord: https://discord.gg/simple-admin

---

## 🎯 Monorepo 部署優勢

1. **統一部署**: 一次構建完成 Frontend + Backend
2. **版本一致**: 避免前後端版本不匹配
3. **簡化 CI/CD**: 單一流水線完成全端測試與部署
4. **原子性發布**: Frontend 和 Backend 功能同步上線
