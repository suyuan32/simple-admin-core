# Proto-First API Generation - Task Allocation & Agent Coordination

**Feature**: Proto-First API Generation
**Branch**: `feature/proto-first-api-generation`
**Total Effort**: 60-80 hours
**Team**: 5 Backend Agents + 1 PM + 1 QA + 1 DevOps
**Tracking**: Notion Tasks Database

## Agent Team Structure

### @pm - Project Manager Agent (全程)
**職責**: 任務追蹤、進度同步、Notion 更新、風險管理
**工時**: 持續參與 (每個階段 2-4h)
**Notion 欄位**:
- Agent: "pm"
- 每個任務完成後更新 Status

### @backend-1 - Plugin Core Developer (核心開發)
**職責**: 插件主架構、Proto 解析、整合邏輯
**工時**: 25-30h
**任務**:
- [PF-001] 建立插件專案結構
- [PF-002] 實作主要 Generator
- [PF-011] 整合所有組件
**Notion 欄位**:
- Agent: "backend-1"
- Priority: "High Priority"

### @backend-2 - HTTP Parser Specialist (HTTP 解析專家)
**職責**: HTTP annotation 解析、路徑轉換
**工時**: 10-12h
**任務**:
- [PF-004] 實作 HTTP Annotation Parser
- [PF-005] 處理 additional_bindings
**Notion 欄位**:
- Agent: "backend-2"
- Priority: "High Priority"

### @backend-3 - Options Parser Specialist (選項解析專家)
**職責**: Go-Zero 自訂選項解析
**工時**: 10-12h
**任務**:
- [PF-003] 定義 Go-Zero Custom Proto Options
- [PF-006] 實作 Options Parser
- [PF-007] 實作選項合併邏輯
**Notion 欄位**:
- Agent: "backend-3"
- Priority: "High Priority"

### @backend-4 - Type Converter Specialist (型別轉換專家)
**職責**: Proto 型別轉換、Go-Zero 型別生成
**工時**: 10-12h
**任務**:
- [PF-008] 實作 Type Converter
- [PF-009] 處理複雜型別 (nested, repeated, optional)
**Notion 欄位**:
- Agent: "backend-4"
- Priority: "Medium Priority"

### @backend-5 - Template & Grouping Specialist (模板與分組專家)
**職責**: 服務分組、模板生成
**工時**: 10-12h
**任務**:
- [PF-010] 實作 Service Grouper
- [PF-012] 實作 Template Generator
**Notion 欄位**:
- Agent: "backend-5"
- Priority: "Medium Priority"

### @qa - Quality Assurance Agent
**職責**: 測試策略、單元測試、整合測試
**工時**: 12-15h
**任務**:
- [PF-013] 撰寫單元測試
- [PF-014] 建立整合測試套件
- [PF-015] Golden file 測試
**Notion 欄位**:
- Agent: "qa"
- Priority: "High Priority"

### @devops - DevOps Agent
**職責**: Makefile 整合、CI/CD 設定
**工時**: 4-6h
**任務**:
- [PF-016] 更新 Makefile
- [PF-017] CI/CD 整合
**Notion 欄位**:
- Agent: "devops"
- Priority: "Medium Priority"

## Task Breakdown (Notion Tasks)

### Phase 1: Setup & Foundation (16-20 hours)

#### [PF-001] 建立插件專案結構
- **Agent**: @backend-1
- **Estimated**: 3-4h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 建立 `tools/protoc-gen-go-zero-api/` 目錄結構
  - 設定 go.mod, 安裝依賴
  - 實作 main.go 入口點
  - 建立基礎目錄: generator/, model/, test/
- **Deliverable**: 可編譯的插件骨架
- **Blocked by**: None
- **Blocking**: [PF-002], [PF-004], [PF-006], [PF-008], [PF-010]

#### [PF-002] 定義內部 Model 結構
- **Agent**: @backend-1
- **Estimated**: 2-3h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 定義 Service, Method, Message, Field 模型
  - 定義 ServerOptions, MethodOptions, HTTPRule 模型
- **Deliverable**: model/ 下的所有結構定義
- **Blocked by**: None
- **Blocking**: [PF-004], [PF-006], [PF-008]

#### [PF-003] 定義 Go-Zero Custom Proto Options
- **Agent**: @backend-3
- **Estimated**: 4-5h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 建立 `rpc/desc/go_zero/options.proto`
  - 定義 service/method/file level extensions
  - 定義 ApiInfo message
  - 產生 Go code: `protoc --go_out=.`
