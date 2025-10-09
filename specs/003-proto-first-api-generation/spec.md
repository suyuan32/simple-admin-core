# Feature Specification: Proto-First API Generation

**Feature Branch**: `feature/proto-first-api-generation`
**Created**: 2025-10-08
**Status**: Draft
**Input**: Current project maintains dual API definitions (Proto for gRPC + .api for REST), causing maintenance overhead and potential inconsistencies. User requested to implement Proto-First approach (inspired by Kratos framework) while retaining Go-Zero's microservice governance capabilities.

## User Scenarios & Testing (mandatory)

### User Story 1 - Single Source of Truth for API Definition (Priority: P1)

As a **backend developer**, I want to **define my API once in Proto files with HTTP annotations**, so that **I don't need to manually maintain separate .api files and can avoid definition inconsistencies**.

**Why this priority**: This is the core value proposition of the feature. Reducing maintenance overhead from dual definitions to single definition directly impacts developer productivity and code quality.

**Independent Test**:
- Developer defines a new CRUD service in Proto with google.api.http annotations
- Run code generation command
- Verify both gRPC service code AND Go-Zero .api file are automatically generated
- Verify generated .api file matches Proto HTTP mappings exactly

**Acceptance Scenarios**:

1. **Given** a Proto file with `google.api.http` annotations defining REST endpoints, **When** developer runs `make gen-proto-api`, **Then** corresponding `.api` files are automatically generated in `api/desc/` directory with correct Go-Zero syntax

2. **Given** a Proto service with multiple HTTP methods (GET/POST/PUT/DELETE), **When** code generation runs, **Then** each method is correctly mapped to Go-Zero API handlers with appropriate HTTP verbs and paths

3. **Given** Proto messages with nested fields and path parameters, **When** .api generation occurs, **Then** request/response types are correctly transformed to Go-Zero format with proper field mappings

4. **Given** an existing .api file and a modified Proto file, **When** regeneration is triggered, **Then** .api file is updated to reflect Proto changes without breaking existing functionality

### User Story 2 - Seamless Integration with Existing Workflow (Priority: P1)

As a **project maintainer**, I want the **Proto-First generation to integrate with existing Makefile commands**, so that **the team can adopt the new workflow without changing their development habits**.

**Why this priority**: Smooth integration ensures zero disruption to current development workflow and allows gradual adoption across the team.

**Independent Test**:
- Modify an existing Proto file (e.g., user.proto)
- Run existing `make gen-rpc` command
- Verify both RPC code and .api files are regenerated
- Verify API service still compiles and runs correctly

**Acceptance Scenarios**:

1. **Given** existing Makefile targets (`gen-rpc`, `gen-api`), **When** new `gen-proto-api` target is added, **Then** it integrates seamlessly with existing targets without breaking backward compatibility

2. **Given** the new Makefile target, **When** developer runs `make gen-all`, **Then** Proto, RPC, and API code are generated in correct order with proper dependencies

3. **Given** multiple Proto files in `rpc/desc/` directory, **When** batch generation runs, **Then** all corresponding .api files are generated in their correct subdirectories

### User Story 3 - Go-Zero Specific Features Support (Priority: P1)

As a **backend developer**, I want **Proto definitions to support Go-Zero specific features like JWT, middleware, and route groups**, so that **generated .api files have the same capabilities as hand-written ones**.

**Why this priority**: Go-Zero's `@server` annotations (jwt, middleware, group) are critical for authentication and authorization. Without these, the generated APIs would be incomplete and non-functional for production use.

**Independent Test**:
- Add custom Proto annotations for JWT and middleware to a service
- Run code generation
- Verify generated .api file contains correct `@server(jwt: Auth, middleware: Authority, group: user)`
- Verify API service compiles and enforces JWT authentication

**Acceptance Scenarios**:

1. **Given** a Proto service with custom option `option (go_zero.jwt) = "Auth"`, **When** .api generation runs, **Then** generated .api file contains `@server(jwt: Auth, group: <service_group>)`

2. **Given** a Proto service with custom option `option (go_zero.middleware) = "Authority,RateLimit"`, **When** .api generation runs, **Then** generated .api file contains `@server(middleware: Authority,RateLimit)`

