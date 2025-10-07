# Technical Plan: Traditional Chinese (zh-TW) Support

**Related Spec**: [spec.md](./spec.md)
**Created**: 2025-10-07
**Status**: Draft
**Estimated Effort**: 20-28 hours

## Architecture Overview

### System Architecture
```
┌─────────────────────────────────────────────────────────────────┐
│                         Browser                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Vben5 Frontend (Vue 3 + TypeScript)                      │  │
│  │  ┌────────────────────────────────────────────────────┐  │  │
│  │  │  Language Selector Component                        │  │  │
│  │  │  - localStorage: user_locale                        │  │  │
│  │  │  - Emits: locale-change event                       │  │  │
│  │  └────────────────────────────────────────────────────┘  │  │
│  │  ┌────────────────────────────────────────────────────┐  │  │
│  │  │  vue-i18n (v9)                                      │  │  │
│  │  │  - src/locales/lang/zh_CN.ts (existing)            │  │  │
│  │  │  - src/locales/lang/en.ts (existing)               │  │  │
│  │  │  - src/locales/lang/zh_TW.ts (NEW)                 │  │  │
│  │  └────────────────────────────────────────────────────┘  │  │
│  │  ┌────────────────────────────────────────────────────┐  │  │
│  │  │  Ant Design Vue Locale                              │  │  │
│  │  │  - import zh_TW from 'ant-design-vue/es/locale/zh_TW' │  │
│  │  └────────────────────────────────────────────────────┘  │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                              │ HTTP Request
                              │ Header: Accept-Language: zh-TW
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    API Service (Port 9100)                       │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  i18n Middleware (api/internal/middleware)               │  │
│  │  - Detect Accept-Language header                         │  │
│  │  - Set context locale                                    │  │
│  └──────────────────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  i18n Package (api/internal/i18n)                        │  │
│  │  ┌────────────────────────────────────────────────────┐ │  │
│  │  │  var.go                                             │ │  │
│  │  │  //go:embed locale/*.json                          │ │  │
│  │  │  var LocaleFS embed.FS                              │ │  │
│  │  └────────────────────────────────────────────────────┘ │  │
│  │  ┌────────────────────────────────────────────────────┐ │  │
│  │  │  locale/zh.json (Simplified Chinese - existing)    │ │  │
│  │  │  locale/en.json (English - existing)               │ │  │
│  │  │  locale/zh-TW.json (Traditional Chinese - NEW)     │ │  │
│  │  └────────────────────────────────────────────────────┘ │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                              │ gRPC Call
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    RPC Service (Port 9101)                       │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Database (PostgreSQL)                                    │  │
│  │  ┌────────────────────────────────────────────────────┐ │  │
│  │  │  sys_users                                          │ │  │
│  │  │  - id, username, password, ...                      │ │  │
│  │  │  - locale VARCHAR(10) NULL (NEW)                    │ │  │
│  │  └────────────────────────────────────────────────────┘ │  │
│  │  ┌────────────────────────────────────────────────────┐ │  │
│  │  │  sys_menus                                          │ │  │
│  │  │  - title: "route.dashboard" (i18n key)             │ │  │
│  │  │  - No change needed                                 │ │  │
│  │  └────────────────────────────────────────────────────┘ │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

### Locale Resolution Flow
```
User Request → Check Accept-Language Header
                    ↓
            Is zh-TW or zh-Hant-TW?
                    ↓ Yes
            Load zh-TW.json
                    ↓
            Translate all messages
                    ↓
            Return localized response

If translation key missing:
zh-TW.json → zh.json (fallback) → en.json (ultimate fallback)
```

## Technology Stack

### Backend
- **Language**: Go 1.25.0
- **Framework**: Go-Zero v1.9.1
- **i18n Library**: Built-in `embed.FS` (Go 1.16+)
- **Translation Tool**: OpenCC 1.1.x (for initial conversion)

### Frontend
- **Framework**: Vue 3.4+ (Vben5)
- **i18n Library**: vue-i18n v9.x
- **UI Framework**: Ant Design Vue 4.x
- **State Management**: Pinia
- **Build Tool**: Vite

### Database
- **DBMS**: PostgreSQL 15+
- **ORM**: Ent 0.14.5
- **Migration**: Ent auto-migration

### Development Tools
- **Conversion**: OpenCC CLI (`opencc -c s2twp.json`)
- **Validation**: JSON Schema validator
- **Testing**: Go test, Playwright (E2E)

## Implementation Details

### Phase 1: Backend API Layer (4-6 hours)

#### Task 1.1: Create zh-TW.json Language File
**File**: `api/internal/i18n/locale/zh-TW.json`

**Approach**:
```bash
# Step 1: Copy simplified Chinese as base
cp api/internal/i18n/locale/zh.json api/internal/i18n/locale/zh-TW.json

