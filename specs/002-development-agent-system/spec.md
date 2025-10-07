# Agent System Specification: Development Efficiency Optimization

**Feature Branch**: `002-dev-agent-system`
**Created**: 2025-10-07
**Status**: Draft
**Input**: Design a multi-agent system to maximize development efficiency for Simple Admin Core project

## User Scenarios & Testing (mandatory)

### User Story 1 - Automated Code Generation Agent (Priority: P1)
A backend developer needs to create a new CRUD module for "Department Management". Instead of manually writing Ent schema, Proto definitions, RPC logic, and API definitions, they invoke the **@generator** agent with a simple description. The agent analyzes the requirements, generates all necessary files following project conventions, and creates a pull request ready for review.

**Why this priority**: Code generation is repetitive and error-prone. Automating this saves 60-70% development time per module.

**Acceptance Scenarios**:
1. **Given** developer provides entity description, **When** invoke `@generator create-module Department`, **Then** generates Ent schema, Proto, RPC logic, API definition
2. **Given** generated code, **When** run `make gen-ent && make gen-rpc && make gen-api`, **Then** all code generation succeeds with zero errors
3. **Given** generated files, **When** code review, **Then** follows project naming conventions (snake_case, i18n keys)

### User Story 2 - Intelligent Code Reviewer Agent (Priority: P1)
A developer completes a feature and creates a pull request. The **@reviewer** agent automatically analyzes the code changes, checks for:
- Go best practices violations
- Missing i18n keys
- Incomplete error handling
- Test coverage gaps
- API/RPC contract mismatches

The agent posts review comments directly on GitHub PR, highlighting issues with suggested fixes.

**Why this priority**: Manual code review is time-consuming. AI-assisted review catches common issues before human review, improving code quality.

**Acceptance Scenarios**:
1. **Given** new PR created, **When** @reviewer triggered, **Then** analyzes all changed files and posts review comments
2. **Given** missing i18n key detected, **When** review complete, **Then** suggests zh-TW.json addition with proper key naming
3. **Given** uncovered code paths, **When** review complete, **Then** identifies missing unit tests and suggests test cases

### User Story 3 - Database Migration Agent (Priority: P2)
When Ent schema changes require database migration beyond auto-migration capabilities (e.g., data transformation, column renaming), the **@migrator** agent generates SQL migration scripts compatible with the project's migration strategy.

**Why this priority**: Complex migrations are risky. Automated script generation with rollback ensures data integrity.

**Acceptance Scenarios**:
1. **Given** Ent schema adds new relationship, **When** invoke @migrator, **Then** generates forward and rollback SQL scripts
2. **Given** generated migration, **When** apply to test database, **Then** executes successfully without data loss
3. **Given** migration script, **When** validate, **Then** includes proper transaction handling and error recovery

### User Story 4 - Documentation Synchronization Agent (Priority: P2)
After API changes, the **@doc-sync** agent automatically:
- Updates Swagger/OpenAPI specs
- Regenerates API documentation
- Updates CLAUDE.md if workflow changes
- Creates changelog entries
- Syncs to Notion project documentation

**Why this priority**: Documentation drift is a major pain point. Automated sync keeps docs always up-to-date.

**Acceptance Scenarios**:
1. **Given** new API endpoint added, **When** @doc-sync runs, **Then** Swagger spec updated, docs regenerated, Notion page created
2. **Given** API contract changed, **When** sync complete, **Then** changelog includes breaking change notice
3. **Given** new Go command added, **When** sync runs, **Then** CLAUDE.md Common Commands section updated

### User Story 5 - Test Generation Agent (Priority: P2)
Developer writes new business logic in RPC layer. The **@test-gen** agent analyzes the function signature, identifies edge cases, and generates comprehensive unit tests and table-driven tests following Go conventions.

**Why this priority**: Writing tests is tedious. Auto-generation ensures consistent test coverage.

**Acceptance Scenarios**:
1. **Given** new RPC logic function, **When** invoke @test-gen, **Then** generates unit tests with happy path, error cases, edge cases
2. **Given** generated tests, **When** run `go test`, **Then** all tests pass with >80% coverage
3. **Given** function with i18n errors, **When** tests generated, **Then** includes tests for different locales

### User Story 6 - Deployment & DevOps Agent (Priority: P1)
Developer completes a sprint. The **@deployer** agent:
- Builds Docker images
- Runs integration tests in staging
- Updates Kubernetes manifests
- Deploys to staging environment
- Monitors deployment health
- Sends Notion notification with deployment summary

**Why this priority**: Deployment is complex and error-prone. Automation reduces deployment time from 30min to 5min.

**Acceptance Scenarios**:
1. **Given** merged to main, **When** @deployer triggered, **Then** builds images, pushes to registry, updates k8s, deploys
2. **Given** deployment failed, **When** health check detects issue, **Then** auto-rollback and sends alert
3. **Given** successful deployment, **When** complete, **Then** Notion page updated with deployment logs and metrics

