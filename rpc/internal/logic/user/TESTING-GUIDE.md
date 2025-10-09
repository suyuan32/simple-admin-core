# User Module RPC Testing Guide

## Overview

This guide provides a comprehensive testing strategy for the User module RPC logic layer, including unit tests, integration tests, and E2E test scenarios.

## Test Coverage Status

### Implemented Unit Tests (3/16) ✅

1. ✅ `login_logic_test.go` - Login authentication tests
   - Success case
   - Invalid username
   - Invalid password
   - Inactive user

2. ✅ `register_logic_test.go` - User registration tests
   - Success case
   - Duplicate email
   - Invalid email

3. ✅ `change_password_logic_test.go` - Password change tests
   - Success case
   - Wrong old password
   - Missing userId context
   - Same password

### Pending Unit Tests (13/16) ⏳

4. ⏳ `login_by_email_logic_test.go` - Email login tests
5. ⏳ `login_by_sms_logic_test.go` - SMS login tests
6. ⏳ `register_by_email_logic_test.go` - Email registration tests
7. ⏳ `register_by_sms_logic_test.go` - SMS registration tests
8. ⏳ `reset_password_by_email_logic_test.go` - Email password reset tests
9. ⏳ `reset_password_by_sms_logic_test.go` - SMS password reset tests
10. ⏳ `get_user_info_logic_test.go` - Get user info tests
11. ⏳ `get_user_perm_code_logic_test.go` - Get permissions tests
12. ⏳ `get_user_profile_logic_test.go` - Get profile tests
13. ⏳ `update_user_profile_logic_test.go` - Update profile tests
14. ⏳ `logout_logic_test.go` - Logout tests
15. ⏳ `refresh_token_logic_test.go` - Token refresh tests
16. ⏳ `access_token_logic_test.go` - Access token tests

## Test Framework

### Dependencies

```go
import (
	"context"
	"testing"

	"github.com/chimerakang/simple-admin-core/rpc/ent/enttest"
	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"
	"github.com/chimerakang/simple-admin-common/utils/encrypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)
```

### Setup Helper Functions

```go
// setupTestDB creates an in-memory test database
func setupTestDB(t *testing.T) *svc.ServiceContext {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	return &svc.ServiceContext{DB: client}
}

// createTestUser creates a test user in the database
func createTestUser(t *testing.T, client *svc.ServiceContext, username string, password string) *ent.User {
	hashedPassword := encrypt.BcryptEncrypt(password)
	userInfo, err := client.DB.User.Create().
		SetUsername(username).
		SetPassword(hashedPassword).
		SetNickname("Test User").
		SetStatus(1).
		SetEmail("test@example.com").
		Save(context.Background())
	require.NoError(t, err)
	return userInfo
}
```

### Test Naming Convention

Tests follow the pattern: `Test<LogicName>_<Method>_<Scenario>`

Examples:
- `TestLoginLogic_Login_Success`
- `TestLoginLogic_Login_InvalidPassword`
- `TestRegisterLogic_Register_DuplicateEmail`

### Test Structure

Each test follows AAA pattern (Arrange, Act, Assert):

```go
func TestLoginLogic_Login_Success(t *testing.T) {
	// Arrange: Setup test database and test data
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()
	createTestUser(t, svcCtx, "testuser", "password123")

	// Act: Execute the logic
	logic := NewLoginLogic(context.Background(), svcCtx)
	resp, err := logic.Login(&core.LoginReq{
		Username: "testuser",
		Password: "password123",
	})

	// Assert: Verify results
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
}
```

## Running Tests

### Run All User Logic Tests

```bash
cd rpc/internal/logic/user
go test -v ./...
```

### Run Specific Test File

```bash
go test -v -run TestLoginLogic_Login_Success
```

### Run with Coverage

```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Expected Coverage

Target: **70%+ code coverage**

Current: **~20%** (3/16 files with tests)

## Test Scenarios Template

### Template for Remaining Tests

```go
func TestLoginByEmailLogic_LoginByEmail_Success(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Create test user with email
	testEmail := "test@example.com"
	userInfo := createTestUser(t, svcCtx, "testuser", "password123")
	svcCtx.DB.User.UpdateOne(userInfo).SetEmail(testEmail).SaveX(context.Background())

	// Test email login
	logic := NewLoginByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.LoginByEmail(&core.LoginByEmailReq{
		Email:   testEmail,
		Captcha: "12345",
	})

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.Code)
}

