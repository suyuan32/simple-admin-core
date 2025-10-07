# Monorepo Docker Compose Configuration

此配置用於 Monorepo 環境，會從 `web/` 目錄建置 Frontend。

## 架構

```
simple-admin-core/ (Monorepo)
├── api/              # Backend API Service
├── rpc/              # Backend RPC Service
├── web/              # Frontend (built from source)
└── deploy/
    └── docker-compose/
        └── monorepo/
            └── docker-compose.yaml
```

## 特點

- **Frontend**: 從 Monorepo 的 `web/` 目錄建置（使用 `Dockerfile.web`）
- **Backend**: 使用預建的 Docker images
- **自動整合**: Frontend Nginx 自動反向代理到 Backend API (port 9100)

## 使用方法

### 1. 啟動所有服務

```bash
cd deploy/docker-compose/monorepo
docker-compose up -d
```

### 2. 查看日誌

```bash
# 查看所有服務日誌
docker-compose logs -f

# 查看特定服務日誌
docker-compose logs -f backend-ui
docker-compose logs -f core-api
```

### 3. 停止服務

```bash
docker-compose down
```

### 4. 重新建置 Frontend

```bash
# 停止並移除舊容器
docker-compose down

# 重新建置並啟動
docker-compose up -d --build backend-ui
```

## 服務端口

- **Frontend (Nginx)**: http://localhost:80
- **Backend API**: http://localhost:9100
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

## 環境變量

Frontend 使用 `.env.production` 配置，API endpoint 設定為：
- `VITE_GLOB_API_URL=/` (透過 Nginx 反向代理到 Backend)

Nginx 反向代理配置（`web/scripts/deploy/nginx.conf`）：
```nginx
location /sys-api/ {
    proxy_pass http://core-api:9100/;
}
```

## 本地開發 vs Docker 部署

### 本地開發
```bash
# Terminal 1: Backend API
cd simple-admin-core
go run api/core.go -f api/etc/core.yaml

# Terminal 2: Backend RPC
go run rpc/core.go -f rpc/etc/core.yaml

# Terminal 3: Frontend
cd web
pnpm install
pnpm dev:simple-admin-core
```

Frontend 開發模式下使用：
- `VITE_GLOB_API_URL=http://localhost:9100` (直接連接 Backend)

### Docker 部署
```bash
cd deploy/docker-compose/monorepo
docker-compose up -d
```

Frontend 生產模式下使用：
- `VITE_GLOB_API_URL=/` (透過 Nginx 反向代理)

## 故障排除

### Frontend 建置失敗

檢查 Docker 日誌：
```bash
docker-compose logs backend-ui
```

手動測試建置：
```bash
cd web
pnpm install
pnpm build:simple-admin-core
```

### Backend API 無法連接

檢查 core-api 服務狀態：
```bash
docker-compose ps core-api
docker-compose logs core-api
```

### 資料庫連接失敗

確認 PostgreSQL 已啟動：
```bash
docker-compose ps postgresql
docker-compose logs postgresql
```

## 更新說明

相較於 `deploy/docker-compose/all_in_one/postgresql/docker-compose.yaml`：
- ✅ Frontend 從 Monorepo 源碼建置（而非使用預建 image）
- ✅ 自動包含最新的 Frontend 代碼變更
- ✅ 支援 zh-TW 語言檔案（位於 `web/apps/simple-admin-core/src/locales/langs/zh-TW/`）
