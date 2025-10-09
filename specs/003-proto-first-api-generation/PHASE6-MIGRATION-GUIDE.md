# Spec-003 Phase 6: Proto-First Migration Guide
## User Module Migration to Auto-Generated .api

**Status**: ⏸️ DEFERRED (Pending strategic decision)
**Date**: 2025-10-10
**Related**: Spec-004 User Module Proto Completion

---

## Executive Summary

Phase 6 was planned to migrate the User module from manually-maintained `user.api` to Proto-First auto-generated `.api` files. During analysis, we discovered that **direct migration requires careful handling of validation rules and comments**, leading to a strategic decision to defer full migration until a migration strategy is finalized.

**Current Status**:
- ✅ Proto definitions: 100% complete (22/22 methods)
- ✅ Generated .api file: Available (`user.api.generated`, 233 lines)
- ✅ Manual .api file: Functional (` user.api`, 419 lines)
- ⏸️ Migration: Deferred (requires validation preservation strategy)

---

## Migration Analysis

### File Comparison

| Aspect | Manual (user.api) | Generated (user.api.generated) |
|--------|-------------------|--------------------------------|
| **Lines of Code** | 419 | 233 |
| **Type Definitions** | 22 types | 22 types |
| **Handlers** | 22 handlers | 22 handlers |
| **Comments** | ✅ Bilingual (EN/ZH) | ❌ None |
| **Validation** | ✅ Comprehensive | ❌ None |
| **Structure** | ✅ BaseUUIDInfo inheritance | ❌ Flat structure |
| **Response Types** | BaseMsgResp, BaseDataInfo | BaseResp, custom types |

### Key Differences

#### 1. Validation Tags

**Manual .api** (Has validation):
```api
type LoginReq {
    // User Name | 用户名
    Username   string `json:"username" validate:"required,alphanum,max=20"`

    // Password | 密码
    Password   string `json:"password" validate:"required,max=30,min=6"`

    // Captcha ID | 验证码编号
    CaptchaId  string `json:"captchaId" validate:"required,len=20"`

    // Captcha | 验证码
    Captcha    string `json:"captcha" validate:"required,len=5"`
}
```

**Generated .api** (No validation):
```api
type LoginReq {
    Username string `json:"username"`
    Password string `json:"password"`
    CaptchaId string `json:"captchaId"`
    Captcha string `json:"captcha"`
}
```

**Impact**: Losing validation means API layer won't validate input, pushing validation to business logic layer (less efficient).

#### 2. Comments and Documentation

**Manual .api** (Bilingual comments):
```api
// The log in information | 登陆返回的数据信息
LoginInfo {
    // User's UUID | 用户的UUID
    UserId       string          `json:"userId"`

    // Token for authorization | 验证身份的token
    Token        string          `json:"token"`
}
```

**Generated .api** (No comments):
```api
LoginInfo {
    UserId string `json:"userId"`
    Token string `json:"token"`
    Expire uint64 `json:"expire"`
}
```

**Impact**: Losing comments reduces code maintainability and developer experience.

#### 3. Type Inheritance

**Manual .api** (Uses inheritance):
```api
UserInfo {
    BaseUUIDInfo  // Inherits: Id, CreatedAt, UpdatedAt

    Status *uint32 `json:"status,optional"`
    Username *string `json:"username,optional"`
}
```

**Generated .api** (Flat structure):
```api
UserInfo {
    Id *string `json:"id,optional"`
    CreatedAt *int64 `json:"createdAt,optional"`
    UpdatedAt *int64 `json:"updatedAt,optional"`
    Status *uint32 `json:"status,optional"`
    Username *string `json:"username,optional"`
}
```

**Impact**: Duplicated fields across types, harder to maintain consistency.

#### 4. Service Grouping

**Manual .api** (Clear grouping):
```api
@server(
    group: publicuser  // Public endpoints (no JWT)
)
service Core {
    @handler login
    post /user/login (LoginReq) returns (LoginResp)
}

@server(
    jwt: Auth
    group: user  // Protected endpoints
    middleware: Authority
)
service Core {
    @handler getUserInfo
    get /user/info returns (UserBaseIDInfoResp)
}
```

**Generated .api** (Different grouping):
```api
@server(
    jwt: Auth
    group: user
    middleware: Authority
)
service Core {
    // All protected endpoints
}

@server(
    group: user
    middleware: Authority
)
service Core {
    // Public endpoints (no JWT, but has middleware)
}
```