- **Deliverable**: go_zero/options.proto, options.pb.go
- **Blocked by**: None
- **Blocking**: [PF-006]

#### [PF-018] Phase 1 PM 追蹤與同步
- **Agent**: @pm
- **Estimated**: 2-3h
- **Priority**: Critical
- **Status**: Not started
- **Description**:
  - 建立 Notion 專案頁面和任務
  - 每日同步 [PF-001] ~ [PF-003] 進度
  - 更新 task status 到 Notion
- **Deliverable**: Notion 任務看板建立完成
- **Blocked by**: None

### Phase 2: Parsers Implementation (20-24 hours)

#### [PF-004] 實作 HTTP Annotation Parser
- **Agent**: @backend-2
- **Estimated**: 5-6h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 解析 google.api.http options
  - 提取 HTTP method, path, body
  - 路徑轉換: {id} → :id
- **Deliverable**: generator/http_parser.go + tests
- **Blocked by**: [PF-001], [PF-002]
- **Blocking**: [PF-011]

#### [PF-005] 處理 additional_bindings
- **Agent**: @backend-2
- **Estimated**: 3-4h
- **Priority**: Medium Priority
- **Status**: Not started
- **Description**:
  - 解析 additional_bindings
  - 支援一個 method 多個 HTTP route
- **Deliverable**: additional_bindings 支援
- **Blocked by**: [PF-004]
- **Blocking**: [PF-011]

#### [PF-006] 實作 Options Parser
- **Agent**: @backend-3
- **Estimated**: 5-6h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 解析 service-level options (jwt, middleware, group)
  - 解析 method-level options (public, middleware)
  - 解析 file-level options (api_info)
- **Deliverable**: generator/options_parser.go + tests
- **Blocked by**: [PF-001], [PF-002], [PF-003]
- **Blocking**: [PF-007], [PF-011]

#### [PF-007] 實作選項合併邏輯
- **Agent**: @backend-3
- **Estimated**: 2-3h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 實作 MergeOptions: service + method options
  - 處理 public endpoint (override JWT)
  - 處理 method-specific middleware
- **Deliverable**: Options merge 功能
- **Blocked by**: [PF-006]
- **Blocking**: [PF-010]

#### [PF-008] 實作 Type Converter
- **Agent**: @backend-4
- **Estimated**: 6-7h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - Proto types → Go types 轉換
  - 處理 basic types, nested messages
  - 處理 optional, repeated fields
  - 產生 .api type definitions
- **Deliverable**: generator/type_converter.go + tests
- **Blocked by**: [PF-001], [PF-002]
- **Blocking**: [PF-011]

#### [PF-009] 處理複雜型別
- **Agent**: @backend-4
- **Estimated**: 3-4h
- **Priority**: Medium Priority
- **Status**: Not started
- **Description**:
  - 處理 oneof, map, Any
  - 處理 enum → string 轉換
- **Deliverable**: 複雜型別支援
- **Blocked by**: [PF-008]
- **Blocking**: [PF-011]

#### [PF-019] Phase 2 PM 追蹤與同步
- **Agent**: @pm
- **Estimated**: 3-4h
- **Priority**: Critical
- **Status**: Not started
- **Description**:
  - 每日更新 [PF-004] ~ [PF-009] 狀態
  - 協調 4 個 backend agents 進度
  - 識別並解決阻塞問題
- **Blocked by**: [PF-018]

### Phase 3: Grouping & Template (12-16 hours)

#### [PF-010] 實作 Service Grouper
- **Agent**: @backend-5
- **Estimated**: 5-6h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 按 @server options 分組 methods
  - 產生多個 service blocks
  - 排序: protected → public
- **Deliverable**: generator/grouper.go + tests
- **Blocked by**: [PF-001], [PF-007]
- **Blocking**: [PF-012]

#### [PF-011] 整合 Generator 組件
- **Agent**: @backend-1
- **Estimated**: 4-5h
- **Priority**: Critical
- **Status**: Not started
- **Description**:
  - 整合所有 parsers
  - 實作 Generator.Generate() 主流程
  - Proto file → AST → .api content
- **Deliverable**: generator/generator.go
- **Blocked by**: [PF-004], [PF-005], [PF-006], [PF-008], [PF-009]
- **Blocking**: [PF-012], [PF-013]