# Step 2: Convert using OpenCC (Simplified → Traditional Taiwan)
opencc -c s2twp.json -i zh.json -o zh-TW.json

# Step 3: Manual review and correction
# Focus areas:
# - Technical terms: 数据库 → 資料庫, 用户 → 使用者
# - Action verbs: 删除 → 刪除, 添加 → 新增
# - Status terms: 成功 → 成功, 失败 → 失敗
```

**Key Translation Rules**:
| Simplified | Traditional | Taiwan Term | Context |
|------------|-------------|-------------|---------|
| 用户 | 用戶 | 使用者 | User (prefer 使用者) |
| 数据库 | 數據庫 | 資料庫 | Database (use 資料庫) |
| 网络 | 網絡 | 網路 | Network (use 網路) |
| 信息 | 信息 | 資訊 | Information (use 資訊) |
| 软件 | 軟件 | 軟體 | Software (use 軟體) |
| 文件 | 文件 | 檔案 | File (prefer 檔案) |
| 登录 | 登錄 | 登入 | Login (use 登入) |
| 注册 | 註冊 | 註冊 | Register |
| 删除 | 刪除 | 刪除 | Delete |
| 添加 | 添加 | 新增 | Add (prefer 新增) |

**Example zh-TW.json structure**:
```json
{
  "common": {
    "success": "成功",
    "failed": "失敗",
    "createSuccess": "建立成功",
    "updateSuccess": "更新成功",
    "deleteSuccess": "刪除成功",
    "operationSuccess": "操作成功",
    "operationFailed": "操作失敗",
    "databaseError": "資料庫錯誤，請稍後再試",
    "permissionDeny": "使用者無權限存取此介面"
  },
  "login": {
    "loginSuccessTitle": "登入成功",
    "loginSuccessDescription": "歡迎回來",
    "wrongUsernameOrPassword": "使用者名稱或密碼錯誤",
    "userBanned": "使用者已被停權",
    "loginFailed": "登入失敗"
  },
  "user": {
    "userNotFound": "使用者不存在",
    "usernameExist": "使用者名稱已存在",
    "emailExist": "電子郵件已被使用"
  }
}
```

#### Task 1.2: Update i18n Loader
**File**: `api/internal/i18n/translator.go` (create if not exists)

```go
package i18n

import (
    "encoding/json"
    "fmt"
    "io/fs"
    "strings"
    "sync"
)

type Translator struct {
    locales map[string]map[string]interface{}
    mu      sync.RWMutex
}

var (
    globalTranslator *Translator
    once             sync.Once
)

// GetTranslator returns singleton translator instance
func GetTranslator() *Translator {
    once.Do(func() {
        globalTranslator = &Translator{
            locales: make(map[string]map[string]interface{}),
        }
        globalTranslator.LoadAll()
    })
    return globalTranslator
}

// LoadAll loads all language files from embedded FS
func (t *Translator) LoadAll() error {
    supportedLocales := []string{"zh", "en", "zh-TW"}

    for _, locale := range supportedLocales {
        fileName := fmt.Sprintf("locale/%s.json", locale)
        data, err := fs.ReadFile(LocaleFS, fileName)
        if err != nil {
            return fmt.Errorf("failed to read %s: %w", fileName, err)
        }

        var messages map[string]interface{}
        if err := json.Unmarshal(data, &messages); err != nil {
            return fmt.Errorf("failed to parse %s: %w", fileName, err)
        }

        t.mu.Lock()
        t.locales[locale] = messages
        t.mu.Unlock()
    }

    return nil
}

// Trans translates a key with optional fallback
func (t *Translator) Trans(locale, key string) string {
    t.mu.RLock()
    defer t.mu.RUnlock()

    // Try requested locale
    if msg := t.getNestedKey(locale, key); msg != "" {
        return msg
    }

    // Fallback to simplified Chinese
    if locale == "zh-TW" {
        if msg := t.getNestedKey("zh", key); msg != "" {
            return msg
        }
    }

    // Ultimate fallback to English
    if msg := t.getNestedKey("en", key); msg != "" {
        return msg
    }

    // Return key itself if all fails
    return key
}

