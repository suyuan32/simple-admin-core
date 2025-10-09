# Technical Plan: User Module Proto Completion and Proto-First Migration

**Related Spec**: [spec.md](./spec.md)
**Created**: 2025-10-10
**Status**: Draft
**Estimated Effort**: 24-28 hours (3.5 days)

---

## Architecture Overview

### System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                      Proto-First Pipeline                       │
│                                                                 │
│  user.proto (Source of Truth)                                  │
│      │                                                          │
│      ├──▶ protoc-gen-go ──────────▶ types/core/*.pb.go        │
│      │                                                          │
│      └──▶ protoc-gen-go-zero-api ─▶ api/desc/core/user.api    │
│                   │                                             │
│                   └───▶ goctls api go ──▶ API Handlers         │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Data Flow:
1. Developer edits user.proto (add/modify endpoint)
2. Run `make gen-rpc` → generates RPC types
3. Run `make gen-proto-api` → generates user.api
4. Run `make gen-api-code` → generates API handlers
5. Implement RPC logic in rpc/internal/logic/user/
6. API handlers automatically call RPC methods
```

### Technology Stack

**Backend**:
- Go 1.25+
- Go-Zero (microservice framework)
- gRPC (RPC communication)
- Protocol Buffers v3

**Database**:
- Ent ORM (schema-first ORM)
- PostgreSQL/MySQL (database)

**Tools**:
- protoc (Protocol Buffer compiler)
- protoc-gen-go (official Go plugin)
- protoc-gen-go-zero-api (custom plugin for .api generation)
- goctls (go-zero CLI tool)

---

## Implementation Details

### Phase 1: Proto Definition (Day 1, 4-6 hours)

#### Task 1.1: Prepare user.proto structure (0.5h)

**File**: `rpc/desc/user.proto`

**Current Structure**:
```protobuf
syntax = "proto3";

message UserInfo { /* 19 fields */ }
message UserListResp { /* ... */ }
message UserListReq { /* ... */ }
message UsernameReq { /* ... */ }

service Core {
  // Only 6 methods
  rpc createUser (UserInfo) returns (BaseUUIDResp);
  rpc updateUser (UserInfo) returns (BaseResp);
  rpc getUserList (UserListReq) returns (UserListResp);
  rpc getUserById (UUIDReq) returns (UserInfo);
  rpc getUserByUsername (UsernameReq) returns (UserInfo);
  rpc deleteUser (UUIDsReq) returns (BaseResp);
}
```

**Target Structure**:
```protobuf
syntax = "proto3";

import "google/api/annotations.proto";
import "go_zero/options.proto";
import "base.proto";

option go_package = "github.com/chimerakang/simple-admin-core/rpc/types/core";

// File-level API metadata
option (go_zero.api_info) = {
  title: "User Management API"
  desc: "User authentication and management services"
  author: "Ryan Su"
  email: "yuansu.china.work@gmail.com"
  version: "v1.0"
};

// ... (all message definitions)

service Core {
  // Service-level default options
  option (go_zero.jwt) = "Auth";
  option (go_zero.middleware) = "Authority";
  option (go_zero.group) = "user";

  // 22 complete methods with HTTP annotations
}
```

#### Task 1.2: Add authentication message types (1h)

**Messages to add**:

```protobuf
// ============= Login Messages =============

message LoginReq {
  string username = 1;
  string password = 2;
  string captcha_id = 3;
  string captcha = 4;
}

message LoginByEmailReq {
  string email = 1;
  string captcha = 2;
}

message LoginBySmsReq {
  string phone_number = 1;
  string captcha = 2;
}

message LoginInfo {
  string user_id = 1;
  string token = 2;
  uint64 expire = 3;
}

message LoginResp {
  uint32 code = 1;
  string msg = 2;
  LoginInfo data = 3;
}

// ============= Register Messages =============

message RegisterReq {
  string username = 1;
  string password = 2;
  string captcha_id = 3;
  string captcha = 4;
  string email = 5;
}

message RegisterByEmailReq {
  string username = 1;
  string password = 2;
  string captcha = 3;
  string email = 4;
}

message RegisterBySmsReq {
  string username = 1;
  string password = 2;
  string captcha = 3;
  string phone_number = 4;
}

// ============= Password Reset Messages =============

message ChangePasswordReq {
  string old_password = 1;
  string new_password = 2;
}

message ResetPasswordByEmailReq {
  string email = 1;
  string captcha = 2;
  string password = 3;
}

message ResetPasswordBySmsReq {
  string phone_number = 1;
  string captcha = 2;
  string password = 3;
}

// ============= User Info Messages =============

message UserBaseIDInfo {
  optional string uuid = 1;
  optional string username = 2;
  optional string nickname = 3;
  optional string avatar = 4;
  optional string home_path = 5;
  optional string description = 6;
  repeated string role_name = 7;
  string department_name = 8;
  optional string locale = 9;
}

message UserBaseIDInfoResp {
  uint32 code = 1;
  string msg = 2;
  UserBaseIDInfo data = 3;
}

message PermCodeResp {
  uint32 code = 1;
  string msg = 2;
  repeated string data = 3;
}

// ============= Profile Messages =============

message ProfileInfo {
  optional string nickname = 1;
  optional string avatar = 2;
  optional string mobile = 3;
  optional string email = 4;
  optional string locale = 5;
}

