# Traditional Chinese (zh-TW) E2E Test Plan

## 測試環境
- **Backend**: simple-admin-core (Go-Zero)
- **Frontend**: simple-admin-vben5-ui (Vue 3 + Vben5)
- **測試工具**: Manual Testing (可選 Playwright/Cypress)

## 測試案例

### TC-001: 語言選擇器顯示

**優先級**: P0
**目標**: 驗證語言選擇器顯示所有語言選項

**前置條件**:
- 系統已部署並可訪問
- 使用者未登入

**測試步驟**:
1. 開啟瀏覽器，訪問系統首頁
2. 點擊右上角語言選擇器圖示

**預期結果**:
- ✅ 語言選擇器下拉選單顯示
- ✅ 包含 3 個選項：
  - 简体中文
  - 繁體中文（台灣）
  - English

---

### TC-002: 切換到繁體中文

**優先級**: P0
**目標**: 驗證語言切換功能正常運作

**前置條件**:
- 系統顯示預設語言（简体中文或 English）

**測試步驟**:
1. 點擊語言選擇器
2. 選擇「繁體中文（台灣）」
3. 等待頁面重新載入

**預期結果**:
- ✅ 頁面自動刷新
- ✅ UI 文字全部顯示為繁體中文
- ✅ localStorage 保存 `locale: "zh-TW"`

---

### TC-003: 登入頁面繁體中文顯示

**優先級**: P0
**目標**: 驗證登入頁面所有文字正確顯示繁體中文

**前置條件**:
- 語言已切換至繁體中文

**測試步驟**:
1. 檢查登入頁面各項文字

**預期結果**:
- ✅ 標題顯示「登入」（非「登录」）
- ✅ 帳號輸入框 placeholder：「請輸入帳號」
- ✅ 密碼輸入框 placeholder：「請輸入密碼」
- ✅ 登入按鈕：「登入」
- ✅ 驗證碼：「取得驗證碼」
- ✅ 忘記密碼：「忘記密碼」

---

### TC-004: 登入後主選單繁體中文顯示

**優先級**: P0
**目標**: 驗證主選單項目使用台灣用語

**前置條件**:
- 以管理員身份登入系統
- 語言設定為繁體中文

**測試步驟**:
1. 檢查左側選單各項目名稱

**預期結果**:
- ✅ 控制台（Dashboard）
- ✅ 使用者管理（非「用户管理」）
- ✅ 角色管理
- ✅ 選單管理（非「菜单管理」）
- ✅ API 管理
- ✅ 字典管理
- ✅ 部門管理

---

### TC-005: CRUD 操作繁體中文訊息

**優先級**: P0
**目標**: 驗證新增、編輯、刪除操作的訊息顯示繁體中文

**前置條件**:
- 已登入系統
- 語言設定為繁體中文

**測試步驟**:
1. 進入「使用者管理」頁面
2. 點擊「新增」按鈕
3. 填寫使用者資訊並儲存
4. 編輯該使用者
5. 刪除該使用者

**預期結果**:
- ✅ 新增按鈕：「新增」
- ✅ 編輯按鈕：「編輯」
- ✅ 刪除按鈕：「刪除」
- ✅ 成功訊息：
  - 「新建成功」（非「新建成功」）
  - 「更新成功」
  - 「刪除成功」
- ✅ 表單欄位標籤：
  - 使用者名稱、電子郵件、暱稱、大頭貼

---

### TC-006: 錯誤訊息繁體中文顯示

**優先級**: P0
**目標**: 驗證錯誤訊息使用繁體中文和台灣用語

**前置條件**:
- 已登入系統
- 語言設定為繁體中文

**測試步驟**:
1. 嘗試新增重複的使用者名稱
2. 嘗試編輯超級管理員（權限不足）
3. 嘗試刪除系統內建角色

**預期結果**:
- ✅ 錯誤訊息顯示繁體中文：
  - 「操作失敗」
  - 「使用者無權限存取此介面」（非「接口」）
  - 「資料庫錯誤」（非「数据库错误」）
  - 「目標不存在」

---

### TC-007: 表格與分頁器繁體中文

**優先級**: P1
**目標**: 驗證 Ant Design 元件顯示繁體中文

**前置條件**:
- 已登入系統
- 語言設定為繁體中文
- 進入任一資料列表頁面

**測試步驟**:
1. 檢查表格元件文字
2. 檢查分頁器文字
3. 檢查日期選擇器

**預期結果**:
- ✅ 分頁器：「共 XX 條資料」
- ✅ 每頁顯示：「每頁顯示」
- ✅ 日期選擇器：
  - 月份：一月、二月...十二月
  - 星期：週日、週一...週六
  - 今天、確定、清空
- ✅ 篩選器：「篩選」、「重設」
- ✅ 搜尋：「搜尋」、「查詢」

---

### TC-008: 語言偏好持久化

**優先級**: P0
**目標**: 驗證語言設定在登出/重新登入後保持