3. **Given** a Proto service without JWT/middleware options, **When** .api generation runs, **Then** generated .api file has basic `@server(group: <service_group>)` only

4. **Given** different Proto methods requiring different middleware, **When** .api generation runs, **Then** methods are grouped into separate service blocks with appropriate `@server` annotations

5. **Given** a Proto file with service-level metadata (title, author, version), **When** .api generation runs, **Then** generated .api file includes proper `info()` section

### User Story 4 - Validation and Error Handling (Priority: P2)

As a **developer**, I want **clear error messages when Proto definitions are invalid or incomplete**, so that **I can quickly fix issues without debugging the generation process**.

**Why this priority**: Good developer experience requires clear feedback. This is P2 because the feature must work correctly first (P1), then provide good UX.

**Independent Test**:
- Create a Proto file with invalid google.api.http syntax
- Run code generation
- Verify clear error message indicating the problem and line number
- Fix the issue and verify successful generation

**Acceptance Scenarios**:

1. **Given** a Proto file missing `google.api.http` annotations, **When** generation runs, **Then** a warning is displayed indicating which methods lack HTTP mappings

2. **Given** a Proto file with invalid HTTP path syntax (e.g., missing braces for path params), **When** generation runs, **Then** a clear error message shows the exact location and expected format

3. **Given** conflicting HTTP paths in different Proto methods, **When** validation runs, **Then** an error is raised identifying the conflict before code generation

### Edge Cases

- What happens when a Proto message has deeply nested fields (3+ levels)?
  → Parser should flatten nested structures or use dot notation for Go-Zero .api format

- How does the system handle Proto services with no HTTP annotations?
  → Skip HTTP generation for those methods, only generate gRPC code, and log a warning

- What if google.api.http defines both `get` and `additional_bindings`?
  → Generate multiple .api handler entries for the same method with different routes

- How to handle Proto3 `optional` fields in Go-Zero .api generation?
  → Map to Go-Zero optional fields or pointer types as appropriate

- What happens when Proto import paths change?
  → Plugin should resolve imports correctly and update .api file paths accordingly

## Requirements (mandatory)

### Functional Requirements

- **FR-001**: System MUST parse Proto files and extract `google.api.http` annotations including GET, POST, PUT, DELETE, PATCH methods
- **FR-002**: System MUST generate valid Go-Zero `.api` format files that compile successfully with `goctl api go` command
- **FR-003**: System MUST support HTTP path parameters mapping (e.g., `/users/{id}` → `/users/:id`)
- **FR-004**: System MUST support HTTP body mapping (`body: "*"`, `body: "field_name"`)
- **FR-005**: System MUST support `additional_bindings` for multiple HTTP routes per RPC method
- **FR-006**: System MUST preserve Proto message field names and types when generating .api request/response types
- **FR-007**: System MUST generate appropriate `@server` group annotations based on Proto service names
- **FR-008**: System MUST generate `@handler` names from Proto method names (CamelCase → snake_case)
- **FR-009**: System MUST validate HTTP path parameters exist in request message fields
- **FR-010**: System MUST support batch generation for multiple Proto files in a directory
- **FR-011**: System MUST integrate with existing Makefile targets without breaking current workflow
- **FR-012**: Generated .api files MUST maintain backward compatibility with existing API service code
- **FR-013**: System MUST support Go-Zero `@server` annotations including `jwt`, `group`, and `middleware` configurations
- **FR-014**: System MUST allow specifying JWT authentication requirements via Proto annotations or configuration
- **FR-015**: System MUST allow specifying middleware (e.g., Authority, RateLimit) via Proto annotations or configuration
- **FR-016**: System MUST support route prefix/group configuration at service level
- **FR-017**: System MUST preserve or generate appropriate `info` section in .api files (title, desc, author, version)

### Non-Functional Requirements

- **NFR-001**: Code generation MUST complete within 5 seconds for a typical Proto file (< 50 methods)
- **NFR-002**: Plugin MUST provide clear error messages with file name and line number on parsing failures
- **NFR-003**: Generated code MUST follow Go-Zero naming conventions and code style
- **NFR-004**: Plugin MUST be implemented as a standard `protoc` plugin for portability

