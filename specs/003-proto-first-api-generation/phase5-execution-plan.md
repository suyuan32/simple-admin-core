# Phase 5 Execution Plan - Proto-First API Generation
## Pilot Migration and Documentation Phase

**Project**: Simple Admin Core - Proto-First API Generation Feature
**Phase**: Phase 5 - Pilot Migration and Documentation
**Feature Branch**: `feature/proto-first-api-generation`
**Start Date**: 2025-10-09
**Estimated Duration**: 12-16 hours (3-4 days)
**Status**: 🚀 **READY TO START**

---

## Executive Summary

Phase 5 focuses on **real-world validation** and **documentation** of the Proto-First API generation system. This phase will migrate the User module as a pilot project, create comprehensive migration guides, update project documentation, and perform E2E testing with real API endpoints.

### Phase 5 Objectives
1. ✅ **Pilot Migration**: Migrate User module from dual .api to Proto-First approach
2. ✅ **Migration Guide**: Create step-by-step migration documentation
3. ✅ **Project Documentation**: Update CLAUDE.md with Proto-First workflows
4. ✅ **E2E Validation**: Test complete pipeline with real API compilation
5. ✅ **Production Readiness**: Verify system ready for team adoption

---

## Phase 5 Task Breakdown

### Task [PF-020]: Migrate User Module (Pilot)
**Agent**: @backend
**Estimated Time**: 4-6 hours
**Priority**: P1 (Critical)
**Dependencies**: Phase 4 completion

#### Objective
Migrate the User module from dual API definitions (Proto + .api) to Proto-First approach, validating the complete workflow with a production module.

#### Success Criteria
- [ ] Go-Zero options added to `rpc/desc/user.proto`
- [ ] Generated .api file matches existing functionality
- [ ] API service compiles successfully
- [ ] All existing User API tests pass
- [ ] Handler code generation works correctly
- [ ] No regression in API behavior

#### Implementation Steps

##### Step 1: Analyze Existing User Module (0.5h)
**Files to Review**:
- `rpc/desc/user.proto` - Current Proto definition
- `api/desc/core/user.api` - Current API definition
- `api/internal/handler/user/` - Generated handlers
- `api/internal/logic/user/` - Business logic

**Analysis Checklist**:
```bash
# Count endpoints
grep -c "rpc " rpc/desc/user.proto

# List all API endpoints
grep "@handler" api/desc/core/user.api

# Check middleware usage
grep "@server" api/desc/core/user.api

# Identify public endpoints (no JWT)
# These need (go_zero.public) = true
```

**Expected Findings**:
- Total RPC methods: ~15-20
- Public endpoints: Login, Register, GetCaptcha
- Protected endpoints: CreateUser, UpdateUser, DeleteUser, GetUser, etc.
- Middleware: Authority (for most endpoints)
- JWT: Auth (for protected endpoints)

##### Step 2: Add Go-Zero Options to user.proto (1.5h)
**File**: `rpc/desc/user.proto`

**Changes**:
```protobuf
syntax = "proto3";

package core.v1;

import "google/api/annotations.proto";
import "go_zero/options.proto";  // ← Add this import

option go_package = "github.com/chimerakang/simple-admin-core/rpc/types/core";

// ← Add file-level API info
option (go_zero.api_info) = {
  title: "User Management API"
  desc: "User management and authentication services"
  author: "Ryan Su"
  email: "yuansu.china.work@gmail.com"
  version: "v1.0"
};

// ← Add service-level options (default for all methods)
service User {
  option (go_zero.jwt) = "Auth";
  option (go_zero.middleware) = "Authority";
  option (go_zero.group) = "user";

  // Protected endpoint (uses service defaults)
  rpc CreateUser(CreateUserReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }

  // Public endpoint (override JWT)
  rpc Login(LoginReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login"
      body: "*"
    };
    option (go_zero.public) = true;  // ← Override: no JWT required
  }

  // Public endpoint
  rpc Register(RegisterReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/register"
      body: "*"
    };
    option (go_zero.public) = true;
  }

  // Public endpoint
  rpc GetCaptcha(EmptyReq) returns (GetCaptchaResp) {
    option (google.api.http) = {
      post: "/user/captcha"
    };
    option (go_zero.public) = true;
  }

  // ... add options for all other methods
}
```

