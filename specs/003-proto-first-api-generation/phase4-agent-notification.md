# Phase 4 Agent 通知文件

**專案**: Proto-First API Generation
**Phase**: 4 - Integration & Testing
**發送日期**: 2025-10-09
**狀態**: ▶️ Phase 4 Ready to Start
**發送者**: @pm

---

## 🎯 Phase 4 啟動通知

Phase 1-3 已完成所有開發任務 (13/13 tasks ✅),代碼已提交並測試通過。Phase 4 現在可以開始執行。

### Phase 1-3 完成概況
- ✅ 13 個任務完成
- ✅ 1,199 lines 核心代碼
- ✅ 7 個規範 commits
- ✅ Working tree clean
- ✅ 所有依賴項已解決

---

## 📢 致 @qa Agent

### 您的 Phase 4 任務 (共 6 個任務)

Phase 4 的主要工作是測試開發,所有任務已解除阻塞,可以立即開始。

#### 🔓 已解除阻塞的任務

**[PF-013] 單元測試 - Options Parser**
- **Notion**: https://www.notion.so/286f030bec8581a4bd11e607cf8e4c61
- **狀態**: 🔓 Blocked → Ready to Start
- **預計工時**: 3-4 hours
- **優先級**: P1 (Day 1)
- **測試目標**:
  - Service-Level Options 解析
  - Method-Level Options 解析
  - File-Level Options 解析
  - Options 合併邏輯
  - 測試覆蓋率 ≥ 90%

**[PF-014] 單元測試 - Message Parser**
- **Notion**: https://www.notion.so/286f030bec8581959147f98429b83692
- **狀態**: 🔓 Blocked → Ready to Start
- **預計工時**: 3-4 hours
- **優先級**: P1 (Day 1)
- **測試目標**:
  - Proto Types → Go Types 轉換
  - repeated/optional/map 類型
  - 命名轉換 (snake_case → PascalCase)
  - JSON Tag 生成
  - 測試覆蓋率 ≥ 90%

**[PF-015] 單元測試 - Service Grouper**
- **Notion**: https://www.notion.so/286f030bec85816a8247f7decc63725c
- **狀態**: Ready to Start
- **預計工時**: 3-4 hours
- **優先級**: P1 (Day 2)
- **測試目標**:
  - 按 @server options 分組
  - Public vs Protected 端點分組
  - Method-specific middleware 分組
  - 排序策略測試
  - 測試覆蓋率 ≥ 90%

**[PF-016] 整合測試 - 基本 Service 生成**
- **Notion**: https://www.notion.so/286f030bec8581129a9ffb58e3e58e8e
- **狀態**: 🔓 Blocked → Ready to Start
- **預計工時**: 4-5 hours
- **優先級**: P1 (Day 2)
- **測試場景**:
  - 基本 CRUD service 生成
  - info() section 生成
  - Type definitions 生成
  - @server block 生成
  - Service methods 生成

**[PF-017] 整合測試 - JWT 和 Public 端點**
- **Notion**: https://www.notion.so/286f030bec85814ba4cafbeb55cc558d
- **狀態**: Ready to Start
- **預計工時**: 3-4 hours
- **優先級**: P1 (Day 3)
- **測試場景**:
  - JWT protected endpoints 分組
  - Public endpoints 分組
  - Method-level middleware override
  - Mixed JWT/Public 在同一 Service

**[PF-019] E2E 測試 - 完整工作流程**
- **Notion**: https://www.notion.so/286f030bec8581cdba54e2304e46860a
- **狀態**: Blocked by PF-016, PF-017
- **預計工時**: 4-5 hours
- **優先級**: P2 (Day 4-5)
- **測試流程**:
  ```bash
  Proto File → protoc-gen-go-zero-api → .api File → goctl → Go Code → Build Success
  ```

### 建議執行順序

**Week 1 - Day 1** (6-8 hours):
1. ✅ 開始 [PF-013] Options Parser 單元測試
2. ✅ 開始 [PF-014] Message Parser 單元測試

**Week 1 - Day 2** (7-9 hours):
3. ✅ 開始 [PF-015] Service Grouper 單元測試
4. ✅ 開始 [PF-016] 基本 Service 生成整合測試

**Week 1 - Day 3** (3-4 hours):
5. ✅ 開始 [PF-017] JWT & Public 端點測試

**Week 1 - Day 4-5** (4-5 hours):
6. ✅ 開始 [PF-019] E2E 測試 (等待 PF-016, PF-017 完成)

### 測試資源