**Impact**: Generated file doesn't distinguish between public and protected groups as clearly.

---

## Migration Options

### Option A: Direct Replacement (NOT RECOMMENDED)

**Approach**: Replace `user.api` with `user.api.generated` directly

**Steps**:
```bash
cd api/desc/core
cp user.api user.api.manual.backup
cp user.api.generated user.api
make gen-api
```

**Pros**:
- ✅ Quick (5 minutes)
- ✅ Proto becomes single source of truth

**Cons**:
- ❌ Loses all validation tags
- ❌ Loses bilingual comments
- ❌ Loses type inheritance
- ❌ Breaking change (API behavior changes)
- ❌ Requires extensive testing

**Risk**: 🔴 **HIGH** - Breaking changes, validation loss

---

### Option B: Hybrid Approach (RECOMMENDED)

**Approach**: Use generated file as base, manually add validation and comments

**Steps**:

1. **Generate Base**:
   ```bash
   cd api/desc/core
   cp user.api.generated user.api.hybrid
   ```

2. **Add Validation Tags**:
   ```api
   // From: LoginReq {
   //     Username string `json:"username"`

   // To:
   LoginReq {
       Username string `json:"username" validate:"required,alphanum,max=20"`
       Password string `json:"password" validate:"required,max=30,min=6"`
       CaptchaId string `json:"captchaId" validate:"required,len=20"`
       Captcha string `json:"captcha" validate:"required,len=5"`
   }
   ```

3. **Add Comments**:
   ```api
   // Add bilingual comments
   // Login request | 登录参数
   LoginReq {
       // User Name | 用户名
       Username string `json:"username" validate:"required,alphanum,max=20"`
   }
   ```

4. **Replace User.api**:
   ```bash
   cp user.api user.api.manual.backup
   cp user.api.hybrid user.api
   make gen-api
   ```

5. **Test Thoroughly**:
   ```bash
   # Run unit tests
   cd api && go test ./...

   # Run E2E tests
   newman run tests/e2e/user-module-e2e.postman_collection.json
   ```

**Pros**:
- ✅ Maintains validation
- ✅ Preserves comments
- ✅ Proto-First benefits
- ✅ Non-breaking change

**Cons**:
- ⏳ Time-consuming (4-6 hours)
- ⏳ Manual work required
- ⏳ Requires careful review

**Risk**: 🟡 **MEDIUM** - Manual work, but controlled

**Estimated Effort**: 4-6 hours

---

### Option C: Tool-Assisted Migration (FUTURE)

**Approach**: Enhance protoc-gen-go-zero-api plugin to support validation and comments

**Steps**:

1. **Extend Proto with Validation Options**:
   ```protobuf
   import "validate/validate.proto";

   message LoginReq {
       string username = 1 [(validate.rules).string = {
           pattern: "^[a-zA-Z0-9]+$",
           max_len: 20
       }];
       string password = 2 [(validate.rules).string = {
           min_len: 6,
           max_len: 30
       }];
   }
   ```

2. **Update Plugin to Generate Validation**:
   ```go
   // In protoc-gen-go-zero-api
   func generateFieldTag(field *protogen.Field) string {
       tag := fmt.Sprintf("`json:\"%s\"", field.Desc.JSONName())

       // Add validation rules from proto
       if hasValidation(field) {
           tag += fmt.Sprintf(" validate:\"%s\"", getValidationRules(field))
       }

       tag += "`"
       return tag
   }
   ```

3. **Support Comments from Proto**:
   ```protobuf
   message LoginReq {
       // User Name | 用户名
       string username = 1;
   }
   ```

   Generated:
   ```api
   // User Name | 用户名
   Username string `json:"username" validate:"required,max=20"`
   ```

**Pros**:
- ✅ Fully automated
- ✅ Single source of truth (Proto)
- ✅ Consistent validation across services
- ✅ Future-proof

**Cons**:
- ⏳ Requires plugin enhancement (20-30 hours)
- ⏳ Needs validate.proto integration
- ⏳ Learning curve for team

**Risk**: 🟢 **LOW** - Once implemented, fully automated

**Estimated Effort**: 20-30 hours (plugin enhancement)

---

### Option D: Gradual Migration (PRAGMATIC)

**Approach**: Migrate new features to Proto-First, keep existing features as-is

**Strategy**:

1. **Existing Endpoints**: Keep manual `user.api`
2. **New Endpoints**: Add to `user.proto`, generate, merge into `user.api`
3. **Refactored Endpoints**: Migrate to Proto-First when refactoring

**Example Workflow**:

```bash
# New Feature: Add "Login with WeChat"