**Method Categorization**:
| Method | JWT Required | Middleware | Notes |
|--------|-------------|------------|-------|
| Login | No (public) | None | Override with `public = true` |
| Register | No (public) | None | Override with `public = true` |
| GetCaptcha | No (public) | None | Override with `public = true` |
| CreateUser | Yes | Authority | Use service defaults |
| UpdateUser | Yes | Authority | Use service defaults |
| DeleteUser | Yes | Authority | Use service defaults |
| GetUser | Yes | Authority | Use service defaults |
| GetUserList | Yes | Authority | Use service defaults |
| ... | ... | ... | ... |

##### Step 3: Generate .api File (0.5h)
**Commands**:
```bash
# Build plugin (if not already built)
make build-proto-plugin

# Generate .api from Proto
make gen-proto-api

# Output location: api/desc/core/user.api.generated
```

**Verification**:
```bash
# Validate generated .api syntax
cd tools/protoc-gen-go-zero-api/testdata/integration
goctls api validate --api ../../api/desc/core/user.api.generated

# Expected output:
# ✅ api format ok
```

##### Step 4: Compare Generated vs Existing .api (1h)
**Comparison Script**:
```bash
# Create backup of existing .api
cp api/desc/core/user.api api/desc/core/user.api.backup

# Compare files
diff -u api/desc/core/user.api.backup api/desc/core/user.api.generated > user_api_diff.txt

# Review differences
cat user_api_diff.txt
```

**Expected Differences**:
- ✅ Minor formatting differences (whitespace, indentation)
- ✅ Order of type definitions may differ
- ✅ Comments may differ
- ❌ No missing endpoints
- ❌ No missing types
- ❌ No missing @server blocks

**Acceptance Criteria**:
- All endpoints present in both files
- All request/response types present
- JWT and middleware configuration matches
- Handler names match (camelCase)

##### Step 5: Replace .api and Regenerate Code (1h)
**Commands**:
```bash
# Replace .api file
mv api/desc/core/user.api.generated api/desc/core/user.api

# Regenerate API code
make gen-api-code

# Output: api/internal/handler/user/*.go (regenerated)
```

**Verification**:
```bash
# Check for compilation errors
go build ./api/...

# Expected output: no errors

# Check handler count
ls -1 api/internal/handler/user/*.go | wc -l
# Expected: Same count as before (~15-20 handlers)
```

##### Step 6: Run Tests and Validate (0.5h)
**Test Commands**:
```bash
# Run User module tests
go test ./api/internal/logic/user/... -v

# Expected: All tests pass

# Run integration tests (if available)
# go test ./api/... -tags=integration

# Manual smoke test: start services
# make run-api
# make run-rpc
# Test login endpoint with curl
```

**Smoke Test**:
```bash
# Test public endpoint (Login)
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"test123"}'

# Expected: 200 OK with token

# Test protected endpoint (GetUser)
curl -X POST http://localhost:9100/user \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"id":1}'

# Expected: 200 OK with user data
```

#### Deliverables
1. ✅ Updated `rpc/desc/user.proto` with Go-Zero options
2. ✅ Generated `api/desc/core/user.api` from Proto
3. ✅ Backup file `api/desc/core/user.api.backup`
4. ✅ Diff report `user_api_diff.txt`
5. ✅ Compilation verification (no errors)
6. ✅ Test results (all passing)
7. ✅ Migration notes documenting any issues

#### Risk Mitigation
**Risk**: Generated .api differs significantly from existing
- **Mitigation**: Detailed diff review, manual reconciliation
- **Rollback**: Use .backup file

**Risk**: API service fails to compile
- **Mitigation**: Incremental testing, validation at each step
- **Rollback**: Restore .backup, regenerate code

