# User Module E2E Testing Guide

## Overview

This guide provides comprehensive End-to-End (E2E) testing procedures for the Simple Admin Core User module, covering authentication, authorization, profile management, and security flows.

## Test Environment Setup

### Prerequisites

1. **Running Services**
   ```bash
   # Start RPC service
   cd rpc
   go run core.go -f etc/core.yaml

   # Start API service (in another terminal)
   cd api
   go run core.go -f etc/core.yaml
   ```

2. **Database**
   - PostgreSQL or MySQL running
   - Test database created and migrated
   - Clean state before test run

3. **Redis**
   - Redis server running for caching and captcha storage

### Environment Variables

```bash
export API_BASE_URL="http://localhost:9100"
export RPC_BASE_URL="http://localhost:9101"
export TEST_USERNAME="e2eTestUser"
export TEST_EMAIL="e2etest@example.com"
export TEST_PASSWORD="TestPass123!"
```

## Test Tools

### Option 1: Postman (Recommended)

**Import Collection**:
1. Open Postman
2. Click **Import**
3. Select `user-module-e2e.postman_collection.json`
4. Update environment variables:
   - `baseUrl`: `http://localhost:9100`
   - `testUsername`: Your test username
   - `testEmail`: Your test email
   - `testPassword`: Your test password

**Run Collection**:
- Click **Run Collection**
- Select all requests
- Click **Run User Module E2E Tests**

**Expected Result**: ✅ All 20 tests pass

### Option 2: cURL

See individual test scenarios below for cURL commands.

### Option 3: Newman (CLI)

```bash
# Install Newman
npm install -g newman

# Run collection
newman run user-module-e2e.postman_collection.json \
  --environment postman-env.json \
  --reporters cli,htmlextra

# Generate HTML report
newman run user-module-e2e.postman_collection.json \
  --reporters htmlextra \
  --reporter-htmlextra-export report.html
```

## Test Scenarios

### 1. Authentication Flow (8 tests)

#### 1.1 Get Captcha ✅

**Endpoint**: `GET /captcha`
**Authentication**: None

```bash
curl -X GET http://localhost:9100/captcha
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.success",
  "data": {
    "captchaId": "abc123...",
    "imgPath": "data:image/png;base64,..."
  }
}
```

**Assertions**:
- Status code: 200
- Response has `captchaId`
- Response has `imgPath`

#### 1.2 Register New User ✅

**Endpoint**: `POST /user/register`
**Authentication**: None

```bash
curl -X POST http://localhost:9100/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "e2eTestUser",
    "password": "TestPass123!",
    "email": "e2etest@example.com",
    "captchaId": "<captcha_id>",
    "captcha": "00000"
  }'
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "login.registerSuccessTitle"
}
```

**Assertions**:
- Status code: 200
- `code` equals 0
- Success message present

#### 1.3 Login with Username ✅

**Endpoint**: `POST /user/login`
**Authentication**: None

```bash
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "e2eTestUser",
    "password": "TestPass123!",
    "captchaId": "<captcha_id>",
    "captcha": "00000"
  }'
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "login.loginSuccessTitle",
  "data": {
    "userId": "uuid-here",
    "token": "eyJhbGciOiJ...",
    "expire": 1728576000000
  }
}
```

**Assertions**:
- Status code: 200
- Token is not empty
- UserId is present
- Expire time is in future

**Save for later**:
- `token` → Use in Authorization header
- `userId` → Use for user-specific operations

#### 1.4 Login with Wrong Password (Should Fail) ❌

**Endpoint**: `POST /user/login`
**Expected Status**: 400 or 401

```bash
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "e2eTestUser",
    "password": "WrongPassword123!",
    "captchaId": "<captcha_id>",
    "captcha": "00000"
  }'
```

**Expected Response**:
```json
{
  "code": 10001,
  "msg": "login.wrongUsernameOrPassword"
}
```

**Assertions**:
- Status code: 400 or 401
- Error message present
- No token returned

#### 1.5 Login with Email ✅

**Endpoint**: `POST /user/login_by_email`

```bash
curl -X POST http://localhost:9100/user/login_by_email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "e2etest@example.com",
    "captcha": "12345"
  }'
```

#### 1.6 Login with SMS ✅

**Endpoint**: `POST /user/login_by_sms`

```bash
curl -X POST http://localhost:9100/user/login_by_sms \
  -H "Content-Type: application/json" \
  -d '{
    "phoneNumber": "+886912345678",
    "captcha": "12345"
  }'
```

#### 1.7 Login with Inactive User (Should Fail) ❌

Requires admin to set user status to 0 (inactive).

**Expected**: `login.userBanned` error

#### 1.8 Concurrent Login ✅

Login from multiple clients simultaneously. All should succeed.

---

### 2. User Information (3 tests)

#### 2.1 Get User Info ✅

**Endpoint**: `GET /user/info`
**Authentication**: Required (JWT)