// getNestedKey retrieves value from nested map using dot notation
func (t *Translator) getNestedKey(locale, key string) string {
    messages, ok := t.locales[locale]
    if !ok {
        return ""
    }

    parts := strings.Split(key, ".")
    var current interface{} = messages

    for _, part := range parts {
        m, ok := current.(map[string]interface{})
        if !ok {
            return ""
        }
        current, ok = m[part]
        if !ok {
            return ""
        }
    }

    if str, ok := current.(string); ok {
        return str
    }

    return ""
}

// NormalizeLocale converts Accept-Language to internal locale code
func NormalizeLocale(acceptLang string) string {
    acceptLang = strings.ToLower(acceptLang)

    // Handle zh-TW, zh-Hant, zh-Hant-TW
    if strings.Contains(acceptLang, "zh-tw") ||
       strings.Contains(acceptLang, "zh-hant") {
        return "zh-TW"
    }

    // Handle zh-CN, zh-Hans, zh
    if strings.Contains(acceptLang, "zh") {
        return "zh"
    }

    // Handle en, en-US, en-GB
    if strings.Contains(acceptLang, "en") {
        return "en"
    }

    // Default to simplified Chinese
    return "zh"
}
```

#### Task 1.3: Update Error Handler
**File**: `rpc/internal/utils/dberrorhandler/error_handler.go`

No changes needed - already uses i18n keys. Verify it imports from correct package.

#### Task 1.4: Add Unit Tests
**File**: `api/internal/i18n/translator_test.go`

```go
package i18n

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTranslator_Trans_ZhTW(t *testing.T) {
    trans := GetTranslator()

    tests := []struct {
        name     string
        locale   string
        key      string
        expected string
    }{
        {
            name:     "Traditional Chinese success message",
            locale:   "zh-TW",
            key:      "common.success",
            expected: "成功",
        },
        {
            name:     "Traditional Chinese login error",
            locale:   "zh-TW",
            key:      "login.wrongUsernameOrPassword",
            expected: "使用者名稱或密碼錯誤",
        },
        {
            name:     "Fallback to simplified Chinese",
            locale:   "zh-TW",
            key:      "some.missing.key",
            expected: "some.missing.key", // Returns key if not found
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := trans.Trans(tt.locale, tt.key)
            assert.Equal(t, tt.expected, result)
        })
    }
}