### Key Entities

- **protoc-gen-go-zero-api Plugin**:
  - Location: `tools/protoc-gen-go-zero-api/`
  - Format: Go executable implementing protoc plugin interface
  - Dependencies: `google.golang.org/protobuf/compiler/protogen`

- **Go-Zero Custom Proto Options** (NEW):
  - Location: `rpc/desc/go_zero/options.proto`
  - Purpose: Define Go-Zero specific annotations for services and methods
  - Example options:
    ```protobuf
    extend google.protobuf.ServiceOptions {
      string jwt = 50001;           // JWT config name (e.g., "Auth")
      string middleware = 50002;    // Middleware list (e.g., "Authority,RateLimit")
      string group = 50003;         // Route group name
      string prefix = 50004;        // Route prefix (e.g., "/api/v1")
    }

    extend google.protobuf.MethodOptions {
      bool public = 50011;          // Mark method as public (no JWT)
      string middleware = 50012;    // Method-specific middleware
    }

    extend google.protobuf.FileOptions {
      ApiInfo api_info = 50021;     // API metadata (title, author, version)
    }

    message ApiInfo {
      string title = 1;
      string desc = 2;
      string author = 3;
      string email = 4;
      string version = 5;
    }
    ```

- **HTTP Annotation Parser**:
  - Purpose: Extract `google.api.http` options from Proto MethodOptions
  - Output: Structured HTTP rule (method, path, body)

- **.api File Generator**:
  - Purpose: Transform Proto service definitions to Go-Zero .api syntax
  - Output: `.api` files in `api/desc/` directory structure

- **Makefile Targets**:
  - `gen-proto-api`: Generate .api from Proto files
  - `gen-all`: Combined target for Ent + RPC + API generation

- **Integration Test Suite**:
  - Location: `tools/protoc-gen-go-zero-api/test/`
  - Purpose: Validate plugin output against expected .api files

## Success Criteria (mandatory)

### Measurable Outcomes

- **SC-001**: Development time for adding new API endpoint reduced by 50% (from 10 min to 5 min)
- **SC-002**: Code duplication between Proto and .api definitions reduced to 0%
- **SC-003**: API definition inconsistency bugs reduced by 100% (currently 2-3 per month)
- **SC-004**: All 15+ existing API modules successfully migrated to Proto-First within 2 weeks
- **SC-005**: Plugin code generation completes in < 3 seconds for all current Proto files

### User Satisfaction Metrics

- **SC-006**: Team survey shows 80%+ developers prefer Proto-First over dual-definition approach
- **SC-007**: Zero regression bugs reported in existing API functionality after migration

### Business Metrics

- **SC-008**: API maintenance overhead reduced by 40% (measured by time spent on API-related bugs/changes)
- **SC-009**: New developer onboarding time reduced by 20% due to simpler API definition process

## Out of Scope

❌ **Automatic migration of existing .api files to Proto** - Manual migration or tool-assisted migration will be handled separately
❌ **Support for non-standard Go-Zero .api extensions** - Only standard Go-Zero syntax is supported
❌ **Bidirectional sync (Proto ← .api)** - This is a one-way generation from Proto to .api
❌ **Custom code generation templates** - Uses fixed template based on Go-Zero conventions
❌ **Support for Protocol Buffers v2** - Only Proto3 is supported
❌ **Automatic API versioning** - API versioning strategy is out of scope

## Dependencies

### Internal Dependencies

- **Existing Proto definitions** in `rpc/desc/**/*.proto`
- **Go-Zero API structure** in `api/desc/` and `api/internal/`
- **Makefile** targets for code generation
- **Google API annotations** Proto files in `third_party/google/api/`

### External Dependencies

- **protoc** (Protocol Buffer Compiler) v3.19.0+
- **google.golang.org/protobuf** v1.31.0+
- **google.golang.org/genproto** (for google.api.http annotations)
- **github.com/zeromicro/go-zero/tools/goctl** v1.6.0+ (for .api compilation)