**前置條件**:
- 已登入系統
- 語言設定為繁體中文

**測試步驟**:
1. 確認當前語言為繁體中文
2. 登出系統
3. 重新登入
4. 檢查語言設定

**預期結果**:
- ✅ 登入後系統自動顯示繁體中文
- ✅ localStorage 保存 `locale: "zh-TW"`
- ✅ 語言選擇器顯示當前選中「繁體中文（台灣）」

---

### TC-009: API 回應語言切換

**優先級**: P1
**目標**: 驗證 API 回應根據 Accept-Language 返回繁體中文

**前置條件**:
- Backend 服務運行中
- 語言設定為繁體中文

**測試步驟**:
1. 開啟瀏覽器開發者工具 (F12)
2. 切換到 Network 標籤
3. 執行任意 API 操作（如查詢使用者列表）
4. 檢查 Request Headers 和 Response

**預期結果**:
- ✅ Request Header 包含：`Accept-Language: zh-TW`
- ✅ API 錯誤訊息顯示繁體中文：
  ```json
  {
    "code": 400,
    "msg": "使用者無權限存取此介面"
  }
  ```
- ✅ 所有 API 回應訊息使用台灣用語

---

### TC-010: 跨頁面語言一致性

**優先級**: P1
**目標**: 驗證所有頁面語言顯示一致

**前置條件**:
- 已登入系統
- 語言設定為繁體中文

**測試步驟**:
1. 依序訪問以下頁面：
   - 控制台
   - 使用者管理
   - 角色管理
   - 選單管理
   - API 管理
   - 字典管理
   - 部門管理
   - 個人資料頁面
2. 檢查每個頁面的文字顯示

**預期結果**:
- ✅ 所有頁面統一使用繁體中文
- ✅ 術語一致（如「使用者」非「用戶」）
- ✅ 按鈕文字一致（確定、取消、儲存、刪除）
- ✅ 無簡體中文殘留

---

## 測試執行記錄

### 執行資訊
- **測試日期**: 2025-10-08
- **測試人員**: @qa
- **環境**: Development
- **版本**:
  - Backend: commit `be16dc9`
  - Frontend: commit `8bc94a930`

### 測試結果摘要

| 測試案例 | 狀態 | 備註 |
|---------|------|------|
| TC-001  | ⚠️ 待測試 | 需部署環境 |
| TC-002  | ⚠️ 待測試 | 需部署環境 |
| TC-003  | ⚠️ 待測試 | 需部署環境 |
| TC-004  | ⚠️ 待測試 | 需部署環境 |
| TC-005  | ⚠️ 待測試 | 需部署環境 |
| TC-006  | ⚠️ 待測試 | 需部署環境 |
| TC-007  | ⚠️ 待測試 | 需部署環境 |
| TC-008  | ⚠️ 待測試 | 需部署環境 |
| TC-009  | ⚠️ 待測試 | 需部署環境 |
| TC-010  | ⚠️ 待測試 | 需部署環境 |

### 發現的問題
_測試執行後填寫_

---

## 自動化測試腳本（可選）

如需自動化測試，可使用以下 Playwright 範例：

```typescript
// tests/e2e/i18n.spec.ts
import { test, expect } from '@playwright/test';

test.describe('Traditional Chinese (zh-TW) i18n', () => {

  test('TC-001: Language selector displays all options', async ({ page }) => {
    await page.goto('http://localhost:5173');

    // Click language selector
    await page.click('[data-testid="language-toggle"]');

    // Check options
    await expect(page.locator('text=简体中文')).toBeVisible();
    await expect(page.locator('text=繁體中文（台灣）')).toBeVisible();
    await expect(page.locator('text=English')).toBeVisible();
  });

  test('TC-002: Switch to Traditional Chinese', async ({ page }) => {
    await page.goto('http://localhost:5173');

    // Switch language
    await page.click('[data-testid="language-toggle"]');
    await page.click('text=繁體中文（台灣）');

    // Wait for reload
    await page.waitForLoadState('networkidle');

    // Verify login button text
    await expect(page.locator('button:has-text("登入")')).toBeVisible();
  });

  test('TC-003: Login page displays zh-TW', async ({ page }) => {
    await page.goto('http://localhost:5173');
    await page.click('[data-testid="language-toggle"]');
    await page.click('text=繁體中文（台灣）');

    await expect(page.locator('text=請輸入帳號')).toBeVisible();
    await expect(page.locator('text=請輸入密碼')).toBeVisible();
    await expect(page.locator('button:has-text("登入")')).toBeVisible();
  });

  // Add more test cases...
});
```

---

## 結論

此測試計劃涵蓋繁體中文 i18n 功能的所有關鍵場景。建議：

1. **優先執行 P0 測試案例**
2. **在 Development 環境先測試**
3. **修復發現的問題後，再部署到 Production**
4. **記錄所有術語問題，供 [ZH-TW-010] 人工 QA 參考**
