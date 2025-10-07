# Ent→gRPC 技術詳解

## 🎯 什麼是 Ent→gRPC？

**Ent→gRPC** 是一種 **Schema-First (架構優先)** 的開發模式，從資料庫 Schema 定義出發，自動生成 gRPC 服務代碼的技術流程。

### 核心理念

```
資料庫架構 (Ent Schema)
    ↓ 自動生成
Protobuf 定義 (.proto)
    ↓ 自動生成
gRPC 服務代碼 (Go)
    ↓ 手動實現
業務邏輯 (Logic Layer)
```

---

## 📚 技術棧對比

### 傳統模式：GORM→HTTP (你的 Hospital ERP)

```go
// 1. 手動定義 GORM Model
type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Email    string
    Password string
}

// 2. 手動定義 HTTP DTO
type CreateUserRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email"`
    Password string `json:"password" binding:"required"`
}

// 3. 手動寫 HTTP Handler
func CreateUser(c *gin.Context) {
    var req CreateUserRequest
    c.ShouldBindJSON(&req)

    user := User{
        Username: req.Username,
        Email:    req.Email,
        Password: req.Password,
    }
    db.Create(&user)

    c.JSON(200, user)
}
```

**特點**：
- ✅ 簡單直觀，適合單體應用
- ✅ RESTful API，前端友好
- ❌ 手動維護 Model、DTO、Handler
- ❌ 類型安全較弱
- ❌ 微服務通信需要額外封裝

---

### 現代模式：Ent→gRPC (Simple Admin)

#### 步驟 1：定義 Ent Schema

```go
// rpc/ent/schema/user.go
package schema

type User struct {
    ent.Schema
}

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("username").Unique().
            Comment("User's login name | 登录名"),
        field.String("password").
            Comment("Password | 密码"),
        field.String("nickname").Unique().
            Comment("Nickname | 昵称"),
        field.String("email").Optional().
            Comment("Email | 邮箱号"),
        field.String("mobile").Optional().
            Comment("Mobile number | 手机号"),
        field.Uint64("department_id").Optional().
            Comment("Department ID | 部门ID"),
    }
}

func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixins.UUIDMixin{},      // 自動添加 UUID ID
        mixins.StatusMixin{},    // 自動添加 status 字段
        SoftDeleteMixin{},       // 自動添加軟刪除
    }
}

func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("departments", Department.Type).Unique(),
        edge.To("roles", Role.Type),         // 多對多關係
        edge.To("positions", Position.Type), // 多對多關係
    }
}
```

#### 步驟 2：生成 Ent ORM 代碼

```bash
make gen-ent
```

這會自動生成：
- `rpc/ent/user.go` - User 實體
- `rpc/ent/user_create.go` - 創建方法
- `rpc/ent/user_update.go` - 更新方法
- `rpc/ent/user_query.go` - 查詢方法
- `rpc/ent/user_delete.go` - 刪除方法

#### 步驟 3：自動生成 Protobuf + gRPC Logic

```bash
make gen-rpc-ent-logic model=User group=user
```

**這一步會自動生成：**

##### 3.1 Protobuf 定義 (`rpc/desc/user.proto`)

```protobuf
syntax = "proto3";

message UserInfo {
  optional string id = 1;
  optional int64 created_at = 2;
  optional int64 updated_at = 3;
  optional uint32 status = 4;
  optional string username = 5;
  optional string password = 6;
  optional string nickname = 7;
  optional string email = 12;
  optional string mobile = 11;
  optional uint64 department_id = 14;
  repeated uint64 role_ids = 10;
  repeated uint64 position_ids = 15;
}

message UserListReq {
  uint64 page = 1;
  uint64 page_size = 2;
  optional string username = 3;
  optional string email = 5;
}

message UserListResp {
  uint64 total = 1;
  repeated UserInfo data = 2;
}

service Core {
  rpc createUser (UserInfo) returns (BaseUUIDResp);
  rpc updateUser (UserInfo) returns (BaseResp);
  rpc getUserList (UserListReq) returns (UserListResp);
  rpc getUserById (UUIDReq) returns (UserInfo);
  rpc deleteUser (UUIDsReq) returns (BaseResp);
}
```

##### 3.2 gRPC Service 實現 (`rpc/internal/logic/user/create_user_logic.go`)

```go
package user

type CreateUserLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
    // ✅ 自動生成的驗證邏輯
    if in.Mobile != nil {
        checkMobile, err := l.svcCtx.DB.User.Query().
            Where(user.MobileEQ(*in.Mobile)).
            Exist(l.ctx)
        if err != nil {
            return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
        }
        if checkMobile {
            return nil, errorx.NewInvalidArgumentError("login.mobileExist")
        }
    }

    // ✅ 自動生成的 Ent ORM 操作
    result, err := l.svcCtx.DB.User.Create().
        SetNotNilUsername(in.Username).
        SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password))).
        SetNotNilNickname(in.Nickname).
        SetNotNilEmail(in.Email).
        SetNotNilMobile(in.Mobile).
        AddRoleIDs(in.RoleIds...).           // ✅ 自動處理關聯
        SetNotNilDepartmentID(in.DepartmentId).
        AddPositionIDs(in.PositionIds...).
        Save(l.ctx)

    if err != nil {
        return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
    }

    return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
```

**特點**：
- ✅ **完全自動生成** - 包含 CRUD 邏輯
- ✅ **類型安全** - Protobuf 強類型
- ✅ **關聯處理** - 自動處理多對多關係
- ✅ **錯誤處理** - 統一的錯誤處理
- ✅ **國際化** - 內置 i18n 支援
- ✅ **微服務友好** - gRPC 高性能

---

## 🔄 完整的代碼生成流程

### Simple Admin 的三段式生成

```bash
# 階段 1：定義並生成 Ent ORM
vim rpc/ent/schema/user.go    # 定義 Schema
make gen-ent                   # 生成 ORM 代碼

# 階段 2：生成 RPC 層 (Proto + Logic)
make gen-rpc-ent-logic model=User group=user

# 階段 3：生成 API 層 (HTTP Gateway)
vim api/desc/core/user.api     # 定義 API 路由
make gen-api                   # 生成 HTTP Handler
```

### 生成的文件結構

```
📦 Simple Admin Core
├── rpc/
│   ├── ent/schema/user.go           # [手動] Schema 定義
│   ├── ent/user.go                  # [自動] ORM 實體
│   ├── ent/user_create.go           # [自動] 創建方法
│   ├── ent/user_query.go            # [自動] 查詢方法
│   ├── desc/user.proto              # [自動] Protobuf 定義
│   ├── types/core/core.pb.go        # [自動] Proto 生成代碼
│   └── internal/logic/user/
│       ├── create_user_logic.go     # [自動] 創建邏輯
│       ├── update_user_logic.go     # [自動] 更新邏輯
│       ├── get_user_list_logic.go   # [自動] 查詢邏輯
│       └── delete_user_logic.go     # [自動] 刪除邏輯
└── api/
    ├── desc/core/user.api           # [手動] API 定義
    ├── internal/handler/user/       # [自動] HTTP Handler
    └── internal/logic/user/         # [半自動] API 邏輯層
        └── create_user_logic.go     # 調用 RPC 服務
```

---

## 🆚 兩種模式深度對比

| 維度 | GORM→HTTP | Ent→gRPC |
|------|-----------|----------|
| **架構** | 單體/分層 | 微服務 |
| **協議** | HTTP/REST | gRPC |
| **ORM** | GORM (反射) | Ent (代碼生成) |
| **類型安全** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **性能** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **代碼生成** | 手動或部分 | 完全自動 |
| **關聯處理** | 手動 Preload | 自動 Edge |
| **遷移** | AutoMigrate | Schema-First |
| **微服務** | 需要封裝 | 原生支援 |
| **學習曲線** | 低 | 中高 |

---

## 💡 Ent 的核心優勢

### 1. **Schema as Code (代碼即架構)**

```go
// Ent Schema 是可執行的 Go 代碼
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("email").
            Unique().
            NotEmpty().
            Validate(func(s string) error {
                if !strings.Contains(s, "@") {
                    return errors.New("invalid email")
                }
                return nil
            }),
    }
}
```

### 2. **類型安全的查詢**

```go
// GORM (運行時錯誤)
db.Where("usrname = ?", "admin").Find(&user)  // 拼寫錯誤！

// Ent (編譯時檢查)
db.User.Query().
    Where(user.UsernameEQ("admin")).  // ✅ 自動補全
    First(ctx)
```

### 3. **自動處理關聯**

```go
// 查詢用戶及其角色、部門
user, err := db.User.Query().
    Where(user.IDEQ(uuid)).
    WithRoles().              // ✅ 自動 JOIN
    WithDepartments().        // ✅ 自動 JOIN
    Only(ctx)