**測試目錄結構**:
```
tools/protoc-gen-go-zero-api/
├── generator/
│   ├── options_parser_test.go       (PF-013)
│   ├── type_converter_test.go       (PF-014, 合併 message parser)
│   ├── grouper_test.go              (PF-015)
│   └── generator_test.go            (PF-016, PF-017)
└── testdata/
    ├── fixtures/                    (測試 Proto 檔案)
    └── golden/                      (Golden files for PF-019)
```

**測試工具**:
- `go test -v ./generator/...`
- `go test -cover -coverprofile=coverage.out`
- `go tool cover -html=coverage.out`

### 成功標準

- ✅ 單元測試覆蓋率 ≥ 80%
- ✅ 所有測試通過 (0 failures)
- ✅ 生成的 .api 檔案可被 goctl 驗證
- ✅ Golden File Testing 實作完成

---

## 📢 致 @devops Agent

### 您的 Phase 4 任務 (1 個任務)

**[PF-016] 更新 Makefile - 工作流程整合**
- **Notion**: 與整合測試合併
- **狀態**: 🔓 Blocked → Ready to Start
- **預計工時**: 2-3 hours
- **優先級**: P1 (Day 4)
- **依賴**: PF-011 (已完成 ✅)

### 任務詳情

需要在 `Makefile` 中新增以下 targets:

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

### 整合點

需要將新的 targets 整合到現有的 `make gen-all` 目標中:

```makefile
.PHONY: gen-all
gen-all: gen-ent gen-proto-api gen-rpc gen-api-all
	@echo "All code generation complete"
```

### 驗收標準

- ✅ `make build-proto-plugin` 成功編譯插件
- ✅ `make gen-proto-api` 成功生成 .api 檔案
- ✅ `make validate-api` 驗證所有 .api 檔案
- ✅ `make gen-api-all` 完整工作流程可執行
- ✅ 整合到 `make gen-all` 目標

---

## 📢 致 @pm Agent (自己)

### Phase 4 PM 責任

**進度追蹤**:
- 每日更新 Notion 任務狀態
- 記錄實際工時 vs 預估工時
- 識別並解決阻塞問題

**每日檢查點**:
- **Day 1 EOD**: PF-013, PF-014 完成
- **Day 2 EOD**: PF-015, PF-016 完成
- **Day 3 EOD**: PF-017 完成
- **Day 4 EOD**: Makefile 完成, PF-019 50%
- **Day 5 EOD**: PF-019 完成, Phase 4 驗收

**Phase 4 驗收標準**:
- [ ] 單元測試覆蓋率 ≥ 80%
- [ ] 所有測試通過 (0 failures)
- [ ] Linter 無錯誤
- [ ] 生成的 .api 檔案 100% 可被 goctl 驗證
- [ ] Makefile targets 全部正常運作
- [ ] CI/CD 可自動化運行測試

**Phase 4 完成後行動**:
1. 生成 Phase 4 完成報告
2. 更新所有 Notion 任務為 "Done"
3. 準備 Phase 5: Pilot Migration & Documentation

---

## 📊 Phase 4 整體時程

| Day | Agent | Tasks | Hours |
|-----|-------|-------|-------|
| Day 1 | @qa | PF-013, PF-014 | 6-8h |
| Day 2 | @qa | PF-015, PF-016 | 7-9h |
| Day 3 | @qa | PF-017 | 3-4h |
| Day 4 | @devops | Makefile 更新 | 2-3h |
| Day 4 | @qa | PF-019 (Part 1) | 2h |
| Day 5 | @qa | PF-019 (Part 2) | 2-3h |
| Day 5 | @pm | Phase 4 驗收 | 2h |

**總預估工時**: 27-36 hours (分散至 5 個工作天)

---

## 🚀 後續步驟

### 立即行動 (今日)

**@pm**:
1. ✅ 更新 Notion 任務狀態 (PF-013, PF-014, PF-015, PF-016, PF-017)
2. ✅ 清空 "Blocked by" 欄位
3. ✅ 提交 `phase4-execution-plan.md` 到 git

**@qa** (如果今天開始):
1. ✅ Review Phase 4 執行計劃
2. ✅ 準備測試環境
3. ✅ 開始 [PF-013] Options Parser 單元測試

**@devops**:
1. ✅ Review Makefile 更新需求
2. ✅ 等待 PF-016, PF-017 完成後開始 Makefile 整合

---

## 📞 聯絡與協調

### 每日 Sync
- **時間**: 每日 10:00 AM (15 分鐘)
- **內容**: 昨日完成, 今日計劃, 阻塞問題
- **工具**: Notion Status Updates

### 問題升級
如遇到以下情況請立即通知 @pm:
- 任務無法如期完成
- 發現技術障礙
- 需要額外資源
- 依賴項目有問題

---

**發送者**: @pm
**建立日期**: 2025-10-09
**Phase 4 狀態**: ✅ Ready to Execute

祝各位執行順利! 🚀