#### [PF-012] 實作 Template Generator
- **Agent**: @backend-5
- **Estimated**: 5-6h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 設計 .api template (text/template)
  - 產生 info() section
  - 產生 type definitions
  - 產生 @server blocks + service definitions
- **Deliverable**: generator/template.go + tests
- **Blocked by**: [PF-010], [PF-011]
- **Blocking**: [PF-014]

#### [PF-020] Phase 3 PM 追蹤與同步
- **Agent**: @pm
- **Estimated**: 2-3h
- **Priority**: Critical
- **Status**: Not started
- **Description**:
  - 更新 [PF-010] ~ [PF-012] 狀態
  - 確認整合進度
  - 準備測試階段
- **Blocked by**: [PF-019]

### Phase 4: Testing & Integration (16-20 hours)

#### [PF-013] 撰寫單元測試
- **Agent**: @qa
- **Estimated**: 6-7h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 每個 component 的單元測試
  - Mock Proto structures
  - 驗證輸出格式
- **Deliverable**: test/*_test.go (80%+ coverage)
- **Blocked by**: [PF-011]
- **Blocking**: [PF-015]

#### [PF-014] 建立整合測試套件
- **Agent**: @qa
- **Estimated**: 5-6h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - End-to-end 測試: Proto → .api
  - 執行 protoc 指令
  - 驗證生成的 .api 可編譯
- **Deliverable**: test/integration_test.go
- **Blocked by**: [PF-012]
- **Blocking**: [PF-015]

#### [PF-015] Golden File 測試
- **Agent**: @qa
- **Estimated**: 3-4h
- **Priority**: Medium Priority
- **Status**: Not started
- **Description**:
  - 建立 fixtures: user.proto, role.proto
  - 建立 expected .api 檔案
  - 對比生成結果與 golden files
- **Deliverable**: test/fixtures/ + golden file tests
- **Blocked by**: [PF-013], [PF-014]

#### [PF-016] 更新 Makefile
- **Agent**: @devops
- **Estimated**: 2-3h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 新增 build-proto-plugin target
  - 新增 gen-proto-api target
  - 新增 gen-api-all target
  - 新增 validate-api target
- **Deliverable**: 更新的 Makefile
- **Blocked by**: [PF-011]
- **Blocking**: [PF-017]

#### [PF-017] CI/CD 整合
- **Agent**: @devops
- **Estimated**: 2-3h
- **Priority**: Medium Priority
- **Status**: Not started
- **Description**:
  - GitHub Actions workflow
  - 自動測試 plugin
  - 驗證生成的 .api files
- **Deliverable**: .github/workflows/proto-api-gen.yml
- **Blocked by**: [PF-016]

#### [PF-021] Phase 4 PM 追蹤與同步
- **Agent**: @pm
- **Estimated**: 3-4h
- **Priority**: Critical
- **Status**: Not started
- **Description**:
  - 更新 [PF-013] ~ [PF-017] 狀態
  - 協調 QA 和 DevOps
  - 驗證測試覆蓋率
- **Blocked by**: [PF-020]

### Phase 5: Migration & Documentation (12-16 hours)

#### [PF-022] User 模組遷移 (Pilot)
- **Agent**: @backend-1
- **Estimated**: 4-5h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 在 user.proto 加入 Go-Zero options
  - 執行 make gen-proto-api
  - 對比生成結果與現有 .api
  - 測試 API service 編譯和運行
- **Deliverable**: User 模組成功遷移
- **Blocked by**: [PF-015], [PF-016]
- **Blocking**: [PF-023]

#### [PF-023] Role 模組遷移
- **Agent**: @backend-2
- **Estimated**: 2-3h
- **Priority**: Medium Priority
- **Status**: Not started
- **Description**:
  - 遷移 role.proto
  - 驗證生成結果
- **Deliverable**: Role 模組成功遷移
- **Blocked by**: [PF-022]

#### [PF-024] 撰寫遷移指南
- **Agent**: @backend-3
- **Estimated**: 3-4h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 完整遷移步驟文檔
  - Troubleshooting 指南
  - 回滾計畫
- **Deliverable**: docs/proto-first-migration-guide.md
- **Blocked by**: [PF-022]

#### [PF-025] 更新專案文檔
- **Agent**: @backend-4
- **Estimated**: 2-3h
- **Priority**: Medium Priority
- **Status**: Not started
- **Description**:
  - 更新 CLAUDE.md
  - 新增 Proto-First 章節
  - 文檔 go_zero.proto options
- **Deliverable**: 更新的 CLAUDE.md
- **Blocked by**: [PF-022]

#### [PF-026] 團隊培訓準備
- **Agent**: @pm
- **Estimated**: 2-3h
- **Priority**: High Priority
- **Status**: Not started
- **Description**:
  - 準備培訓材料
  - 建立範例 demos
  - FAQ 文檔
- **Deliverable**: 培訓材料套組
- **Blocked by**: [PF-024], [PF-025]

## Parallel Execution Strategy

### 可並行執行的任務組:

**Week 1 - Foundation (並行 3 組)**:
```
Group A: @backend-1 執行 [PF-001], [PF-002]
Group B: @backend-3 執行 [PF-003]
Group C: @pm 執行 [PF-018]
```

**Week 2 - Parsers (並行 4 組)**:
```
Group A: @backend-2 執行 [PF-004] → [PF-005]
Group B: @backend-3 執行 [PF-006] → [PF-007]
Group C: @backend-4 執行 [PF-008] → [PF-009]
Group D: @pm 執行 [PF-019]
```
預期加速: **40% faster** (24h → 14h)

**Week 3 - Integration (並行 3 組)**:
```
Group A: @backend-1 執行 [PF-011]
Group B: @backend-5 執行 [PF-010] → [PF-012]
Group C: @pm 執行 [PF-020]
```

**Week 4 - Testing (並行 3 組)**:
```
Group A: @qa 執行 [PF-013], [PF-014], [PF-015]
Group B: @devops 執行 [PF-016] → [PF-017]
Group C: @pm 執行 [PF-021]
```
預期加速: **50% faster** (20h → 10h)

**Week 5 - Migration (並行 4 組)**:
```
Group A: @backend-1 執行 [PF-022]
Group B: @backend-2 執行 [PF-023] (after [PF-022])
Group C: @backend-3 執行 [PF-024]
Group D: @backend-4 執行 [PF-025]
Then: @pm 執行 [PF-026]
```

## Notion Workflow Protocol

### 任務狀態更新流程:

1. **任務開始時**:
   ```
   @agent: Update [PF-XXX] status to "In progress"
   @pm: Record start time in Notion
   ```

2. **開發過程中**:
   ```
   @agent: Update task Description with progress notes
   @pm: Monitor blocked tasks, update Dependencies
   ```

3. **Git Commit 後**:
   ```
   @agent: Commit code with message
   @pm: Add commit hash to Notion task description
   @pm: Update task progress percentage
   ```

4. **任務完成時**:
   ```
   @agent: Verify acceptance criteria met
   @pm: Update status to "Done"
   @pm: Record actual hours
   @pm: Unblock dependent tasks
   ```

5. **遇到阻塞時**:
   ```
   @agent: Report blocker
   @pm: Update "Blocked by" field in Notion
   @pm: Escalate if critical
   ```

## Success Metrics

### Development Velocity:
- **Solo**: 60-80 hours (3-4 weeks single developer)
- **5 Agents Parallel**: **35-45 hours** (1.5-2 weeks wall time)
- **加速比**: **1.7x - 2.2x faster** ⚡

### Quality Metrics:
- Test coverage: 80%+
- All .api files compile: 100%
- Zero regression bugs
- Pilot migration success: 100%

### Notion Tracking:
- Task completion rate tracked daily
- Blocker resolution time < 4 hours
- Agent utilization > 85%

## Risk Management

### Risk: Agent 之間溝通開銷
**Mitigation**:
- 清楚定義介面契約 (model structures)
- @pm 每日 sync meeting (15 min)
- Notion 作為 single source of truth

### Risk: 並行開發整合衝突
**Mitigation**:
- 每個 agent 獨立模組 (http_parser, options_parser 等)
- @backend-1 負責最終整合 (clear owner)
- 頻繁 integration testing

### Risk: Notion 更新延遲
**Mitigation**:
- @pm 專職更新 Notion (real-time)
- 自動化: git hook → Notion webhook (future enhancement)

## Next Steps

1. **@pm**: 在 Notion 建立所有 26 個任務
2. **@pm**: 設定任務依賴關係 (Blocked by)
3. **Team**: Review task allocation
4. **@pm**: Kick-off meeting, 分配 Week 1 任務
5. **Start**: 5 agents 開始並行開發 🚀