message ProfileResp {
  uint32 code = 1;
  string msg = 2;
  ProfileInfo data = 3;
}

// ============= Token Messages =============

message RefreshTokenInfo {
  string token = 1;
  int64 expired_at = 2;
}

message RefreshTokenResp {
  uint32 code = 1;
  string msg = 2;
  RefreshTokenInfo data = 3;
}
```

#### Task 1.3: Add public endpoints (login, register) with annotations (1h)

**Public Endpoints Group**:

```protobuf
service Core {
  // ============= Public Endpoints (no JWT) =============

  rpc login(LoginReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login"
      body: "*"
    };
    option (go_zero.public) = true;  // Override service-level JWT
  }

  rpc loginByEmail(LoginByEmailReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login_by_email"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  rpc loginBySms(LoginBySmsReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login_by_sms"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  rpc register(RegisterReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/register"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  rpc registerByEmail(RegisterByEmailReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/register_by_email"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  rpc registerBySms(RegisterBySmsReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/register_by_sms"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  rpc resetPasswordByEmail(ResetPasswordByEmailReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/reset_password_by_email"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  rpc resetPasswordBySms(ResetPasswordBySmsReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/reset_password_by_sms"
      body: "*"
    };
    option (go_zero.public) = true;
  }
}
```

**Key Points**:
- All use `option (go_zero.public) = true` to override service-level JWT requirement
- HTTP method: POST (matches existing API)
- Request body: `body: "*"` (entire request as body)

#### Task 1.4: Add protected endpoints with annotations (1.5h)

**Protected Endpoints Group**:

```protobuf
service Core {
  // Service-level defaults (applies to all unless overridden)
  option (go_zero.jwt) = "Auth";
  option (go_zero.middleware) = "Authority";
  option (go_zero.group) = "user";

  // ============= Protected Endpoints (require JWT + Authority) =============

  // Existing CRUD methods (add annotations)
  rpc createUser(UserInfo) returns (BaseUUIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }

  rpc updateUser(UserInfo) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/update"
      body: "*"
    };
  }

  rpc deleteUser(UUIDsReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/delete"
      body: "*"
    };
  }

  rpc getUserList(UserListReq) returns (UserListResp) {
    option (google.api.http) = {
      post: "/user/list"
      body: "*"
    };
  }

  rpc getUserById(UUIDReq) returns (UserInfo) {
    option (google.api.http) = {
      post: "/user"
      body: "*"
    };
  }

  // New methods
  rpc changePassword(ChangePasswordReq) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/change_password"
      body: "*"
    };
  }

  rpc getUserInfo(EmptyReq) returns (UserBaseIDInfoResp) {
    option (google.api.http) = {
      get: "/user/info"
    };
  }

  rpc getUserPermCode(EmptyReq) returns (PermCodeResp) {
    option (google.api.http) = {
      get: "/user/perm"
    };
  }

  rpc getUserProfile(EmptyReq) returns (ProfileResp) {
    option (google.api.http) = {
      get: "/user/profile"
    };
  }

  rpc updateUserProfile(ProfileInfo) returns (BaseResp) {
    option (google.api.http) = {
      post: "/user/profile"
      body: "*"
    };
  }

  rpc logout(EmptyReq) returns (BaseResp) {
    option (google.api.http) = {
      get: "/user/logout"
    };
  }

  rpc refreshToken(EmptyReq) returns (RefreshTokenResp) {
    option (google.api.http) = {
      get: "/user/refresh_token"
    };
  }

  rpc accessToken(EmptyReq) returns (RefreshTokenResp) {
    option (google.api.http) = {
      get: "/user/access_token"
    };
  }
}
```

**Summary Table**:

| Method | HTTP | Route | JWT | Middleware | Public |
|--------|------|-------|-----|------------|--------|
| login | POST | /user/login | No | No | ✅ |
| loginByEmail | POST | /user/login_by_email | No | No | ✅ |
| loginBySms | POST | /user/login_by_sms | No | No | ✅ |
| register | POST | /user/register | No | No | ✅ |
| registerByEmail | POST | /user/register_by_email | No | No | ✅ |
| registerBySms | POST | /user/register_by_sms | No | No | ✅ |
| resetPasswordByEmail | POST | /user/reset_password_by_email | No | No | ✅ |
| resetPasswordBySms | POST | /user/reset_password_by_sms | No | No | ✅ |
| createUser | POST | /user/create | Yes | Authority | No |
| updateUser | POST | /user/update | Yes | Authority | No |
| deleteUser | POST | /user/delete | Yes | Authority | No |
| getUserList | POST | /user/list | Yes | Authority | No |
| getUserById | POST | /user | Yes | Authority | No |
| changePassword | POST | /user/change_password | Yes | Authority | No |
| getUserInfo | GET | /user/info | Yes | Authority | No |
| getUserPermCode | GET | /user/perm | Yes | Authority | No |
| getUserProfile | GET | /user/profile | Yes | Authority | No |
| updateUserProfile | POST | /user/profile | Yes | Authority | No |
| logout | GET | /user/logout | Yes | Authority | No |
| refreshToken | GET | /user/refresh_token | Yes | Authority | No |
| accessToken | GET | /user/access_token | Yes | Authority | No |

**Total**: 22 endpoints

#### Task 1.5: Validate Proto syntax (0.5h)

**Commands**:
```bash
# Check if go_zero/options.proto exists
ls -la rpc/desc/go_zero/options.proto

