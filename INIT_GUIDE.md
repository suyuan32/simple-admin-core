# Simple Admin Core 初始化指南

## ✅ 初始化已完成！

系統已成功初始化，所有資料庫表格和預設資料已建立。

**初始化結果**：
- ✅ 核心資料庫 (Core): 15 tables
- ✅ 工作排程 (Job): +3 tables
- ✅ 訊息中心 (MCMS): +3 tables
- 📊 **總計: 21 tables**

---

## 🔐 登入系統

### 預設管理員帳號
```
使用者名：admin
密碼：simple-admin
```

### 訪問地址
- 🖥️ **前端管理後台**: http://localhost:8888
- 🔌 **API 服務**: http://localhost:9100
- 🔧 **RPC 服務**: localhost:9101

### 登入步驟
1. 訪問 http://localhost:8888
2. 輸入帳號：`admin`
3. 輸入密碼：`simple-admin`
4. 完成圖形驗證碼
5. 點擊登入

## 🎉 系統已完全就緒

已完成初始化，包含：
1. ✅ 創建所有資料庫表（21個表）
2. ✅ 創建預設管理員帳號
3. ✅ 創建預設角色和權限
4. ✅ 創建預設選單
5. ✅ 創建預設部門和職位
6. ✅ 創建基礎字典數據

---

## 🔍 初始化驗證結果

### 資料庫表統計

執行命令：
```bash
docker exec -e PGPASSWORD=simple-admin. postgresql psql -U postgres -d simple_admin -c "\dt"
```

**已建立的 21 個表**：

**核心模組 (Core - 15 tables)**:
- `sys_users` - 使用者表
- `sys_roles` - 角色表
- `sys_menus` - 選單表
- `sys_apis` - API 權限表
- `sys_departments` - 部門表
- `sys_positions` - 職位表
- `sys_dictionaries` - 字典表
- `sys_dictionary_details` - 字典明細表
- `sys_configuration` - 系統設定表
- `sys_oauth_providers` - OAuth 提供者表
- `sys_tokens` - Token 管理表
- `casbin_rules` - Casbin 權限規則表
- `user_roles` - 使用者角色關聯表
- `user_positions` - 使用者職位關聯表
- `role_menus` - 角色選單關聯表

**工作排程模組 (Job - 3 tables)**
**訊息中心模組 (MCMS - 3 tables)**

### 初始化 API 執行記錄

```bash
# 核心資料庫初始化
$ curl http://localhost:9100/core/init/database
{"code":0,"msg":"成功"}

# 工作排程資料庫初始化
$ curl http://localhost:9100/core/init/job_database
{"code":0,"msg":"成功"}

# 訊息中心資料庫初始化
$ curl http://localhost:9100/core/init/mcms_database
{"code":0,"msg":"成功"}
```

### 問題排查歷程

**之前遇到的錯誤**：`{"code":13,"msg":"数据库错误"}`

**根本原因**：
- 之前手動建立的 `sys_users` 表格與 Ent ORM 自動產生的 schema 不相容
- `created_at` 欄位類型衝突：手動建立的是 `timestamp`，Ent 需要 `timestamp with time zone`

**解決方案**：
```sql
-- 刪除衝突的表格
DROP TABLE IF EXISTS sys_users CASCADE;
DROP TABLE IF EXISTS casbin_rules CASCADE;

-- 讓系統自動建立正確的表格結構
curl http://localhost:9100/core/init/database
```

---

## 🛠️ 服務狀態

所有服務運行正常：

```bash
$ docker ps
CONTAINER ID   IMAGE                                    STATUS              PORTS
simple-admin-ui   ryanpower/backend-ui-vben5:v1.7.0     Up 30+ minutes      0.0.0.0:8888->80/tcp
core-api          ryanpower/core-api:v1.7.4             Up 24+ minutes      0.0.0.0:9100->9100/tcp
core-rpc          ryanpower/core-rpc:v1.7.4             Up 25+ minutes      9101/tcp
job-rpc           ryanpower/job-rpc:v1.1.7              Up 25+ minutes      9105/tcp
mcms-rpc          ryanpower/mcms-rpc:v1.1.5             Up 25+ minutes      9106/tcp
fms-api           ryanpower/fms-api:v1.1.8              Up 25+ minutes      9102/tcp
postgresql        bitnami/postgresql:18.0               Up 25+ minutes      0.0.0.0:5432->5432/tcp
redis-server      redis:7.2-alpine                      Up 25+ minutes      0.0.0.0:6379->6379/tcp
```

---

## 📊 系統功能測試指南

### 可測試的核心功能

登入後台管理系統 (http://localhost:8888)，您可以測試：

1. **RBAC 權限控制**
   - 角色管理：新增/修改/刪除角色
   - 權限分配：為角色分配選單和 API 權限
   - Casbin 規則驗證

2. **使用者管理**
   - 新增使用者
   - 分配角色
   - 設定部門和職位
   - 使用者狀態管理

3. **選單配置**
   - 動態選單管理
   - 多層級選單結構
   - 選單權限綁定

4. **API 權限管理**
   - API 註冊與管理
   - API 與角色綁定
   - 路徑匹配規則

5. **組織架構**
   - 部門管理（樹狀結構）
   - 職位管理
   - 人員分配

6. **系統設定**
   - 字典管理
   - 系統配置
   - OAuth 提供者設定

---

## 🔧 資料庫連線資訊

- **類型**: PostgreSQL 18.0
- **主機**: localhost:5432
- **資料庫**: simple_admin
- **使用者**: postgres
- **密碼**: simple-admin.
- **狀態**: ✅ 運行中，21 個表已建立

## 🎯 完成的初始化步驟

```
✅ 1. 啟動 Docker Compose (PostgreSQL + Redis + 所有服務)
✅ 2. 等待服務就緒
✅ 3. 刪除衝突的手動建立表格
✅ 4. 執行核心資料庫初始化 API
✅ 5. 執行工作排程資料庫初始化 API
✅ 6. 執行訊息中心資料庫初始化 API
✅ 7. 驗證 21 個表格已正確建立
━━━━━━━━━━━━━━━━━━━━━━━━━
🎉 系統已完全初始化並就緒！
```

---

## ✅ 下一步：開始評估系統

### 立即登入測試

訪問：http://localhost:8888
- 帳號：`admin`
- 密碼：`simple-admin`

### 相關文件

- [系統訪問指南](./SYSTEM_ACCESS_GUIDE.md) - 詳細的功能介紹
- [Ent→gRPC 技術詳解](./ENT_GRPC_EXPLAINED.md) - 技術架構說明
- [部署指南](./DEPLOYMENT.md) - Docker Compose 部署說明
- [驗證報告](./VERIFICATION_REPORT.md) - 專案驗證結果

---

**🎉 Simple Admin Core 已完全部署並初始化！現在可以開始評估系統穩定性了！**
