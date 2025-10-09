# Feature Specification: User Module Proto Completion and Proto-First Migration

**Feature Branch**: `feature/user-module-proto-completion`
**Created**: 2025-10-10
**Status**: Draft
**Input**: Complete User module migration to Proto-First approach, discovered incomplete RPC definitions during Phase 5

**Related**:
- Spec-003: Proto-First API Generation (completed Phase 1-4)
- Phase 5 blocked due to incomplete User module Proto definitions

---

## User Scenarios & Testing

### User Story 1 - Developer Maintains User Module with Proto-First (Priority: P0)

As a backend developer, I want to define all User module endpoints in Proto only, so that I don't need to maintain duplicate API definitions and can avoid definition drift.

**Why this priority**: This is blocking Phase 5 of Proto-First feature and affects developer productivity immediately.

**Independent Test**: Can be tested by modifying a User endpoint (e.g., add a field to Login request) and verifying only Proto needs to be updated.

**Acceptance Scenarios**:
1. **Given** User module has complete Proto definitions with HTTP annotations
   **When** Developer runs `make gen-api-all`
   **Then** All 22 User API endpoints are generated correctly in .api file

2. **Given** Developer modifies a Login request field in user.proto
   **When** Developer regenerates code with `make gen-api-all`
   **Then** API handlers reflect the change without manual .api editing

3. **Given** All User endpoints have RPC implementations
   **When** API service calls any User endpoint
   **Then** Request is properly routed through RPC layer (except local auth logic)

### User Story 2 - Authentication Flow Works Correctly (Priority: P0)

As an end user, I want to login/register/reset password through API endpoints, and all functionality should work exactly as before the migration.

**Acceptance Scenarios**:
1. **Given** User has valid credentials
   **When** User calls POST /user/login with username, password, captcha
   **Then** User receives JWT token and user info

2. **Given** New user wants to register
   **When** User calls POST /user/register with valid data
   **Then** User account is created and receives success message

3. **Given** User forgot password
   **When** User calls POST /user/reset_password_by_email with valid captcha
   **Then** Password is reset and user receives confirmation

### User Story 3 - JWT and Permission Control Works (Priority: P1)

As a system administrator, I want protected endpoints to require authentication and check permissions, while public endpoints remain accessible.

**Acceptance Scenarios**:
1. **Given** User is not logged in
   **When** User calls protected endpoint like /user/list
   **Then** System returns 401 Unauthorized

2. **Given** User is logged in but has no permission
   **When** User calls /user/create
   **Then** Authority middleware blocks request with 403 Forbidden

3. **Given** User is logged in with valid permissions
   **When** User calls /user/create
   **Then** Request succeeds and user is created

### Edge Cases

- What happens when Proto has HTTP annotation but no RPC implementation?
  - **Answer**: Compilation should fail early with clear error

- How does system handle endpoints that only need local logic (e.g., logout)?
  - **Answer**: Define RPC method but implementation can be minimal/stub

- What if generated .api format differs from existing?
  - **Answer**: Use diff tool to validate, manual reconciliation if needed

---

## Requirements

### Functional Requirements

**Proto Definitions**:
- **FR-001**: User.proto MUST define all 22 API endpoints with complete HTTP annotations
- **FR-002**: Each RPC method MUST have corresponding Request/Response message types
- **FR-003**: Public endpoints (login, register, etc.) MUST be marked with `(go_zero.public) = true`
- **FR-004**: Protected endpoints MUST have JWT and middleware configuration at service level
- **FR-005**: HTTP routes MUST match existing API routes exactly (backward compatible)

**RPC Implementation**:
- **FR-006**: All new RPC methods MUST be implemented in `rpc/internal/logic/user/`
- **FR-007**: RPC logic MUST handle business logic (database operations, validation)
- **FR-008**: Authentication logic (captcha, JWT generation) MAY remain in API layer
- **FR-009**: RPC methods MUST return proper error codes and i18n messages

**Code Generation**:
- **FR-010**: `make gen-api-all` MUST generate complete user.api file from user.proto
- **FR-011**: Generated .api file MUST be valid and compile without errors
- **FR-012**: API handlers MUST be regenerated and call correct RPC methods

**Testing**:
- **FR-013**: All existing User API tests MUST pass after migration
- **FR-014**: New RPC methods MUST have unit tests (>70% coverage)
- **FR-015**: E2E tests MUST verify complete authentication flow

### Non-Functional Requirements

- **NFR-001**: API response time MUST NOT increase by more than 10ms
- **NFR-002**: Migration MUST NOT break existing client integrations
- **NFR-003**: Rollback to previous version MUST be possible within 5 minutes

---

## Key Entities

### Proto Files

