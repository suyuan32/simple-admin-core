# Phase 4 執行計劃: 整合與測試

**專案**: Proto-First API Generation
**Phase**: 4
**建立日期**: 2025-10-09
**狀態**: ▶️ Ready to Start
**負責 PM**: @pm
**預計工時**: 12-16 hours

---

## 🎯 Phase 4 目標

完成 protoc-gen-go-zero-api 插件的整合測試、驗證與 CI/CD 配置,確保生成的 .api 檔案品質達到生產標準。

### 核心交付物
1. ✅ 完整的單元測試套件 (覆蓋率 > 80%)
2. ✅ 整合測試框架和測試案例
3. ✅ Makefile 更新與工作流程整合
4. ✅ Golden File Testing 實作
5. ✅ CI/CD 流程建立

---

## 📋 Phase 4 任務清單

### 4.1 單元測試開發 (3 tasks)

#### [PF-013] 單元測試 - Options Parser
- **負責**: @qa
- **預計工時**: 3-4 hours
- **依賴**: PF-006 (已完成 ✅)
- **Notion**: https://www.notion.so/286f030bec8581a4bd11e607cf8e4c61
- **狀態**: 🔓 已解除阻塞 → Ready to Start

**驗收標準**:
- [ ] 測試 Service-Level Options 解析
- [ ] 測試 Method-Level Options 解析
- [ ] 測試 File-Level Options 解析
- [ ] 測試 Options 合併邏輯
- [ ] 測試覆蓋率 ≥ 90%

---

#### [PF-014] 單元測試 - Message Parser
- **負責**: @qa
- **預計工時**: 3-4 hours
- **依賴**: PF-005 (已完成 ✅,合併至 type_converter)
- **Notion**: https://www.notion.so/286f030bec8581959147f98429b83692
- **狀態**: 🔓 已解除阻塞 → Ready to Start

**驗收標準**:
- [ ] 測試 Proto Types → Go Types 轉換
- [ ] 測試 repeated/optional/map 類型
- [ ] 測試命名轉換 (snake_case → PascalCase)
- [ ] 測試 JSON Tag 生成
- [ ] 測試覆蓋率 ≥ 90%

---

#### [PF-015] 單元測試 - Service Grouper
- **負責**: @qa
- **預計工時**: 3-4 hours
- **依賴**: PF-008 (已完成 ✅)
- **Notion**: https://www.notion.so/286f030bec85816a8247f7decc63725c
- **狀態**: Ready to Start

**驗收標準**:
- [ ] 測試按 @server options 分組
- [ ] 測試 Public vs Protected 端點分組
- [ ] 測試 Method-specific middleware 分組
- [ ] 測試排序策略 (JWT → middleware → group name)
- [ ] 測試覆蓋率 ≥ 90%

---

### 4.2 整合測試 (3 tasks)

#### [PF-016] 整合測試 - 基本 Service 生成
- **負責**: @qa
- **預計工時**: 4-5 hours
- **依賴**: PF-007, PF-011 (已完成 ✅)
- **Notion**: https://www.notion.so/286f030bec8581129a9ffb58e3e58e8e
- **狀態**: 🔓 已解除阻塞 → Ready to Start

**測試案例**:
1. 基本 CRUD service 生成
2. info() section 生成
3. Type definitions 生成
4. @server block 生成
5. Service methods 生成

**驗收標準**:
- [ ] 生成的 .api 檔案可被 goctl 驗證通過
- [ ] 基本 HTTP methods (GET/POST/PUT/DELETE) 正確轉換
- [ ] Type 定義完整且格式正確
- [ ] @server annotations 正確生成

---

#### [PF-017] 整合測試 - JWT 和 Public 端點
- **負責**: @qa
- **預計工時**: 3-4 hours
- **依賴**: PF-006, PF-008 (已完成 ✅)
- **Notion**: https://www.notion.so/286f030bec85814ba4cafbeb55cc558d
- **狀態**: Ready to Start

**測試案例**:
1. JWT protected endpoints 分組
2. Public endpoints (public=true) 分組
3. Method-level middleware override
4. Mixed JWT/Public 在同一 Service