### Team Dependencies

- **Backend Developer** for plugin implementation (40 hours estimated)
- **DevOps Engineer** for CI/CD integration (4 hours estimated)
- **Tech Lead** for code review and architecture validation (8 hours estimated)

## Risks & Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Generated .api files don't compile with goctl | High | Medium | Implement comprehensive test suite with golden file testing; validate output against known-good .api files |
| Complex Proto types (oneof, map, Any) fail to convert to .api format | High | Medium | Document unsupported types; provide clear error messages; implement fallback to manual .api definition |
| Team resistance to changing workflow | Medium | Low | Provide comprehensive documentation; offer training sessions; allow gradual migration |
| Plugin performance degrades with large Proto files (100+ methods) | Medium | Low | Implement streaming parsing; optimize AST traversal; add performance benchmarks |
| Breaking changes in future Go-Zero .api format | Medium | Low | Version lock Go-Zero; monitor Go-Zero releases; implement adapter pattern for format changes |
| Existing .api files have custom modifications that can't be represented in Proto | High | Medium | Provide migration guide; support custom annotations or fallback to manual .api |
| google.api.http annotations are insufficient for all Go-Zero features | Medium | High | Document limitations; provide extension mechanism for Go-Zero-specific features (middleware, JWT, etc.) |

## Implementation Notes

### High-Level Phases (Technical details go in plan.md)

**Phase 1: Plugin Development (Week 1)**
- Implement protoc plugin framework
- Parse google.api.http annotations
- Generate basic .api file structure

**Phase 2: Integration (Week 2)**
- Integrate with Makefile
- Migrate 1-2 pilot modules (User, Role)
- Validate generated code compiles and runs

**Phase 3: Full Migration (Week 3)**
- Migrate remaining modules
- Update documentation
- Team training

### Migration Strategy

1. **Gradual Migration**: Migrate one module at a time, starting with least critical
2. **Dual Running**: Keep both .api and Proto-generated versions temporarily
3. **Validation**: Compare API behavior before/after migration
4. **Rollback Plan**: Keep original .api files in git history for quick rollback

### Design Decisions

1. **How to handle Go-Zero specific features (JWT, middleware, route groups)?**
   - ✅ **Decision**: Define custom Proto options in `rpc/desc/go_zero/options.proto`
   - **Rationale**:
     - Proto-native solution, type-safe and validated at compile time
     - Follows Google's extension pattern for custom options
     - IDE support for autocomplete and validation
   - **Example**:
     ```protobuf
     service UserService {
       option (go_zero.jwt) = "Auth";
       option (go_zero.middleware) = "Authority";
       option (go_zero.group) = "user";

       rpc GetUser(GetUserReq) returns (GetUserResp) {
         option (google.api.http) = {get: "/user/{id}"};
       }
     }
     ```

2. **Should we support importing .api fragments into generated files?**
   - ✅ **Decision**: Phase 2 enhancement, not in initial scope
   - **Rationale**: Keep initial implementation simple; assess need after migration

3. **How to version the .api generation format if Go-Zero .api syntax changes?**
   - ✅ **Decision**: Plugin version tied to Go-Zero version in go.mod
   - **Rationale**: Ensures compatibility; plugin tests validate against current Go-Zero version

4. **How to handle methods that need different JWT/middleware settings?**
   - ✅ **Decision**: Generate multiple service blocks in .api file
   - **Rationale**: Go-Zero .api format supports multiple service blocks with different `@server` configs
   - **Example output**:
     ```api
     @server(jwt: Auth, group: user)
     service Core {
       @handler getUser
       get /user/:id (GetUserReq) returns (GetUserResp)
     }

     @server(group: user)  // No JWT - public endpoint
     service Core {
       @handler loginUser
       post /user/login (LoginReq) returns (LoginResp)
     }
     ```

### Example: Complete Proto Definition with Go-Zero Features

