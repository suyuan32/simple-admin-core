# Traditional Chinese (zh-TW) Feature - 部署指南

## 📋 部署架構

Simple Admin 採用 **前後端分離架構**：

```
┌─────────────────────────────────────────────┐
│              Nginx / Load Balancer          │
│              (Port 80/443)                  │
└───────┬─────────────────────────────┬───────┘
        │                             │
        ▼                             ▼
┌───────────────┐           ┌──────────────────┐
│   Frontend    │           │     Backend      │
│   (Vue 3)     │    API    │   (Go-Zero)      │
│   Port 5173   │◄─────────►│   Port 9100/9101 │
│               │           │                  │
│ simple-admin- │           │ simple-admin-    │
│ vben5-ui      │           │ core             │
└───────────────┘           └──────────────────┘
                                    │
                                    ▼
                            ┌──────────────────┐
                            │   Database       │
                            │ MySQL/PostgreSQL │
                            │ + Redis          │
                            └──────────────────┘
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

## 🚀 部署方式選擇

### 方式 1: Docker Compose（推薦用於開發測試）

適合快速測試 zh-TW 功能。

### 方式 2: 獨立部署（推薦用於生產環境）

前後端分別部署，更靈活。

### 方式 3: Kubernetes（推薦用於大規模生產）

適合需要高可用性的場景。

---

## 📦 方式 1: Docker Compose 部署（最簡單）

### 步驟 1: 準備 Docker Compose 檔案

在 `D:\Projects\` 建立 `docker-compose-dev.yaml`：

```yaml
version: '3.8'

services:
  # PostgreSQL 資料庫
  postgres:
    image: postgres:16-alpine
    container_name: simple-admin-postgres
    environment:
      POSTGRES_DB: simple_admin
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: simple-admin.
      TZ: Asia/Taipei
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - simple-admin-net

  # Redis 快取
  redis:
    image: redis:7-alpine
    container_name: simple-admin-redis
    command: redis-server --requirepass simple-admin.
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    networks:
      - simple-admin-net

  # Backend RPC Service
  core-rpc:
    build:
      context: ./simple-admin-core
      dockerfile: Dockerfile-rpc
    container_name: simple-admin-core-rpc
    environment:
      TZ: Asia/Taipei
    ports:
      - "9101:9101"
    depends_on:
      - postgres
      - redis
    networks:
      - simple-admin-net
    restart: unless-stopped

  # Backend API Service
  core-api:
    build:
      context: ./simple-admin-core
      dockerfile: Dockerfile-api
    container_name: simple-admin-core-api
    environment:
      TZ: Asia/Taipei
    ports:
      - "9100:9100"
    depends_on:
      - core-rpc
      - postgres
      - redis
    networks:
      - simple-admin-net
    restart: unless-stopped

  # Frontend
  frontend:
    build:
      context: ./simple-admin-vben5-ui
      dockerfile: Dockerfile
    container_name: simple-admin-frontend
    ports:
      - "80:80"
    depends_on:
      - core-api
    networks:
      - simple-admin-net
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:

networks:
  simple-admin-net:
    driver: bridge
```

### 步驟 2: 修改 Backend 配置

編輯 `simple-admin-core/api/etc/core.yaml`：

```yaml
Name: core.api
Host: 0.0.0.0
Port: 9100
Timeout: 30000

Auth:
  AccessSecret: "your-secret-key"
  AccessExpire: 259200

# Database Configuration
DatabaseConf:
  Type: postgres
  Host: postgres
  Port: 5432
  DBName: simple_admin
  Username: postgres
  Password: simple-admin.
  MaxOpenConns: 100
  SSLMode: disable
  CacheTime: 5

# Redis Configuration
RedisConf:
  Host: redis:6379
  Type: node
  Pass: simple-admin.

# RPC Configuration
CoreRpc:
  Target: core-rpc:9101

# i18n Configuration (確保有這一段)
I18nConf:
  Dir: etc/locale

# Casbin Configuration
CasbinConf:
  ModelText: |
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [role_definition]
    g = _, _

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
```

編輯 `simple-admin-core/rpc/etc/core.yaml`：

```yaml
Name: core.rpc
ListenOn: 0.0.0.0:9101

# Database Configuration
DatabaseConf:
  Type: postgres
  Host: postgres
  Port: 5432
  DBName: simple_admin
  Username: postgres
  Password: simple-admin.
  MaxOpenConns: 100
  SSLMode: disable

# Redis Configuration
RedisConf:
  Host: redis:6379
  Type: node
  Pass: simple-admin.

# Casbin Configuration
CasbinConf:
  ModelText: |
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [role_definition]
    g = _, _

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
```

### 步驟 3: 修改 Frontend 配置

編輯 `simple-admin-vben5-ui/apps/simple-admin-core/.env.production`：

```env
# API 基礎 URL
VITE_GLOB_API_URL=http://localhost:9100

