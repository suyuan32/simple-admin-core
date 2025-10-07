# zh-TW 功能本地開發測試指南

## 📋 前置需求

### 必要軟體
- Go 1.25+
- Node.js 20+
- pnpm 9+
- PostgreSQL 16 或 MySQL 8+
- Redis 7+

### 專案結構
```
D:\Projects\
├── simple-admin-core\          # Backend (本專案)
└── simple-admin-vben5-ui\      # Frontend (已 clone)
```

---

## 🚀 快速啟動（推薦）

### 選項 A：使用現有 Docker Compose（需更新映像）

**問題**：現有的 `deploy/docker-compose/all_in_one/postgresql/docker-compose.yaml` 使用 v1.7.0 映像，**不包含 zh-TW 功能**。

**解決方案**：
1. 本地構建新映像
2. 使用選項 B（本地直接運行）

---

### 選項 B：本地運行服務（推薦用於開發測試）

#### 步驟 1：啟動基礎設施

使用現有的 PostgreSQL + Redis Docker Compose：

```bash
cd D:\Projects\simple-admin-core\deploy\docker-compose\postgresql_redis
docker-compose up -d
```

這會啟動：
- PostgreSQL (Port 5432)
- Redis (Port 6379)

#### 步驟 2：配置 RPC 服務

創建 PostgreSQL 配置檔案：

```bash
# 複製並編輯 RPC 配置
cd D:\Projects\simple-admin-core\rpc\etc
cp core.yaml core-postgres.yaml
```

編輯 `rpc/etc/core-postgres.yaml`：

```yaml
Name: core.rpc
ListenOn: 0.0.0.0:9101

DatabaseConf:
  Type: postgres                    # 改為 postgres
  Host: 127.0.0.1
  Port: 5432                        # PostgreSQL port
  DBName: simple_admin
  Username: postgres                # PostgreSQL username
  Password: simple-admin.           # PostgreSQL password
  MaxOpenConn: 100
  SSLMode: disable
  CacheTime: 5

RedisConf:
  Host: 127.0.0.1:6379

# ... 其他配置保持不變
```

#### 步驟 3：配置 API 服務

創建 API 配置檔案：

```bash
cd D:\Projects\simple-admin-core\api\etc
cp core.yaml core-local.yaml
```

編輯 `api/etc/core-local.yaml`：

```yaml
Name: core.api
Host: 0.0.0.0
Port: 9100

Auth:
  AccessSecret: jS6VKDtsJf3z1n2VKDtsJf3z1n2
  AccessExpire: 259200

DatabaseConf:
  Type: postgres                    # 改為 postgres
  Host: 127.0.0.1
  Port: 5432
  DBName: simple_admin
  Username: postgres
  Password: simple-admin.
  MaxOpenConn: 100
  SSLMode: disable
  CacheTime: 5

RedisConf:
  Host: 127.0.0.1:6379

CoreRpc:
  Target: 127.0.0.1:9101           # 改為本地 RPC
  Enabled: true

JobRpc:
  Enabled: false                    # 停用 Job RPC

McmsRpc:
  Enabled: false                    # 停用 MCMS RPC

I18nConf:
  Dir: ./api/internal/i18n/locale   # zh-TW 語言檔案路徑

# ... 其他配置保持不變
```

#### 步驟 4：啟動 Backend 服務

**終端 1 - RPC 服務**：
```bash
cd D:\Projects\simple-admin-core
go run rpc/core.go -f rpc/etc/core-postgres.yaml
```

**終端 2 - API 服務**：
```bash
cd D:\Projects\simple-admin-core
go run api/core.go -f api/etc/core-local.yaml
```

驗證服務啟動：
```bash
# 檢查 RPC
curl http://localhost:9101

# 檢查 API
curl http://localhost:9100/health
```

#### 步驟 5：配置前端

編輯 `simple-admin-vben5-ui/apps/simple-admin-core/.env.development`：

```bash
VITE_GLOB_API_URL=http://localhost:9100
```

#### 步驟 6：啟動前端

```bash
cd D:\Projects\simple-admin-vben5-ui
pnpm install
pnpm run dev:core
```

前端啟動後訪問：http://localhost:5555

---

## ✅ 測試 zh-TW 功能

### 1. 語言選擇器測試

1. 開啟瀏覽器訪問 http://localhost:5555
2. 點擊右上角語言選擇器（地球圖示）
3. 選擇「繁體中文（台灣）」
4. 驗證頁面文字變更為台灣用語

