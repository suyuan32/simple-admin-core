# Simple Admin Core RPC 部署指南 (PostgreSQL)

## 前置需求

### 1. 安裝並啟動 Docker Desktop

由於專案需要 PostgreSQL 和 Redis，建議使用 Docker 運行：

1. 啟動 Docker Desktop
2. 確認 Docker 正在運行：
   ```bash
   docker ps
   ```

### 2. 啟動資料庫服務

使用專案提供的 docker-compose 配置：

```bash
# 啟動 PostgreSQL 和 Redis
docker-compose -f docker-compose-dev.yaml up -d

# 檢查服務狀態
docker-compose -f docker-compose-dev.yaml ps

# 查看日誌
docker-compose -f docker-compose-dev.yaml logs -f
```

服務資訊：
- **PostgreSQL**
  - Port: `5432`
  - Database: `simple_admin`
  - Username: `postgres`
  - Password: `simple-admin.`

- **Redis**
  - Port: `6379`
  - 無密碼

### 3. 初始化資料庫（可選）

服務會自動創建資料庫表結構（Ent 自動遷移），但如果需要手動執行：

```bash
# 連接到 PostgreSQL
docker exec -it simple-admin-postgresql psql -U postgres -d simple_admin

# 查看表
\dt

# 退出
\q
```

## 運行 RPC 服務

### 方式一：直接運行（開發模式）

```bash
# 使用 PostgreSQL 配置運行
go run rpc/core.go -f rpc/etc/core-postgres.yaml
```

### 方式二：編譯後運行

```bash
# Windows
make build-win
./core_rpc.exe -f rpc/etc/core-postgres.yaml

# Linux
make build-linux
./core_rpc -f rpc/etc/core-postgres.yaml

# macOS
make build-mac
./core_rpc -f rpc/etc/core-postgres.yaml
```

## 驗證服務

### 1. 檢查服務監聽

```bash
# Windows
netstat -an | findstr "9101"

# Linux/macOS
netstat -an | grep 9101
```

應該看到：
```
TCP    0.0.0.0:9101    LISTENING
```

### 2. 檢查 Prometheus 指標

訪問：http://localhost:4001/metrics

應該看到服務的監控指標

### 3. 檢查日誌

日誌位置：`/home/data/logs/core/rpc/` 或查看控制台輸出

## 配置說明

配置文件：`rpc/etc/core-postgres.yaml`

關鍵配置項：
```yaml
DatabaseConf:
  Type: postgres          # 資料庫類型
  Host: 127.0.0.1        # 資料庫地址
  Port: 5432             # 資料庫端口
  DBName: simple_admin   # 資料庫名稱
  Username: postgres     # 用戶名
  Password: simple-admin. # 密碼

RedisConf:
  Host: 127.0.0.1:6379   # Redis 地址
```

## 停止服務

### 停止 RPC 服務
按 `Ctrl+C` 停止運行的服務

### 停止資料庫服務
```bash
# 停止容器
docker-compose -f docker-compose-dev.yaml down

# 停止並刪除資料
docker-compose -f docker-compose-dev.yaml down -v
```

## 常見問題

### 1. Docker 連接錯誤
```
Error: cannot connect to Docker daemon
```
**解決方案**：啟動 Docker Desktop

### 2. 端口被占用
```
Error: port 5432 already in use
```
**解決方案**：
- 停止本地的 PostgreSQL 服務
- 或修改 docker-compose-dev.yaml 中的端口映射

### 3. 資料庫連接失敗
```
Error: connection refused
```
**解決方案**：
- 檢查 Docker 容器是否運行：`docker ps`
- 檢查資料庫連接資訊是否正確
- 等待容器完全啟動（約 10-20 秒）

### 4. Ent Schema 遷移失敗
```
Error: migration failed
```
**解決方案**：
- 手動刪除資料庫並重新創建
- 檢查 PostgreSQL 用戶權限

## 下一步

完成 RPC 服務部署後，可以：

1. 運行 API 服務（需要修改 `api/etc/core.yaml` 配置）
2. 測試 gRPC 接口
3. 整合到現有的 Hospital ERP 專案

## 技術架構驗證

✅ **已驗證功能**：
- PostgreSQL 資料庫支援
- Ent ORM 自動遷移
- Casbin 權限框架整合
- Redis 快取支援
- gRPC 服務架構
- 字典管理模組
- 使用者/角色/部門/選單管理

✅ **與報告對比**：
- ✅ 確認支援 PostgreSQL（報告未提及）
- ✅ 確認使用 Ent ORM（報告正確）
- ✅ 確認無使用 GORM（報告正確）
- ✅ 確認 Casbin 完整整合（報告正確）