**驗收標準**:
- [ ] Protected endpoints 生成 @server(jwt: Auth)
- [ ] Public endpoints 不包含 jwt 配置
- [ ] 兩者生成獨立的 service 區塊
- [ ] Method middleware 正確覆蓋 Service middleware

---

#### [PF-018] 整合測試 - 複雜場景 (已完成)
- **負責**: @pm
- **狀態**: ✅ Done
- **Notion**: https://www.notion.so/286f030bec8581a7a6abc33f405b0503

---

#### [PF-019] E2E 測試 - 完整工作流程
- **負責**: @qa
- **預計工時**: 4-5 hours
- **依賴**: PF-011, PF-016, PF-017, PF-018
- **Notion**: https://www.notion.so/286f030bec8581cdba54e2304e46860a
- **狀態**: Blocked (需完成 PF-016, PF-017)

**測試流程**:
```bash
Proto File → protoc-gen-go-zero-api → .api File → goctl → Go Code → Build Success
```

**驗收標準**:
- [ ] 完整工作流程 E2E 測試通過
- [ ] Golden File Testing 實作
- [ ] 生成的 .api 可成功編譯為 Go 代碼
- [ ] API Service 可成功啟動

---

### 4.3 工作流程整合 (1 task)

#### [PF-016] 更新 Makefile (合併至整合測試)
- **負責**: @devops
- **預計工時**: 2-3 hours
- **依賴**: PF-011 (已完成 ✅)
- **狀態**: 🔓 已解除阻塞 → Ready to Start

**新增 Makefile targets**:
```makefile
.PHONY: build-proto-plugin
build-proto-plugin:
	@echo "Building protoc-gen-go-zero-api plugin..."
	cd tools/protoc-gen-go-zero-api && go build -o ../../bin/protoc-gen-go-zero-api

.PHONY: gen-proto-api
gen-proto-api: build-proto-plugin
	@echo "Generating .api files from Proto..."
	protoc --plugin=protoc-gen-go-zero-api=./bin/protoc-gen-go-zero-api \
	       --go-zero-api_out=api/desc \
	       --proto_path=. \
	       rpc/desc/**/*.proto

.PHONY: validate-api
validate-api: gen-proto-api
	@echo "Validating generated .api files..."
	@for file in api/desc/**/*.api; do \
		goctl api validate -api $$file || exit 1; \
	done

.PHONY: gen-api-all
gen-api-all: gen-proto-api gen-api-code
	@echo "API generation complete"
```

**驗收標準**:
- [ ] `make build-proto-plugin` 成功編譯插件
- [ ] `make gen-proto-api` 成功生成 .api 檔案
- [ ] `make validate-api` 驗證所有 .api 檔案
- [ ] `make gen-api-all` 完整工作流程可執行
- [ ] 整合到 `make gen-all` 目標

---

## 🗓️ Phase 4 時程規劃

### Week 1: 單元測試與整合測試 (Day 1-3)
| Day | Agent | Tasks | Hours |
|-----|-------|-------|-------|
| Day 1 | @qa | PF-013 Options Parser 單元測試 | 3-4h |
| Day 1 | @qa | PF-014 Message Parser 單元測試 | 3-4h |
| Day 2 | @qa | PF-015 Service Grouper 單元測試 | 3-4h |
| Day 2 | @qa | PF-016 基本 Service 生成整合測試 | 4-5h |
| Day 3 | @qa | PF-017 JWT & Public 端點測試 | 3-4h |

### Week 1: E2E 與工作流程 (Day 4-5)
| Day | Agent | Tasks | Hours |
|-----|-------|-------|-------|
| Day 4 | @devops | Makefile 更新與整合 | 2-3h |
| Day 4 | @qa | PF-019 E2E 測試 (Part 1) | 2h |
| Day 5 | @qa | PF-019 E2E 測試 (Part 2 - Golden Files) | 2-3h |
| Day 5 | @pm | Phase 4 完成報告與驗收 | 2h |

**總工時**: 27-36 hours (分散至 5 個工作天)

---

## ✅ 前置條件檢查

