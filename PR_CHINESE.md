<!--
感谢您发送拉取请求！以下是一些提示：

1. 如果您想**更快**获得 PR 审查，请阅读：https://github.com/kubesphere/community/blob/master/developer-guide/development/the-pr-author-guide-to-getting-through-code-review.md
2. 如果您想了解您的 PR 是如何被审查的，请阅读：https://github.com/kubesphere/community/blob/master/developer-guide/development/code-review-guide.md
3. KubeSphere 社区遵循的编码约定：https://github.com/kubesphere/community/blob/master/developer-guide/development/coding-conventions.md
-->

### 这是一个什么样的 PR？
/kind feature
/kind api-change

### 这个 PR 做什么 / 为什么需要它：

为 simple-admin-core 项目添加完整的库存管理系统，包括仓库管理、产品管理、库存跟踪和库存变动记录功能。

#### 核心功能：
- **仓库管理**：完整的 CRUD 操作，支持仓库位置和容量管理
- **产品管理**：产品目录管理，包括 SKU、价格、分类等信息
- **库存管理**：按仓库跟踪库存水平，支持最低/最高库存预警
- **库存变动**：完整的库存变动审计记录，支持入库和出库操作

#### 技术实现：
- **数据库层**：使用 Ent ORM，自动生成实体和查询代码
- **API 层**：RESTful API 接口，包含 16 个端点
- **gRPC 服务**：Protobuf 定义的服务接口
- **业务逻辑**：清晰的分层架构，逻辑与处理分离

### 这个 PR 修复了哪些问题：
修复 #

### 给审查者的特别说明：

#### 数据库架构：
```
仓库(Warehouse) 1:N 库存(Inventory) N:1 产品(Product)
库存变动(StockMovement) N:1 产品(Product)
库存变动(StockMovement) N:1 仓库(Warehouse)
```

#### API 端点：
- **仓库**: `/warehouse/*` (4 个端点)
- **产品**: `/product/*` (4 个端点)
- **库存**: `/inventory/*` (4 个端点)
- **库存变动**: `/stock_movement/*` (4 个端点)

### 这个 PR 是否引入了面向用户的变更？
```release-note
添加完整的库存管理系统，包含仓库、产品、库存和库存变动管理功能，支持 REST API 和 gRPC 服务调用。
```

### 其他文档、使用文档等：

### 手动验证步骤：

#### 1. 环境准备
```bash
# 确保 MySQL 和 Redis 正在运行
brew services start mysql
brew services start redis

# 创建数据库
mysql -u root -e "CREATE DATABASE simple_admin;"
```

#### 2. 启动服务
```bash
# 启动 RPC 服务
cd rpc && go run . &

# 启动 API 服务
cd api && go run . &
```

#### 3. 测试仓库管理
```bash
# 创建仓库
curl -X POST http://localhost:9100/warehouse/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "主仓库",
    "location": "123 工业街",
    "capacity": 10000,
    "description": "主要存储仓库"
  }'

# 查询仓库列表
curl -X POST http://localhost:9100/warehouse/list \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'
```

#### 4. 测试产品管理
```bash
# 创建产品
curl -X POST http://localhost:9100/product/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "笔记本电脑",
    "sku": "LT-001",
    "price": 1299.99,
    "category": "电子产品",
    "brand": "TechCorp"
  }'

# 查询产品列表
curl -X POST http://localhost:9100/product/list \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'
```

#### 5. 测试库存管理
```bash
# 创建库存记录
curl -X POST http://localhost:9100/inventory/create \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "warehouse_id": 1,
    "quantity": 100,
    "min_stock_level": 10,
    "max_stock_level": 500
  }'

# 查询库存列表
curl -X POST http://localhost:9100/inventory/list \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'
```

#### 6. 测试库存变动
```bash
# 记录入库操作
curl -X POST http://localhost:9100/stock_movement/create \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "warehouse_id": 1,
    "movement_type": "IN",
    "quantity": 25,
    "reference": "PO-2024-001",
    "notes": "收到新库存"
  }'

# 查询库存变动记录
curl -X POST http://localhost:9100/stock_movement/list \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'
```

### 验证结果：
- ✅ 所有 API 端点响应正常
- ✅ 数据库记录正确创建
- ✅ 数据关联关系正确维护
- ✅ CRUD 操作完全正常
- ✅ 业务逻辑正确执行

### 文件变更摘要：
- **新增文件**: 126 个文件
- **修改文件**: 20 个文件
- **新增代码行**: 28,891 行
- **功能模块**: 4 个主要实体 + 16 个 API 端点
- **技术栈**: Ent ORM + gRPC + REST API