func TestNormalizeLocale(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"zh-TW", "zh-TW"},
        {"zh-Hant", "zh-TW"},
        {"zh-Hant-TW", "zh-TW"},
        {"zh-CN", "zh"},
        {"zh", "zh"},
        {"en-US", "en"},
        {"en", "en"},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            result := NormalizeLocale(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### Phase 2: Frontend Implementation (8-10 hours)

#### Task 2.1: Create zh_TW Language File
**File**: `src/locales/lang/zh_TW.ts` (path may vary based on Vben5 structure)

```typescript
export default {
  common: {
    okText: '確定',
    closeText: '關閉',
    cancelText: '取消',
    loadingText: '載入中...',
    saveText: '儲存',
    delText: '刪除',
    resetText: '重設',
    searchText: '搜尋',
    queryText: '查詢',

    inputText: '請輸入',
    chooseText: '請選擇',

    redo: '重新整理',
    back: '返回',

    light: '淺色主題',
    dark: '深色主題',
  },

  login: {
    backSignIn: '返回登入',
    signInFormTitle: '登入',
    mobileSignInFormTitle: '行動裝置登入',
    qrSignInFormTitle: 'QR Code 登入',
    signUpFormTitle: '註冊',
    forgetFormTitle: '重設密碼',

    signInTitle: '後台管理系統',
    signInDesc: '輸入您的個人資訊以登入',
    policy: '我同意服務條款與隱私權政策',
    scanSign: `掃描 QR Code 以登入`,

    loginButton: '登入',
    registerButton: '註冊',
    rememberMe: '記住我',
    forgetPassword: '忘記密碼？',
    otherSignIn: '其他登入方式',

    // Placeholder
    accountPlaceholder: '請輸入使用者名稱',
    passwordPlaceholder: '請輸入密碼',
    smsPlaceholder: '請輸入簡訊驗證碼',
    mobilePlaceholder: '請輸入手機號碼',
    policyPlaceholder: '註冊即表示同意',
    diffPwd: '兩次密碼輸入不一致',

    userName: '使用者名稱',
    password: '密碼',
    confirmPassword: '確認密碼',
    email: '電子郵件',
    smsCode: '簡訊驗證碼',
    mobile: '手機號碼',
  },

  route: {
    dashboard: '控制台',
    about: '關於',

    // System Management
    system: '系統管理',
    user: '使用者管理',
    role: '角色管理',
    menu: '選單管理',
    dept: '部門管理',
    dict: '字典管理',
    api: 'API 管理',
    token: 'Token 管理',

    // Other modules
    moduleStore: '模組商店',
  },

  user: {
    title: '使用者管理',
    username: '使用者名稱',
    nickname: '暱稱',
    email: '電子郵件',
    mobile: '手機號碼',
    status: '狀態',
    createTime: '建立時間',
    updateTime: '更新時間',

    addUser: '新增使用者',
    editUser: '編輯使用者',
    deleteUser: '刪除使用者',

    statusNormal: '正常',
    statusBanned: '停權',
  },

  role: {
    title: '角色管理',
    roleName: '角色名稱',
    roleCode: '角色代碼',
    roleDesc: '角色描述',
    sort: '排序',
    status: '狀態',

    addRole: '新增角色',
    editRole: '編輯角色',
    deleteRole: '刪除角色',
    assignPermission: '分配權限',
  },

  menu: {
    title: '選單管理',
    menuName: '選單名稱',
    menuType: '選單類型',
    icon: '圖示',
    path: '路徑',
    component: '元件',
    permission: '權限標識',
    sort: '排序',
    visible: '顯示',
    status: '狀態',

    menuTypeDir: '目錄',
    menuTypeMenu: '選單',
    menuTypeButton: '按鈕',
  },
};
```

#### Task 2.2: Register zh_TW in i18n
**File**: `src/locales/index.ts`

```typescript
import { createI18n } from 'vue-i18n';
import zh_CN from './lang/zh_CN';
import en from './lang/en';
import zh_TW from './lang/zh_TW'; // NEW

const messages = {
  zh_CN,
  en,
  zh_TW, // NEW
};

// Get locale from localStorage or browser
const getLocale = (): string => {
  const localLocale = localStorage.getItem('locale');
  if (localLocale && ['zh_CN', 'en', 'zh_TW'].includes(localLocale)) {
    return localLocale;
  }

  // Browser locale detection
  const browserLang = navigator.language;
  if (browserLang.includes('zh-TW') || browserLang.includes('zh-Hant')) {
    return 'zh_TW';
  }
  if (browserLang.includes('zh')) {
    return 'zh_CN';
  }
  if (browserLang.includes('en')) {
    return 'en';
  }

  return 'zh_CN'; // Default
};

export const i18n = createI18n({
  legacy: false,
  locale: getLocale(),
  fallbackLocale: 'zh_CN',
  messages,
  globalInjection: true,
});

export default i18n;
```

#### Task 2.3: Update Language Selector Component
**File**: `src/components/LanguageSelector/index.vue`

```vue
<template>
  <a-dropdown :trigger="['click']">
    <div class="language-selector">
      <GlobalOutlined />
      <span>{{ currentLanguageName }}</span>
    </div>
    <template #overlay>
      <a-menu @click="handleLanguageChange">
        <a-menu-item key="zh_CN">
          <span>简体中文</span>
        </a-menu-item>
        <a-menu-item key="en">
          <span>English</span>
        </a-menu-item>
        <a-menu-item key="zh_TW">
          <span>繁體中文（台灣）</span>
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { GlobalOutlined } from '@ant-design/icons-vue';
import { useUserStore } from '@/store/modules/user';

const { locale } = useI18n();
const userStore = useUserStore();

const languageMap = {
  zh_CN: '简体中文',
  en: 'English',
  zh_TW: '繁體中文',
};

const currentLanguageName = computed(() => {
  return languageMap[locale.value as keyof typeof languageMap] || '简体中文';
});

const handleLanguageChange = async ({ key }: { key: string }) => {
  // Update i18n locale
  locale.value = key;

  // Save to localStorage
  localStorage.setItem('locale', key);

  // Save to user profile (API call)
  try {
    await userStore.updateUserLocale(key);
  } catch (error) {
    console.error('Failed to save locale preference:', error);
  }

  // Update Ant Design locale
  updateAntdLocale(key);

  // Reload current route to apply translations
  window.location.reload();
};

const updateAntdLocale = (locale: string) => {
  // This will be handled in App.vue with ConfigProvider
};
</script>

<style scoped lang="less">
.language-selector {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 8px 12px;

  &:hover {
    background-color: rgba(0, 0, 0, 0.04);
  }
}
</style>
```

#### Task 2.4: Update Ant Design ConfigProvider
**File**: `src/App.vue`

```vue
<template>
  <a-config-provider :locale="antdLocale">
    <router-view />
  </a-config-provider>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import enUS from 'ant-design-vue/es/locale/en_US';
import zhTW from 'ant-design-vue/es/locale/zh_TW'; // NEW

const { locale } = useI18n();

const antdLocale = computed(() => {
  const localeMap = {
    zh_CN: zhCN,
    en: enUS,
    zh_TW: zhTW, // NEW
  };
  return localeMap[locale.value as keyof typeof localeMap] || zhCN;
});
</script>
```

### Phase 3: Database & User Preference (4-6 hours)

#### Task 3.1: Update Ent Schema
**File**: `rpc/ent/schema/user.go`

```go
// Add to User schema Fields()
func (User) Fields() []ent.Field {
    return []ent.Field{
        // ... existing fields ...

        field.String("locale").
            Comment("User's preferred language: zh-CN, en, zh-TW").
            Optional().
            Nillable().
            MaxLen(10).
            Default("zh-CN"),
    }
}
```

#### Task 3.2: Generate Ent Code
```bash
make gen-ent
```

#### Task 3.3: Update User RPC Logic
**File**: `rpc/internal/logic/user/update_user_logic.go`

```go
// Add locale field handling in UpdateUser
func (l *UpdateUserLogic) UpdateUser(in *core.UserInfo) (*core.BaseResp, error) {
    query := l.svcCtx.DB.User.UpdateOneID(in.Id).
        SetUsername(in.Username).
        SetNickname(in.Nickname).
        SetEmail(in.Email)
        // ... other fields ...

    // Handle locale update if provided
    if in.Locale != nil {
        query.SetLocale(*in.Locale)
    }

    err := query.Exec(l.ctx)
    if err != nil {
        return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
    }

    return &core.BaseResp{Msg: l.svcCtx.Trans.Trans(l.ctx, "common.updateSuccess")}, nil
}
```

#### Task 3.4: Add Proto Field
**File**: `rpc/desc/user.proto`

```protobuf
message UserInfo {
  uint64 id = 1;
  string username = 2;
  // ... other fields ...
  optional string locale = 20; // NEW: User's preferred language
}
```

Run `make gen-rpc` after updating proto.

#### Task 3.5: Update API Layer
**File**: `api/desc/core/user.api`

```go
type UserInfo {
    Id        uint64  `json:"id"`
    Username  string  `json:"username"`
    // ... other fields ...
    Locale    *string `json:"locale,optional"` // NEW
}
```

Run `make gen-api` after updating.

### Phase 4: Testing & QA (4-6 hours)

#### Task 4.1: Unit Tests
```bash
# Backend tests
cd api/internal/i18n
go test -v -cover

# Frontend tests
cd frontend
npm run test:unit
```

#### Task 4.2: E2E Tests
**File**: `tests/e2e/i18n.spec.ts`

```typescript
import { test, expect } from '@playwright/test';

test.describe('i18n Traditional Chinese', () => {
  test('should switch to Traditional Chinese', async ({ page }) => {
    await page.goto('http://localhost:3000/login');

    // Click language selector
    await page.click('.language-selector');

    // Select Traditional Chinese
    await page.click('text=繁體中文（台灣）');

    // Verify login button text changed
    await expect(page.locator('button[type="submit"]')).toHaveText('登入');

    // Verify placeholder text
    await expect(page.locator('input[placeholder]').first()).toHaveAttribute(
      'placeholder',
      '請輸入使用者名稱'
    );
  });

  test('should persist language preference', async ({ page, context }) => {
    await page.goto('http://localhost:3000/login');

    // Switch to zh-TW
    await page.click('.language-selector');
    await page.click('text=繁體中文（台灣）');

    // Reload page
    await page.reload();

    // Verify language persisted
    const localStorage = await page.evaluate(() => window.localStorage.getItem('locale'));
    expect(localStorage).toBe('zh_TW');
  });

  test('should display Traditional Chinese error messages', async ({ page }) => {
    await page.goto('http://localhost:3000/login');

    // Switch to zh-TW
    await page.click('.language-selector');
    await page.click('text=繁體中文（台灣）');

    // Submit with wrong credentials
    await page.fill('input[name="username"]', 'wronguser');
    await page.fill('input[name="password"]', 'wrongpass');
    await page.click('button[type="submit"]');

    // Verify error message in Traditional Chinese
    await expect(page.locator('.ant-message')).toContainText('使用者名稱或密碼錯誤');
  });
});
```

#### Task 4.3: Manual QA Checklist
```markdown
## Manual Testing Checklist

### Login Flow
- [ ] Language selector visible on login page
- [ ] Can select 繁體中文（台灣）
- [ ] All login form labels in Traditional Chinese
- [ ] Error messages in Traditional Chinese
- [ ] Remember language after logout/login

### Backend UI
- [ ] All menu items in Traditional Chinese
- [ ] Breadcrumbs in Traditional Chinese
- [ ] Table column headers in Traditional Chinese
- [ ] Action buttons (新增, 編輯, 刪除) in Traditional Chinese
- [ ] Modal dialogs in Traditional Chinese
- [ ] Form validation messages in Traditional Chinese
- [ ] Success/error notifications in Traditional Chinese

### Data Display
- [ ] Dates formatted as YYYY/MM/DD
- [ ] Numbers with comma separator (1,234)
- [ ] Time in 24-hour format
- [ ] Status labels in Traditional Chinese

### User Preference
- [ ] Language preference saved to database
- [ ] Preference persists across devices
- [ ] Can change language from user profile

### Edge Cases
- [ ] Missing translation key shows fallback (zh-CN)
- [ ] Browser locale zh-TW auto-selects Traditional Chinese
- [ ] No console errors when using zh-TW
- [ ] No layout breaking with longer Chinese text
```

## Performance Considerations

### Bundle Size Impact
- **zh-TW.json**: ~5KB (gzipped)
- **zh_TW.ts**: ~8KB (gzipped)
- **Total Impact**: <15KB additional bundle size

### Runtime Performance
- **Language switch**: Target <100ms
- **Translation lookup**: O(1) with map structure
- **Memory footprint**: ~50KB additional for zh-TW locale

### Optimization Strategies
1. **Lazy loading**: Load zh-TW only when selected
2. **Code splitting**: Separate zh-TW into async chunk
3. **Caching**: Cache translation files with service worker
4. **CDN**: Serve language files from CDN with long cache

## Deployment Strategy

### Rollout Plan
1. **Week 1**: Deploy to development environment
2. **Week 2**: Internal testing with Taiwan team members
3. **Week 3**: Beta release to 10 Taiwan pilot users
4. **Week 4**: Production release to all users

### Feature Flag
```go
// api/etc/core.yaml
I18n:
  EnableTraditionalChinese: true  # Feature flag
```

### Monitoring
- Track language selection distribution
- Monitor translation key misses
- Alert on fallback usage spikes
- User feedback collection

## Rollback Plan

If critical issues found:
1. Set `EnableTraditionalChinese: false` in config
2. Redeploy API service (no database rollback needed)
3. Frontend will hide zh-TW option
4. Existing users fall back to zh-CN

## Documentation Updates

### User Documentation
- [ ] Update user manual with language switching guide
- [ ] Create screenshots showing zh-TW interface
- [ ] Document Taiwan-specific terminology choices

### Developer Documentation
- [ ] Update CLAUDE.md with i18n workflow
- [ ] Document translation key naming convention
- [ ] Add troubleshooting guide for i18n issues

### API Documentation
- [ ] Update Swagger with locale field
- [ ] Document Accept-Language header usage
- [ ] Add i18n examples to API guide

## Success Metrics

### Technical Metrics
- ✅ Zero translation key errors in production
- ✅ Language switch < 100ms
- ✅ 100% test coverage for i18n code
- ✅ Zero accessibility violations

### Business Metrics
- ✅ 30%+ users select zh-TW (Taiwan market)
- ✅ 60% reduction in language-related support tickets
- ✅ 4.5+/5.0 terminology accuracy rating from Taiwan users

## Timeline

| Phase | Duration | Dependencies |
|-------|----------|--------------|
| Backend API | 4-6 hours | OpenCC tool |
| Frontend | 8-10 hours | Backend complete |
| Database | 4-6 hours | Ent schema design |
| Testing | 4-6 hours | All phases complete |
| **Total** | **20-28 hours** | - |

## Team Assignment

| Role | Responsibility | Hours |
|------|----------------|-------|
| Backend Developer | zh-TW.json, API i18n logic | 6h |
| Frontend Developer | Vue i18n, language selector | 10h |
| Full-stack Developer | Database schema, user preference | 6h |
| QA Engineer | Testing, manual QA | 4h |
| Taiwan Native Speaker | Translation review, QA | 4h |

---

**Next Action**: Review this plan with team, get approval, create sprint tasks