# Validate Proto syntax
cd rpc/desc
protoc --proto_path=. \
  --proto_path=../../tools/protoc-gen-go-zero-api/options \
  --go_out=paths=source_relative:../types/core \
  user.proto

# Expected: No syntax errors
```

**Deliverables**:
- ✅ Complete user.proto with 22 RPC methods
- ✅ All message types defined
- ✅ HTTP annotations for all methods
- ✅ Go-Zero options (jwt, middleware, group, public)
- ✅ Proto syntax validated

---

### Phase 2: RPC Implementation (Day 1-2, 12-16 hours)

#### Task 2.1: Generate RPC server stubs (0.5h)

**Commands**:
```bash
# Generate RPC types and server interface
make gen-rpc

# This will create:
# - rpc/types/core/user.pb.go (protobuf types)
# - rpc/coreclient/core.go (RPC client)
# - rpc/internal/server/coreserver.go (server interface, NEW METHODS ADDED)
```

**Verify**:
```bash
# Check that new methods are in server interface
grep -E "login|register|logout" rpc/internal/server/coreserver.go

# Expected output: 16 new method signatures
```

#### Task 2.2: Implement authentication RPC methods (4-5h)

##### 2.2.1: Login method (1h)

**File**: `rpc/internal/logic/user/login_logic.go`

**Approach**:
- Verify user exists by username
- Check password with bcrypt
- Return user info (NO JWT generation - that stays in API layer)

**Code Example**:
```go
package user