**Risk**: Handler behavior changes
- **Mitigation**: Run all existing tests, manual smoke testing
- **Rollback**: Full revert of changes

---

### Task [PF-021]: Create Migration Guide
**Agent**: @doc
**Estimated Time**: 3-4 hours
**Priority**: P1 (Critical)
**Dependencies**: [PF-020] completion

#### Objective
Create comprehensive migration guide to help developers migrate existing modules from dual API definitions to Proto-First approach.

#### Success Criteria
- [ ] Complete migration guide document created
- [ ] Step-by-step instructions for all common scenarios
- [ ] Troubleshooting section with common issues
- [ ] Example code snippets for all Go-Zero options
- [ ] Rollback procedures documented

#### Document Structure
**File**: `specs/003-proto-first-api-generation/migration-guide.md`

**Sections**:
1. **Overview** (0.5h)
   - Benefits of Proto-First approach
   - When to migrate
   - Prerequisites

2. **Migration Steps** (1.5h)
   - Step 1: Add Go-Zero imports
   - Step 2: Add file-level API info
   - Step 3: Add service-level options
   - Step 4: Add method-level overrides
   - Step 5: Generate and validate
   - Step 6: Replace and test

3. **Go-Zero Options Reference** (0.5h)
   - File-level: `api_info`
   - Service-level: `jwt`, `middleware`, `group`, `prefix`
   - Method-level: `public`, `middleware` (override)
   - Examples for each option

4. **Troubleshooting** (0.5h)
   - Generated .api doesn't compile
   - Missing endpoints
   - Wrong JWT configuration
   - Middleware not applied
   - Handler name mismatch

5. **Rollback Procedures** (0.5h)
   - How to revert changes
   - Restoring backup .api files
   - Regenerating code

6. **Appendix** (0.5h)
   - Comparison table: Proto-First vs Dual API
   - Common patterns
   - FAQs

#### Deliverables
1. ✅ Complete migration guide (markdown)
2. ✅ Code examples for all scenarios
3. ✅ Troubleshooting flowchart
4. ✅ FAQ section

---

### Task [PF-022]: Update Project Documentation
**Agent**: @doc
**Estimated Time**: 2-3 hours
**Priority**: P2 (High)
**Dependencies**: [PF-020], [PF-021] completion

#### Objective
Update CLAUDE.md and other project documentation to include Proto-First API generation workflows.

#### Success Criteria
- [ ] CLAUDE.md updated with Proto-First section
- [ ] New make commands documented
- [ ] Development workflow updated
- [ ] Examples added to documentation

#### Updates Required

##### CLAUDE.md Updates (1.5h)
**Section to Add**: "Proto-First API Generation (New)"

**Content**:
```markdown
## Proto-First API Generation (New)

This project now supports **Proto-First API generation**, eliminating the need to maintain separate `.api` files. Define your API once in Proto, and the system automatically generates Go-Zero compatible .api files.

### Quick Start

#### 1. Define API in Proto with Go-Zero Options
```protobuf
service User {
  option (go_zero.jwt) = "Auth";
  option (go_zero.group) = "user";

  rpc CreateUser(CreateUserReq) returns (BaseIDResp) {
    option (google.api.http) = {
      post: "/user/create"
      body: "*"
    };
  }

  rpc Login(LoginReq) returns (LoginResp) {
    option (google.api.http) = {
      post: "/user/login"
      body: "*"
    };
    option (go_zero.public) = true;  // Public endpoint
  }
}
```

#### 2. Generate .api and Code
```bash
make gen-api-all  # Proto → .api → Go code
```

### Available Commands

#### Generation Commands
```bash
# Generate .api from Proto
make gen-proto-api

# Complete API generation pipeline
make gen-api-all

# Validate generated .api files
make validate-api