# 1. Add to user.proto
rpc loginByWechat(LoginByWechatReq) returns (LoginResp) {
    option (google.api.http) = {
        post: "/user/login_by_wechat"
        body: "*"
    };
}

# 2. Generate
make gen-rpc
protoc --go-zero-api_out=. user.proto

# 3. Extract new endpoint from generated file
grep -A 10 "loginByWechat" user.api.generated > new_endpoint.api

# 4. Manually add validation and comments
# 5. Append to user.api
cat new_endpoint.api >> user.api

# 6. Regenerate API code
make gen-api
```

**Pros**:
- ✅ No disruption to existing code
- ✅ Gradual adoption
- ✅ Low risk
- ✅ Team can learn incrementally

**Cons**:
- ⏳ Dual maintenance temporarily
- ⏳ Inconsistent approach
- ⏳ Longer migration timeline

**Risk**: 🟢 **LOW** - Incremental, reversible

**Estimated Effort**: 1-2 hours per new feature

---

## Recommended Migration Path

### Phase 6.1: Preparation (Week 1)

**Goal**: Set up foundation for Proto-First adoption

**Tasks**:
1. ✅ Complete Spec-004 (Proto definitions)
2. ✅ Generate user.api.generated
3. ⏳ Create validation mapping document
4. ⏳ Document migration strategy
5. ⏳ Get team buy-in

**Deliverables**:
- ✅ 100% Proto coverage
- ✅ Generated .api file
- ⏳ Migration guide (this document)
- ⏳ Team training materials

### Phase 6.2: Pilot Migration (Week 2)

**Goal**: Migrate one endpoint as proof of concept

**Tasks**:
1. Choose simple endpoint (e.g., `getUserProfile`)
2. Extract from generated file
3. Add validation and comments
4. Replace in user.api
5. Test thoroughly
6. Document learnings

**Success Criteria**:
- ✅ Endpoint works identically
- ✅ Validation preserved
- ✅ Comments maintained
- ✅ Team confident in process

### Phase 6.3: Full Migration (Week 3-4)

**Goal**: Migrate all 22 endpoints

**Approach**: Option B (Hybrid) or Option D (Gradual)

**Tasks**:
1. Migrate public endpoints (8 endpoints)
2. Test authentication flow
3. Migrate protected endpoints (14 endpoints)
4. Comprehensive E2E testing
5. Performance validation
6. Documentation update

**Success Criteria**:
- ✅ All 22 endpoints migrated
- ✅ Zero regression bugs
- ✅ E2E tests pass (100%)
- ✅ Performance maintained

### Phase 6.4: Continuous Improvement (Ongoing)

**Goal**: Optimize Proto-First workflow

**Tasks**:
1. Gather team feedback
2. Identify pain points
3. Consider Option C (tool enhancement)
4. Update templates and tooling
5. Train new team members

---

## Migration Checklist

### Pre-Migration

- [ ] Spec-004 complete (Proto definitions 100%)
- [ ] Generated .api file reviewed
- [ ] Validation mapping documented
- [ ] Team trained on Proto-First
- [ ] Backup of current user.api created
- [ ] Test environment prepared

### During Migration

- [ ] For each endpoint:
  - [ ] Extract from generated file
  - [ ] Add validation tags
  - [ ] Add bilingual comments
  - [ ] Preserve response types
  - [ ] Update service grouping
  - [ ] Run unit tests
  - [ ] Run E2E tests

### Post-Migration

- [ ] All endpoints migrated
- [ ] Validation preserved
- [ ] Comments maintained
- [ ] E2E tests pass (100%)
- [ ] Performance benchmarked
- [ ] Documentation updated
- [ ] Team retrospective

---

## Risk Mitigation

### Risk 1: Validation Loss

**Mitigation**:
- Option B: Manually add validation
- Option C: Enhance plugin to support validation
- Automated tests to catch validation issues

### Risk 2: Breaking Changes

**Mitigation**:
- Comprehensive E2E testing before deployment
- Feature flag for gradual rollout
- Quick rollback plan (keep manual .api in git)

### Risk 3: Team Resistance

**Mitigation**:
- Clear communication of benefits
- Training and documentation
- Pilot migration to build confidence
- Gradual adoption (Option D)

### Risk 4: Performance Degradation

**Mitigation**:
- Benchmark before and after migration
- Monitor production metrics
- Load testing with same scenarios

---

## Success Metrics

### Technical Metrics

- **Proto Coverage**: 100% (22/22 methods) ✅
- **Code Duplication**: Reduce from 100% to 0% (Proto as single source)
- **Maintenance Time**: Reduce by 40% (no dual updates)
- **Bug Rate**: 0 regressions after migration

### Process Metrics

- **Migration Time**: Target 2-4 weeks
- **Team Confidence**: 80%+ comfortable with Proto-First
- **Documentation**: Complete migration guide + training

### Business Metrics

- **API Consistency**: 100% (Proto guarantees consistency)
- **Development Speed**: 50% faster for new endpoints
- **Onboarding Time**: 20% reduction (simpler workflow)

---

## Tooling and Automation

### Validation Extraction Script

```bash
#!/bin/bash
# extract-validation.sh
# Extracts validation tags from manual .api file