### User Story 7 - Spec-Kit Workflow Agent (Priority: P1)
PM wants to plan a new feature. The **@spec-writer** agent:
- Interviews PM via interactive prompts
- Generates spec.md following spec-kit format
- Analyzes technical feasibility
- Creates plan.md with implementation details
- Breaks down into tasks and uploads to Notion
- Links all documentation together

**Why this priority**: Spec writing is time-consuming. Guided generation ensures completeness and consistency.

**Acceptance Scenarios**:
1. **Given** PM provides feature idea, **When** @spec-writer runs, **Then** generates complete spec.md with user stories, requirements, success criteria
2. **Given** spec approved, **When** continue, **Then** generates plan.md with architecture, timeline, code examples
3. **Given** plan complete, **When** finalize, **Then** creates 10-20 Notion tasks with proper dependencies

### Edge Cases
- What if generated code conflicts with existing code? Agent should detect conflicts and suggest resolution strategies
- How to handle agent errors or hallucinations? Implement validation layer that runs tests/linters before committing
- What if multiple agents modify same file? Implement agent coordination system with file locks
- How to prevent agents from making unauthorized changes? All agent actions require approval or run in dry-run mode first
- What if Notion API rate limits hit? Implement exponential backoff and batch operations

## Requirements (mandatory)

### Functional Requirements

#### Agent Infrastructure
- **FR-001**: System MUST support 7 specialized agents: @generator, @reviewer, @migrator, @doc-sync, @test-gen, @deployer, @spec-writer
- **FR-002**: Each agent MUST have clear role definition and non-overlapping responsibilities
- **FR-003**: Agents MUST communicate via standardized message protocol
- **FR-004**: System MUST log all agent actions for audit trail
- **FR-005**: Agents MUST support both CLI invocation and GitHub Actions triggers

#### Code Generation Agent (@generator)
- **FR-006**: MUST generate Ent schema with proper field types, indexes, relationships
- **FR-007**: MUST generate Proto messages matching Ent schema
- **FR-008**: MUST generate API definitions (.api files) following go-zero conventions
- **FR-009**: MUST generate RPC logic with proper error handling and i18n
- **FR-010**: MUST support templates for common patterns (CRUD, tree structure, nested resources)

#### Code Review Agent (@reviewer)
- **FR-011**: MUST check Go code against `golangci-lint` rules
- **FR-012**: MUST verify i18n keys exist in all language files (zh.json, en.json, zh-TW.json)
- **FR-013**: MUST detect missing error handling (unchecked errors)
- **FR-014**: MUST verify test coverage >70% for new code
- **FR-015**: MUST check API/RPC contract consistency (proto ↔ api definitions)

#### Migration Agent (@migrator)
- **FR-016**: MUST generate forward and rollback SQL scripts
- **FR-017**: MUST support PostgreSQL and MySQL syntax
- **FR-018**: MUST include transaction handling in migrations
- **FR-019**: MUST detect backward-incompatible schema changes
- **FR-020**: MUST generate migration version numbers (timestamp-based)

#### Documentation Agent (@doc-sync)
- **FR-021**: MUST update Swagger specs via `make gen-swagger`
- **FR-022**: MUST sync CLAUDE.md when project structure changes
- **FR-023**: MUST create Notion pages for new features
- **FR-024**: MUST generate changelog entries following Keep a Changelog format
- **FR-025**: MUST support versioned documentation (v1.7, v1.8, etc.)

#### Test Generation Agent (@test-gen)
- **FR-026**: MUST generate table-driven tests for Go functions
- **FR-027**: MUST identify edge cases (nil inputs, empty strings, boundary values)
- **FR-028**: MUST generate mock objects for dependencies
- **FR-029**: MUST include test data fixtures
- **FR-030**: MUST follow AAA pattern (Arrange, Act, Assert)

#### Deployment Agent (@deployer)
- **FR-031**: MUST build Docker images with proper tags (version, git sha)
- **FR-032**: MUST run smoke tests before deployment
- **FR-033**: MUST support blue-green deployment strategy
- **FR-034**: MUST auto-rollback on health check failures
- **FR-035**: MUST send deployment notifications to Notion and Slack

#### Spec-Kit Workflow Agent (@spec-writer)
- **FR-036**: MUST generate spec.md following GitHub spec-kit template
- **FR-037**: MUST validate all required sections (User Scenarios, Requirements, Success Criteria)
- **FR-038**: MUST generate plan.md with implementation phases
- **FR-039**: MUST create Notion tasks with proper metadata (priority, due date, assignee)
- **FR-040**: MUST link spec ↔ plan ↔ tasks bidirectionally

### Key Entities

#### Agent
- **Name**: String (e.g., "generator", "reviewer")
- **Description**: String (role and responsibilities)
- **Trigger**: Enum (CLI, GitHub Action, Webhook, Schedule)
- **RequiredTools**: Array of tool names
- **Configuration**: JSON (agent-specific settings)