**預期結果**：
- ✅ 選單、按鈕、表單標籤顯示繁體中文
- ✅ 使用台灣術語：使用者、資料庫、選單、檔案

### 2. API 錯誤訊息測試

測試後端 i18n：

```bash
# 測試 zh-TW locale
curl -H "Accept-Language: zh-TW" http://localhost:9100/api/v1/core/user/list

# 測試 zh-Hant (應正規化為 zh-TW)
curl -H "Accept-Language: zh-Hant" http://localhost:9100/api/v1/core/user/list

# 測試 zh-Hant-TW (應正規化為 zh-TW)
curl -H "Accept-Language: zh-Hant-TW" http://localhost:9100/api/v1/core/user/list
```

**預期結果**：
- ✅ 錯誤訊息顯示繁體中文（台灣）
- ✅ zh-Hant, zh-Hant-TW 自動轉換為 zh-TW

### 3. 使用者偏好設定測試

1. 登入系統
2. 切換語言為「繁體中文（台灣）」
3. 登出並重新登入
4. 驗證系統記住語言偏好

---

## 🐛 常見問題排除

### 問題 1：RPC 連線失敗

**錯誤**：`rpc error: code = Unavailable`

**解決**：
- 確認 RPC 服務運行中
- 檢查 `api/etc/core-local.yaml` 的 `CoreRpc.Target: 127.0.0.1:9101`

### 問題 2：資料庫連線失敗

**錯誤**：`failed to connect to database`

**解決**：
```bash
# 檢查 PostgreSQL 運行狀態
docker ps | grep postgres

# 檢查資料庫是否創建
docker exec -it <postgres-container> psql -U postgres -l
```

### 問題 3：語言檔案找不到

**錯誤**：`failed to load locale file`

**解決**：
- 確認 `api/etc/core-local.yaml` 設定 `I18nConf.Dir: ./api/internal/i18n/locale`
- 確認 `api/internal/i18n/locale/zh-TW.json` 存在

### 問題 4：前端語言選項沒有 zh-TW

**解決**：
```bash
# 確認前端有最新的 zh-TW 分支程式碼
cd D:\Projects\simple-admin-vben5-ui
git log --oneline | grep zh-TW

# 重新安裝依賴並清除快取
pnpm install
pnpm run clean
pnpm run dev:core
```

---

## 📊 測試檢查清單

### Backend 測試
- [ ] RPC 服務正常啟動 (Port 9101)
- [ ] API 服務正常啟動 (Port 9100)
- [ ] `curl` 測試 Accept-Language: zh-TW 返回繁體中文錯誤訊息
- [ ] Locale 正規化功能正常（zh-Hant → zh-TW）
- [ ] 資料庫 users 表的 locale 欄位存在

### Frontend 測試
- [ ] 前端啟動成功 (Port 5555)
- [ ] 語言選擇器顯示「繁體中文（台灣）」選項
- [ ] 切換語言後頁面文字變更
- [ ] 登入頁面使用台灣用語
- [ ] 主選單和側邊欄使用台灣用語
- [ ] 資料表格和表單使用台灣用語

### 整合測試
- [ ] 前端 → API → RPC 完整流程正常
- [ ] 語言偏好設定可儲存至資料庫
- [ ] 登出後重新登入保持語言偏好

---

## 🎯 下一步

測試完成後：

1. **執行單元測試**：
   ```bash
   cd D:\Projects\simple-admin-core
   go test ./api/internal/i18n/...
   ```

2. **執行 E2E 測試**：參考 `docs/zh-TW-E2E-TEST-PLAN.md`

3. **人工 QA**：參考 `docs/zh-TW-MANUAL-QA-CHECKLIST.md`

4. **構建生產映像**：
   ```bash
   # Backend
   make docker DOCKER_USERNAME=your-username

   # Frontend
   cd D:\Projects\simple-admin-vben5-ui
   docker build -t your-username/backend-ui-vben5:zh-tw .
   ```

---

## 📝 備註

- 本地開發不需要修改 Dockerfile，直接運行 Go 和 pnpm 即可
- 如需生產部署，需要構建新的 Docker 映像包含 zh-TW 功能
- 現有的 `deploy/docker-compose/all_in_one/` 配置使用 v1.7.0 映像，不包含 zh-TW 功能
- 建議使用本地開發方式測試，確認功能正常後再構建生產映像