**user.proto** (Location: `rpc/desc/user.proto`)
- Current state: 6 basic CRUD RPC methods
- Target state: 22 complete RPC methods with HTTP annotations
- Format: Protocol Buffers v3
- Dependencies: `google/api/annotations.proto`, `go_zero/options.proto`, `base.proto`

**Message Types to Add**:
- LoginReq, LoginResp, LoginInfo
- RegisterReq, RegisterByEmailReq, RegisterBySmsReq
- ChangePasswordReq
- ResetPasswordByEmailReq, ResetPasswordBySmsReq
- ProfileInfo, ProfileResp
- PermCodeResp
- RefreshTokenInfo, RefreshTokenResp
- UserBaseIDInfo, UserBaseIDInfoResp

### RPC Logic Files

**New files to create in `rpc/internal/logic/user/`**:
- `login_logic.go` - Login with username/password
- `login_by_email_logic.go` - Email login
- `login_by_sms_logic.go` - SMS login
- `register_logic.go` - Basic registration (may reuse createUser)
- `register_by_email_logic.go` - Email registration
- `register_by_sms_logic.go` - SMS registration
- `change_password_logic.go` - Change password
- `reset_password_by_email_logic.go` - Reset via email
- `reset_password_by_sms_logic.go` - Reset via SMS
- `get_user_info_logic.go` - Get basic user info
- `get_user_perm_code_logic.go` - Get permission codes
- `get_user_profile_logic.go` - Get user profile
- `update_user_profile_logic.go` - Update profile
- `logout_logic.go` - Logout (minimal implementation)
- `refresh_token_logic.go` - Refresh JWT token
- `access_token_logic.go` - Get access token

### API Files (Generated)

**user.api** (Location: `api/desc/core/user.api`)
- Current: Manually maintained with 22 endpoints
- Target: Auto-generated from user.proto
- Will be replaced by generated version

---

## Success Criteria

### Technical Success Criteria

- **SC-001**: All 22 User endpoints defined in user.proto with HTTP annotations
- **SC-002**: `make gen-api-all` generates user.api successfully (0 errors)
- **SC-003**: Generated user.api compiles with goctls (0 errors)
- **SC-004**: All 16 new RPC methods implemented with unit tests
- **SC-005**: Test coverage for user module RPC logic ≥ 70%
- **SC-006**: All existing User API tests pass (100% pass rate)
- **SC-007**: E2E test covering login → get user info → logout succeeds

### Performance Success Criteria

- **SC-008**: Login endpoint response time < 150ms (P95)
- **SC-009**: Get user list response time < 200ms (P95)
- **SC-010**: No memory leaks in RPC service (24h stress test)

### User Satisfaction Metrics

- **SC-011**: Zero breaking changes for existing API clients
- **SC-012**: Developer feedback: "Easier to maintain" (survey)
- **SC-013**: Time to add new User endpoint reduced by 50%

---

## Architecture Design

### Current Architecture (Before)

```
┌─────────────────────────────────────────────────────────┐
│                    API Layer                            │
│  ┌────────────────┐            ┌──────────────────┐   │
│  │  user.api      │────────────▶  Handlers        │   │
│  │  (22 endpoints)│   manual    │  + Logic        │   │
│  └────────────────┘   maintain  └──────────────────┘   │
│                                          │               │
│                                          ▼               │
│                                  ┌──────────────────┐   │
│                                  │  Auth Logic      │   │
│                                  │  (Local)         │   │
│                                  └──────────────────┘   │
│                                          │               │
└──────────────────────────────────────────┼───────────────┘
                                           │ gRPC
                                           ▼
┌──────────────────────────────────────────────────────────┐
│                    RPC Layer                             │
│  ┌────────────────┐            ┌──────────────────┐    │
│  │  user.proto    │────────────▶  6 RPC Methods   │    │
│  │  (6 methods)   │   generate  │  (CRUD only)    │    │
│  └────────────────┘             └──────────────────┘    │
│                                          │               │
│                                          ▼               │
│                                  ┌──────────────────┐   │
│                                  │  Database (Ent)  │   │
│                                  └──────────────────┘   │
└──────────────────────────────────────────────────────────┘

❌ Problem: API definitions (22) != Proto definitions (6)
❌ Problem: Manual maintenance of user.api required
```

### Target Architecture (After)

