# Feature Specification: Traditional Chinese (zh-TW) Support

**Feature Branch**: `001-zh-tw-i18n`
**Created**: 2025-10-07
**Status**: Draft
**Input**: Add Traditional Chinese language support for Simple Admin Core backend and frontend to serve Taiwan market users

## User Scenarios & Testing (mandatory)

### User Story 1 - Taiwan Users Can Select Traditional Chinese UI (Priority: P1)
A Taiwan user visits the Simple Admin backend for the first time. Before logging in, they see a language selector showing "简体中文", "English", and "繁體中文（台灣）". They select Traditional Chinese, and the login page immediately displays all text in Traditional Chinese with Taiwan-specific terminology (e.g., "登入" instead of "登录", "網路" instead of "网络").

**Why this priority**: Taiwan represents approximately 50% of potential Asia-Pacific users. Traditional Chinese support is mandatory for Taiwan government and enterprise procurement requirements.

**Independent Test**:
1. Deploy system with zh-TW language files
2. Access login page without authentication
3. Verify language selector includes zh-TW option
4. Verify all UI elements render in Traditional Chinese

**Acceptance Scenarios**:
1. **Given** user on login page, **When** select zh-TW from language dropdown, **Then** login button displays "登入", username field shows "使用者名稱", password field shows "密碼"
2. **Given** user logged in with zh-CN selected, **When** switch to zh-TW, **Then** all menus, breadcrumbs, and buttons update to Traditional Chinese without page reload
3. **Given** user selected zh-TW, **When** logout and login again, **Then** system remembers zh-TW preference

### User Story 2 - Error Messages Display Taiwan Terminology (Priority: P1)
When a Taiwan user encounters system errors (database errors, permission denied, validation failures), all error messages display in Traditional Chinese using Taiwan-specific technical terms (e.g., "資料庫" not "数据库", "權限" not "权限").

**Why this priority**: Error messages are critical for user support and troubleshooting. Simplified Chinese technical terms confuse Taiwan IT staff.

**Independent Test**:
1. Trigger various error scenarios (wrong password, no permission, network timeout)
2. Verify all error messages use Traditional Chinese
3. Verify terminology matches Taiwan IT conventions

**Acceptance Scenarios**:
1. **Given** user enters wrong password, **When** submit login form, **Then** error message displays "使用者名稱或密碼錯誤" (not "用户名或密码错误")
2. **Given** user without permission access protected API, **When** request sent, **Then** error shows "使用者無權限存取此介面" (not "用户无权限访问此接口")
3. **Given** database connection fails, **When** any operation attempted, **Then** error displays "資料庫錯誤，請稍後再試" (not "数据库错误，请稍后再试")

### User Story 3 - Menu and Navigation Use Taiwan Terms (Priority: P2)
The system menu structure (dashboard, system management, user management, role management, etc.) displays using Taiwan business terminology that matches local conventions and expectations.

**Why this priority**: Consistent terminology improves user adoption and reduces training time.

**Independent Test**:
1. Login as admin user with zh-TW selected
2. Navigate through all menu items
3. Verify all menu labels use Taiwan terminology
4. Compare with Taiwan government digital service terminology standards

**Acceptance Scenarios**:
1. **Given** admin user logged in with zh-TW, **When** view sidebar menu, **Then** displays "控制台" (dashboard), "系統管理" (system management), "使用者管理" (user management)
2. **Given** viewing user list page, **When** check column headers, **Then** displays "使用者名稱", "電子郵件", "狀態", "建立時間" (not mainland terms)
3. **Given** on any CRUD page, **When** check action buttons, **Then** displays "新增", "編輯", "刪除", "搜尋" (Taiwan conventions)

### User Story 4 - Date and Number Formats Follow Taiwan Standards (Priority: P2)
Dates display in Taiwan standard format (YYYY/MM/DD or 民國年), numbers use Taiwan punctuation (comma for thousands), and time displays in 24-hour format.

**Why this priority**: Date format mismatches cause user confusion and data entry errors.

**Independent Test**:
1. Create test data with various dates and numbers
2. View data in zh-TW locale
3. Verify formatting matches Taiwan standards

**Acceptance Scenarios**:
1. **Given** viewing user creation date, **When** locale is zh-TW, **Then** date displays as "2025/10/07 14:30" or "民國114/10/07 14:30"
2. **Given** viewing numeric data like user count, **When** value exceeds 1000, **Then** displays as "1,234" (comma separator)
3. **Given** system timestamps, **When** displayed in zh-TW, **Then** uses 24-hour format "14:30:45" not "2:30:45 PM"

