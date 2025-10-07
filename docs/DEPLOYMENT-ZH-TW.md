# zh-TW 功能部署指南

## 📋 部署概覽

Simple Admin zh-TW 功能已完成開發，本文檔說明如何部署並測試此功能。

---

## 🏗️ 部署架構

Simple Admin 採用**前後端分離**架構：

```
Backend Repository: simple-admin-core
├── API Service (Port 9100)      # REST API 閘道
└── RPC Service (Port 9101)      # gRPC 業務邏輯

Frontend Repository: simple-admin-vben5-ui
└── Vben5 App (Port 80/5555)     # Vue 3 前端應用
```

---

## 🗂️ 專案目錄結構

```
D:\Projects\
├── simple-admin-core\          # Backend (已完成 zh-TW)
│   ├── api\
│   │   └── internal\
│   │       └── i18n\
│   │           └── locale\
│   │               ├── zh.json
│   │               ├── en.json
│   │               └── zh-TW.json  ✅ 新增
│   ├── rpc\
│   └── deploy\
│
└── simple-admin-vben5-ui\      # Frontend (已完成 zh-TW)
    └── apps\
        └── simple-admin-core\
            └── src\
                └── locales\
                    └── langs\
                        ├── zh-CN\
                        ├── en-US\
                        └── zh-TW\  ✅ 新增
                            ├── common.json
                            ├── sys.json
                            ├── component.json
                            ├── fms.json
                            ├── mcms.json
                            └── page.json
```

---

## 🚀 部署方法

### 方法一：本地開發測試（推薦）⭐

**適用場景**：開發測試、功能驗證、QA 測試

**優點**：
- ✅ 快速啟動，無需構建 Docker 映像
- ✅ 程式碼變更即時生效
- ✅ 方便除錯

**步驟**：詳見 [`LOCAL-DEVELOPMENT-ZH-TW.md`](./LOCAL-DEVELOPMENT-ZH-TW.md)

**快速啟動**：
```bash
# 1. 啟動 PostgreSQL + Redis
cd D:\Projects\simple-admin-core\deploy\docker-compose\postgresql_redis
docker-compose up -d

# 2. 啟動 RPC 服務
cd D:\Projects\simple-admin-core
go run rpc/core.go -f rpc/etc/core-postgres.yaml

# 3. 啟動 API 服務（新終端）
go run api/core.go -f api/etc/core-local.yaml

# 4. 啟動前端（新終端）
cd D:\Projects\simple-admin-vben5-ui
pnpm run dev:core
```

訪問：http://localhost:5555

---

### 方法二：Docker Compose 部署（需更新映像）

**適用場景**：生產預覽、整合測試

**問題**：現有 Docker 映像（v1.7.0）**不包含 zh-TW 功能**

**解決方案**：構建包含 zh-TW 的新映像

#### 步驟 1: 使用現有 Docker Compose 配置

專案已包含完整配置：`deploy/docker-compose/all_in_one/postgresql/docker-compose.yaml`

但映像版本為 v1.7.0，**不包含 zh-TW 功能**。

#### 步驟 2: 構建包含 zh-TW 的映像

```bash
cd D:\Projects\simple-admin-core

# 構建 RPC 映像
docker build -f Dockerfile-rpc -t your-username/core-rpc:zh-tw .

# 構建 API 映像
docker build -f Dockerfile-api -t your-username/core-api:zh-tw .
```

```bash
cd D:\Projects\simple-admin-vben5-ui

# 構建前端映像
docker build -t your-username/backend-ui-vben5:zh-tw .
```

#### 步驟 3: 修改 Docker Compose 配置

編輯 `deploy/docker-compose/all_in_one/postgresql/docker-compose.yaml`：

```yaml
services:
  # ... (PostgreSQL, Redis 保持不變)

  core-rpc:
    image: your-username/core-rpc:zh-tw  # 改為你的映像
    # ...

  core-api:
    image: your-username/core-api:zh-tw  # 改為你的映像
    # ...

  backend-ui:
    image: your-username/backend-ui-vben5:zh-tw  # 改為你的映像
    # ...
```