# 應用標題
VITE_GLOB_APP_TITLE=Simple Admin

# 預設語言
VITE_GLOB_LOCALE=zh-CN
```

或建立 `.env.development`:

```env
VITE_GLOB_API_URL=http://localhost:9100
VITE_GLOB_APP_TITLE=Simple Admin (Dev)
VITE_GLOB_LOCALE=zh-TW
```

### 步驟 4: 構建並啟動

```bash
# 在 D:\Projects\ 目錄下執行

# 1. 構建 Backend (必須先構建二進制檔案)
cd simple-admin-core
make build-win  # Windows
# 或
make build-linux  # Linux (Docker 內使用)

# 2. 啟動所有服務
cd ..
docker-compose -f docker-compose-dev.yaml up -d

# 3. 查看日誌
docker-compose -f docker-compose-dev.yaml logs -f

# 4. 初始化資料庫（第一次運行）
# Backend 會自動執行 Ent 遷移
```

### 步驟 5: 訪問系統

- **Frontend**: http://localhost
- **Backend API**: http://localhost:9100
- **Backend RPC**: localhost:9101

### 步驟 6: 測試 zh-TW 功能

1. 開啟瀏覽器訪問 http://localhost
2. 點擊右上角語言選擇器
3. 選擇「繁體中文（台灣）」
4. 驗證 UI 顯示繁體中文
5. 執行 CRUD 操作，檢查訊息顯示

---

## 📦 方式 2: 獨立部署

### A. Backend 部署

#### 1. 構建二進制檔案

```bash
cd D:\Projects\simple-admin-core

# Windows
make build-win

# Linux
make build-linux

# macOS
make build-mac
```

輸出檔案：
- `core_api.exe` / `core_api` (API 服務)
- `core_rpc.exe` / `core_rpc` (RPC 服務)

#### 2. 準備部署目錄

```
deploy/
├── core_api(.exe)
├── core_rpc(.exe)
├── api/
│   └── etc/
│       ├── core.yaml
│       └── locale/
│           ├── zh.json
│           ├── en.json
│           └── zh-TW.json  ✅
└── rpc/
    └── etc/
        └── core.yaml
```

#### 3. 啟動服務

```bash
# 啟動 RPC 服務（先啟動）
./core_rpc -f rpc/etc/core.yaml

# 啟動 API 服務
./core_api -f api/etc/core.yaml
```

#### 4. 使用 systemd（Linux 推薦）

建立 `/etc/systemd/system/simple-admin-rpc.service`：

```ini
[Unit]
Description=Simple Admin Core RPC Service
After=network.target postgresql.service redis.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/simple-admin
ExecStart=/opt/simple-admin/core_rpc -f /opt/simple-admin/rpc/etc/core.yaml
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

建立 `/etc/systemd/system/simple-admin-api.service`：

```ini
[Unit]
Description=Simple Admin Core API Service
After=network.target simple-admin-rpc.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/simple-admin
ExecStart=/opt/simple-admin/core_api -f /opt/simple-admin/api/etc/core.yaml
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

啟動服務：

```bash
sudo systemctl daemon-reload
sudo systemctl enable simple-admin-rpc
sudo systemctl enable simple-admin-api
sudo systemctl start simple-admin-rpc
sudo systemctl start simple-admin-api

# 查看狀態
sudo systemctl status simple-admin-rpc
sudo systemctl status simple-admin-api
```

### B. Frontend 部署

#### 1. 構建前端

```bash
cd D:\Projects\simple-admin-vben5-ui

# 安裝依賴
pnpm install

# 構建生產版本
pnpm run build

# 構建產物位於：apps/simple-admin-core/dist/
```

#### 2. 使用 Nginx 部署

安裝 Nginx：

```bash
# Ubuntu/Debian
sudo apt install nginx

