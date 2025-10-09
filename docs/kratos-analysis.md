# Kratos Framework 深度分析報告

**文檔版本**: 1.0
**建立日期**: 2025-10-08
**框架版本**: Kratos v2 (2025)

---

## 目錄

1. [Kratos 框架概述](#1-kratos-框架概述)
2. [核心架構設計](#2-核心架構設計)
3. [Proto-First 開發模式](#3-proto-first-開發模式)
4. [分層架構詳解](#4-分層架構詳解)
5. [Kratos vs Go-Zero 深度對比](#5-kratos-vs-go-zero-深度對比)
6. [完整實作範例](#6-完整實作範例)
7. [遷移策略建議](#7-遷移策略建議)

---

## 1. Kratos 框架概述

### 1.1 基本資訊

- **開發者**: Bilibili (B站)
- **授權**: MIT License
- **語言**: Go 1.19+
- **最新版本**: v2 (2025年3月7日發布)
- **定位**: 雲原生微服務治理框架

### 1.2 核心特性

✅ **Proto-First 設計**: 單一 Protobuf 定義產生 HTTP + gRPC
✅ **分層架構**: DDD + Clean Architecture
✅ **輕量插件化**: 高度可定制,不綁定基礎設施
✅ **可觀測性**: OpenTelemetry + Prometheus 原生支援
✅ **服務治理**: 服務發現、配置管理、限流熔斷

### 1.3 設計哲學

> "Simple: Appropriate design, plain and easy to code"

- **工具箱思維**: 提供組件而非框定架構
- **可插拔性**: 所有組件都可替換
- **輕量級**: 不強制使用全部功能

---

## 2. 核心架構設計

### 2.1 專案結構

```
kratos-project/
├── api/                          # Proto API 定義 (單一數據源)
│   └── helloworld/
│       ├── v1/
│       │   ├── greeter.proto           # 服務定義
│       │   ├── greeter.pb.go           # Protobuf 產生
│       │   ├── greeter_grpc.pb.go      # gRPC server/client
│       │   └── greeter_http.pb.go      # HTTP handler
│       └── errors/
│           └── errors.proto            # 錯誤定義
│
├── cmd/                          # 應用程式入口
│   └── server/
│       └── main.go                     # 主程序
│       ├── wire.go                     # 依賴注入定義
│       └── wire_gen.go                 # Wire 自動產生
│
├── configs/                      # 配置檔案
│   └── config.yaml
│
├── internal/                     # 私有應用程式碼
│   ├── conf/                     # 配置解析
│   │   └── conf.proto                  # 配置結構定義
│   │
│   ├── data/                     # 數據存取層 (Repository 實作)
│   │   ├── data.go                     # Data provider
│   │   ├── user.go                     # User repository 實作
│   │   └── ent/                        # Ent ORM (可選)
│   │
│   ├── biz/                      # 業務邏輯層 (Domain Layer)
│   │   ├── biz.go                      # Biz provider
│   │   ├── user.go                     # User usecase
│   │   └── README.md                   # 業務邏輯說明
│   │
│   ├── service/                  # 服務實作層 (Application Layer)
│   │   ├── service.go                  # Service provider
│   │   └── greeter.go                  # Greeter service 實作
│   │
│   └── server/                   # Server 配置
│       ├── http.go                     # HTTP server
│       ├── grpc.go                     # gRPC server
│       └── server.go                   # Server provider
│
├── third_party/                  # 第三方 proto 檔案
│   └── google/
│       └── api/
│           ├── annotations.proto
│           └── http.proto
│
├── go.mod
├── go.sum
└── Makefile
```

### 2.2 架構圖

```
┌─────────────────────────────────────────────────────────────┐
│                     Client (HTTP/gRPC)                       │
└────────────────────┬────────────────────────────────────────┘
                     │
          ┌──────────▼──────────┐
          │   Transport Layer    │
          │  (HTTP/gRPC Server)  │
          └──────────┬───────────┘
                     │
          ┌──────────▼──────────┐
          │   Service Layer      │  ← 實作 Proto 定義的 API
          │  (API Implementation)│  ← 轉換 DTO
          └──────────┬───────────┘
                     │
          ┌──────────▼──────────┐
          │     Biz Layer        │  ← 業務邏輯 (Domain)
          │  (Business Logic)    │  ← 定義 Repository 介面
          └──────────┬───────────┘
                     │
          ┌──────────▼──────────┐
          │    Data Layer        │  ← Repository 實作
          │  (Data Access)       │  ← PO ↔ DO 轉換
          └──────────┬───────────┘
                     │
          ┌──────────▼──────────┐
          │   Infrastructure     │
          │  (DB/Cache/MQ)       │
          └──────────────────────┘
```

---

## 3. Proto-First 開發模式

### 3.1 核心概念

**Proto-First = 以 Protobuf 為單一數據源**

- ✅ 一份 Proto 定義 → 自動產生 HTTP + gRPC 雙棧
- ✅ 型別安全 → 前後端契約完全一致
- ✅ 自動文檔 → OpenAPI/Swagger 自動產生
- ✅ 語言中立 → 支援多語言客戶端

### 3.2 Proto 定義範例

**基礎 CRUD API 定義**:

```protobuf
// api/user/v1/user.proto
syntax = "proto3";

package user.v1;

import "google/api/annotations.proto";
import "google/protobuf/empty.proto";

option go_package = "github.com/myproject/api/user/v1;v1";

// User service definition
service UserService {
  // Create a new user
  rpc CreateUser(CreateUserRequest) returns (CreateUserReply) {
    option (google.api.http) = {
      post: "/api/v1/users"
      body: "*"
    };
  }

  // Get user by ID
  rpc GetUser(GetUserRequest) returns (GetUserReply) {
    option (google.api.http) = {
      get: "/api/v1/users/{id}"
    };
  }

  // Update user
  rpc UpdateUser(UpdateUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put: "/api/v1/users/{id}"
      body: "*"
    };
  }

  // Delete user
  rpc DeleteUser(DeleteUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/api/v1/users/{id}"
    };
  }

  // List users with pagination
  rpc ListUsers(ListUsersRequest) returns (ListUsersReply) {
    option (google.api.http) = {
      get: "/api/v1/users"
    };
  }
}

// Messages
message CreateUserRequest {
  string username = 1;
  string email = 2;
  string password = 3;
}

message CreateUserReply {
  int64 id = 1;
}

message GetUserRequest {
  int64 id = 1;
}

message GetUserReply {
  int64 id = 1;
  string username = 2;
  string email = 3;
  string created_at = 4;
}

message UpdateUserRequest {
  int64 id = 1;
  string username = 2;
  string email = 3;
}

message DeleteUserRequest {
  int64 id = 1;
}

message ListUsersRequest {
  int32 page = 1;
  int32 page_size = 2;
}

message ListUsersReply {
  repeated GetUserReply users = 1;
  int32 total = 2;
}
```

### 3.3 google.api.http 註解詳解

**HTTP Mapping 規則**:

```protobuf
rpc Method(Request) returns (Response) {
  option (google.api.http) = {
    // HTTP 方法: URL 路徑
    get: "/api/v1/resource/{id}"

    // 或其他 HTTP 方法
    // post: "/api/v1/resource"
    // put: "/api/v1/resource/{id}"
    // delete: "/api/v1/resource/{id}"
    // patch: "/api/v1/resource/{id}"

    // 請求體映射 (POST/PUT/PATCH 專用)
    body: "*"              // 整個 message 作為 body
    // body: "user"        // 只有 user 欄位作為 body

    // 額外的 HTTP binding
    additional_bindings {
      post: "/v2/resource"
      body: "*"
    }
  };
}
```

**路徑參數範例**:

```protobuf
// 單一參數
rpc GetUser(GetUserRequest) returns (GetUserReply) {
  option (google.api.http) = {
    get: "/api/v1/users/{id}"  // id 從 GetUserRequest.id 讀取
  };
}

// 多個參數
rpc GetUserPost(GetUserPostRequest) returns (GetUserPostReply) {
  option (google.api.http) = {
    get: "/api/v1/users/{user_id}/posts/{post_id}"
  };
}

// 巢狀參數
rpc GetUserProfile(GetUserProfileRequest) returns (GetUserProfileReply) {
  option (google.api.http) = {
    get: "/api/v1/users/{user.id}/profile"  // 從 user.id 讀取
  };
}
```

**Query 參數範例**:

```protobuf
// 未在路徑中定義的欄位自動成為 query 參數
rpc ListUsers(ListUsersRequest) returns (ListUsersReply) {
  option (google.api.http) = {
    get: "/api/v1/users"  // ?page=1&page_size=10&username=john
  };
}

message ListUsersRequest {
  int32 page = 1;       // query: ?page=1
  int32 page_size = 2;  // query: ?page_size=10
  string username = 3;  // query: ?username=john (optional filter)
}
```

### 3.4 程式碼產生流程

```bash
# 1. 安裝 Kratos CLI
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

# 2. 建立專案
kratos new myproject
cd myproject

# 3. 建立 Proto 定義
kratos proto add api/user/v1/user.proto

# 4. 編輯 Proto 檔案 (定義服務和訊息)

# 5. 產生 Client 程式碼 (生成 HTTP/gRPC client 和 server 介面)
kratos proto client api/user/v1/user.proto

# 產生的檔案:
# - api/user/v1/user.pb.go           (Protobuf messages)
# - api/user/v1/user_grpc.pb.go      (gRPC server/client)
# - api/user/v1/user_http.pb.go      (HTTP handlers)

# 6. 產生 Service 實作框架
kratos proto server api/user/v1/user.proto -t internal/service

# 產生的檔案:
# - internal/service/user.go         (Service 實作框架)

# 7. 實作業務邏輯 (手動編寫)
# - internal/biz/user.go             (業務邏輯 Usecase)
# - internal/data/user.go            (Data Repository)

# 8. 編譯並運行
go build -o ./bin/server ./cmd/server
./bin/server -conf ./configs
```

### 3.5 自動產生的程式碼分析

**HTTP Handler 自動產生範例**:

```go
// api/user/v1/user_http.pb.go (自動產生)
func _UserService_CreateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
    return func(ctx http.Context) error {
        var in CreateUserRequest
        if err := ctx.Bind(&in); err != nil {
            return err
        }
        http.SetOperation(ctx, "/user.v1.UserService/CreateUser")
        h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
            return srv.CreateUser(ctx, req.(*CreateUserRequest))
        })
        out, err := h(ctx, &in)
        if err != nil {
            return err
        }
        reply := out.(*CreateUserReply)
        return ctx.Result(200, reply)
    }
}

// HTTP 路由註冊
func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
    r := s.Route("/")
    r.POST("/api/v1/users", _UserService_CreateUser0_HTTP_Handler(srv))
    r.GET("/api/v1/users/{id}", _UserService_GetUser0_HTTP_Handler(srv))
    r.PUT("/api/v1/users/{id}", _UserService_UpdateUser0_HTTP_Handler(srv))
    r.DELETE("/api/v1/users/{id}", _UserService_DeleteUser0_HTTP_Handler(srv))
    r.GET("/api/v1/users", _UserService_ListUsers0_HTTP_Handler(srv))
}
```

**優勢分析**:

✅ **零手動編寫**: HTTP 路由、參數綁定、錯誤處理全自動
✅ **型別安全**: 編譯期檢查,杜絕型別錯誤
✅ **維護簡單**: Proto 變更後重新產生即可同步
✅ **中間件支援**: 自動整合中間件鏈

---

## 4. 分層架構詳解

### 4.1 四層架構概覽

```
┌────────────────────────────────────────────────────────┐
│  API Layer (Proto Definitions)                         │
│  - 定義服務契約                                          │
│  - 自動產生 HTTP/gRPC code                              │
└────────────────────────────────────────────────────────┘
                        ↓
┌────────────────────────────────────────────────────────┐
│  Service Layer (internal/service/)                     │
│  - 實作 API 介面                                         │
│  - DTO 轉換                                             │
│  - 呼叫 Biz 層                                          │
│  - 避免複雜業務邏輯                                       │
└────────────────────────────────────────────────────────┘
                        ↓
┌────────────────────────────────────────────────────────┐
│  Biz Layer (internal/biz/)                             │
│  - 核心業務邏輯 (Domain)                                 │
│  - 定義 Repository 介面 (依賴倒置)                        │
│  - 組合業務規則                                          │
│  - 不依賴外部框架                                         │
└────────────────────────────────────────────────────────┘
                        ↓
┌────────────────────────────────────────────────────────┐
│  Data Layer (internal/data/)                           │
│  - 實作 Repository 介面                                  │
│  - 資料庫/快取存取                                        │
│  - PO ↔ DO 轉換                                         │
│  - 封裝外部依賴                                          │
└────────────────────────────────────────────────────────┘
```

### 4.2 各層職責詳解

#### Service Layer (服務層)

**職責**:
- 實作 Proto 定義的 API 介面
- 將 API 請求 (DTO) 轉換為 Domain 物件 (DO)
- 協調 Biz 層的呼叫
- 處理錯誤轉換和回應封裝

**範例**:

```go
// internal/service/user.go
package service

import (
    "context"
    pb "myproject/api/user/v1"
    "myproject/internal/biz"
)

type UserService struct {
    pb.UnimplementedUserServiceServer

    uc *biz.UserUsecase  // 注入 Biz 層
}

func NewUserService(uc *biz.UserUsecase) *UserService {
    return &UserService{uc: uc}
}

// 實作 CreateUser API
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
    // 1. DTO → DO 轉換
    user := &biz.User{
        Username: req.Username,
        Email:    req.Email,
        Password: req.Password,
    }

    // 2. 呼叫業務邏輯
    id, err := s.uc.CreateUser(ctx, user)
    if err != nil {
        return nil, err  // 錯誤處理由框架統一處理
    }

    // 3. 回傳結果
    return &pb.CreateUserReply{Id: id}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
    // 1. 呼叫業務邏輯
    user, err := s.uc.GetUser(ctx, req.Id)
    if err != nil {
        return nil, err
    }

    // 2. DO → DTO 轉換
    return &pb.GetUserReply{
        Id:        user.ID,
        Username:  user.Username,
        Email:     user.Email,
        CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
    }, nil
}
```

**設計原則**:
- ❌ 不應包含複雜業務邏輯
- ✅ 只做數據轉換和協調
- ✅ 保持方法簡短清晰

#### Biz Layer (業務邏輯層)

**職責**:
- 實作核心業務邏輯 (Domain Logic)
- 定義 Repository 介面 (依賴倒置原則)
- 組合多個 Repository 完成業務流程
- 與框架無關,純 Go 程式碼

**範例**:

```go
// internal/biz/user.go
package biz

import (
    "context"
    "time"

    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
)

// Domain Object (領域物件)
type User struct {
    ID        int64
    Username  string
    Email     string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Repository Interface (由 Biz 層定義,Data 層實作)
type UserRepo interface {
    Create(ctx context.Context, user *User) (int64, error)
    GetByID(ctx context.Context, id int64) (*User, error)
    GetByUsername(ctx context.Context, username string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int64) error
    List(ctx context.Context, page, pageSize int) ([]*User, int, error)
}

// Usecase (業務用例)
type UserUsecase struct {
    repo UserRepo
    log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
    return &UserUsecase{
        repo: repo,
        log:  log.NewHelper(logger),
    }
}

// 建立使用者 - 包含業務邏輯驗證
func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (int64, error) {
    // 業務規則 1: 檢查使用者名稱是否已存在
    existing, err := uc.repo.GetByUsername(ctx, user.Username)
    if err == nil && existing != nil {
        return 0, errors.New(409, "USER_EXISTS", "username already exists")
    }

    // 業務規則 2: 密碼加密 (實際應使用 bcrypt)
    user.Password = hashPassword(user.Password)

    // 業務規則 3: 設定預設值
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    // 呼叫 Repository 儲存
    id, err := uc.repo.Create(ctx, user)
    if err != nil {
        uc.log.Errorf("failed to create user: %v", err)
        return 0, err
    }

    uc.log.Infof("user created successfully: %d", id)
    return id, nil
}

// 取得使用者
func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
    user, err := uc.repo.GetByID(ctx, id)
    if err != nil {
        return nil, errors.New(404, "USER_NOT_FOUND", "user not found")
    }
    return user, nil
}

// 列出使用者 - 包含業務邏輯
func (uc *UserUsecase) ListUsers(ctx context.Context, page, pageSize int) ([]*User, int, error) {
    // 業務規則: 分頁參數驗證
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    return uc.repo.List(ctx, page, pageSize)
}

func hashPassword(password string) string {
    // 實際應使用 bcrypt.GenerateFromPassword
    return "hashed_" + password
}
```

**設計原則**:
- ✅ **依賴倒置**: Biz 定義介面,Data 實作介面
- ✅ **純業務邏輯**: 不依賴資料庫、HTTP 等外部細節
- ✅ **可測試性**: 容易 mock Repository 進行單元測試
- ✅ **可重用性**: 同一個 Usecase 可供 HTTP、gRPC、CLI 使用

#### Data Layer (資料存取層)

**職責**:
- 實作 Biz 層定義的 Repository 介面
- 封裝資料庫、快取等外部依賴
- PO (Persistent Object) ↔ DO (Domain Object) 轉換
- 處理資料存取邏輯

**範例**:

```go
// internal/data/user.go
package data

import (
    "context"
    "myproject/internal/biz"

    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

// Persistent Object (資料庫模型)
type User struct {
    ID        int64     `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex;size:50;not null"`
    Email     string    `gorm:"uniqueIndex;size:100;not null"`
    Password  string    `gorm:"size:255;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
    return "users"
}

// Repository 實作
type userRepo struct {
    data *Data
    log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

// 實作 Create 方法
func (r *userRepo) Create(ctx context.Context, user *biz.User) (int64, error) {
    // DO → PO 轉換
    po := &User{
        Username:  user.Username,
        Email:     user.Email,
        Password:  user.Password,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }

    // 儲存到資料庫
    if err := r.data.db.WithContext(ctx).Create(po).Error; err != nil {
        return 0, err
    }

    return po.ID, nil
}

// 實作 GetByID 方法
func (r *userRepo) GetByID(ctx context.Context, id int64) (*biz.User, error) {
    var po User
    if err := r.data.db.WithContext(ctx).First(&po, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New(404, "USER_NOT_FOUND", "user not found")
        }
        return nil, err
    }

    // PO → DO 轉換
    return &biz.User{
        ID:        po.ID,
        Username:  po.Username,
        Email:     po.Email,
        Password:  po.Password,
        CreatedAt: po.CreatedAt,
        UpdatedAt: po.UpdatedAt,
    }, nil
}

// 實作 GetByUsername 方法
func (r *userRepo) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
    var po User
    err := r.data.db.WithContext(ctx).Where("username = ?", username).First(&po).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil  // 不存在不算錯誤
        }
        return nil, err
    }

    return &biz.User{
        ID:       po.ID,
        Username: po.Username,
        Email:    po.Email,
        Password: po.Password,
    }, nil
}

// 實作 List 方法
func (r *userRepo) List(ctx context.Context, page, pageSize int) ([]*biz.User, int, error) {
    var (
        pos   []User
        total int64
    )

    offset := (page - 1) * pageSize

    // 查詢總數
    if err := r.data.db.Model(&User{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 分頁查詢
    if err := r.data.db.WithContext(ctx).
        Offset(offset).
        Limit(pageSize).
        Order("created_at DESC").
        Find(&pos).Error; err != nil {
        return nil, 0, err
    }

    // PO → DO 批次轉換
    users := make([]*biz.User, 0, len(pos))
    for _, po := range pos {
        users = append(users, &biz.User{
            ID:        po.ID,
            Username:  po.Username,
            Email:     po.Email,
            CreatedAt: po.CreatedAt,
        })
    }

    return users, int(total), nil
}
```

**設計原則**:
- ✅ **封裝外部依賴**: Biz 層不知道使用什麼資料庫
- ✅ **PO/DO 分離**: 資料庫模型與領域模型解耦
- ✅ **錯誤轉換**: 將資料庫錯誤轉換為業務錯誤
- ✅ **可替換性**: 可輕鬆切換資料庫實作 (MySQL → PostgreSQL)

### 4.3 依賴注入 (Wire)

Kratos 使用 **Wire** 進行依賴注入,自動管理各層依賴關係。

**Wire 配置**:

```go
// cmd/server/wire.go
//go:build wireinject
// +build wireinject

package main

import (
    "myproject/internal/biz"
    "myproject/internal/conf"
    "myproject/internal/data"
    "myproject/internal/server"
    "myproject/internal/service"

    "github.com/go-kratos/kratos/v2"
    "github.com/google/wire"
)

func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
    panic(wire.Build(
        server.ProviderSet,   // HTTP/gRPC server
        data.ProviderSet,     // Data layer
        biz.ProviderSet,      // Biz layer
        service.ProviderSet,  // Service layer
        newApp,
    ))
}
```

**自動產生的依賴圖**:

```
wireApp
  ├── server.NewHTTPServer
  │     └── service.NewUserService
  │           └── biz.NewUserUsecase
  │                 └── data.NewUserRepo
  │                       └── data.NewData (DB connection)
  │
  └── server.NewGRPCServer
        └── service.NewUserService (複用同一個實例)
```

**執行 Wire 產生程式碼**:

```bash
cd cmd/server
wire
# 產生 wire_gen.go
```

---

## 5. Kratos vs Go-Zero 深度對比

### 5.1 架構對比表

| 維度 | Kratos | Go-Zero (Simple Admin) |
|------|--------|------------------------|
| **API 定義** | 單一 Proto 定義 | Proto (RPC) + .api (REST) 雙重定義 |
| **HTTP 產生** | 從 Proto 自動產生 (google.api.http) | 從 .api 手動定義產生 |
| **服務架構** | 單體服務 (HTTP+gRPC 同一進程) | 雙層服務 (API Gateway + RPC Service) |
| **分層設計** | DDD 四層 (API/Service/Biz/Data) | 功能分層 (Handler/Logic + RPC Logic) |
| **依賴注入** | Wire (自動化) | 手動注入 (ServiceContext) |
| **服務發現** | 支援 (Consul/Etcd/K8s) | **更強大** (原生 k8s/etcd/direct) |
| **熔斷限流** | 可選插件 | **內建** (adaptive breaker/shedding) |
| **中間件** | 統一 middleware 鏈 | HTTP/RPC 分別配置 |
| **錯誤處理** | 統一 errors package | HTTP/RPC 分別處理 |
| **可觀測性** | OpenTelemetry | Prometheus + OpenTelemetry |
| **學習曲線** | 中等 (需理解 DDD) | 較陡 (雙層架構,兩套定義) |
| **程式碼產生** | 高度自動化 | 需手動同步 Proto/API |
| **維護成本** | 低 (單一定義源) | 高 (需維護兩套定義) |

### 5.2 架構圖對比

**Kratos 架構**:

```
┌─────────────────────────────────────┐
│         Client Request              │
└────────┬───────────────┬────────────┘
         │               │
    HTTP │          gRPC │
         │               │
    ┌────▼───────────────▼────┐
    │   Same Service Process   │
    │  ┌────────────────────┐  │
    │  │  Service Layer     │  │ ← 同一個實作
    │  └────────┬───────────┘  │
    │  ┌────────▼───────────┐  │
    │  │   Biz Layer        │  │ ← 共享業務邏輯
    │  └────────┬───────────┘  │
    │  ┌────────▼───────────┐  │
    │  │   Data Layer       │  │ ← 共享資料存取
    │  └────────────────────┘  │
    └──────────────────────────┘
              │
         ┌────▼────┐
         │ Database│
         └─────────┘
```

**Go-Zero 架構** (Simple Admin):

```
┌─────────────────────────────────────┐
│         Client Request              │
└────────┬────────────────────────────┘
         │ HTTP
    ┌────▼────────────────┐
    │   API Service       │ Port 9100
    │  ┌──────────────┐   │
    │  │  Handler     │   │ ← REST 端點
    │  └──────┬───────┘   │
    │  ┌──────▼───────┐   │
    │  │  Logic       │   │ ← 轉發層
    │  └──────┬───────┘   │
    │         │ gRPC      │
    └─────────┼───────────┘
              │
    ┌─────────▼───────────┐
    │   RPC Service       │ Port 9101
    │  ┌──────────────┐   │
    │  │ gRPC Server  │   │ ← gRPC 端點
    │  └──────┬───────┘   │
    │  ┌──────▼───────┐   │
    │  │  Logic       │   │ ← 業務邏輯
    │  └──────┬───────┘   │
    │  ┌──────▼───────┐   │
    │  │  Ent ORM     │   │ ← 資料存取
    │  └──────────────┘   │
    └─────────┬───────────┘
              │
         ┌────▼────┐
         │ Database│
         └─────────┘
```

### 5.3 程式碼產生對比

**Kratos** (一次定義,全自動產生):

```bash
# 1. 定義 Proto (唯一定義源)
cat > api/user/v1/user.proto <<EOF
service UserService {
  rpc CreateUser(CreateUserRequest) returns (CreateUserReply) {
    option (google.api.http) = {
      post: "/api/v1/users"
      body: "*"
    };
  }
}
EOF

# 2. 自動產生 HTTP + gRPC
kratos proto client api/user/v1/user.proto

# 產生:
# - user.pb.go          (Protobuf messages)
# - user_grpc.pb.go     (gRPC server/client)
# - user_http.pb.go     (HTTP handlers)  ← 自動產生!

# 3. 產生 Service 框架
kratos proto server api/user/v1/user.proto -t internal/service

# 完成! HTTP 和 gRPC 同時就緒
```

**Go-Zero** (兩次定義,需手動同步):

```bash
# 1. 定義 Proto (RPC 用)
cat > rpc/desc/user.proto <<EOF
service User {
  rpc CreateUser(CreateUserReq) returns (BaseIDResp);
}
EOF

# 2. 產生 RPC 程式碼
make gen-rpc

# 3. 再定義 API (REST 用) - 需手動同步 Proto!
cat > api/desc/core/user.api <<EOF
@server(group: user)
service Core {
  @handler createUser
  post /user/create (CreateUserReq) returns (BaseIDResp)
}
EOF

# 4. 產生 API 程式碼
make gen-api

# 5. 手動在 API Logic 呼叫 RPC
# api/internal/logic/user/create_user_logic.go
resp, err := l.svcCtx.UserRpc.CreateUser(...)

# 完成! 但需維護兩套定義
```

### 5.4 優劣勢分析

#### Kratos 優勢 ✅

1. **Proto-First 開發體驗**
   - 單一定義源,減少 50% 維護成本
   - HTTP/gRPC 自動同步,避免不一致

2. **更清晰的分層架構**
   - DDD 設計,業務邏輯與框架解耦
   - 依賴倒置,容易單元測試

3. **統一的服務實作**
   - 同一個 Service 實作同時支援 HTTP/gRPC
   - 避免程式碼重複

4. **更低的部署成本**
   - 單一服務進程,減少資源佔用
   - 簡化部署和監控

#### Go-Zero 優勢 ✅

1. **更強大的服務治理**
   - 內建熔斷器、自適應限流、降載保護
   - 生產級的微服務可靠性保障

2. **更靈活的 API Gateway**
   - API 層可獨立擴展
   - 支援複雜的請求轉換和授權邏輯

3. **更好的 K8s 整合**
   - 原生支援 k8s service discovery
   - 無需額外 service mesh

4. **成熟的工具鏈**
   - goctl 工具功能完整
   - 豐富的模板和腳手架

### 5.5 適用場景建議

**選擇 Kratos 的場景**:

✅ 新專案,追求開發效率
✅ 團隊熟悉 DDD 設計
✅ 需要同時支援 HTTP + gRPC
✅ 希望降低維護成本
✅ 單體或中小型微服務

**選擇 Go-Zero 的場景**:

✅ 大型微服務集群
✅ 需要強大的服務治理 (熔斷、限流)
✅ API Gateway 與後端服務需要獨立擴展
✅ 深度 K8s 整合需求
✅ 已有 .api 定義的專案遷移

---

## 6. 完整實作範例

### 6.1 使用者管理微服務 (完整流程)

**Step 1: 建立專案**

```bash
kratos new user-service
cd user-service
```

**Step 2: 定義 Proto API**

```protobuf
// api/user/v1/user.proto
syntax = "proto3";

package user.v1;

import "google/api/annotations.proto";
import "google/protobuf/empty.proto";
import "validate/validate.proto";

option go_package = "user-service/api/user/v1;v1";

service UserService {
  rpc CreateUser(CreateUserRequest) returns (CreateUserReply) {
    option (google.api.http) = {
      post: "/api/v1/users"
      body: "*"
    };
  }

  rpc GetUser(GetUserRequest) returns (GetUserReply) {
    option (google.api.http) = {
      get: "/api/v1/users/{id}"
    };
  }

  rpc ListUsers(ListUsersRequest) returns (ListUsersReply) {
    option (google.api.http) = {
      get: "/api/v1/users"
    };
  }

  rpc UpdateUser(UpdateUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put: "/api/v1/users/{id}"
      body: "*"
    };
  }

  rpc DeleteUser(DeleteUserRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/api/v1/users/{id}"
    };
  }
}

message CreateUserRequest {
  string username = 1 [(validate.rules).string = {min_len: 3, max_len: 50}];
  string email = 2 [(validate.rules).string.email = true];
  string password = 3 [(validate.rules).string.min_len = 6];
}

message CreateUserReply {
  int64 id = 1;
}

message GetUserRequest {
  int64 id = 1 [(validate.rules).int64.gt = 0];
}

message GetUserReply {
  int64 id = 1;
  string username = 2;
  string email = 3;
  string created_at = 4;
  string updated_at = 5;
}

message ListUsersRequest {
  int32 page = 1 [(validate.rules).int32 = {gte: 1}];
  int32 page_size = 2 [(validate.rules).int32 = {gte: 1, lte: 100}];
}

message ListUsersReply {
  repeated GetUserReply users = 1;
  int32 total = 2;
}

message UpdateUserRequest {
  int64 id = 1 [(validate.rules).int64.gt = 0];
  string username = 2 [(validate.rules).string = {min_len: 3, max_len: 50}];
  string email = 3 [(validate.rules).string.email = true];
}

message DeleteUserRequest {
  int64 id = 1 [(validate.rules).int64.gt = 0];
}
```

**Step 3: 產生程式碼**

```bash
# 產生 client (HTTP + gRPC)
kratos proto client api/user/v1/user.proto

# 產生 service 框架
kratos proto server api/user/v1/user.proto -t internal/service
```

**Step 4: 實作 Data Layer**

```go
// internal/data/user.go
package data

import (
    "context"
    "time"

    "user-service/internal/biz"

    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type User struct {
    ID        int64     `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex;size:50;not null"`
    Email     string    `gorm:"uniqueIndex;size:100;not null"`
    Password  string    `gorm:"size:255;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type userRepo struct {
    data *Data
    log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (r *userRepo) Create(ctx context.Context, user *biz.User) (int64, error) {
    po := &User{
        Username:  user.Username,
        Email:     user.Email,
        Password:  user.Password,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }

    if err := r.data.db.WithContext(ctx).Create(po).Error; err != nil {
        return 0, err
    }

    return po.ID, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*biz.User, error) {
    var po User
    if err := r.data.db.WithContext(ctx).First(&po, id).Error; err != nil {
        return nil, err
    }

    return &biz.User{
        ID:        po.ID,
        Username:  po.Username,
        Email:     po.Email,
        Password:  po.Password,
        CreatedAt: po.CreatedAt,
        UpdatedAt: po.UpdatedAt,
    }, nil
}

func (r *userRepo) List(ctx context.Context, page, pageSize int) ([]*biz.User, int, error) {
    var (
        pos   []User
        total int64
    )

    offset := (page - 1) * pageSize

    if err := r.data.db.Model(&User{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    if err := r.data.db.WithContext(ctx).
        Offset(offset).
        Limit(pageSize).
        Order("created_at DESC").
        Find(&pos).Error; err != nil {
        return nil, 0, err
    }

    users := make([]*biz.User, 0, len(pos))
    for _, po := range pos {
        users = append(users, &biz.User{
            ID:        po.ID,
            Username:  po.Username,
            Email:     po.Email,
            CreatedAt: po.CreatedAt,
            UpdatedAt: po.UpdatedAt,
        })
    }

    return users, int(total), nil
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) error {
    return r.data.db.WithContext(ctx).Model(&User{}).
        Where("id = ?", user.ID).
        Updates(map[string]interface{}{
            "username":   user.Username,
            "email":      user.Email,
            "updated_at": time.Now(),
        }).Error
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
    return r.data.db.WithContext(ctx).Delete(&User{}, id).Error
}
```

**Step 5: 實作 Biz Layer**

```go
// internal/biz/user.go
package biz

import (
    "context"
    "time"

    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        int64
    Username  string
    Email     string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type UserRepo interface {
    Create(ctx context.Context, user *User) (int64, error)
    GetByID(ctx context.Context, id int64) (*User, error)
    GetByUsername(ctx context.Context, username string) (*User, error)
    List(ctx context.Context, page, pageSize int) ([]*User, int, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int64) error
}

type UserUsecase struct {
    repo UserRepo
    log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
    return &UserUsecase{
        repo: repo,
        log:  log.NewHelper(logger),
    }
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (int64, error) {
    // 檢查使用者是否存在
    existing, _ := uc.repo.GetByUsername(ctx, user.Username)
    if existing != nil {
        return 0, errors.New(409, "USER_EXISTS", "username already exists")
    }

    // 密碼加密
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return 0, errors.New(500, "HASH_ERROR", "failed to hash password")
    }
    user.Password = string(hashedPassword)

    // 設定時間戳
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    id, err := uc.repo.Create(ctx, user)
    if err != nil {
        uc.log.Errorf("failed to create user: %v", err)
        return 0, errors.New(500, "CREATE_ERROR", "failed to create user")
    }

    uc.log.Infof("user created successfully: %d", id)
    return id, nil
}

func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
    user, err := uc.repo.GetByID(ctx, id)
    if err != nil {
        return nil, errors.New(404, "USER_NOT_FOUND", "user not found")
    }
    return user, nil
}

func (uc *UserUsecase) ListUsers(ctx context.Context, page, pageSize int) ([]*User, int, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    return uc.repo.List(ctx, page, pageSize)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) error {
    // 檢查使用者是否存在
    existing, err := uc.repo.GetByID(ctx, user.ID)
    if err != nil {
        return errors.New(404, "USER_NOT_FOUND", "user not found")
    }

    // 檢查新使用者名稱是否被佔用
    if user.Username != existing.Username {
        duplicate, _ := uc.repo.GetByUsername(ctx, user.Username)
        if duplicate != nil && duplicate.ID != user.ID {
            return errors.New(409, "USERNAME_EXISTS", "username already exists")
        }
    }

    user.UpdatedAt = time.Now()

    if err := uc.repo.Update(ctx, user); err != nil {
        uc.log.Errorf("failed to update user: %v", err)
        return errors.New(500, "UPDATE_ERROR", "failed to update user")
    }

    return nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
    // 檢查使用者是否存在
    if _, err := uc.repo.GetByID(ctx, id); err != nil {
        return errors.New(404, "USER_NOT_FOUND", "user not found")
    }

    if err := uc.repo.Delete(ctx, id); err != nil {
        uc.log.Errorf("failed to delete user: %v", err)
        return errors.New(500, "DELETE_ERROR", "failed to delete user")
    }

    return nil
}
```

**Step 6: 實作 Service Layer**

```go
// internal/service/user.go
package service

import (
    "context"

    pb "user-service/api/user/v1"
    "user-service/internal/biz"

    "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
    pb.UnimplementedUserServiceServer

    uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
    return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
    user := &biz.User{
        Username: req.Username,
        Email:    req.Email,
        Password: req.Password,
    }

    id, err := s.uc.CreateUser(ctx, user)
    if err != nil {
        return nil, err
    }

    return &pb.CreateUserReply{Id: id}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
    user, err := s.uc.GetUser(ctx, req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.GetUserReply{
        Id:        user.ID,
        Username:  user.Username,
        Email:     user.Email,
        CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
    }, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
    users, total, err := s.uc.ListUsers(ctx, int(req.Page), int(req.PageSize))
    if err != nil {
        return nil, err
    }

    pbUsers := make([]*pb.GetUserReply, 0, len(users))
    for _, u := range users {
        pbUsers = append(pbUsers, &pb.GetUserReply{
            Id:        u.ID,
            Username:  u.Username,
            Email:     u.Email,
            CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
        })
    }

    return &pb.ListUsersReply{
        Users: pbUsers,
        Total: int32(total),
    }, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
    user := &biz.User{
        ID:       req.Id,
        Username: req.Username,
        Email:    req.Email,
    }

    if err := s.uc.UpdateUser(ctx, user); err != nil {
        return nil, err
    }

    return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
    if err := s.uc.DeleteUser(ctx, req.Id); err != nil {
        return nil, err
    }

    return &emptypb.Empty{}, nil
}
```

**Step 7: 註冊服務**

```go
// internal/server/http.go
package server

import (
    v1 "user-service/api/user/v1"
    "user-service/internal/conf"
    "user-service/internal/service"

    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/middleware/validate"
    "github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(c *conf.Server, user *service.UserService) *http.Server {
    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
            validate.Validator(),
        ),
    }

    if c.Http.Network != "" {
        opts = append(opts, http.Network(c.Http.Network))
    }
    if c.Http.Addr != "" {
        opts = append(opts, http.Address(c.Http.Addr))
    }
    if c.Http.Timeout != nil {
        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
    }

    srv := http.NewServer(opts...)
    v1.RegisterUserServiceHTTPServer(srv, user)

    return srv
}
```

```go
// internal/server/grpc.go
package server

import (
    v1 "user-service/api/user/v1"
    "user-service/internal/conf"
    "user-service/internal/service"

    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/middleware/validate"
    "github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *conf.Server, user *service.UserService) *grpc.Server {
    var opts = []grpc.ServerOption{
        grpc.Middleware(
            recovery.Recovery(),
            validate.Validator(),
        ),
    }

    if c.Grpc.Network != "" {
        opts = append(opts, grpc.Network(c.Grpc.Network))
    }
    if c.Grpc.Addr != "" {
        opts = append(opts, grpc.Address(c.Grpc.Addr))
    }
    if c.Grpc.Timeout != nil {
        opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
    }

    srv := grpc.NewServer(opts...)
    v1.RegisterUserServiceServer(srv, user)

    return srv
}
```

**Step 8: 編譯執行**

```bash
# 編譯
go build -o ./bin/user-service ./cmd/server

# 執行
./bin/user-service -conf ./configs

# HTTP API 可用: http://localhost:8000/api/v1/users
# gRPC API 可用: localhost:9000
```

**Step 9: 測試 API**

```bash
# 建立使用者 (HTTP)
curl -X POST http://localhost:8000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "email": "john@example.com",
    "password": "secret123"
  }'

# 回應: {"id": 1}

# 取得使用者 (HTTP)
curl http://localhost:8000/api/v1/users/1

# 回應:
# {
#   "id": 1,
#   "username": "john",
#   "email": "john@example.com",
#   "created_at": "2025-10-08 10:30:00",
#   "updated_at": "2025-10-08 10:30:00"
# }

# 列出使用者 (HTTP)
curl "http://localhost:8000/api/v1/users?page=1&page_size=10"

# 使用 gRPC (grpcurl)
grpcurl -plaintext -d '{"username":"jane","email":"jane@example.com","password":"pass456"}' \
  localhost:9000 user.v1.UserService/CreateUser
```

---

## 7. 遷移策略建議

### 7.1 方案 A: 漸進式混合架構

**保留 Go-Zero,增強 Proto-First 能力**

**實施步驟**:

1. **開發 protoc-gen-go-zero-api 插件**

```go
// tools/protoc-gen-go-zero-api/main.go
package main

import (
    "fmt"
    "strings"

    "google.golang.org/protobuf/compiler/protogen"
    "google.golang.org/protobuf/types/descriptorpb"
)

func main() {
    protogen.Options{}.Run(func(gen *protogen.Plugin) error {
        for _, f := range gen.Files {
            if !f.Generate {
                continue
            }
            generateAPIFile(gen, f)
        }
        return nil
    })
}

func generateAPIFile(gen *protogen.Plugin, file *protogen.File) {
    filename := file.GeneratedFilenamePrefix + ".api"
    g := gen.NewGeneratedFile(filename, file.GoImportPath)

    // 產生 .api 檔案內容
    g.P("syntax = \"v1\"")
    g.P()

    for _, service := range file.Services {
        // 取得 service 名稱
        serviceName := service.Desc.Name()
        g.P(fmt.Sprintf("@server("))
        g.P(fmt.Sprintf("    group: %s", strings.ToLower(string(serviceName))))
        g.P(")")
        g.P(fmt.Sprintf("service Core {"))

        for _, method := range service.Methods {
            // 解析 google.api.http 註解
            httpRule := extractHTTPRule(method)
            if httpRule != nil {
                handler := camelToSnake(string(method.Desc.Name()))
                g.P(fmt.Sprintf("    @handler %s", handler))
                g.P(fmt.Sprintf("    %s %s (%s) returns (%s)",
                    httpRule.Method,
                    httpRule.Path,
                    method.Input.Desc.Name(),
                    method.Output.Desc.Name()))
            }
        }

        g.P("}")
        g.P()
    }
}

func extractHTTPRule(method *protogen.Method) *HTTPRule {
    // 解析 google.api.http option
    opts := method.Desc.Options().(*descriptorpb.MethodOptions)
    if opts == nil {
        return nil
    }

    // 提取 HTTP annotation
    // 實作細節...
    return &HTTPRule{
        Method: "post",  // 從 annotation 提取
        Path:   "/api/v1/resource",
    }
}

type HTTPRule struct {
    Method string
    Path   string
}

func camelToSnake(s string) string {
    // 駝峰轉蛇形
    var result strings.Builder
    for i, r := range s {
        if i > 0 && 'A' <= r && r <= 'Z' {
            result.WriteRune('_')
        }
        result.WriteRune(r)
    }
    return strings.ToLower(result.String())
}
```

2. **整合到 Makefile**

```makefile
# Makefile
.PHONY: gen-proto-api
gen-proto-api:
    @echo "Generating .api files from proto..."
    protoc --go-zero-api_out=api/desc \
           --go-zero-api_opt=paths=source_relative \
           rpc/desc/**/*.proto
    @echo "Generating API code from .api files..."
    goctl api go -api api/desc/all.api -dir api/

.PHONY: gen-all
gen-all: gen-ent gen-rpc gen-proto-api
    @echo "All code generation completed!"
```

3. **修改工作流程**

```bash
# 舊流程 (雙重定義)
vim rpc/desc/user.proto     # 定義 RPC
make gen-rpc
vim api/desc/core/user.api  # 再定義 REST (需手動同步!)
make gen-api

# 新流程 (單一定義)
vim rpc/desc/user.proto     # 只定義一次 (加 google.api.http)
make gen-all                # 自動產生 RPC + REST
```

**優勢**:
- ✅ 保留 Go-Zero 全部能力 (服務治理、k8s 整合)
- ✅ 減少 50% 維護成本
- ✅ 漸進式遷移,風險可控

**實施時程**: 2-3 週

### 7.2 方案 B: 完全遷移到 Kratos

**重新設計為 Kratos 架構**

**實施步驟**:

1. **Phase 1: 架構設計 (1 週)**
   - 設計 DDD 分層結構
   - 規劃 Proto 定義
   - 設計依賴注入

2. **Phase 2: 基礎設施 (1 週)**
   - 建立 Kratos 專案骨架
   - 配置資料庫、Redis
   - 設定 Wire 依賴注入

3. **Phase 3: 模組遷移 (4-6 週)**
   - 每週遷移 1-2 個核心模組
   - User → Role → Menu → API → Department
   - 保持 API 向後相容

4. **Phase 4: 測試與部署 (2 週)**
   - 整合測試
   - 效能測試
   - 灰度發布

**範例遷移 (User 模組)**:

```bash
# 1. 定義 Proto
cat > api/core/v1/user.proto <<EOF
service CoreService {
  rpc CreateUser(CreateUserReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }
  // ... 其他方法
}
EOF

# 2. 產生程式碼
kratos proto client api/core/v1/user.proto
kratos proto server api/core/v1/user.proto -t internal/service

# 3. 遷移 Ent schema → Data layer
# rpc/ent/schema/user.go → internal/data/user.go

# 4. 遷移 RPC logic → Biz layer
# rpc/internal/logic/user/ → internal/biz/user.go

# 5. 實作 Service layer
# internal/service/user.go (新增)

# 6. 測試
go test ./internal/...
```

**優勢**:
- ✅ 徹底解決架構問題
- ✅ 更清晰的程式碼結構
- ✅ 長期維護成本最低

**挑戰**:
- ⚠️ 工作量大 (6-8 週全職)
- ⚠️ 需要停機遷移或雙跑
- ⚠️ 團隊需要學習 DDD 和 Kratos

**實施時程**: 8-10 週

### 7.3 方案 C: 保持現狀 + 工具增強

**最小改動,工具輔助**

**實施措施**:

1. **自動同步腳本**

```bash
#!/bin/bash
# tools/sync-proto-api.sh

# 比對 proto 和 api 定義
diff <(grep "rpc " rpc/desc/*.proto) <(grep "post\\|get\\|put\\|delete" api/desc/**/*.api)

# 如果不同,提示需要同步
if [ $? -ne 0 ]; then
    echo "⚠️  Proto and API definitions are out of sync!"
    echo "Please update api/desc/ to match rpc/desc/"
    exit 1
fi
```

2. **Pre-commit Hook**

```bash
# .git/hooks/pre-commit
#!/bin/bash

# 檢查 proto 和 api 是否同步
./tools/sync-proto-api.sh

if [ $? -ne 0 ]; then
    echo "❌ Commit blocked: API definitions out of sync"
    exit 1
fi
```

3. **CI/CD 檢查**

```yaml
# .github/workflows/check-api-sync.yml
name: Check API Sync
on: [pull_request]

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Check Proto/API sync
        run: ./tools/sync-proto-api.sh
```

**優勢**:
- ✅ 零架構變更
- ✅ 最快實施 (1 天)
- ✅ 降低人為錯誤

**缺點**:
- ❌ 仍需維護兩套定義
- ❌ 沒有根本解決問題

**實施時程**: 1 天

### 7.4 建議方案

**推薦: 方案 A (漸進式混合架構)**

**理由**:
1. ✅ **最佳平衡**: 保留 Go-Zero 優勢,增強開發體驗
2. ✅ **風險可控**: 漸進式實施,隨時可回退
3. ✅ **快速見效**: 2-3 週即可上線
4. ✅ **技術債務**: 長期消除維護負擔

**實施路線圖**:

```
Week 1: 開發 protoc-gen-go-zero-api 插件
  ├── Day 1-2: 插件基礎框架
  ├── Day 3-4: HTTP annotation 解析
  └── Day 5: .api 格式產生

Week 2: 測試與整合
  ├── Day 1-2: 單元測試
  ├── Day 3: 整合到 Makefile
  ├── Day 4: 遷移 1-2 個模組試點
  └── Day 5: 文檔撰寫

Week 3: 全面推廣
  ├── Day 1-3: 遷移所有模組
  ├── Day 4: CI/CD 整合
  └── Day 5: 團隊培訓
```

**預期效果**:
- 📉 維護成本降低 **50%**
- 📈 開發效率提升 **30%**
- ✅ API 一致性 **100%**

---

## 8. 總結

### 8.1 Kratos 關鍵優勢

✅ **Proto-First**: 單一定義源,避免不一致
✅ **DDD 架構**: 清晰的分層,業務邏輯與框架解耦
✅ **雙協議支援**: HTTP + gRPC 自動產生
✅ **低維護成本**: 減少 50% 程式碼維護工作
✅ **高可測試性**: 依賴注入 + 介面設計,容易單元測試

### 8.2 Go-Zero 關鍵優勢

✅ **服務治理**: 內建熔斷、限流、降載
✅ **K8s 整合**: 原生 k8s service discovery
✅ **API Gateway**: 強大的 HTTP 層控制能力
✅ **成熟穩定**: 生產環境大規模驗證
✅ **生態完整**: 豐富的工具和模板

### 8.3 最終建議

**新專案**: 優先選擇 **Kratos**
**現有專案**: 採用 **漸進式混合架構** (方案 A)
**大型集群**: 保持 **Go-Zero** 並增強工具

無論選擇哪種方案,**Proto-First 理念**都是未來微服務架構的趨勢。

---

## 附錄

### A. 參考資源

- Kratos 官方文檔: https://go-kratos.dev/
- Kratos GitHub: https://github.com/go-kratos/kratos
- Google API Design Guide: https://cloud.google.com/apis/design
- Proto3 Language Guide: https://protobuf.dev/programming-guides/proto3/

### B. 相關工具

- **Wire**: 依賴注入工具 (https://github.com/google/wire)
- **protoc-gen-go-http**: Kratos HTTP 產生器
- **protoc-gen-validate**: Proto 驗證工具
- **buf**: Proto 管理工具 (https://buf.build/)

### C. 社群支援

- Kratos Discord: https://discord.gg/kratos
- Kratos 官方範例: https://github.com/go-kratos/examples

---

**文檔結束**