FILE="user.api"
OUTPUT="validation-mapping.txt"

grep -B 1 "validate:" $FILE | \
    sed 's/--//' | \
    paste - - > $OUTPUT

echo "Validation mapping saved to $OUTPUT"
```

### Comment Extraction Script

```bash
#!/bin/bash
# extract-comments.sh
# Extracts bilingual comments from manual .api file

FILE="user.api"
OUTPUT="comments.txt"

grep "\/\/" $FILE | \
    grep "|" > $OUTPUT

echo "Comments saved to $OUTPUT"
```

### Diff Tool

```bash
# Compare manual vs generated
diff -u user.api user.api.generated > migration.diff
```

---

## Timeline

### Immediate (Week 1)

- [x] Complete Spec-004
- [x] Generate user.api.generated
- [x] Analyze differences
- [x] Document migration options
- [x] Create this guide

### Short-term (Week 2-4)

- [ ] Team training
- [ ] Pilot migration (1 endpoint)
- [ ] Full migration decision
- [ ] Execute chosen option

### Medium-term (Month 2-3)

- [ ] Migrate all User endpoints
- [ ] Migrate other modules (Role, Menu, etc.)
- [ ] Refine Proto-First workflow

### Long-term (Quarter 2)

- [ ] Consider Option C (tool enhancement)
- [ ] Standardize across all modules
- [ ] Continuous improvement

---

## Conclusion

**Phase 6 Status**: ⏸️ **DEFERRED** (Strategic decision pending)

**Recommendation**: **Option D (Gradual Migration)** for immediate adoption, plan for **Option C (Tool Enhancement)** in Q2 2025.

**Rationale**:
- ✅ Spec-004 unblocked Proto-First for User module (100% coverage)
- ✅ Generated .api file available and validated
- ⏸️ Direct replacement risks validation loss
- 🎯 Hybrid/Gradual approach balances risk and benefit
- 🔮 Tool enhancement provides long-term solution

**Next Steps**:
1. Review this guide with team
2. Decide on migration option (B or D recommended)
3. Schedule pilot migration (1 endpoint)
4. Gather feedback and iterate

---

**Document Version**: 1.0
**Created**: 2025-10-10
**Status**: Draft (Pending Team Review)
**Related Documents**:
- `specs/004-user-module-proto-completion/completion-report.md`
- `specs/004-user-module-proto-completion/final-integration-report.md`
- `specs/003-proto-first-api-generation/spec.md`

---

## Appendix: Quick Reference

### File Locations

- Manual .api: `api/desc/core/user.api` (419 lines)
- Generated .api: `api/desc/core/user.api.generated` (233 lines)
- Proto source: `rpc/desc/user.proto` (22 methods)
- Plugin: `tools/protoc-gen-go-zero-api/`

### Key Commands

```bash
# Generate .api from Proto
protoc --go-zero-api_out=api/desc/core \
    --proto_path=rpc/desc \
    user.proto

# Regenerate API code
cd api && make gen-api

# Run E2E tests
newman run tests/e2e/user-module-e2e.postman_collection.json

# Backup current .api
cp api/desc/core/user.api api/desc/core/user.api.backup.$(date +%Y%m%d)
```

### Contact

- **Tech Lead**: [Name]
- **PM**: @pm Agent
- **Questions**: Create issue in GitHub repo