#### AgentJob
- **ID**: UUID
- **AgentName**: String (foreign key to Agent)
- **Status**: Enum (pending, running, success, failed, rolled_back)
- **Input**: JSON (job parameters)
- **Output**: JSON (job results)
- **Logs**: Text (execution logs)
- **CreatedAt**: Timestamp
- **CompletedAt**: Timestamp

#### AgentAction
- **JobID**: UUID (foreign key to AgentJob)
- **ActionType**: Enum (file_create, file_modify, file_delete, api_call, command_execute)
- **Target**: String (file path, API endpoint, command)
- **Changes**: JSON (diff or payload)
- **RequiresApproval**: Boolean
- **ApprovedBy**: String (user ID or "auto")

## Success Criteria (mandatory)

### Measurable Outcomes
- **SC-001**: Code generation agent reduces module creation time from 4 hours to <1 hour (75% time savings)
- **SC-002**: Review agent catches 80%+ of common issues before human review
- **SC-003**: Documentation sync agent maintains docs freshness >95% (doc age <7 days)
- **SC-004**: Test generation achieves >70% code coverage automatically
- **SC-005**: Deployment agent reduces deployment time from 30min to <5min (83% time savings)
- **SC-006**: Spec-writer agent reduces spec creation time from 8 hours to <2 hours (75% time savings)

### User Satisfaction Metrics
- **SC-007**: Developer survey: agent system rated ≥4.5/5.0 for usefulness
- **SC-008**: Zero production incidents caused by agent-generated code (after 3-month period)
- **SC-009**: Agent-generated PRs accepted without major revisions ≥80% of time

### Business Metrics
- **SC-010**: Overall development velocity increases 40% (measured by story points/sprint)
- **SC-011**: Code review cycle time decreases 50% (PR open to merge)
- **SC-012**: Documentation maintenance effort decreases 60%

## Out of Scope

❌ General-purpose AI chat assistant (focus on specialized task agents)<br>
❌ Agents for non-code tasks (design, marketing, sales)<br>
❌ Autonomous agents that can push to production without approval<br>
❌ Agents that modify user business logic without explicit instruction<br>
❌ Cross-project agents (focus on Simple Admin Core only)<br>
❌ Real-time pair programming assistant (different use case)

## Dependencies

### Internal Dependencies
- Simple Admin Core v1.7.x codebase
- Existing Makefile targets (gen-ent, gen-rpc, gen-api)
- CLAUDE.md project documentation
- spec-kit format templates

### External Dependencies
- **LLM API**: Claude 3.5 Sonnet (via Anthropic API)
- **Code Analysis**: golangci-lint, go-vet
- **Git Integration**: GitHub API, git CLI
- **CI/CD**: GitHub Actions
- **Notification**: Notion API, Slack Webhooks
- **Container**: Docker, Kubernetes
- **Testing**: Go testing framework, testify

### Team Dependencies
- DevOps engineer to configure GitHub Actions (4 hours)
- Backend lead to define code generation templates (8 hours)
- QA lead to define review criteria (4 hours)

## Risks & Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Agent generates incorrect code | High | Medium | Multi-layer validation: linter → tests → human review |
| LLM API rate limits or downtime | Medium | Low | Implement caching, fallback to template-based generation |
| Over-reliance on agents reduces developer skills | Medium | Medium | Use agents for boilerplate only, developers own business logic |
| Agent coordination conflicts | Medium | Low | Implement file-level locking, sequential execution for related tasks |
| Security: agents leak sensitive data in logs | High | Low | Sanitize all logs, restrict log access, no secrets in prompts |
| Cost: LLM API usage spirals | Medium | Medium | Set monthly budget caps, optimize prompts, cache common responses |

## Implementation Notes

### Phase 1: Core Infrastructure (Week 1-2)
1. Design agent protocol and message format
2. Implement agent registry and job queue
3. Create CLI for agent invocation
4. Set up logging and monitoring

### Phase 2: Priority Agents (Week 3-4)
1. @generator - Code generation
2. @reviewer - Code review
3. @deployer - Deployment automation

### Phase 3: Supporting Agents (Week 5-6)
1. @doc-sync - Documentation sync
2. @test-gen - Test generation
3. @migrator - Migration scripts

### Phase 4: Advanced Agents (Week 7-8)
1. @spec-writer - Spec-kit workflow
2. Agent coordination system
3. Performance optimization

**Total Estimated Effort**: 120-150 hours (6-8 weeks, 2-3 developers)

## Review Checklist

- [ ] All 7 agents have clear responsibilities
- [ ] No agent responsibility overlaps
- [ ] All agents integrate with existing Simple Admin Core workflow
- [ ] Security concerns addressed (approval, logging, secrets)
- [ ] Cost model defined (LLM API usage, infrastructure)
- [ ] Rollback strategy for each agent type
- [ ] User training plan for agent usage
- [ ] Metrics collection plan for measuring success

---

**Next Steps**:
1. Review with development team
2. Get approval from tech lead
3. Create technical plan (plan.md)
4. Prototype @generator agent first
5. Iterate based on feedback