func TestLoginByEmailLogic_LoginByEmail_EmailNotFound(t *testing.T) {
	svcCtx := setupTestDB(t)
	defer svcCtx.DB.Close()

	// Test login with non-existent email
	logic := NewLoginByEmailLogic(context.Background(), svcCtx)
	resp, err := logic.LoginByEmail(&core.LoginByEmailReq{
		Email:   "nonexistent@example.com",
		Captcha: "12345",
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)
}
```

## Common Test Scenarios

### Authentication Tests

- ✅ Valid credentials → Success
- ✅ Invalid credentials → Error
- ✅ Inactive user → Error
- ⏳ Rate limiting → Error (if implemented)
- ⏳ Concurrent logins → Success

### Registration Tests

- ✅ Valid data → Success
- ✅ Duplicate email → Error
- ⏳ Duplicate username → Error
- ⏳ Invalid email format → Error (API layer validates)
- ⏳ Weak password → Error (API layer validates)

### Password Management Tests

- ✅ Valid old password → Success
- ✅ Invalid old password → Error
- ✅ Missing userId → Error
- ⏳ Password history validation → Error (if implemented)

### Profile Management Tests

- ⏳ Valid update → Success
- ⏳ Update non-existent user → Error
- ⏳ Update with invalid data → Error

### Token Management Tests

- ⏳ Valid token → Success
- ⏳ Expired token → Error
- ⏳ Invalid token → Error
- ⏳ Token refresh → Success
- ⏳ Token blacklist → Error

## Integration Testing

### Test with Real Database

```bash
# Set test database connection
export TEST_DB_DSN="postgres://user:password@localhost:5432/testdb?sslmode=disable"

# Run tests
go test -v -tags=integration ./...
```

### Mock External Services

For services that depend on external APIs (SMS, Email):

```go
type MockSMSProvider struct{}

func (m *MockSMSProvider) SendSMS(phone, message string) error {
	// Mock implementation
	return nil
}
```

## Performance Testing

### Benchmark Tests

```go
func BenchmarkLoginLogic_Login(b *testing.B) {
	svcCtx := setupTestDB(&testing.T{})
	defer svcCtx.DB.Close()

	createTestUser(&testing.T{}, svcCtx, "testuser", "password123")
	logic := NewLoginLogic(context.Background(), svcCtx)

	req := &core.LoginReq{
		Username: "testuser",
		Password: "password123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logic.Login(req)
	}
}
```

### Run Benchmarks

```bash
go test -bench=. -benchmem
```

## Test Data Management

### Fixtures

Create reusable test data:

```go
var (
	testUsers = []struct {
		username string
		password string
		email    string
	}{
		{"user1", "password1", "user1@example.com"},
		{"user2", "password2", "user2@example.com"},
		{"admin", "admin123", "admin@example.com"},
	}
)

func seedTestUsers(t *testing.T, svcCtx *svc.ServiceContext) {
	for _, user := range testUsers {
		createTestUser(t, svcCtx, user.username, user.password)
	}
}
```

### Test Data Cleanup

```go
func teardownTestDB(t *testing.T, svcCtx *svc.ServiceContext) {
	// Clean up test data
	svcCtx.DB.User.Delete().ExecX(context.Background())
	svcCtx.DB.Close()
}
```

## Continuous Integration

### GitHub Actions Workflow

```yaml
name: RPC Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: '1.25'

      - name: Run Tests
        run: |
          cd rpc/internal/logic/user
          go test -v -coverprofile=coverage.out ./...

      - name: Upload Coverage
        uses: codecov/codecov-action@v2
        with:
          file: ./coverage.out
```

## Best Practices

### 1. Test Isolation

- Each test should be independent
- Use `setupTestDB` to create fresh database for each test
- Don't rely on test execution order

### 2. Clear Test Names

- Use descriptive names that explain the scenario
- Follow convention: `Test<Logic>_<Method>_<Scenario>`

### 3. Comprehensive Assertions

```go
// Good: Multiple specific assertions
assert.NoError(t, err)
assert.NotNil(t, resp)
assert.Equal(t, uint32(0), resp.Code)
assert.NotEmpty(t, resp.Data.UserId)

// Bad: Single vague assertion
assert.NotNil(t, resp)
```

### 4. Test Edge Cases

- Null values
- Empty strings
- Boundary values
- Concurrent operations

### 5. Test Error Paths

- Invalid input
- Missing required fields
- Permission denied
- Resource not found

## Troubleshooting

### Common Issues

**Issue**: `panic: runtime error: invalid memory address`
**Solution**: Check if `svcCtx.DB` is properly initialized

**Issue**: `tests pass locally but fail in CI`
**Solution**: Ensure test database is properly set up in CI environment

**Issue**: `tests are slow`
**Solution**: Use in-memory SQLite for unit tests, real DB only for integration tests

## Next Steps

1. **Complete Unit Tests**: Implement remaining 13 test files
   - Estimated effort: 4-6 hours
   - Target coverage: 70%+

2. **Add Integration Tests**: Test with real PostgreSQL database
   - Estimated effort: 2-3 hours

3. **Add E2E Tests**: Test complete API → RPC → Database flow
   - Estimated effort: 2-3 hours

4. **Set Up CI/CD**: Automate test execution on every commit
   - Estimated effort: 1-2 hours

## Resources

- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Ent Testing Guide](https://entgo.io/docs/testing/)
- [Go-Zero Testing Best Practices](https://go-zero.dev/docs/tasks/test/)

---

**Status**: 3/16 unit tests implemented (18.75%)
**Target**: 16/16 unit tests (100%)
**Coverage**: ~20% → Target 70%+

**Created**: 2025-10-10
**Updated**: 2025-10-10