### Edge Cases
- What happens when zh-TW translation key is missing? System should fallback to zh-CN, then en
- How does system handle mixed content (user-generated data in simplified Chinese)? Display as-is, don't auto-convert
- What if user's browser locale is zh-TW but they never explicitly selected language? Use browser locale as default
- How to handle dynamic content from APIs that don't support i18n? Show in English with [EN] prefix
- What about third-party modules (FMS, Job, MCMS) without zh-TW support? Display warning banner "部分模組尚未支援繁體中文"

## Requirements (mandatory)

### Functional Requirements

#### Backend API Layer
- **FR-001**: System MUST provide zh-TW.json language file in `api/internal/i18n/locale/` containing all translation keys from zh.json
- **FR-002**: Backend MUST support Accept-Language header value "zh-TW" and "zh-Hant-TW"
- **FR-003**: i18n middleware MUST detect zh-TW requests and return Traditional Chinese responses
- **FR-004**: All error messages MUST use i18n keys (e.g., i18n.DatabaseError) not hardcoded strings
- **FR-005**: API response errors MUST include translated message based on request language

#### Frontend Layer
- **FR-006**: Frontend MUST include zh_TW.ts language file in locales directory
- **FR-007**: Language selector MUST display three options: "简体中文", "English", "繁體中文（台灣）"
- **FR-008**: System MUST persist language preference in localStorage and user profile
- **FR-009**: All Vue components MUST use $t() function, not hardcoded strings
- **FR-010**: Ant Design Vue components MUST load zh_TW locale for date pickers, tables, modals

#### Data Layer
- **FR-011**: Database menu records MUST continue using i18n keys (e.g., "route.dashboard") not translated text
- **FR-012**: User table MUST store language preference column (user.locale)
- **FR-013**: System MUST support language preference priority: user profile > localStorage > browser locale > default (zh-CN)

#### Translation Quality
- **FR-014**: All translations MUST use Taiwan terminology standards (e.g., "軟體" not "软件", "網路" not "网络")
- **FR-015**: Technical terms MUST match Taiwan Ministry of Education terminology database
- **FR-016**: UI text MUST use polite form suitable for business context
- **FR-017**: Translation coverage MUST reach 100% for core module, 80%+ for optional modules

#### Testing & Validation
- **FR-018**: System MUST include automated tests verifying zh-TW language file format validity
- **FR-019**: E2E tests MUST cover language switching without page reload
- **FR-020**: Translation keys MUST have no duplicates or conflicts

### Key Entities

#### LanguageFile (Backend)
- **Location**: `api/internal/i18n/locale/zh-TW.json`
- **Format**: JSON with nested keys matching zh.json structure
- **Size**: Estimated 314 lines (matching zh.json)
- **Encoding**: UTF-8 with BOM
- **Example Structure**:
```json
{
  "common": {
    "success": "成功",
    "failed": "失敗",
    "databaseError": "資料庫錯誤"
  },
  "login": {
    "loginSuccessTitle": "登入成功",
    "wrongUsernameOrPassword": "使用者名稱或密碼錯誤"
  }
}
```

#### LanguageFile (Frontend)
- **Location**: `src/locales/lang/zh_TW.ts`
- **Framework**: vue-i18n compatible
- **Dependencies**: Ant Design Vue zh_TW locale
- **Coverage**: All routes, menus, forms, validation messages, notifications

#### UserPreference
- **Table**: sys_users
- **Column**: locale (VARCHAR(10), nullable, default: NULL)
- **Values**: "zh-CN" | "en" | "zh-TW"
- **Migration**: Add column via Ent schema update

#### LocaleConfig
- **Backend Config**: `api/etc/core.yaml` - add SupportedLocales list
- **Frontend Config**: `src/locales/index.ts` - register zh_TW
- **Fallback Chain**: zh-TW → zh-CN → en

## Success Criteria (mandatory)

### Measurable Outcomes
- **SC-001**: Translation coverage reaches 100% for zh-TW.json (all 200+ keys translated)
- **SC-002**: Language switching completes in <100ms without page reload
- **SC-003**: Zero console errors or warnings when using zh-TW locale
- **SC-004**: All automated tests pass with zh-TW locale enabled