// 訪問關聯數據
for _, role := range user.Edges.Roles {
    fmt.Println(role.Name)
}
```

### 4. **內置軟刪除**

```go
// Schema 定義
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        SoftDeleteMixin{},  // 自動添加 deleted_at
    }
}

// 自動過濾已刪除記錄
users := db.User.Query().All(ctx)  // 不包含已刪除
```

---

## 🔧 實際應用場景

### 場景 1：新增欄位

**GORM→HTTP 模式：**
```go
// 1. 修改 Model
type User struct {
    PhoneVerified bool  // 新增
}

// 2. 修改 DTO
type UserDTO struct {
    PhoneVerified bool  // 新增
}

// 3. 修改 Handler
func CreateUser() {
    // 手動處理新欄位
}

// 4. 運行遷移
db.AutoMigrate(&User{})
```

**Ent→gRPC 模式：**
```go
// 1. 修改 Schema
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Bool("phone_verified").Default(false),  // 新增
    }
}

// 2. 重新生成（自動更新所有層）
make gen-ent
make gen-rpc-ent-logic model=User group=user
make gen-api

// ✅ Proto、Logic、Handler 全部自動更新
```

### 場景 2：複雜查詢

```go
// Ent 的優雅語法
users, err := db.User.Query().
    Where(
        user.And(
            user.StatusEQ(1),
            user.Or(
                user.UsernameContains("admin"),
                user.EmailContains("@admin.com"),
            ),
        ),
    ).
    WithRoles(func(q *ent.RoleQuery) {
        q.Where(role.NameIn("Admin", "Manager"))
    }).
    Order(ent.Desc(user.FieldCreatedAt)).
    Offset(10).
    Limit(20).
    All(ctx)
```

---

## 🎓 學習路徑建議

### 對於 Hospital ERP 專案

#### 階段 1：理解概念（1-2 天）
- ✅ 閱讀 Ent 官方文檔
- ✅ 理解 Schema-First 思想
- ✅ 研究 Simple Admin 範例

#### 階段 2：小範圍試驗（1 週）
- 🔧 建立獨立的測試專案
- 🔧 用 Ent 重寫一個簡單模組（如 Dictionary）
- 🔧 比較開發效率

#### 階段 3：選擇性整合（2-3 週）
**推薦方案：混合架構**

```
Hospital ERP
├── 醫療核心業務（保持 GORM + HTTP）
│   ├── patient/
│   ├── doctor/
│   └── appointment/
│
└── 通用管理模組（整合 Ent + gRPC）
    ├── user/         ← 從 Simple Admin 整合
    ├── role/         ← 從 Simple Admin 整合
    ├── permission/   ← Casbin 整合
    └── dictionary/   ← 從 Simple Admin 整合
```

---

## 📊 整合建議矩陣

| 功能模組 | 推薦技術 | 理由 |
|---------|---------|------|
| 使用者管理 | Ent + gRPC | 可直接引用 Simple Admin |
| 權限控制 | Casbin | 兩邊都支援 |
| 字典管理 | Ent + gRPC | 可直接引用 Simple Admin |
| 患者管理 | GORM + HTTP | 保持現有架構 |
| 醫生管理 | GORM + HTTP | 保持現有架構 |
| 預約系統 | GORM + HTTP | 業務邏輯複雜 |

---

## 🚀 總結

### Ent→gRPC 的本質

**從資料庫架構自動推導出整個服務層的開發範式**

```
一次定義 Schema
    ↓
自動生成 90% 的代碼
    ↓
只需實現 10% 的業務邏輯
```

### 適合的場景

✅ **推薦使用 Ent→gRPC**
- 微服務架構
- 需要高性能
- 團隊規模較大
- CRUD 密集型應用
- 需要強類型安全

❌ **不推薦使用**
- 小型單體應用
- 快速原型開發
- 團隊對 gRPC 不熟悉
- 業務邏輯極其複雜

### 對 Hospital ERP 的建議

**漸進式整合策略**：
1. 🎯 保持醫療核心業務不動（GORM + HTTP）
2. 🎯 新的通用模組使用 Ent + gRPC
3. 🎯 逐步遷移成熟的管理功能
4. 🎯 兩種模式並存，互相調用

---

## 📚 延伸學習資源

- [Ent 官方文檔](https://entgo.io/docs/getting-started)
- [gRPC Go 快速開始](https://grpc.io/docs/languages/go/quickstart/)
- [Go-Zero 微服務框架](https://go-zero.dev/)
- [Simple Admin 視頻教程](https://space.bilibili.com/9872669)