```bash
curl -X GET http://localhost:9100/user/info \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.success",
  "data": {
    "userId": "uuid",
    "username": "e2eTestUser",
    "nickname": "Test User",
    "avatar": null,
    "homePath": "/dashboard",
    "roleName": ["user"],
    "departmentName": "",
    "locale": "en-US"
  }
}
```

**Assertions**:
- Status code: 200
- User data matches registered user
- Username is correct

#### 2.2 Get User Permissions ✅

**Endpoint**: `GET /user/perm`
**Authentication**: Required

```bash
curl -X GET http://localhost:9100/user/perm \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.success",
  "data": ["user:view", "user:edit", "profile:edit"]
}
```

**Assertions**:
- Status code: 200
- Permission array returned
- At least one permission present

#### 2.3 Get User Profile ✅

**Endpoint**: `GET /user/profile`
**Authentication**: Required

```bash
curl -X GET http://localhost:9100/user/profile \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.success",
  "data": {
    "nickname": "Test User",
    "avatar": null,
    "mobile": null,
    "email": "e2etest@example.com",
    "locale": "en-US"
  }
}
```

---

### 3. Profile Management (2 tests)

#### 3.1 Update User Profile ✅

**Endpoint**: `POST /user/profile`
**Authentication**: Required

```bash
curl -X POST http://localhost:9100/user/profile \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "Updated Test User",
    "mobile": "+886912345678",
    "locale": "zh-TW"
  }'
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.updateSuccess"
}
```

**Assertions**:
- Status code: 200
- Update successful

#### 3.2 Verify Profile Update ✅

**Endpoint**: `GET /user/profile`
**Expected**: Updated values persisted

```bash
curl -X GET http://localhost:9100/user/profile \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "data": {
    "nickname": "Updated Test User",
    "mobile": "+886912345678",
    "locale": "zh-TW"
  }
}
```

**Assertions**:
- Nickname changed to "Updated Test User"
- Mobile number updated
- Locale changed to "zh-TW"

---

### 4. Password Management (3 tests)

#### 4.1 Change Password ✅

**Endpoint**: `POST /user/change_password`
**Authentication**: Required

```bash
curl -X POST http://localhost:9100/user/change_password \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "oldPassword": "TestPass123!",
    "newPassword": "NewTestPass456!"
  }'
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "user.passwordChangeSuccess"
}
```

**Assertions**:
- Status code: 200
- Success message

#### 4.2 Login with New Password ✅

**Endpoint**: `POST /user/login`
**Use new password to verify change**

```bash
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "e2eTestUser",
    "password": "NewTestPass456!",
    "captchaId": "<captcha_id>",
    "captcha": "00000"
  }'
```

**Expected**: Login successful with new password

#### 4.3 Login with Old Password (Should Fail) ❌

**Endpoint**: `POST /user/login`
**Use old password**

**Expected**: Login fails with error

#### 4.4 Reset Password by Email ✅

**Endpoint**: `POST /user/reset_password_by_email`

```bash
curl -X POST http://localhost:9100/user/reset_password_by_email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "e2etest@example.com",
    "captcha": "12345",
    "password": "ResetPass789!"
  }'
```

#### 4.5 Reset Password by SMS ✅

**Endpoint**: `POST /user/reset_password_by_sms`

```bash
curl -X POST http://localhost:9100/user/reset_password_by_sms \
  -H "Content-Type: application/json" \
  -d '{
    "phoneNumber": "+886912345678",
    "captcha": "12345",
    "password": "ResetPass789!"
  }'
```

---

### 5. Token Management (3 tests)

#### 5.1 Refresh Token ✅

**Endpoint**: `GET /user/refresh_token`
**Authentication**: Required

```bash
curl -X GET http://localhost:9100/user/refresh_token \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.success",
  "data": {
    "token": "eyJhbGciOiJ...",
    "expiredAt": 1728579600000
  }
}
```

**Assertions**:
- New token received
- Expired time is in future
- Old token still valid until expiry

#### 5.2 Get Access Token ✅

**Endpoint**: `GET /user/access_token`
**Authentication**: Required

```bash
curl -X GET http://localhost:9100/user/access_token \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "common.success",
  "data": {
    "token": "eyJhbGciOiJ...",
    "expiredAt": 1728576000000
  }
}
```

#### 5.3 Use Expired Token (Should Fail) ❌

Wait for token to expire or manipulate JWT manually.

**Expected**: 401 Unauthorized

---

### 6. Logout (2 tests)

#### 6.1 Logout ✅

**Endpoint**: `GET /user/logout`
**Authentication**: Required

```bash
curl -X GET http://localhost:9100/user/logout \
  -H "Authorization: Bearer <token>"
```

**Expected Response**:
```json
{
  "code": 0,
  "msg": "login.logoutSuccessTitle"
}
```

**Assertions**:
- Status code: 200
- Token should be invalidated

#### 6.2 Access Protected Endpoint After Logout (Should Fail) ❌

**Endpoint**: `GET /user/info`
**Use logged-out token**

```bash
curl -X GET http://localhost:9100/user/info \
  -H "Authorization: Bearer <logged_out_token>"
```

**Expected**:
- Status code: 401
- Error: Token invalid or expired

---