# Build plugin (usually not needed manually)
make build-proto-plugin
```

### Go-Zero Custom Options

See `rpc/desc/go_zero/options.proto` for all available options:

#### File-Level Options
```protobuf
option (go_zero.api_info) = {
  title: "API Title"
  desc: "API Description"
  author: "Your Name"
  email: "your.email@example.com"
  version: "v1.0"
};
```

#### Service-Level Options
- `(go_zero.jwt)` - JWT config name (e.g., "Auth")
- `(go_zero.middleware)` - Comma-separated middleware list
- `(go_zero.group)` - Route group name
- `(go_zero.prefix)` - Route prefix (optional)

#### Method-Level Options
- `(go_zero.public)` - Mark method as public (no JWT)
- `(go_zero.middleware)` - Override service-level middleware

### Migration

See [Migration Guide](../specs/003-proto-first-api-generation/migration-guide.md) for migrating existing modules.

### Development Workflow

#### Traditional Workflow (OLD)
1. Edit `rpc/desc/user.proto`
2. Run `make gen-rpc`
3. **Manually edit** `api/desc/core/user.api` (keep in sync!)
4. Run `make gen-api`
5. Implement logic

#### Proto-First Workflow (NEW)
1. Edit `rpc/desc/user.proto` (add Go-Zero options)
2. Run `make gen-api-all` (generates both RPC and API code)
3. Implement logic

**Result**: 50% less manual work, zero API definition drift!
```

##### Makefile Documentation Update (0.5h)
**Section to Add**: Command reference

```markdown
### Code Generation Commands

#### Proto-First API Generation (New)
- `make gen-api-all` - Complete pipeline: Proto → .api → Go code
- `make gen-proto-api` - Generate .api files from Proto
- `make validate-api` - Validate generated .api files
- `make build-proto-plugin` - Build protoc-gen-go-zero-api plugin

#### Traditional Generation
- `make gen-rpc` - Generate RPC code from Proto
- `make gen-api` - Generate API code from .api files
- `make gen-ent` - Generate Ent code from schema
```

##### README.md Updates (Optional, 0.5h)
Add quick start section for new developers.

#### Deliverables
1. ✅ Updated CLAUDE.md with Proto-First section
2. ✅ Command reference documentation
3. ✅ Workflow comparison table
4. ✅ Quick start examples

---

### Task [PF-023]: E2E Testing and Validation
**Agent**: @qa
**Estimated Time**: 3-4 hours
**Priority**: P2 (High)
**Dependencies**: [PF-020] completion

#### Objective
Perform comprehensive E2E testing with real API endpoints, validating the complete Proto → .api → Go code → Running API pipeline.

#### Success Criteria
- [ ] Generated API handlers compile successfully
- [ ] API service starts without errors
- [ ] All User endpoints respond correctly
- [ ] JWT authentication works
- [ ] Middleware is applied correctly
- [ ] Public endpoints work without JWT
- [ ] Protected endpoints require JWT
- [ ] Response format matches expectations

#### Test Plan

##### Test 1: Compilation Validation (0.5h)
**Objective**: Verify generated code compiles without errors

**Steps**:
```bash
# Clean build
rm -rf api/internal/handler/user/*.go
rm -rf api/internal/logic/user/*.go

# Regenerate from Proto
make gen-api-all

# Compile API service
go build -o bin/core-api ./api/core.go

# Expected: No compilation errors
echo $?  # Should be 0
```

**Success Criteria**:
- ✅ No compilation errors
- ✅ Binary created successfully
- ✅ All handlers generated

##### Test 2: Service Startup (0.5h)
**Objective**: Verify API service starts correctly

**Steps**:
```bash
# Start RPC service (dependency)
go run rpc/core.go -f rpc/etc/core.yaml &
RPC_PID=$!

# Wait for RPC to start
sleep 3

# Start API service
go run api/core.go -f api/etc/core.yaml &
API_PID=$!

# Wait for API to start
sleep 3

# Check if processes are running
ps -p $RPC_PID && echo "RPC running" || echo "RPC failed"
ps -p $API_PID && echo "API running" || echo "API failed"

# Check logs for errors
tail -n 20 api/logs/core.log
```