#### 步驟 4: 啟動服務

```bash
cd D:\Projects\simple-admin-core\deploy\docker-compose\all_in_one\postgresql
docker-compose up -d
```

訪問：http://localhost

---

### 方法三：Kubernetes 部署

**適用場景**：生產環境

**前置需求**：
- 已構建包含 zh-TW 的 Docker 映像
- Kubernetes 叢集可用
- kubectl 已配置

#### 步驟 1: 更新 K8s 配置

編輯 `deploy/k8s/core-rpc.yaml`, `deploy/k8s/core-api.yaml`, `deploy/k8s/backend-ui.yaml`：

```yaml
spec:
  containers:
  - name: core-rpc
    image: your-username/core-rpc:zh-tw  # 更新映像版本
```

#### 步驟 2: 部署

```bash
cd D:\Projects\simple-admin-core\deploy\k8s

# 部署所有服務
kubectl apply -f .

# 檢查部署狀態
kubectl get pods
kubectl get svc
```

---

## ✅ 部署驗證

### 1. 檢查服務狀態

**Docker Compose**：
```bash
docker-compose ps
```

**Kubernetes**：
```bash
kubectl get pods
kubectl get svc
```

**本地運行**：
```bash
# 檢查 RPC
curl http://localhost:9101

# 檢查 API
curl http://localhost:9100/health
```

### 2. 驗證 zh-TW 功能

#### 前端測試
1. 訪問前端 URL
2. 點擊語言選擇器（右上角地球圖示）
3. 選擇「繁體中文（台灣）」
4. 驗證頁面文字變更

**預期結果**：
- ✅ 語言選擇器顯示「繁體中文（台灣）」
- ✅ 頁面文字使用台灣術語：使用者、資料庫、選單、檔案

#### 後端 API 測試

測試 Accept-Language header：

```bash
# 測試 zh-TW
curl -H "Accept-Language: zh-TW" http://localhost:9100/api/v1/core/user/login

# 測試 locale 正規化（zh-Hant → zh-TW）
curl -H "Accept-Language: zh-Hant" http://localhost:9100/api/v1/core/user/login

# 測試 zh-Hant-TW → zh-TW
curl -H "Accept-Language: zh-Hant-TW" http://localhost:9100/api/v1/core/user/login
```

**預期結果**：
- ✅ API 錯誤訊息顯示繁體中文（台灣）
- ✅ zh-Hant, zh-Hant-TW 自動正規化為 zh-TW

---

## 🧪 完整測試流程

### E2E 測試

詳見 [`zh-TW-E2E-TEST-PLAN.md`](./zh-TW-E2E-TEST-PLAN.md)

**測試案例**：
- TC-001: 語言選擇器顯示
- TC-002: 切換到繁體中文
- TC-003: 頁面文字檢查
- TC-004: API 錯誤訊息檢查
- TC-005: 使用者偏好設定儲存
- TC-006~TC-010: 其他功能測試

### 人工 QA 測試

詳見 [`zh-TW-MANUAL-QA-CHECKLIST.md`](./zh-TW-MANUAL-QA-CHECKLIST.md)

**檢查項目**：
- 台灣術語正確性（使用者 vs 用户）
- UI 文字自然度
- 錯誤訊息適當性
- 日期時間格式

---

## 🐛 常見問題

### 問題 1：語言選擇器沒有 zh-TW 選項

**原因**：前端程式碼版本不包含 zh-TW

**解決**：
```bash
# 確認前端有 zh-TW 相關 commits
cd D:\Projects\simple-admin-vben5-ui
git log --oneline | grep -i zh-tw

# 如果沒有，pull 最新 develop 分支
git pull origin develop
```

### 問題 2：後端 API 返回簡體中文

**原因**：
1. zh-TW.json 未包含在 Docker 映像中
2. I18nConf.Dir 配置錯誤