```
┌─────────────────────────────────────────────────────────┐
│                    API Layer                            │
│  ┌────────────────┐   generate  ┌──────────────────┐   │
│  │  user.api      │◀────────────│  Handlers        │   │
│  │  (GENERATED)   │             │  + Logic         │   │
│  └────────────────┘             └──────────────────┘   │
│         ▲                                │               │
│         │ Proto-First                    ▼               │
│         │ Generation              ┌──────────────────┐  │
│         │                         │  Auth Logic      │  │
│         │                         │  (Local: Captcha,│  │
│         │                         │   JWT generation)│  │
│         │                         └──────────────────┘  │
│         │                                │               │
└─────────┼────────────────────────────────┼───────────────┘
          │                                │ gRPC
          │                                ▼
┌─────────┼────────────────────────────────────────────────┐
│         │              RPC Layer                         │
│  ┌──────┴─────────┐            ┌──────────────────┐    │
│  │  user.proto    │────────────▶ 22 RPC Methods   │    │
│  │  (22 methods)  │   generate  │  (Complete)     │    │
│  │  + HTTP        │             └──────────────────┘    │
│  │  + Go-Zero     │                      │               │
│  │    Options     │                      ▼               │
│  └────────────────┘              ┌──────────────────┐   │
│         │                        │  Database (Ent)  │   │
│         │                        └──────────────────┘   │
│         │                                                │
│         └────────────────────────────────────────────────┤
│                 Single Source of Truth                   │
└──────────────────────────────────────────────────────────┘

✅ Solution: Proto-First - Define once, generate everywhere
✅ Solution: All 22 endpoints have Proto definitions
```

### Responsibility Split (Important)

| Concern | Layer | Reason |
|---------|-------|--------|
| Captcha validation | API | Infrastructure concern, not business logic |
| JWT generation | API | Authentication mechanism, not business data |
| Password encryption | RPC | Business logic, needs to be consistent |
| User CRUD | RPC | Core business logic |
| Permission check | API (middleware) | Authorization policy enforcement |
| Database operations | RPC | Data persistence layer |

---

## Out of Scope

❌ **Not included in this spec**:
- Migrating other modules (Role, Menu, etc.) - will be separate specs
- OAuth/SSO integration - future enhancement
- Multi-factor authentication (MFA) - future enhancement
- User profile picture upload - separate feature
- Audit log for user operations - separate feature
- Rate limiting for login attempts - can be added later

---

## Dependencies

### Internal Dependencies
- ✅ Proto-First toolchain (Spec-003 Phase 1-4 completed)
- ✅ Existing User RPC methods (6/22 already implemented)
- ✅ Ent schema for User entity
- ⚠️ go_zero/options.proto (needs to be available)

### External Dependencies
- Go 1.25+
- protoc (Protocol Buffer compiler)
- protoc-gen-go-zero-api (custom plugin, already built)
- goctls (go-zero CLI tool)

### Team Dependencies
- @backend: RPC implementation (16 new methods) - 16 hours
- @backend: Proto definition with annotations - 4 hours
- @qa: Testing and validation - 6 hours
- @doc: Update documentation - 2 hours

**Total estimated effort**: 28 hours (3.5 days)

---

## Risks & Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Generated .api differs significantly from existing | High | Medium | Create detailed diff, manual reconciliation plan |
| Breaking existing API clients | Critical | Low | Maintain exact route paths, request/response formats |
| RPC layer performance degradation | Medium | Low | Benchmark before/after, optimize if needed |
| Authentication logic split between layers | Medium | Medium | Clear documentation of responsibility split |
| JWT token format changes | High | Low | Maintain exact token structure, test compatibility |
| Ent schema doesn't support new fields | Medium | Low | Add migrations if needed |
| Time estimate too low | Medium | Medium | Break into smaller milestones, re-estimate after Day 1 |

---

## Implementation Phases

### Phase 1: Proto Definition (Day 1, 4-6 hours)
- Add all 22 RPC method signatures to user.proto
- Define all required message types
- Add HTTP annotations for each method
- Add Go-Zero options (jwt, middleware, group, public)

### Phase 2: RPC Implementation (Day 1-2, 12-16 hours)
- Implement 16 new RPC logic files
- Handle password validation, encryption
- Database operations with Ent
- Error handling and i18n messages

### Phase 3: Code Generation & Testing (Day 3, 4-6 hours)
- Generate user.api from user.proto
- Compare with existing user.api
- Regenerate API handlers
- Run unit tests and integration tests

### Phase 4: E2E Validation (Day 3-4, 4-6 hours)
- Test complete authentication flow
- Verify JWT and permission control
- Performance benchmarking
- Manual QA testing

---

## Review Checklist

- [ ] All 22 user stories have clear acceptance criteria
- [ ] All functional requirements are testable
- [ ] Success criteria are measurable
- [ ] Dependencies identified and available
- [ ] Risks have mitigation strategies
- [ ] Architecture design addresses current problems
- [ ] Responsibility split between API/RPC is clear
- [ ] Backward compatibility is ensured
- [ ] Rollback plan is defined

---

**Spec Status**: 🟡 Draft - Ready for Review
**Next Step**: Create technical plan (plan.md)