**Success Criteria**:
- ✅ RPC service starts successfully
- ✅ API service starts successfully
- ✅ No errors in logs
- ✅ Listening on configured port

##### Test 3: Public Endpoint Testing (0.5h)
**Objective**: Verify public endpoints work without JWT

**Test Cases**:
```bash
# Test 1: GetCaptcha (no auth required)
curl -X POST http://localhost:9100/user/captcha \
  -H "Content-Type: application/json"

# Expected: 200 OK with captcha data

# Test 2: Login (no auth required)
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "simple-admin",
    "captchaId": "test",
    "captcha": "test"
  }'

# Expected: 200 OK with JWT token
TOKEN=$(curl -s -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"simple-admin"}' | jq -r '.data.token')

echo "Token: $TOKEN"

# Test 3: Register (no auth required)
curl -X POST http://localhost:9100/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "Test123456",
    "email": "test@example.com"
  }'

# Expected: 200 OK with user ID
```

**Success Criteria**:
- ✅ All public endpoints accessible without JWT
- ✅ Response format matches expectations
- ✅ HTTP status codes correct (200, 400, etc.)

##### Test 4: Protected Endpoint Testing (0.5h)
**Objective**: Verify protected endpoints require JWT

**Test Cases**:
```bash
# Test 1: GetUserList without JWT (should fail)
curl -X POST http://localhost:9100/user/list \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'

# Expected: 401 Unauthorized

# Test 2: GetUserList with JWT (should succeed)
curl -X POST http://localhost:9100/user/list \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"page": 1, "pageSize": 10}'

# Expected: 200 OK with user list

# Test 3: CreateUser with JWT
curl -X POST http://localhost:9100/user/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "username": "newuser",
    "password": "Pass123456",
    "email": "new@example.com"
  }'

# Expected: 200 OK with user ID

# Test 4: GetUser with JWT
curl -X POST http://localhost:9100/user \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"id": 1}'

# Expected: 200 OK with user details
```

**Success Criteria**:
- ✅ Protected endpoints reject requests without JWT (401)
- ✅ Protected endpoints accept requests with valid JWT (200)
- ✅ Invalid JWT returns 401
- ✅ Expired JWT returns 401

##### Test 5: Middleware Validation (0.5h)
**Objective**: Verify Authority middleware is applied

**Test Cases**:
```bash
# Test 1: User with insufficient permissions
# Create a limited user
curl -X POST http://localhost:9100/user/create \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "limited",
    "password": "Pass123456",
    "roleIds": [2]
  }'

# Login as limited user
LIMITED_TOKEN=$(curl -s -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"limited","password":"Pass123456"}' | jq -r '.data.token')

# Try to create user (should fail due to Authority middleware)
curl -X POST http://localhost:9100/user/create \
  -H "Authorization: Bearer $LIMITED_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "test",
    "password": "Pass123456"
  }'

# Expected: 403 Forbidden (insufficient permissions)
```

**Success Criteria**:
- ✅ Authority middleware checks permissions
- ✅ Unauthorized actions return 403
- ✅ Authorized actions succeed

##### Test 6: Response Format Validation (0.5h)
**Objective**: Verify response format matches Proto definitions

**Test Cases**:
```bash
# Test response structure
RESPONSE=$(curl -s -X POST http://localhost:9100/user \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"id": 1}')

echo $RESPONSE | jq .

# Verify response fields match Proto definition
echo $RESPONSE | jq '.data | has("id")'      # Should be true
echo $RESPONSE | jq '.data | has("username")'  # Should be true
echo $RESPONSE | jq '.data | has("email")'     # Should be true
```

**Success Criteria**:
- ✅ Response structure matches Proto definition
- ✅ All required fields present
- ✅ Field types correct (string, int, etc.)
- ✅ JSON tags match Proto json_name

#### Test Report Template
**File**: `specs/003-proto-first-api-generation/e2e-test-report.md`