# Windows
# 下載 nginx.exe from http://nginx.org/en/download.html
```

Nginx 配置 `/etc/nginx/sites-available/simple-admin`:

```nginx
server {
    listen 80;
    server_name localhost;

    # Frontend 靜態檔案
    root /var/www/simple-admin/dist;
    index index.html;

    # Gzip 壓縮
    gzip on;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;

    # Frontend 路由
    location / {
        try_files $uri $uri/ /index.html;
    }

    # Backend API 代理
    location /api/ {
        proxy_pass http://localhost:9100/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # CORS Headers (如果需要)
        add_header Access-Control-Allow-Origin *;
        add_header Access-Control-Allow-Methods 'GET, POST, PUT, DELETE, OPTIONS';
        add_header Access-Control-Allow-Headers 'DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization';

        if ($request_method = 'OPTIONS') {
            return 204;
        }
    }

    # 靜態資源快取
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

部署前端檔案：

```bash
# 複製構建產物
sudo mkdir -p /var/www/simple-admin
sudo cp -r D:\Projects\simple-admin-vben5-ui\apps\simple-admin-core\dist\* /var/www/simple-admin/

# 設定權限
sudo chown -R www-data:www-data /var/www/simple-admin
sudo chmod -R 755 /var/www/simple-admin

# 啟用網站
sudo ln -s /etc/nginx/sites-available/simple-admin /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

---

## 🐳 方式 3: Docker 獨立容器部署

### Backend Docker

```bash
cd D:\Projects\simple-admin-core

# 構建 RPC 鏡像
docker build -t simple-admin-core-rpc:latest -f Dockerfile-rpc .

# 構建 API 鏡像
docker build -t simple-admin-core-api:latest -f Dockerfile-api .

# 運行 RPC
docker run -d \
  --name simple-admin-rpc \
  -p 9101:9101 \
  --network simple-admin-net \
  simple-admin-core-rpc:latest

# 運行 API
docker run -d \
  --name simple-admin-api \
  -p 9100:9100 \
  --network simple-admin-net \
  simple-admin-core-api:latest
```

### Frontend Docker

```bash
cd D:\Projects\simple-admin-vben5-ui

# 構建前端鏡像
docker build -t simple-admin-frontend:latest .

# 運行前端
docker run -d \
  --name simple-admin-frontend \
  -p 80:80 \
  --network simple-admin-net \
  simple-admin-frontend:latest
```

---

## ✅ 部署驗證

### 1. Backend 健康檢查

```bash
# 檢查 API 服務
curl http://localhost:9100/health

# 檢查 zh-TW 語言檔案
curl -H "Accept-Language: zh-TW" http://localhost:9100/api/v1/init/database
```

### 2. Frontend 驗證

開啟瀏覽器：http://localhost

檢查項目：
- ✅ 頁面正常載入
- ✅ 語言選擇器顯示 3 個選項
- ✅ 可切換到「繁體中文（台灣）」
- ✅ UI 文字顯示繁體中文
- ✅ Console 無錯誤

### 3. 整合測試

執行之前建立的測試計劃：
- `docs/zh-TW-E2E-TEST-PLAN.md`
- `docs/zh-TW-MANUAL-QA-CHECKLIST.md`

---

## 🔧 常見問題排查

### 問題 1: Frontend 連不上 Backend

**症狀**: 前端顯示網路錯誤

**解決**:
1. 檢查 Backend 是否運行：`curl http://localhost:9100/health`
2. 檢查前端 API URL 配置：`.env.production` 中的 `VITE_GLOB_API_URL`
3. 檢查 CORS 設定（如前後端不同域名）

### 問題 2: zh-TW 語言檔案不生效

**症狀**: 切換語言後仍顯示簡體中文

**Backend 檢查**:
```bash
# 確認 locale 目錄存在
ls simple-admin-core/api/internal/i18n/locale/
# 應該看到 zh-TW.json
```

**Frontend 檢查**:
```bash
# 確認語言檔案存在
ls simple-admin-vben5-ui/apps/simple-admin-core/src/locales/langs/zh-TW/
# 應該看到 6 個 .json 檔案
```

### 問題 3: Docker 容器啟動失敗

**檢查日誌**:
```bash
docker logs simple-admin-api
docker logs simple-admin-rpc
docker logs simple-admin-frontend
```

**常見原因**:
- 資料庫連線失敗（檢查 `core.yaml` 中的資料庫設定）
- Redis 連線失敗
- 端口被占用

---

## 📊 效能優化建議

### Backend
- 使用 PostgreSQL 連線池
- 啟用 Redis 快取
- 使用 CDN 分發靜態資源

### Frontend
- 開啟 Nginx gzip 壓縮
- 設定靜態資源快取
- 使用 HTTP/2

---

## 🔐 生產環境安全建議

1. **使用 HTTPS**（Let's Encrypt 免費證書）
2. **修改預設密碼**（資料庫、Redis）
3. **設定防火牆規則**
4. **定期備份資料庫**
5. **啟用日誌監控**

---

## 📝 部署檢查清單

部署前：
- [ ] Backend 程式碼已提交（commit: `00c0d59`+）
- [ ] Frontend 程式碼已提交（commit: `8bc94a930`+）
- [ ] 配置檔案已修改（資料庫、Redis 連線）
- [ ] 環境變數已設定

部署後：
- [ ] Backend 服務啟動成功
- [ ] Frontend 服務啟動成功
- [ ] 資料庫遷移完成
- [ ] 健康檢查通過
- [ ] zh-TW 功能測試通過

---

## 📞 需要協助？

- 查看 [E2E 測試計劃](./zh-TW-E2E-TEST-PLAN.md)
- 查看 [人工 QA 檢查清單](./zh-TW-MANUAL-QA-CHECKLIST.md)
- 提交 GitHub Issue
- 查閱 Simple Admin 官方文檔

---

**祝部署順利！🎉**