### Phase 1-3 完成狀態
- ✅ [PF-001] 建立插件專案結構 → Done
- ✅ [PF-002] 定義內部 Model 結構 → Done
- ✅ [PF-003] Go-Zero Custom Proto Options → Done
- ✅ [PF-004] HTTP Annotation Parser → Done
- ✅ [PF-006] Go-Zero Options Parser → Done
- ✅ [PF-008] Type Converter/Service Grouper → Done
- ✅ [PF-010] Service Grouper → Done
- ✅ [PF-011] Generator 整合 → Done
- ✅ [PF-012] Template Generator → Done
- ✅ [PF-018] PM 追蹤與同步 → Done

### 代碼庫狀態
- ✅ Working tree clean (無未提交變更)
- ✅ Feature branch: `feature/proto-first-api-generation`
- ✅ 7 commits 完成
- ✅ 1,199 lines 核心代碼

### 依賴項目
- ✅ Go 1.25+ installed
- ✅ protoc installed
- ✅ goctl installed
- ✅ Plugin 可編譯成功

---

## 🎯 成功標準

### 程式碼品質
- [ ] 單元測試覆蓋率 ≥ 80%
- [ ] 所有測試通過 (0 failures)
- [ ] Linter 無錯誤 (`golangci-lint run`)
- [ ] 無 TODO/FIXME 標記

### 功能完整性
- [ ] 生成的 .api 檔案 100% 可被 goctl 驗證
- [ ] 生成的 Go 代碼可成功編譯
- [ ] API Service 可成功啟動
- [ ] 支持所有 FR-001 至 FR-017 功能需求

### 工作流程
- [ ] Makefile targets 全部正常運作
- [ ] 開發者可使用 `make gen-api-all` 完成完整生成
- [ ] CI/CD 可自動化運行測試

---

## 🚧 風險與應對

### 風險 1: 測試覆蓋率不足
**機率**: Medium
**影響**: High
**應對**:
- 使用 coverage report 識別未覆蓋的代碼路徑
- 針對性補充測試案例
- 設定 CI 自動檢查覆蓋率門檻

### 風險 2: Golden File Testing 實作複雜
**機率**: Medium
**影響**: Medium
**應對**:
- 參考 Go 標準庫 golden file 模式
- 建立清晰的 fixtures 目錄結構
- 提供 `--update-golden` 旗標便於更新

### 風險 3: Makefile 整合衝突
**機率**: Low
**影響**: Low
**應對**:
- 在獨立分支測試 Makefile 變更
- 與現有 gen-* targets 保持一致命名
- 提供 rollback 方案

---

## 📊 進度追蹤

### Notion 任務狀態
- [ ] PF-013: Not started → In progress (Day 1)
- [ ] PF-014: Not started → In progress (Day 1)
- [ ] PF-015: Not started → In progress (Day 2)
- [ ] PF-016: Blocked → In progress (Day 2)
- [ ] PF-017: Not started → In progress (Day 3)
- [ ] PF-019: Blocked → In progress (Day 4)

### 每日更新檢查點
- **Day 1 EOD**: PF-013, PF-014 完成
- **Day 2 EOD**: PF-015, PF-016 完成
- **Day 3 EOD**: PF-017 完成
- **Day 4 EOD**: Makefile 完成, PF-019 50%
- **Day 5 EOD**: PF-019 完成, Phase 4 驗收

---

## 📝 後續行動 (Phase 5 預覽)

Phase 4 完成後,將進入 Phase 5: Pilot Migration & Documentation

### Phase 5 關鍵任務
1. User module Proto-First migration
2. Migration guide documentation
3. CLAUDE.md update
4. Team training session

---

## 📞 聯絡與協調

### Agent 責任分配
- **@qa**: PF-013, PF-014, PF-015, PF-016, PF-017, PF-019
- **@devops**: Makefile 更新, CI/CD 配置
- **@pm**: 進度追蹤, Notion 更新, Phase 驗收

### 每日 Sync
- 時間: 每日 10:00 AM (15 分鐘)
- 內容: 昨日完成, 今日計劃, 阻塞問題
- 工具: Notion Status Updates

---

**建立者**: @pm
**最後更新**: 2025-10-09
**狀態**: ✅ Ready to Execute