```markdown
# E2E Test Report - User Module Migration

**Date**: 2025-10-09
**Tester**: @qa
**Module**: User API
**Status**: ✅ PASSED / ❌ FAILED

## Test Summary

| Test Case | Status | Duration | Notes |
|-----------|--------|----------|-------|
| Compilation | ✅ PASS | 15s | No errors |
| Service Startup | ✅ PASS | 5s | Both services started |
| Public Endpoints | ✅ PASS | 2s | All 3 endpoints work |
| Protected Endpoints | ✅ PASS | 3s | JWT validation works |
| Middleware | ✅ PASS | 2s | Authority checks work |
| Response Format | ✅ PASS | 1s | Matches Proto |

## Detailed Results

### Test 1: Compilation
- **Status**: ✅ PASS
- **Duration**: 15 seconds
- **Output**: Binary created at `bin/core-api`
- **Notes**: No warnings or errors

### Test 2: Service Startup
- **Status**: ✅ PASS
- **RPC Port**: 9101 (listening)
- **API Port**: 9100 (listening)
- **Notes**: Both services healthy

### Test 3: Public Endpoints
- **GetCaptcha**: ✅ PASS (200 OK)
- **Login**: ✅ PASS (200 OK, token received)
- **Register**: ✅ PASS (200 OK, user created)

### Test 4: Protected Endpoints
- **Without JWT**: ✅ PASS (401 Unauthorized)
- **With Valid JWT**: ✅ PASS (200 OK)
- **With Invalid JWT**: ✅ PASS (401 Unauthorized)

### Test 5: Middleware
- **Authority Check**: ✅ PASS (403 for unauthorized)
- **Casbin Integration**: ✅ PASS

### Test 6: Response Format
- **Field Presence**: ✅ PASS (all fields present)
- **Field Types**: ✅ PASS (types correct)
- **JSON Tags**: ✅ PASS (matches Proto)

## Issues Found

None

## Recommendations

1. ✅ User module migration successful
2. ✅ Ready to migrate additional modules
3. ✅ Proto-First approach validated
```

#### Deliverables
1. ✅ E2E test report
2. ✅ Test scripts (bash)
3. ✅ Performance metrics
4. ✅ Regression test results

---

## Phase 5 Timeline

| Day | Tasks | Agent | Hours |
|-----|-------|-------|-------|
| **Day 1** | [PF-020] Migrate User Module (Steps 1-3) | @backend | 2.5h |
| **Day 1** | [PF-020] Migrate User Module (Steps 4-6) | @backend | 2.5h |
| **Day 2** | [PF-021] Create Migration Guide (Sections 1-3) | @doc | 2.5h |
| **Day 2** | [PF-021] Create Migration Guide (Sections 4-6) | @doc | 1.5h |
| **Day 3** | [PF-022] Update Project Documentation | @doc | 2.5h |
| **Day 3** | [PF-023] E2E Testing (Tests 1-3) | @qa | 1.5h |
| **Day 4** | [PF-023] E2E Testing (Tests 4-6) | @qa | 2h |
| **Day 4** | Phase 5 Completion Report | @pm | 1h |
| **Total** | | | **16h** |

---

## Success Metrics

### Technical Metrics
- ✅ User module successfully migrated
- ✅ Generated .api file compiles without errors
- ✅ All existing tests pass (100% pass rate)
- ✅ API service starts successfully
- ✅ All endpoints respond correctly
- ✅ JWT authentication works
- ✅ Middleware applied correctly

### Documentation Metrics
- ✅ Migration guide complete (5+ sections)
- ✅ CLAUDE.md updated with Proto-First section
- ✅ Example code for all options
- ✅ Troubleshooting guide with 5+ scenarios

### Quality Metrics
- ✅ Zero regression bugs
- ✅ E2E test coverage: 100% of critical paths
- ✅ Response times unchanged (<10ms overhead)

---

## Risk Management