import (
	"context"

	"github.com/chimerakang/simple-admin-common/utils/encrypt"
	"github.com/chimerakang/simple-admin-core/rpc/ent/user"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *core.LoginReq) (*core.LoginResp, error) {
	// 1. Query user by username
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.UsernameEQ(in.Username)).
		Only(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
	}

	// 2. Check if user is active
	if userInfo.Status != 1 {
		return nil, errorx.NewCodeAbortedError("login.userBanned")
	}

	// 3. Verify password
	if !encrypt.BcryptCheck(in.Password, userInfo.Password) {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongUsernameOrPassword")
	}

	// 4. Return user info (API layer will generate JWT)
	// Note: LoginResp contains LoginInfo, but we don't fill token/expire here
	// API layer will handle that after calling this RPC
	return &core.LoginResp{
		Code: 0,
		Msg:  "login.loginSuccessTitle",
		Data: &core.LoginInfo{
			UserId: userInfo.ID.String(),
			// Token and Expire will be set by API layer
		},
	}, nil
}
```

**Testing**:
```go
// rpc/internal/logic/user/login_logic_test.go
func TestLoginLogic_Login(t *testing.T) {
	// Test case 1: Valid credentials
	// Test case 2: Wrong password
	// Test case 3: User not found
	// Test case 4: User banned
}
```

##### 2.2.2: LoginByEmail and LoginBySms (1h)

Similar structure to Login, but query by email/phone instead of username.

##### 2.2.3: Register methods (1h)

**File**: `rpc/internal/logic/user/register_logic.go`

**Approach**:
- Validate email/phone uniqueness
- Call existing createUser method
- Return success message

**Code Example**:
```go
func (l *RegisterLogic) Register(in *core.RegisterReq) (*core.BaseResp, error) {
	// 1. Check if email exists
	exists, err := l.svcCtx.DB.User.Query().
		Where(user.EmailEQ(in.Email)).
		Exist(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exists {
		return nil, errorx.NewInvalidArgumentError("login.signupUserExist")
	}

	// 2. Check if username exists
	exists, err = l.svcCtx.DB.User.Query().
		Where(user.UsernameEQ(in.Username)).
		Exist(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exists {
		return nil, errorx.NewInvalidArgumentError("login.userAlreadyExists")
	}

	// 3. Create user (reuse existing createUser logic)
	createLogic := NewCreateUserLogic(l.ctx, l.svcCtx)
	_, err = createLogic.CreateUser(&core.UserInfo{
		Username: &in.Username,
		Password: &in.Password,
		Email:    &in.Email,
		Status:   pointy.GetPointer(uint32(1)),
		// Set defaults from config
	})
	if err != nil {
		return nil, err
	}

	return &core.BaseResp{
		Code: 0,
		Msg:  "login.registerSuccessTitle",
	}, nil
}
```

##### 2.2.4: Password reset methods (1h)

**Files**:
- `reset_password_by_email_logic.go`
- `reset_password_by_sms_logic.go`

**Approach**:
- Find user by email/phone
- Update password (with bcrypt encryption)
- Return success

#### Task 2.3: Implement user info RPC methods (3-4h)

##### 2.3.1: ChangePassword (0.5h)

**File**: `rpc/internal/logic/user/change_password_logic.go`

**Approach**:
- Get user ID from context (passed by API layer)
- Verify old password
- Update to new password

##### 2.3.2: GetUserInfo (0.5h)

**File**: `rpc/internal/logic/user/get_user_info_logic.go`

**Approach**:
- Get user ID from context
- Query user with role names, department name
- Return UserBaseIDInfo

**Code Example**:
```go
func (l *GetUserInfoLogic) GetUserInfo(in *core.EmptyReq) (*core.UserBaseIDInfoResp, error) {
	// Get user ID from metadata (set by API layer JWT middleware)
	userId := l.ctx.Value("userId").(string)

	// Query user with eager loading of roles and department
	userInfo, err := l.svcCtx.DB.User.Query().
		Where(user.IDEQ(uuid.MustParse(userId))).
		WithRoles().
		WithDepartment().
		Only(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// Build role names
	roleNames := make([]string, 0, len(userInfo.Edges.Roles))
	for _, role := range userInfo.Edges.Roles {
		roleNames = append(roleNames, role.Name)
	}

	// Build department name
	var deptName string
	if userInfo.Edges.Department != nil {
		deptName = userInfo.Edges.Department.Name
	}

	return &core.UserBaseIDInfoResp{
		Code: 0,
		Msg:  "common.success",
		Data: &core.UserBaseIDInfo{
			Uuid:           &userInfo.ID.String(),
			Username:       &userInfo.Username,
			Nickname:       &userInfo.Nickname,
			Avatar:         &userInfo.Avatar,
			HomePath:       &userInfo.HomePath,
			Description:    &userInfo.Description,
			RoleName:       roleNames,
			DepartmentName: deptName,
			Locale:         &userInfo.Locale,
		},
	}, nil
}
```

##### 2.3.3: GetUserPermCode (1h)

**File**: `rpc/internal/logic/user/get_user_perm_code_logic.go`

**Approach**:
- Get user ID from context
- Query all menus/APIs the user has access to (via roles)
- Extract permission codes
- Return string array

##### 2.3.4: Profile methods (1h)

**Files**:
- `get_user_profile_logic.go`
- `update_user_profile_logic.go`

**Approach**:
- GetProfile: Query and return user's nickname, avatar, mobile, email, locale
- UpdateProfile: Update only profile fields (not username, roles, etc.)

#### Task 2.4: Implement token management methods (2-3h)

##### 2.4.1: Logout (0.5h)

**File**: `rpc/internal/logic/user/logout_logic.go`

**Approach**:
- Add token to blacklist (Redis)
- Or do nothing (API layer will clear local token)

**Note**: This might be a minimal implementation since JWT is stateless.

##### 2.4.2: RefreshToken and AccessToken (1.5h)

**Files**:
- `refresh_token_logic.go`
- `access_token_logic.go`

**Approach**:
- Verify current token (via API layer)
- Generate new token with extended expiry
- Return new token info

**Note**: JWT generation might stay in API layer, RPC just validates user status.

#### Task 2.5: Write unit tests (3-4h)

**Test Coverage Target**: ≥70%

**Test Files**:
```bash
rpc/internal/logic/user/
├── login_logic_test.go
├── login_by_email_logic_test.go
├── login_by_sms_logic_test.go
├── register_logic_test.go
├── register_by_email_logic_test.go
├── register_by_sms_logic_test.go
├── change_password_logic_test.go
├── reset_password_by_email_logic_test.go
├── reset_password_by_sms_logic_test.go
├── get_user_info_logic_test.go
├── get_user_perm_code_logic_test.go
├── get_user_profile_logic_test.go
├── update_user_profile_logic_test.go
├── logout_logic_test.go
├── refresh_token_logic_test.go
└── access_token_logic_test.go
```

**Test Template**:
```go
package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
)

func TestLoginLogic_Login(t *testing.T) {
	ctx := context.Background()
	svcCtx := svc.NewServiceContext(/* test config */)

	t.Run("valid credentials", func(t *testing.T) {
		logic := NewLoginLogic(ctx, svcCtx)
		resp, err := logic.Login(&core.LoginReq{
			Username: "testuser",
			Password: "testpass",
		})

		require.NoError(t, err)
		assert.Equal(t, uint32(0), resp.Code)
		assert.NotEmpty(t, resp.Data.UserId)
	})

	t.Run("wrong password", func(t *testing.T) {
		logic := NewLoginLogic(ctx, svcCtx)
		_, err := logic.Login(&core.LoginReq{
			Username: "testuser",
			Password: "wrongpass",
		})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "wrongUsernameOrPassword")
	})

	t.Run("user not found", func(t *testing.T) {
		logic := NewLoginLogic(ctx, svcCtx)
		_, err := logic.Login(&core.LoginReq{
			Username: "nonexistent",
			Password: "testpass",
		})

		require.Error(t, err)
	})
}
```

**Run Tests**:
```bash
# Run all user RPC tests
go test ./rpc/internal/logic/user/... -v -cover

# Check coverage
go test ./rpc/internal/logic/user/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Deliverables**:
- ✅ 16 new RPC logic files implemented
- ✅ All business logic (DB operations, validation) in RPC
- ✅ Auth logic (captcha, JWT) remains in API layer (by design)
- ✅ Unit tests with ≥70% coverage
- ✅ Error handling with i18n messages

---

### Phase 3: Code Generation & Testing (Day 3, 4-6 hours)

#### Task 3.1: Generate .api from Proto (0.5h)

**Commands**:
```bash
# Build protoc-gen-go-zero-api plugin (if not already built)
cd tools/protoc-gen-go-zero-api
go build -o ../../bin/protoc-gen-go-zero-api ./cmd/protoc-gen-go-zero-api

# Generate .api file from user.proto
cd ../../rpc/desc
protoc --proto_path=. \
  --proto_path=../../tools/protoc-gen-go-zero-api/options \
  --go-zero-api_out=../../api/desc/core \
  user.proto

# Output: api/desc/core/user.api.generated
```

**Verify**:
```bash
# Check generated file exists
ls -lh api/desc/core/user.api.generated

# Count endpoints in generated file
grep -c "@handler" api/desc/core/user.api.generated
# Expected: 22
```

#### Task 3.2: Compare generated vs existing .api (1-2h)

**Create Backup**:
```bash
cp api/desc/core/user.api api/desc/core/user.api.backup
```

**Compare Files**:
```bash
# Generate diff
diff -u api/desc/core/user.api.backup api/desc/core/user.api.generated > user_api_diff.txt

# View diff
cat user_api_diff.txt
```

**Expected Differences**:
- ✅ Whitespace/indentation
- ✅ Comment format
- ✅ Type definition order
- ❌ Should NOT differ: endpoint paths, types, middleware config

**Manual Reconciliation** (if needed):
```bash
# If generated file is missing custom types or has wrong config
# Manually merge the differences

# Example: Add custom validation tags
# In user.api.generated, find:
#   Username string `json:"username"`
# Change to:
#   Username string `json:"username" validate:"required,alphanum,max=20"`
```

**Decision Point**:
- If diff is acceptable → proceed to Task 3.3
- If diff has issues → fix protoc-gen-go-zero-api plugin or Proto annotations

#### Task 3.3: Replace .api and regenerate API code (1h)

**Commands**:
```bash
# Replace original with generated
mv api/desc/core/user.api.generated api/desc/core/user.api

# Regenerate API handlers and logic
cd api
goctls api go -api desc/all.api -dir . -style go_zero

# This will regenerate:
# - api/internal/handler/user/*.go (22 handlers)
# - api/internal/handler/publicuser/*.go (8 handlers)
# - api/internal/logic/user/*.go (14 logic files)
# - api/internal/logic/publicuser/*.go (8 logic files)
# - api/internal/types/types.go (all types)
```

**Verify Compilation**:
```bash
# Compile API service
go build -o bin/core-api ./api/core.go

# Expected: No errors
echo $?
# Should output: 0
```

**Check Generated Files**:
```bash
# Count handlers
ls -1 api/internal/handler/user/*.go | wc -l
# Expected: 14

ls -1 api/internal/handler/publicuser/*.go | wc -l
# Expected: 8

# Total: 22 handlers ✅
```

#### Task 3.4: Update API logic to call new RPC methods (1-2h)

**Example**: Update login_logic.go in API layer

**Before** (existing):
```go
// api/internal/logic/publicuser/login_logic.go
func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// ... captcha validation ...

	user, err := l.svcCtx.CoreRpc.GetUserByUsername(l.ctx, &core.UsernameReq{
		Username: req.Username,
	})
	// ... password check ...
	// ... JWT generation ...
}
```

**After** (with new RPC method):
```go
// api/internal/logic/publicuser/login_logic.go
func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// 1. Captcha validation (local)
	if l.svcCtx.Config.ProjectConf.LoginVerify != "captcha" {
		return nil, errorx.NewCodeAbortedError("login.loginTypeForbidden")
	}

	if ok := l.svcCtx.Captcha.Verify(config.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true); !ok {
		return nil, errorx.NewCodeInvalidArgumentError("login.wrongCaptcha")
	}

	// 2. Call new RPC Login method (password check happens in RPC)
	loginResp, err := l.svcCtx.CoreRpc.Login(l.ctx, &core.LoginReq{
		Username:  req.Username,
		Password:  req.Password,
		CaptchaId: req.CaptchaId,
		Captcha:   req.Captcha,
	})
	if err != nil {
		return nil, err
	}

	// 3. Generate JWT (local)
	token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, jwt.WithOption("userId", loginResp.Data.UserId))
	if err != nil {
		return nil, err
	}

	// 4. Return response with token
	return &types.LoginResp{
		BaseDataInfo: types.BaseDataInfo{Code: 0, Msg: loginResp.Msg},
		Data: types.LoginInfo{
			UserId: loginResp.Data.UserId,
			Token:  token,
			Expire: uint64(time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).Unix()),
		},
	}, nil
}
```

**Update All API Logic Files**:
```bash
# Files to update (call new RPC methods):
api/internal/logic/publicuser/
├── login_logic.go                    ✅ Call Login RPC
├── login_by_email_logic.go           ✅ Call LoginByEmail RPC
├── login_by_sms_logic.go             ✅ Call LoginBySms RPC
├── register_logic.go                 ✅ Call Register RPC
├── register_by_email_logic.go        ✅ Call RegisterByEmail RPC
├── register_by_sms_logic.go          ✅ Call RegisterBySms RPC
├── reset_password_by_email_logic.go  ✅ Call ResetPasswordByEmail RPC
└── reset_password_by_sms_logic.go    ✅ Call ResetPasswordBySms RPC

api/internal/logic/user/
├── change_password_logic.go          ✅ Call ChangePassword RPC
├── get_user_info_logic.go            ✅ Call GetUserInfo RPC
├── get_user_perm_code_logic.go       ✅ Call GetUserPermCode RPC
├── get_user_profile_logic.go         ✅ Call GetUserProfile RPC
├── update_user_profile_logic.go      ✅ Call UpdateUserProfile RPC
├── logout_logic.go                   ✅ Call Logout RPC
├── refresh_token_logic.go            ✅ Call RefreshToken RPC
└── access_token_logic.go             ✅ Call AccessToken RPC
```

#### Task 3.5: Run tests (0.5h)

**Unit Tests**:
```bash
# Run API layer tests
go test ./api/internal/logic/user/... -v
go test ./api/internal/logic/publicuser/... -v

# Run RPC layer tests
go test ./rpc/internal/logic/user/... -v
```

**Integration Tests** (if available):
```bash
# Start RPC service
go run rpc/core.go -f rpc/etc/core.yaml &
RPC_PID=$!

# Start API service
go run api/core.go -f api/etc/core.yaml &
API_PID=$!

# Wait for services to start
sleep 5

# Run integration tests
go test ./api/... -tags=integration -v

# Cleanup
kill $RPC_PID $API_PID
```

**Deliverables**:
- ✅ Generated user.api from Proto
- ✅ Diff report (user_api_diff.txt)
- ✅ API handlers regenerated (22 files)
- ✅ API logic updated to call new RPC methods
- ✅ Compilation successful (no errors)
- ✅ All tests passing

---

### Phase 4: E2E Validation (Day 3-4, 4-6 hours)

#### Task 4.1: Start services (0.5h)

**Prerequisites**:
- PostgreSQL/MySQL running
- Redis running

**Start Services**:
```bash
# Terminal 1: Start RPC service
go run rpc/core.go -f rpc/etc/core.yaml

# Terminal 2: Start API service
go run api/core.go -f api/etc/core.yaml

# Check logs for errors
tail -f api/logs/core.log
tail -f rpc/logs/core.log
```

**Verify Services**:
```bash
# Check RPC service
grpcurl -plaintext localhost:9101 list

# Check API service
curl http://localhost:9100/health
```

#### Task 4.2: Test public endpoints (1h)

**Test Script** (`test_public_endpoints.sh`):
```bash
#!/bin/bash

BASE_URL="http://localhost:9100"

# Test 1: Get Captcha
echo "Test 1: Get Captcha"
CAPTCHA_RESP=$(curl -s -X POST "$BASE_URL/user/captcha")
echo $CAPTCHA_RESP | jq .

# Extract captcha ID (if your API provides one)
# CAPTCHA_ID=$(echo $CAPTCHA_RESP | jq -r '.data.captchaId')

# Test 2: Register
echo "Test 2: Register"
curl -s -X POST "$BASE_URL/user/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser_'$(date +%s)'",
    "password": "Test123456",
    "email": "test'$(date +%s)'@example.com",
    "captchaId": "test",
    "captcha": "test"
  }' | jq .

# Test 3: Login
echo "Test 3: Login"
LOGIN_RESP=$(curl -s -X POST "$BASE_URL/user/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "simple-admin",
    "captchaId": "test",
    "captcha": "test"
  }')
echo $LOGIN_RESP | jq .

# Extract token
TOKEN=$(echo $LOGIN_RESP | jq -r '.data.token')
echo "Token: $TOKEN"

# Test 4: Login by Email
echo "Test 4: Login by Email"
curl -s -X POST "$BASE_URL/user/login_by_email" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "captcha": "test"
  }' | jq .

echo "✅ Public endpoints test complete"
```

**Run Tests**:
```bash
chmod +x test_public_endpoints.sh
./test_public_endpoints.sh
```

**Expected Results**:
- ✅ All requests return 200 OK
- ✅ Login returns valid JWT token
- ✅ Register creates new user

#### Task 4.3: Test protected endpoints (1h)

**Test Script** (`test_protected_endpoints.sh`):
```bash
#!/bin/bash

BASE_URL="http://localhost:9100"

# Step 1: Login to get token
echo "Step 1: Login"
LOGIN_RESP=$(curl -s -X POST "$BASE_URL/user/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "simple-admin",
    "captchaId": "test",
    "captcha": "test"
  }')

TOKEN=$(echo $LOGIN_RESP | jq -r '.data.token')
echo "Token: $TOKEN"

# Test 1: GetUserInfo
echo "Test 1: Get User Info"
curl -s -X GET "$BASE_URL/user/info" \
  -H "Authorization: Bearer $TOKEN" | jq .

# Test 2: GetUserPermCode
echo "Test 2: Get Permission Codes"
curl -s -X GET "$BASE_URL/user/perm" \
  -H "Authorization: Bearer $TOKEN" | jq .

# Test 3: GetUserProfile
echo "Test 3: Get User Profile"
curl -s -X GET "$BASE_URL/user/profile" \
  -H "Authorization: Bearer $TOKEN" | jq .

# Test 4: UpdateUserProfile
echo "Test 4: Update User Profile"
curl -s -X POST "$BASE_URL/user/profile" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "Updated Nickname",
    "locale": "zh-TW"
  }' | jq .

# Test 5: GetUserList
echo "Test 5: Get User List"
curl -s -X POST "$BASE_URL/user/list" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "page": 1,
    "pageSize": 10
  }' | jq .

# Test 6: CreateUser
echo "Test 6: Create User"
curl -s -X POST "$BASE_URL/user/create" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser_'$(date +%s)'",
    "password": "Pass123456",
    "email": "newuser'$(date +%s)'@example.com",
    "nickname": "New User"
  }' | jq .

# Test 7: ChangePassword
echo "Test 7: Change Password"
curl -s -X POST "$BASE_URL/user/change_password" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "oldPassword": "simple-admin",
    "newPassword": "NewPass123456"
  }' | jq .

# Test 8: RefreshToken
echo "Test 8: Refresh Token"
curl -s -X GET "$BASE_URL/user/refresh_token" \
  -H "Authorization: Bearer $TOKEN" | jq .

# Test 9: Logout
echo "Test 9: Logout"
curl -s -X GET "$BASE_URL/user/logout" \
  -H "Authorization: Bearer $TOKEN" | jq .

echo "✅ Protected endpoints test complete"
```

**Run Tests**:
```bash
chmod +x test_protected_endpoints.sh
./test_protected_endpoints.sh
```

**Expected Results**:
- ✅ All requests with valid token return 200 OK
- ✅ Requests without token return 401 Unauthorized
- ✅ Requests with invalid token return 401 Unauthorized

#### Task 4.4: Test JWT and permission control (1h)

**Test Script** (`test_auth.sh`):
```bash
#!/bin/bash

BASE_URL="http://localhost:9100"

# Test 1: Access protected endpoint without JWT
echo "Test 1: No JWT (should fail)"
curl -s -X POST "$BASE_URL/user/list" \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}' | jq .
# Expected: 401 Unauthorized

# Test 2: Access protected endpoint with invalid JWT
echo "Test 2: Invalid JWT (should fail)"
curl -s -X POST "$BASE_URL/user/list" \
  -H "Authorization: Bearer invalid_token_here" \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}' | jq .
# Expected: 401 Unauthorized

# Test 3: Login and access with valid JWT
echo "Test 3: Valid JWT (should succeed)"
LOGIN_RESP=$(curl -s -X POST "$BASE_URL/user/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "simple-admin",
    "captchaId": "test",
    "captcha": "test"
  }')
TOKEN=$(echo $LOGIN_RESP | jq -r '.data.token')

curl -s -X POST "$BASE_URL/user/list" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}' | jq .
# Expected: 200 OK with user list

# Test 4: Create limited user and test permission denial
echo "Test 4: Permission denial"
# (This requires setting up a user with limited permissions)
# Create user with minimal role
# Try to call admin endpoint
# Expected: 403 Forbidden

echo "✅ Auth test complete"
```

#### Task 4.5: Performance benchmarking (1h)

**Benchmark Tool**: Apache Bench (ab) or wrk

**Benchmark Script** (`benchmark.sh`):
```bash
#!/bin/bash

BASE_URL="http://localhost:9100"

# Get token first
TOKEN=$(curl -s -X POST "$BASE_URL/user/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"simple-admin","captchaId":"test","captcha":"test"}' \
  | jq -r '.data.token')

# Benchmark 1: Login endpoint (public)
echo "Benchmark 1: Login endpoint"
ab -n 1000 -c 10 -p login.json -T "application/json" "$BASE_URL/user/login"

# Benchmark 2: GetUserList (protected)
echo "Benchmark 2: GetUserList endpoint"
ab -n 1000 -c 10 \
  -H "Authorization: Bearer $TOKEN" \
  -p userlist.json -T "application/json" \
  "$BASE_URL/user/list"

# Benchmark 3: GetUserInfo (protected, lightweight)
echo "Benchmark 3: GetUserInfo endpoint"
ab -n 1000 -c 10 \
  -H "Authorization: Bearer $TOKEN" \
  "$BASE_URL/user/info"

# Expected metrics:
# - Login: < 150ms (P95)
# - GetUserList: < 200ms (P95)
# - GetUserInfo: < 100ms (P95)
```

**Prepare JSON payloads**:
```bash
# login.json
echo '{"username":"admin","password":"simple-admin","captchaId":"test","captcha":"test"}' > login.json

# userlist.json
echo '{"page":1,"pageSize":10}' > userlist.json
```

**Run Benchmarks**:
```bash
chmod +x benchmark.sh
./benchmark.sh
```

**Analyze Results**:
- Check P50, P95, P99 response times
- Verify no errors (all 2xx responses)
- Compare with baseline (before migration)

#### Task 4.6: Manual QA testing (1h)

**QA Checklist**:

**Public Endpoints**:
- [ ] Login with valid credentials succeeds
- [ ] Login with wrong password fails
- [ ] Login with non-existent user fails
- [ ] Register with valid data creates user
- [ ] Register with existing email fails
- [ ] Register with weak password fails
- [ ] LoginByEmail sends email and works
- [ ] LoginBySms sends SMS and works
- [ ] ResetPasswordByEmail resets password
- [ ] ResetPasswordBySms resets password

**Protected Endpoints**:
- [ ] GetUserList returns paginated results
- [ ] GetUserById returns user details
- [ ] CreateUser creates new user
- [ ] UpdateUser updates user info
- [ ] DeleteUser soft-deletes user
- [ ] ChangePassword changes password
- [ ] GetUserInfo returns current user info
- [ ] GetUserPermCode returns permission codes
- [ ] GetUserProfile returns profile
- [ ] UpdateUserProfile updates profile
- [ ] Logout invalidates token (if implemented)
- [ ] RefreshToken issues new token
- [ ] AccessToken issues short-lived token

**Security**:
- [ ] Public endpoints accessible without JWT
- [ ] Protected endpoints require JWT
- [ ] Invalid JWT returns 401
- [ ] Expired JWT returns 401
- [ ] Authority middleware checks permissions
- [ ] Unauthorized actions return 403

**Deliverables**:
- ✅ E2E test scripts
- ✅ All tests passing
- ✅ Performance benchmarks (within targets)
- ✅ QA checklist completed
- ✅ Test report documenting results

---

## Performance Considerations

### Bundle Size Impact
- Proto file size increase: ~500 lines (22 methods + messages)
- Generated .api file: ~800 lines (similar to current)
- Binary size increase: < 1MB (negligible)

### Runtime Performance
- RPC call overhead: ~1-2ms per call (network + serialization)
- No change for endpoints that already call RPC
- New endpoints (login, etc.) now call RPC: +1-2ms latency
- Overall impact: < 10ms increase (acceptable per NFR-001)

### Optimization Strategies
1. **Connection Pooling**: Ensure gRPC client uses connection pooling
2. **Caching**: Cache permission codes in Redis (reduce RPC calls)
3. **Batch Operations**: For getUserPermCode, use batch query instead of N+1
4. **Proto Optimization**: Use `optional` only when needed to reduce marshaling

---

## Deployment Strategy

### Rollout Plan

**Week 1: Development & Testing**
- Day 1-2: Phase 1 & 2 (Proto definition + RPC implementation)
- Day 3: Phase 3 (Code generation & testing)
- Day 4: Phase 4 (E2E validation)

**Week 2: Staging Deployment**
- Deploy to staging environment
- Run load tests
- Invite team for UAT (User Acceptance Testing)

**Week 3: Production Deployment**
- Deploy RPC service first (backward compatible)
- Deploy API service (calls new RPC methods)
- Monitor for 24 hours
- Rollback if issues detected

### Feature Flag

Not applicable - this is a migration, not a new feature. All changes are backward compatible.

### Monitoring

**Metrics to Track**:
- RPC call latency (by method)
- API endpoint response times
- Error rates (by endpoint)
- JWT token generation time
- Database query performance

**Alerts to Configure**:
- Alert if login endpoint P95 > 200ms
- Alert if error rate > 1%
- Alert if RPC service unavailable

**Dashboards**:
- User API metrics dashboard (Grafana)
- RPC service health dashboard
- Error log dashboard (Kibana)

---

## Rollback Plan

### Scenario 1: Generated .api is incorrect

**Rollback Steps**:
```bash
# Restore backup
cp api/desc/core/user.api.backup api/desc/core/user.api

# Regenerate API code
cd api
goctls api go -api desc/all.api -dir . -style go_zero

# Rebuild
go build -o bin/core-api ./api/core.go
```

### Scenario 2: RPC service fails to start

**Rollback Steps**:
```bash
# Revert Proto file
git checkout HEAD~1 -- rpc/desc/user.proto

# Regenerate RPC code
make gen-rpc

# Rebuild RPC service
go build -o bin/core-rpc ./rpc/core.go
```

### Scenario 3: API service fails after deployment

**Rollback Steps**:
```bash
# Revert all changes
git revert <commit-hash>

# Regenerate all code
make gen-rpc
make gen-api-code

# Rebuild services
make build-linux

# Redeploy
```

### Scenario 4: Complete rollback

**Rollback Steps**:
```bash
# Revert to previous version
git reset --hard <previous-commit>

# Restore all backups
cp api/desc/core/user.api.backup api/desc/core/user.api
cp rpc/desc/user.proto.backup rpc/desc/user.proto

# Regenerate everything
make gen-all

# Rebuild and redeploy
make build-linux
# Deploy to production
```

**Recovery Time Objective (RTO)**: < 5 minutes

---

## Success Metrics

### Technical Metrics
✅ All 22 User endpoints defined in user.proto
✅ `make gen-api-all` succeeds (0 errors)
✅ Test coverage ≥ 70%
✅ All existing tests pass (100%)

### Performance Metrics
✅ Login response time < 150ms (P95)
✅ GetUserList response time < 200ms (P95)
✅ No performance regression > 10ms

### Business Metrics
✅ Zero breaking changes for clients
✅ Developer time to add endpoint reduced by 50%
✅ Code maintainability improved (no duplicate definitions)

---

## Timeline

| Phase | Duration | Days | Agent |
|-------|----------|------|-------|
| Phase 1: Proto Definition | 4-6h | Day 1 | @backend |
| Phase 2: RPC Implementation | 12-16h | Day 1-2 | @backend |
| Phase 3: Code Generation | 4-6h | Day 3 | @backend |
| Phase 4: E2E Validation | 4-6h | Day 3-4 | @qa |
| **Total** | **24-34h** | **3.5-4 days** | |

---

## Team Assignment

| Role | Responsibility | Hours |
|------|----------------|-------|
| @backend | Proto definition, RPC implementation, code generation | 20-28h |
| @qa | Testing, validation, performance benchmarking | 4-6h |
| @doc | Update documentation (CLAUDE.md, migration guide) | 2h |
| @pm | Progress tracking, Notion updates, status reports | 2h |

---

**Plan Status**: 🟡 Draft - Ready for Review
**Next Step**: Review plan with team, get approval, start Phase 1