### User Satisfaction Metrics
- **SC-005**: Taiwan pilot users rate terminology accuracy ≥4.5/5.0
- **SC-006**: Zero user-reported translation errors in core workflows (login, user management, role management)
- **SC-007**: Documentation includes zh-TW setup guide for administrators

### Business Metrics
- **SC-008**: Feature enables Taiwan market entry (prerequisite for local partnerships)
- **SC-009**: Reduces support tickets from Taiwan users by 60% (language barrier removal)

## Compliance & Standards

### Taiwan Terminology Standards
- Follow Taiwan Ministry of Education National Academy for Educational Research terminology database
- Use Taiwan-specific IT terms from "資訊科技術語彙編"
- Apply Traditional Chinese written form per Taiwan education standards

### Accessibility
- Language selector must be keyboard accessible (Tab navigation)
- Screen readers must announce language changes
- Text must maintain minimum contrast ratios per WCAG 2.1 AA

## Out of Scope

The following are explicitly NOT included in this specification:
- ❌ Cantonese (Hong Kong zh-HK) support - separate feature
- ❌ Auto-translation of user-generated content
- ❌ OCR or image text translation
- ❌ Real-time translation API integration
- ❌ Third-party module translations (FMS, Job, MCMS) - handled by respective module teams
- ❌ Mobile app i18n (different codebase)
- ❌ Email template translations (future iteration)

## Dependencies

### Internal Dependencies
- Simple Admin Core v1.7.x or higher (current version)
- Vben5 frontend framework with vue-i18n
- Existing i18n infrastructure (embed.FS, locale files)

### External Dependencies
- OpenCC library (for initial simplified → traditional conversion)
- Ant Design Vue zh_TW locale package
- Taiwan terminology reference databases

### Team Dependencies
- Frontend developer for Vue i18n implementation (8 hours)
- Backend developer for Go i18n extension (4 hours)
- Native Taiwan speaker for translation review (8 hours)
- QA engineer for testing (4 hours)

## Risks & Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Terminology inconsistency between backend/frontend | High | Medium | Create shared terminology glossary, single source of truth |
| OpenCC auto-conversion produces incorrect terms | Medium | High | Manual review of all auto-converted text before commit |
| Third-party libraries lack zh-TW support | Medium | Low | Contribute zh-TW locale upstream or maintain fork |
| Translation key missing fallback causes blank UI | High | Low | Implement fallback chain zh-TW → zh-CN → en |
| User switches language mid-workflow loses progress | Medium | Medium | Persist form state across language switches |

## Implementation Notes

### Phase 1: Backend (4-6 hours)
1. Copy zh.json → zh-TW.json
2. Run through OpenCC converter (Simplified → Traditional, Taiwan standard)
3. Manual review of terminology (focus on technical terms)
4. Update var.go to recognize zh-TW locale
5. Add unit tests for zh-TW language file

### Phase 2: Frontend (8-10 hours)
1. Create zh_TW.ts based on existing zh_CN structure
2. Translate all UI strings using Taiwan terminology
3. Import Ant Design Vue zh_TW locale
4. Update language selector component
5. Add language switch E2E tests

### Phase 3: Integration (4-6 hours)
1. Add user.locale column via Ent schema
2. Update user profile API to save/retrieve locale preference
3. Implement locale priority logic (user > localStorage > browser > default)
4. Add migration script for existing users
5. Update API documentation

### Phase 4: QA & Polish (4-6 hours)
1. Manual testing of all core workflows
2. Native speaker review session
3. Fix reported terminology issues
4. Performance testing (language switch speed)
5. Update user documentation

**Total Estimated Effort**: 20-28 hours

## Review Checklist

Before marking this specification as "Ready for Planning", verify:

- [ ] All user stories have clear acceptance criteria
- [ ] All functional requirements are testable
- [ ] Success criteria are measurable
- [ ] Dependencies are identified and available
- [ ] Risks have mitigation strategies
- [ ] Out-of-scope items are explicitly listed
- [ ] Terminology standards are referenced
- [ ] Translation coverage target is realistic
- [ ] Timeline aligns with project roadmap
- [ ] Stakeholders have reviewed and approved

---

**Next Steps**:
1. Review this specification with team
2. Get approval from product owner
3. Create technical plan (plan.md)
4. Break down into tasks
5. Assign to sprint