```protobuf
// rpc/desc/user.proto
syntax = "proto3";

package core.v1;

import "google/api/annotations.proto";
import "go_zero/options.proto";

option go_package = "github.com/chimerakang/simple-admin-core/rpc/types/core";

// File-level API metadata
option (go_zero.api_info) = {
  title: "User Management API"
  desc: "User management endpoints for authentication and CRUD"
  author: "Ryan Su"
  email: "yuansu.china.work@gmail.com"
  version: "v1.0"
};

// User service with JWT and middleware
service User {
  // Service-level Go-Zero options
  option (go_zero.jwt) = "Auth";           // Require JWT authentication
  option (go_zero.middleware) = "Authority"; // Apply Authority middleware
  option (go_zero.group) = "user";         // API route group

  // Protected endpoint: Create user (requires JWT + Authority)
  rpc CreateUser(CreateUserReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }

  // Protected endpoint: Get user by ID
  rpc GetUser(GetUserReq) returns (UserInfo) {
    option (google.api.http) = {
      post: "/user"  // Note: current project uses POST for queries
    };
  }

  // Public endpoint: Login (no JWT required)
  rpc Login(LoginReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login"
      body: "*"
    };
    option (go_zero.public) = true;  // Override service-level JWT requirement
  }

  // Endpoint with additional rate limiting
  rpc UpdateUser(UpdateUserReq) returns (BaseMsgResp) {
    option (google.api.http) = {
      post: "/user/update"
      body: "*"
    };
    option (go_zero.middleware) = "Authority,RateLimit";  // Method-specific middleware
  }
}

message CreateUserReq {
  string username = 1;
  string password = 2;
  string email = 3;
}

message GetUserReq {
  uint64 id = 1;
}

message UserInfo {
  uint64 id = 1;
  string username = 2;
  string email = 3;
}

message LoginReq {
  string username = 1;
  string password = 2;
}

message LoginResp {
  string token = 1;
  string expire = 2;
}
```

**Generated .api file** (`api/desc/core/user.api`):

```api
syntax = "v1"

info(
    title: "User Management API"
    desc: "User management endpoints for authentication and CRUD"
    author: "Ryan Su"
    email: "yuansu.china.work@gmail.com"
    version: "v1.0"
)

import "../base.api"

// Types (generated from Proto messages)
type (
    CreateUserReq {
        Username string `json:"username"`
        Password string `json:"password"`
        Email    string `json:"email"`
    }

    GetUserReq {
        Id uint64 `json:"id"`
    }

    UserInfo {
        Id       uint64 `json:"id"`
        Username string `json:"username"`
        Email    string `json:"email"`
    }

    LoginReq {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    LoginResp {
        Token  string `json:"token"`
        Expire string `json:"expire"`
    }
)

// Protected endpoints (with JWT + Authority middleware)
@server(
    jwt: Auth
    group: user
    middleware: Authority
)
service Core {
    @handler createUser
    post /user/create (CreateUserReq) returns (BaseIDResp)

    @handler getUser
    post /user (GetUserReq) returns (UserInfo)
}

// Endpoint with additional rate limiting
@server(
    jwt: Auth
    group: user
    middleware: Authority,RateLimit
)
service Core {
    @handler updateUser
    post /user/update (UpdateUserReq) returns (BaseMsgResp)
}

// Public endpoint (no JWT)
@server(
    group: user
)
service Core {
    @handler loginUser
    post /user/login (LoginReq) returns (LoginResp)
}
```

## Review Checklist

- [x] All user stories have clear acceptance criteria (4 stories: P1, P1, P1, P2)
- [x] All functional requirements are testable (FR-001 to FR-017)
- [x] Success criteria are measurable (9 specific metrics)
- [x] Dependencies identified and available (Proto, Go-Zero, protoc)
- [x] Risks have mitigation strategies (7 risks with mitigations)
- [x] Out of scope items clearly defined (6 items)
- [x] Implementation phases outlined (3 weeks)
- [x] **Go-Zero specific features documented** (JWT, middleware, route groups)
- [x] **Custom Proto options defined** (go_zero.proto with service/method/file options)
- [x] **Complete example provided** (Proto → .api transformation)
- [x] **Design decisions documented** (4 key decisions with rationale)
- [ ] Stakeholder approval pending
- [ ] Technical plan (plan.md) to be created after approval