**解決**：
```bash
# 檢查 Dockerfile-api 是否複製 locale 目錄
grep "locale" Dockerfile-api
# 應該看到: COPY ./api/internal/i18n/locale/ ./etc/locale/

# 檢查 API 配置
grep -A2 "I18nConf" api/etc/core-local.yaml
# 應設定: Dir: ./api/internal/i18n/locale
```

### 問題 3：zh-Hant 未正規化為 zh-TW

**原因**：translator.go 的 NormalizeLocale() 未生效

**解決**：
```bash
# 測試 NormalizeLocale 函數
cd D:\Projects\simple-admin-core
go test ./api/internal/i18n -run TestNormalizeLocale -v
```

### 問題 4：Docker 映像構建失敗

**常見錯誤**：
```
ERROR: failed to solve: failed to compute cache key
```

**解決**：
```bash
# 清除 Docker 快取
docker builder prune

# 重新構建（不使用快取）
docker build --no-cache -f Dockerfile-api -t your-username/core-api:zh-tw .
```

---

## 📊 部署檢查清單

### 部署前檢查
- [ ] Backend 程式碼包含 zh-TW.json
- [ ] Frontend 程式碼包含 6 個 zh-TW 語言檔案
- [ ] Ent schema 包含 locale 欄位
- [ ] Proto 定義包含 locale 欄位
- [ ] API 定義包含 Locale 欄位

### 部署後驗證
- [ ] 前端語言選擇器顯示「繁體中文（台灣）」
- [ ] 切換語言後頁面文字變更
- [ ] API 錯誤訊息返回 zh-TW
- [ ] Locale 正規化功能正常
- [ ] 使用者語言偏好可儲存

### E2E 測試
- [ ] 執行完整 E2E 測試計劃（10 個測試案例）
- [ ] 所有測試案例通過

### 人工 QA
- [ ] 台灣母語者審查術語正確性
- [ ] UI 文字自然度評分 ≥ 4/5
- [ ] 錯誤訊息適當性評分 ≥ 4/5

---

## 🎯 推薦部署策略

### 開發/測試階段
**使用方法一**：本地開發測試
- 快速迭代
- 方便除錯
- 詳見 [`LOCAL-DEVELOPMENT-ZH-TW.md`](./LOCAL-DEVELOPMENT-ZH-TW.md)

### 整合測試階段
**使用方法二**：Docker Compose
- 構建包含 zh-TW 的映像
- 使用 `docker-compose up` 快速部署
- 執行完整 E2E 測試

### 生產環境
**使用方法三**：Kubernetes
- 構建生產級映像
- 發布到 Docker Hub 或私有 Registry
- 使用 K8s 進行滾動更新

---

## 📝 下一步

部署完成後：

1. **執行單元測試**
   ```bash
   go test ./api/internal/i18n/... -v
   ```

2. **執行 E2E 測試**
   - 參考 `zh-TW-E2E-TEST-PLAN.md`

3. **人工 QA 測試**
   - 參考 `zh-TW-MANUAL-QA-CHECKLIST.md`
   - 需要台灣母語者協助

4. **準備發布**
   - 所有測試通過後
   - 發布新版本映像（例如 v1.7.1-zh-tw）
   - 更新生產環境

---

## 🔗 相關文檔

- [本地開發測試指南](./LOCAL-DEVELOPMENT-ZH-TW.md) - 推薦優先閱讀
- [E2E 測試計劃](./zh-TW-E2E-TEST-PLAN.md)
- [人工 QA 檢查清單](./zh-TW-MANUAL-QA-CHECKLIST.md)
- [Specification](../specs/001-traditional-chinese-i18n/spec.md)
- [Technical Plan](../specs/001-traditional-chinese-i18n/plan.md)

---

**部署支援**：
- GitHub Issues: https://github.com/suyuan32/simple-admin-core/issues
- Discord: https://discord.gg/simple-admin