### Risk Matrix

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Generated .api differs significantly | High | Medium | Detailed diff review, manual reconciliation |
| API service fails to compile | High | Low | Incremental testing, validation at each step |
| Handler behavior changes | Medium | Low | Comprehensive testing, rollback plan |
| Performance degradation | Low | Low | Benchmark before/after, optimize if needed |
| Team adoption resistance | Medium | Medium | Clear documentation, training, support |

### Mitigation Strategies

1. **Incremental Migration**: Start with User module, validate before proceeding
2. **Backup Strategy**: Keep .backup files, version control for rollback
3. **Testing Strategy**: Unit tests → Integration tests → E2E tests → Manual testing
4. **Documentation**: Comprehensive guides, examples, troubleshooting
5. **Support**: Dedicated support channel for questions during migration

---

## Rollback Plan

### Scenario 1: Generated .api is incorrect
```bash
# Restore backup
cp api/desc/core/user.api.backup api/desc/core/user.api

# Regenerate code
make gen-api-code
```

### Scenario 2: API service fails to compile
```bash
# Revert all changes
git checkout develop -- rpc/desc/user.proto
git checkout develop -- api/desc/core/user.api

# Regenerate code
make gen-rpc
make gen-api-code
```

### Scenario 3: Complete rollback of Phase 5
```bash
# Revert to Phase 4 completion
git revert <phase5-commits>

# Restore all backups
find . -name "*.backup" -exec bash -c 'cp "$0" "${0%.backup}"' {} \;

# Regenerate all code
make gen-all
```

---

## Deliverables Checklist

### Code Changes
- [ ] `rpc/desc/user.proto` updated with Go-Zero options
- [ ] `api/desc/core/user.api` generated from Proto
- [ ] Backup files created
- [ ] All changes committed to git

### Documentation
- [ ] Migration guide (`migration-guide.md`)
- [ ] CLAUDE.md updated
- [ ] E2E test report
- [ ] Phase 5 completion report

### Testing
- [ ] Compilation validation passed
- [ ] Service startup validated
- [ ] Public endpoint tests passed
- [ ] Protected endpoint tests passed
- [ ] Middleware tests passed
- [ ] Response format tests passed

### Project Management
- [ ] All Notion tasks updated to "Done"
- [ ] Test reports uploaded to Notion
- [ ] Completion report generated
- [ ] Phase 5 branch merged (if applicable)

---

## Next Steps After Phase 5

### Immediate (Post-Phase 5)
1. ✅ Team demo of Proto-First workflow
2. ✅ Collect feedback from pilot migration
3. ✅ Address any issues discovered
4. ✅ Update documentation based on feedback

### Short-term (1-2 weeks)
1. Migrate additional modules (Role, Menu, Department)
2. Create team training materials
3. Schedule training session
4. Monitor adoption metrics

### Long-term (1-2 months)
1. Migrate all 15 core modules
2. Deprecate dual API approach
3. Update CI/CD pipeline
4. Measure efficiency gains

---

## Appendix A: Command Reference

### Phase 5 Commands

```bash
# Migrate User Module
cd /d/Projects/simple-admin-core

# Step 1: Build plugin
make build-proto-plugin

# Step 2: Generate .api from Proto
make gen-proto-api

# Step 3: Validate generated .api
goctls api validate --api api/desc/core/user.api.generated

# Step 4: Backup existing .api
cp api/desc/core/user.api api/desc/core/user.api.backup

# Step 5: Replace with generated
mv api/desc/core/user.api.generated api/desc/core/user.api

# Step 6: Regenerate API code
make gen-api-code

# Step 7: Compile
go build ./api/...

# Step 8: Test
go test ./api/internal/logic/user/... -v
```

### E2E Testing Commands

```bash
# Start services
go run rpc/core.go -f rpc/etc/core.yaml &
go run api/core.go -f api/etc/core.yaml &

# Test public endpoint
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"simple-admin"}'

# Test protected endpoint
curl -X POST http://localhost:9100/user/list \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"page":1,"pageSize":10}'
```

---

**Plan Created**: 2025-10-09
**Plan Status**: 🚀 **READY TO START**
**Next Action**: Begin [PF-020] User Module Migration