## Security Testing

### 7. Authorization Tests

#### 7.1 Access Without Token ❌

```bash
curl -X GET http://localhost:9100/user/info
```

**Expected**: 401 Unauthorized

#### 7.2 Access with Invalid Token ❌

```bash
curl -X GET http://localhost:9100/user/info \
  -H "Authorization: Bearer invalid_token_here"
```

**Expected**: 401 Unauthorized

#### 7.3 Access with Malformed Token ❌

```bash
curl -X GET http://localhost:9100/user/info \
  -H "Authorization: Bearer eyJhbGciOiJ.invalid"
```

**Expected**: 401 Unauthorized

#### 7.4 SQL Injection Test ❌

```bash
curl -X POST http://localhost:9100/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin'--",
    "password": "anything"
  }'
```

**Expected**: Login fails (SQL injection prevented)

#### 7.5 XSS Test ❌

```bash
curl -X POST http://localhost:9100/user/profile \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "<script>alert(1)</script>"
  }'
```

**Expected**: Script tags sanitized or escaped

---

## Performance Testing

### 8. Load Tests

#### 8.1 Concurrent Login (100 users)

```bash
# Using Apache Bench
ab -n 1000 -c 100 -p login.json \
  -T application/json \
  http://localhost:9100/user/login
```

**Expected**:
- 95% requests < 500ms
- 99% requests < 1000ms
- 0% error rate

#### 8.2 Token Refresh Under Load

```bash
ab -n 1000 -c 50 -H "Authorization: Bearer <token>" \
  http://localhost:9100/user/refresh_token
```

**Expected**:
- All requests succeed
- Response time consistent

---

## Test Execution Checklist

### Pre-Test Setup

- [ ] Services running (API + RPC)
- [ ] Database clean state
- [ ] Redis running
- [ ] Test user doesn't exist
- [ ] Postman collection imported
- [ ] Environment variables set

### Test Execution

- [ ] Authentication Flow (8 tests)
- [ ] User Information (3 tests)
- [ ] Profile Management (2 tests)
- [ ] Password Management (5 tests)
- [ ] Token Management (3 tests)
- [ ] Logout (2 tests)
- [ ] Security Tests (5 tests)
- [ ] Performance Tests (2 tests)

**Total**: 30 test scenarios

### Post-Test Cleanup

- [ ] Delete test user
- [ ] Clear test tokens from database
- [ ] Reset database to clean state
- [ ] Verify no test data remaining

---

## Test Results Documentation

### Test Run Report Template

```
# E2E Test Run Report

**Date**: 2025-10-10
**Environment**: Local Development
**Tester**: [Name]

## Summary

- Total Tests: 30
- Passed: 28 ✅
- Failed: 2 ❌
- Skipped: 0
- Duration: 5m 32s

## Failed Tests

1. **5.3 Use Expired Token**
   - Status: ❌ Failed
   - Expected: 401 Unauthorized
   - Actual: 200 OK
   - Issue: Token expiry not enforced
   - Ticket: #123

2. **7.5 XSS Test**
   - Status: ❌ Failed
   - Expected: Script sanitized
   - Actual: Script executed
   - Issue: Input sanitization missing
   - Ticket: #124

## Coverage

- Authentication: 100% (8/8)
- User Info: 100% (3/3)
- Profile Mgmt: 100% (2/2)
- Password Mgmt: 80% (4/5)
- Token Mgmt: 67% (2/3)
- Logout: 100% (2/2)
- Security: 60% (3/5)
- Performance: 100% (2/2)

## Recommendations

1. Fix token expiry enforcement
2. Implement input sanitization for XSS
3. Add rate limiting for authentication endpoints
4. Increase test coverage for edge cases
```

---

## Continuous Integration

### GitHub Actions Workflow

```yaml
name: E2E Tests

on: [push, pull_request]

jobs:
  e2e:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:14
        env:
          POSTGRES_PASSWORD: postgres
      redis:
        image: redis:7

    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
      - uses: actions/setup-node@v2

      - name: Start Services
        run: |
          cd rpc && go run core.go -f etc/core.yaml &
          cd api && go run core.go -f etc/core.yaml &
          sleep 10

      - name: Run E2E Tests
        run: |
          npm install -g newman
          newman run tests/e2e/user-module-e2e.postman_collection.json \
            --reporters cli,junit \
            --reporter-junit-export results.xml

      - name: Publish Test Results
        uses: EnricoMi/publish-unit-test-result-action@v1
        if: always()
        with:
          files: results.xml
```

---

## Troubleshooting

### Common Issues

**Issue**: Captcha validation fails
**Solution**: Use test captcha "00000" or disable captcha in test environment

**Issue**: Token expired
**Solution**: Generate fresh token before running tests

**Issue**: User already exists
**Solution**: Clean database or use unique username per test run

**Issue**: Services not responding
**Solution**: Check if API and RPC services are running on correct ports

---

**Status**: 20/30 test scenarios documented ✅
**Estimated Execution Time**: ~10 minutes (manual), ~5 minutes (automated)
**Created**: 2025-10-10
