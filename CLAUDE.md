# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Simple Admin Core is a powerful microservice-based backend management system built with:
- **Go-Zero**: Microservice framework
- **Ent**: ORM and schema management
- **Casbin**: RBAC permission control
- **gRPC**: RPC communication
- **MySQL/PostgreSQL**: Database support
- **Redis**: Caching and distributed locking

The system follows a dual-service architecture:
1. **API Service** (`api/`): REST API gateway that handles HTTP requests
2. **RPC Service** (`rpc/`): gRPC backend service containing business logic

## Architecture

### Service Structure

```
api/                        # REST API Service (Port 9100)
├── desc/                   # API definitions in go-zero format (.api files)
│   ├── all.api            # Main aggregator importing all API files
│   ├── core/              # Core business APIs (user, role, menu, etc.)
│   ├── job/               # Job scheduler APIs
│   └── mcms/              # Message center APIs
├── internal/
│   ├── handler/           # HTTP request handlers
│   ├── logic/             # Business logic layer
│   ├── middleware/        # HTTP middleware
│   ├── svc/               # Service context
│   └── types/             # Generated types from API definitions
└── core.go                # API service entry point

rpc/                        # gRPC RPC Service (Port 9101)
├── desc/                   # Proto definitions organized by feature
├── ent/
│   ├── schema/            # Ent schema definitions (database models)
│   │   ├── user.go
│   │   ├── role.go
│   │   ├── menu.go
│   │   └── mixins/        # Reusable schema mixins
│   └── template/          # Custom Ent templates
├── internal/
│   ├── logic/             # RPC business logic
│   ├── server/            # gRPC server implementation
│   ├── svc/               # Service context
│   └── utils/             # Utility functions
├── types/                  # Generated protobuf types
└── core.proto             # Main proto file

deploy/
├── docker-compose/         # Docker Compose configurations
│   ├── all_in_one/        # Complete stack deployment
│   ├── core-only/         # Core services only
│   └── mysql_redis/       # Infrastructure services
└── k8s/                    # Kubernetes deployment manifests
```

### Core Concepts

**API Layer (api/)**: Acts as the HTTP gateway. Handlers receive requests, validate them, and forward to RPC service for processing. Uses JWT authentication and Casbin for authorization.

**RPC Layer (rpc/)**: Contains the actual business logic and database operations. Uses Ent ORM for database access. All CRUD operations happen here.

**Code Generation Flow**:
1. Define schema in `rpc/ent/schema/*.go`
2. Define protobuf in `rpc/desc/*.proto`
3. Define API in `api/desc/**/*.api`
4. Generate code using Makefile targets

**Service Communication**: The API service calls RPC service via gRPC. Service discovery is handled by go-zero (supports direct, etcd, k8s endpoints).

## Common Commands

### Development

```bash
# Install required tools
make tools

# Format code
make fmt

# Run linter
make lint

# Run tests
make test
```

### Code Generation

```bash
# Generate Ent code (after modifying schema files)
make gen-ent

# Generate RPC code (after modifying .proto files)
make gen-rpc

# Generate API code (after modifying .api files)
make gen-api

# Generate RPC logic from Ent schema
# Requires: model=<ModelName> group=<GroupName>
# Example: make gen-rpc-ent-logic model=User group=user
make gen-rpc-ent-logic model=<ModelName> group=<GroupName>

# Generate/update Swagger documentation
make gen-swagger

# Serve Swagger UI (port 36666)
make serve-swagger
```

### Building

```bash
# Build for Windows
make build-win

# Build for Linux
make build-linux

# Build for macOS
make build-mac

# Build Docker images
# Requires: DOCKER_USERNAME environment variable
make docker

# Publish Docker images
# Requires: DOCKER_USERNAME, DOCKER_PASSWORD, REPO environment variables
make publish-docker
```

### Running Services

```bash
# Run API service
go run api/core.go -f api/etc/core.yaml

# Run RPC service
go run rpc/core.go -f rpc/etc/core.yaml
```

## Configuration

### API Service Configuration
- File: `api/etc/core.yaml`
- Port: 9100 (default)
- Key settings:
  - `DatabaseConf`: Database connection
  - `RedisConf`: Redis connection
  - `Auth`: JWT settings
  - `CoreRpc.Target`: RPC service endpoint (supports k8s://, etcd://, or direct)
  - `CasbinConf`: RBAC model definition
  - `ProjectConf`: Business logic settings (default roles, departments, verification types)

### RPC Service Configuration
- File: `rpc/etc/core.yaml`
- Port: 9101 (default)
- Key settings:
  - `DatabaseConf`: Database connection
  - `RedisConf`: Redis connection
  - `CasbinConf`: RBAC model definition

## Development Workflow

### Spec-Driven Development (Recommended for New Features)

This project adopts **Spec-Driven Development** methodology inspired by GitHub's spec-kit. For complex features (3+ steps, non-trivial changes), follow this structured approach:

#### 1. Write Specification (spec.md)

Create a specification document in `specs/<feature-number>-<feature-name>/spec.md` using the spec-kit format:

**Required Sections**:
- **User Scenarios & Testing**: User stories with acceptance criteria (Given-When-Then format)
- **Requirements**: Functional requirements with testable criteria (FR-001, FR-002, etc.)
- **Key Entities**: Data models, files, and system components
- **Success Criteria**: Measurable outcomes (SC-001, SC-002, etc.)

**Example Structure**:
```markdown
# Feature Specification: [Feature Name]

**Feature Branch**: `<branch-name>`
**Created**: YYYY-MM-DD
**Status**: Draft

## User Scenarios & Testing (mandatory)

### User Story 1 - [Title] (Priority: P1)
[Description of user journey]

**Acceptance Scenarios**:
1. **Given** [state], **When** [action], **Then** [outcome]
2. **Given** [state], **When** [action], **Then** [outcome]

## Requirements (mandatory)

### Functional Requirements
- **FR-001**: System MUST [specific requirement]
- **FR-002**: User MUST be able to [capability]

### Key Entities
- **EntityName**: [description, location, attributes]

## Success Criteria (mandatory)

- **SC-001**: [Measurable outcome with target metric]
- **SC-002**: [User satisfaction metric]
```

**Reference Template**: See `specs/001-traditional-chinese-i18n/spec.md` for a complete example.

#### 2. Create Technical Plan (plan.md)

After spec approval, create `specs/<feature-number>-<feature-name>/plan.md`:

**Required Sections**:
- **Architecture Overview**: System diagrams, data flow
- **Technology Stack**: Languages, frameworks, libraries
- **Implementation Details**: Phase-by-phase breakdown with code examples
- **Testing Strategy**: Unit tests, E2E tests, manual QA
- **Performance Considerations**: Bundle size, runtime impact, optimizations
- **Deployment Strategy**: Rollout plan, feature flags, monitoring

**Example Structure**:
```markdown
# Technical Plan: [Feature Name]

**Related Spec**: [spec.md](./spec.md)
**Estimated Effort**: X-Y hours

## Architecture Overview
[System architecture diagram and data flow]

## Implementation Details

### Phase 1: [Component] (X hours)
[Detailed tasks with code examples]

### Phase 2: [Component] (Y hours)
[Detailed tasks with code examples]

## Timeline
| Phase | Duration | Dependencies |
```

**Reference Template**: See `specs/001-traditional-chinese-i18n/plan.md` for a complete example.

#### 3. Implement Following Plan

Execute the implementation phases defined in plan.md, following this sequence:

### Traditional Implementation Workflow (For Simple Changes)

For small bug fixes, configuration changes, or simple updates, use the direct implementation approach:

1. **Define Ent Schema** (if new entity needed):
   - Create/modify schema in `rpc/ent/schema/<entity>.go`
   - Run `make gen-ent`

2. **Define Proto Messages**:
   - Add message definitions to `rpc/desc/<feature>.proto` or `rpc/core.proto`
   - Define RPC service methods
   - Run `make gen-rpc`

3. **Generate RPC Logic** (optional, for CRUD):
   - Run `make gen-rpc-ent-logic model=<Model> group=<group>`
   - This auto-generates basic CRUD logic in `rpc/internal/logic/<group>/`

4. **Define API Endpoints**:
   - Create/modify API definition in `api/desc/core/<feature>.api`
   - Import it in `api/desc/all.api`
   - Run `make gen-api`

5. **Implement Logic**:
   - RPC logic: `rpc/internal/logic/<group>/<feature>_logic.go`
   - API logic: `api/internal/logic/<group>/<feature>_logic.go` (calls RPC)

6. **Update Documentation**:
   - Run `make gen-swagger` to regenerate Swagger docs

### When to Use Spec-Driven vs Traditional

**Use Spec-Driven Development for**:
- ✅ New features with multiple components (frontend + backend + database)
- ✅ Changes requiring user experience design
- ✅ Features that need stakeholder approval
- ✅ Complex refactoring affecting multiple services
- ✅ Features requiring cross-team coordination

**Use Traditional Workflow for**:
- ✅ Bug fixes
- ✅ Simple configuration changes
- ✅ Minor UI tweaks
- ✅ Dependency updates
- ✅ Performance optimizations

### Testing

- Test files should be in `api/internal/` or `rpc/internal/` subdirectories
- Run `make test` to execute all tests
- Tests cover both API and RPC layers

## Tools Required

- Go 1.25+
- `goctls`: go-zero code generation tool
- `swagger`: Swagger documentation generator
- `golangci-lint`: Code linter
- `protoc`: Protocol buffer compiler (for proto files)

## Database Migrations

Ent supports automatic migrations. The system will create/update tables on startup based on schema definitions in `rpc/ent/schema/`.

## Specification Document Format (spec-kit)

### Document Structure

All specification documents follow the [spec-kit format](https://github.com/github/spec-kit) to ensure clarity, testability, and alignment across the team.

#### spec.md Template

```markdown
# Feature Specification: [Feature Name]

**Feature Branch**: `<number>-<feature-slug>`
**Created**: YYYY-MM-DD
**Status**: Draft | Ready for Planning | In Progress | Completed
**Input**: [Brief description of what user/stakeholder requested]

## User Scenarios & Testing (mandatory)

### User Story 1 - [Brief Title] (Priority: P1)
[Describe user journey in plain language from user perspective]

**Why this priority**: [Explain business value and urgency]

**Independent Test**:
[How this can be tested independently without other features]

**Acceptance Scenarios**:
1. **Given** [initial state], **When** [user action], **Then** [expected outcome]
2. **Given** [initial state], **When** [user action], **Then** [expected outcome]

### Edge Cases
- What happens when [boundary condition]?
- How does system handle [error scenario]?
- What if [unexpected input]?

## Requirements (mandatory)

### Functional Requirements
- **FR-001**: System MUST [specific, testable capability]
- **FR-002**: Users MUST be able to [user-facing capability]
- **FR-003**: System MUST [data or behavior requirement]

### Key Entities
- **EntityName**: [What it represents, key attributes, relationships]
  - Location: `path/to/file`
  - Format: [file format or data structure]
  - Dependencies: [what it depends on]

## Success Criteria (mandatory)

### Measurable Outcomes
- **SC-001**: [Metric] reaches [target value]
- **SC-002**: [Performance metric] < [threshold]

### User Satisfaction Metrics
- **SC-003**: User feedback rating ≥ [target]

### Business Metrics
- **SC-004**: [Business impact metric]

## Out of Scope
[Explicitly list what is NOT included to prevent scope creep]
- ❌ [Feature/capability not included]
- ❌ [Future iteration item]

## Dependencies

### Internal Dependencies
- [Existing system component required]

### External Dependencies
- [Third-party library/service needed]
- Version requirements

### Team Dependencies
- [Role] for [task] ([hours] estimated)

## Risks & Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| [Risk description] | High/Medium/Low | High/Medium/Low | [How to prevent/handle] |

## Implementation Notes
[High-level phases without technical details - those go in plan.md]

## Review Checklist
- [ ] All user stories have clear acceptance criteria
- [ ] All functional requirements are testable
- [ ] Success criteria are measurable
- [ ] Dependencies identified and available
- [ ] Risks have mitigation strategies
```

#### plan.md Template

```markdown
# Technical Plan: [Feature Name]

**Related Spec**: [spec.md](./spec.md)
**Created**: YYYY-MM-DD
**Status**: Draft
**Estimated Effort**: X-Y hours

## Architecture Overview

### System Architecture
[ASCII diagram or description showing:
- Which services are affected (API/RPC)
- Data flow between components
- New files/modules to be created]

### Technology Stack
- **Backend**: [Language, framework, libraries]
- **Frontend**: [Framework, libraries]
- **Database**: [DBMS, ORM, migration strategy]
- **Tools**: [Development and build tools]

## Implementation Details

### Phase 1: [Component Name] (X-Y hours)

#### Task 1.1: [Task Name]
**File**: `path/to/file.go`

**Approach**:
[Detailed explanation of how to implement]

**Code Example**:
```go
// Actual code showing implementation
```

**Testing**:
```go
// Unit test examples
```

### Phase 2: [Component Name] (X-Y hours)
[Repeat structure]

## Performance Considerations
- **Bundle Size Impact**: [Impact on build size]
- **Runtime Performance**: [Expected performance characteristics]
- **Optimization Strategies**: [How to optimize if needed]

## Deployment Strategy

### Rollout Plan
1. Week 1: [Environment/stage]
2. Week 2: [Environment/stage]

### Feature Flag
```yaml
# Configuration example
FeatureName:
  Enabled: true
```

### Monitoring
- [Metrics to track]
- [Alerts to configure]

## Rollback Plan
1. [Step to safely rollback]
2. [Step to restore previous state]

## Success Metrics
✅ [Technical metric]
✅ [Performance metric]
✅ [Business metric]

## Timeline
| Phase | Duration | Dependencies |
|-------|----------|--------------|
| Phase 1 | X hours | [Dependency] |
| Phase 2 | Y hours | Phase 1 complete |

## Team Assignment
| Role | Responsibility | Hours |
|------|----------------|-------|
| [Role] | [Task] | Xh |
```

### Specification Naming Convention

- Specifications are numbered sequentially: `001`, `002`, `003`, etc.
- Directory naming: `specs/<number>-<feature-slug>/`
- Examples:
  - `specs/001-traditional-chinese-i18n/`
  - `specs/002-user-profile-upload/`
  - `specs/003-rbac-enhancement/`

### Key Principles

1. **User-Centric**: Start with user scenarios, not technical implementation
2. **Testable**: Every requirement must be verifiable
3. **Measurable**: Success criteria must have concrete targets
4. **Clear Scope**: Explicitly define what's in and out of scope
5. **Risk-Aware**: Identify risks upfront with mitigation strategies

### Integration with Project Management

- Specifications are tracked in Notion under project workspace
- Each spec.md maps to a Notion page for team collaboration
- Technical plan (plan.md) becomes child page of spec
- Status updates flow: Draft → Ready for Planning → In Progress → Completed

### Example Reference

See `specs/001-traditional-chinese-i18n/` for a complete, production-ready example of both spec.md and plan.md documents.

## AI Agent Collaboration System

When working on tasks in this project, Claude Code should adopt specialized AI agent roles to maximize efficiency and expertise. Each agent has specific responsibilities and expertise areas.

### Agent Roles

#### @backend - Backend Development Agent
**Expertise**: Go, Go-Zero, gRPC, Ent ORM, PostgreSQL, Redis, Casbin
**Responsibilities**:
- Design and implement RPC services (`rpc/internal/logic/`)
- Create and modify Ent schemas (`rpc/ent/schema/`)
- Define protobuf messages (`rpc/desc/*.proto`)
- Implement business logic and database operations
- Write unit tests for backend components
- Handle database migrations and optimizations
- Implement RBAC policies with Casbin

**When to Use**:
- Creating new CRUD modules
- Modifying database schemas
- Implementing gRPC services
- Writing backend business logic
- Performance optimization of queries

**Example Invocation**:
```
@backend: Please create a new Department entity with CRUD operations
@backend: Optimize the user query performance in the RPC layer
@backend: Implement RBAC policy for the new module
```

#### @api - API Gateway Agent
**Expertise**: REST API design, Go-Zero API, HTTP middleware, JWT, request validation
**Responsibilities**:
- Design and implement REST API endpoints (`api/desc/**/*.api`)
- Create API handlers (`api/internal/handler/`)
- Implement API-layer business logic (`api/internal/logic/`)
- Design middleware (auth, logging, rate limiting)
- Handle request/response transformation
- Integrate with RPC services
- Manage API versioning

**When to Use**:
- Designing REST API endpoints
- Implementing HTTP handlers
- Creating middleware
- API request validation
- Rate limiting implementation

**Example Invocation**:
```
@api: Create REST endpoints for the Department module
@api: Add rate limiting middleware for authentication endpoints
@api: Design API contract for file upload feature
```

#### @frontend - Frontend Development Agent
**Expertise**: Vue 3, Vben5, TypeScript, Ant Design Vue, i18n, Pinia
**Responsibilities**:
- Implement UI components
- Design page layouts and navigation
- Integrate with backend APIs
- Handle state management with Pinia
- Implement internationalization (i18n)
- Create reactive forms with validation
- Optimize frontend performance

**When to Use**:
- Creating new pages or components
- Implementing UI features
- Integrating APIs with frontend
- Adding i18n support
- State management

**Example Invocation**:
```
@frontend: Create a Department management page with CRUD operations
@frontend: Add Traditional Chinese language support to the login page
@frontend: Implement a data table with pagination for users
```

#### @ui-ux - UI/UX Design Agent
**Expertise**: User interface design, user experience, accessibility, responsive design, Ant Design patterns
**Responsibilities**:
- Design user-friendly interfaces
- Create consistent UI patterns
- Ensure accessibility (WCAG compliance)
- Design responsive layouts
- Create design mockups and wireframes
- Define color schemes and typography
- Validate user flows

**When to Use**:
- Designing new features from UX perspective
- Improving existing UI layouts
- Ensuring accessibility
- Creating design specifications
- User flow optimization

**Example Invocation**:
```
@ui-ux: Design the user flow for multi-step form submission
@ui-ux: Review the Dashboard layout for mobile responsiveness
@ui-ux: Create a design spec for the new notification system
```

#### @devops - DevOps & Infrastructure Agent
**Expertise**: Docker, Kubernetes, CI/CD, monitoring, deployment, infrastructure
**Responsibilities**:
- Manage Docker configurations (`deploy/docker-compose/`)
- Create and maintain Kubernetes manifests (`deploy/k8s/`)
- Set up CI/CD pipelines
- Configure monitoring and logging
- Handle deployment strategies
- Manage environment configurations
- Optimize build processes

**When to Use**:
- Setting up deployment pipelines
- Configuring Docker/K8s
- Implementing monitoring
- Optimizing build times
- Infrastructure scaling

**Example Invocation**:
```
@devops: Create a GitHub Actions workflow for automated testing
@devops: Update Docker Compose configuration for development environment
@devops: Set up Prometheus monitoring for the RPC service
```

#### @qa - Quality Assurance Agent
**Expertise**: Testing strategies, test automation, E2E testing, integration testing
**Responsibilities**:
- Design test strategies
- Write integration tests
- Create E2E test scenarios
- Perform code quality reviews
- Validate acceptance criteria
- Test i18n implementations
- Performance testing

**When to Use**:
- Creating test plans
- Writing integration tests
- Validating feature completeness
- Testing edge cases
- Quality reviews

**Example Invocation**:
```
@qa: Create integration tests for the Department CRUD operations
@qa: Design E2E test scenarios for the login flow
@qa: Review the Traditional Chinese i18n implementation
```

#### @doc - Documentation Agent
**Expertise**: Technical writing, API documentation, code documentation, user guides
**Responsibilities**:
- Write and maintain documentation
- Generate/update Swagger documentation
- Create user guides and tutorials
- Document API contracts
- Maintain CLAUDE.md and README.md
- Create architectural diagrams
- Write code comments

**When to Use**:
- Documenting new features
- Updating API documentation
- Creating user guides
- Writing architectural docs
- Code documentation

**Example Invocation**:
```
@doc: Update Swagger documentation for the new Department API
@doc: Create a user guide for the RBAC configuration
@doc: Document the Traditional Chinese i18n implementation process
```

#### @pm - Project Manager Agent
**Expertise**: Task tracking, workflow management, Notion integration, git operations, status reporting
**Responsibilities**:
- Track task progress in Notion Tasks database
- Update task status after code commits
- Ensure git commit messages follow conventions
- Synchronize task status between codebase and Notion
- Generate progress reports and status updates
- Validate task completion against acceptance criteria
- Manage task dependencies and blocking relationships
- Create and update project documentation in Notion

**When to Use**:
- Starting work on a Notion task
- After completing a task (update status to "Done")
- After git commit (sync status to Notion)
- When checking project progress
- When generating status reports
- When task is blocked or needs escalation

**Workflow Protocol**:
1. **Task Start**: When beginning work on a Notion task
   - Update task status from "Not started" → "In progress"
   - Record start time in task properties
   - Verify task dependencies are met

2. **During Implementation**: While working on the task
   - Keep task description updated with progress notes
   - Update estimated hours if needed
   - Flag blockers immediately by updating "Blocked by" field

3. **After Code Commit**: When code is committed to git
   - Verify commit message follows convention (feat/fix/refactor/etc.)
   - Add commit hash to task description
   - Update task progress percentage
   - Link related PRs to task

4. **Task Completion**: When task is fully implemented
   - Verify all acceptance criteria are met
   - Update task status to "Done"
   - Record actual hours spent
   - Update related documentation links
   - Create handoff notes for dependent tasks

5. **Status Reporting**: Generate progress updates
   - Daily: Update task progress for active tasks
   - Weekly: Generate progress report for project
   - On-demand: Create status summary when requested

**Example Invocation**:
```
@pm: I'm starting work on [ZH-TW-001], please update Notion status
@pm: I've completed [ZH-TW-001] with commit abc123, update task status to Done
@pm: Generate a progress report for Traditional Chinese i18n feature
@pm: Check if [ZH-TW-005] can start (verify dependencies)
@pm: Task [ZH-TW-003] is blocked by missing test data, update Notion
```

**Integration Points**:
- **Notion API**: Direct integration with Tasks database
- **Git**: Read commit history, verify commit messages
- **Project Tracking**: Links tasks to specs, plans, and PRs
- **Status Sync**: Bidirectional sync between code and Notion

**Key Metrics Tracked**:
- Task completion rate
- Estimated vs actual hours
- Blocked task count
- Tasks by status (Not started, In progress, Done, Cancelled)
- Tasks by agent assignment
- Critical path analysis

### Multi-Agent Collaboration

For complex tasks requiring multiple specializations, agents should coordinate:

**Example 1: Full-Stack Feature Implementation**
```
User: Implement a Department management module

Coordination Plan:
1. @backend: Create Ent schema, Proto definitions, RPC logic
2. @api: Design REST endpoints, implement handlers
3. @frontend: Create Department CRUD page with data table
4. @ui-ux: Review and improve the UI layout
5. @qa: Create integration and E2E tests
6. @doc: Update API documentation and user guide
```

**Example 2: Internationalization Feature with PM Tracking**
```
User: Add Traditional Chinese support

@pm: Starting Traditional Chinese i18n feature, update Notion tasks

Coordination Plan:
1. @backend: Update Ent schema with locale field, implement i18n logic
   @pm: Mark [ZH-TW-007] as "In progress"

2. @api: Add locale parameter to API contracts
   @pm: Update [ZH-TW-007] to "Done", start [ZH-TW-008]

3. @frontend: Create language files, implement language selector
   @pm: Mark [ZH-TW-004], [ZH-TW-005], [ZH-TW-006] as "In progress"

4. @ui-ux: Review terminology and UI text placement
   @pm: Update [ZH-TW-006] with UI feedback

5. @qa: Test all UI text and language switching
   @pm: Start [ZH-TW-009], [ZH-TW-010]

6. @doc: Document i18n usage and translation process
   @pm: Generate final completion report, update all task statuses to "Done"
```

**Example 3: PM-Driven Task Workflow**
```
User: I want to start implementing [ZH-TW-001]

@pm: Check task dependencies and status
→ Verifies: No blocking dependencies
→ Updates Notion: Status "Not started" → "In progress"
→ Records: Start time, assigned agent (@backend)
→ Responds: "Task [ZH-TW-001] is ready. No blockers. Estimated: 4-6h"

@backend: Implement zh-TW.json language file
→ Creates file, commits code
→ Commit message: "feat: add Traditional Chinese language file (zh-TW.json)"

@pm: Task [ZH-TW-001] completed with commit abc123
→ Verifies: All acceptance criteria met
→ Updates Notion: Status "In progress" → "Done"
→ Records: Commit hash, actual hours (5h), completion timestamp
→ Checks: If dependent tasks can now start
→ Notifies: "[ZH-TW-002] is now unblocked and ready"
```

### Agent Selection Guidelines

**Single Agent Tasks** (use one specialized agent):
- Backend-only changes (schema, RPC logic) → @backend
- API-only changes (endpoints, middleware) → @api
- Frontend-only changes (UI, components) → @frontend
- Documentation updates → @doc
- Deployment configuration → @devops
- Project tracking and status updates → @pm

**Multi-Agent Tasks** (coordinate multiple agents):
- New feature with full-stack implementation → @pm coordinates @backend, @api, @frontend
- UI/UX redesign with implementation → @pm coordinates @ui-ux, @frontend
- Performance optimization across layers → @pm coordinates all technical agents
- Major refactoring affecting multiple components → @pm tracks progress across agents

**PM Always Involved When**:
- Starting any Notion task
- Completing any Notion task
- Making git commits
- Generating progress reports
- Managing task dependencies
- Escalating blockers

### Agent Communication Protocol

When multiple agents need to collaborate:

1. **Task Decomposition**: @pm breaks down the task by responsibility area and creates Notion tasks
2. **Dependency Order**: @pm ensures tasks execute in proper order (backend → api → frontend)
3. **Interface Agreement**: Agents must agree on contracts (Proto, API spec), @pm documents agreements
4. **Progress Tracking**: @pm updates Notion task status after each phase completion
5. **Integration Testing**: @qa validates integration, @pm tracks test results
6. **Documentation**: @doc creates final documentation, @pm links docs to tasks
7. **Completion**: @pm verifies all acceptance criteria and closes tasks

**Example Collaboration with PM Orchestration**:
```
User: Implement file upload feature with preview

@pm: Create feature tasks in Notion
→ [FILE-001] Backend: File entity and RPC service
→ [FILE-002] API: REST endpoint for file upload
→ [FILE-003] Frontend: Upload component with preview
→ [FILE-004] QA: E2E testing
→ [FILE-005] Doc: API documentation

@pm: Start [FILE-001], assign @backend
@backend: Create File entity schema and RPC service for upload
@backend: Commit with message "feat: add file upload RPC service"
@pm: [FILE-001] completed (commit: abc123), update Notion to "Done"
↓ (Produces: Proto contract)

@pm: Start [FILE-002], assign @api
@api: Implement REST API endpoint for file upload
@api: Commit with message "feat: add file upload API endpoint"
@pm: [FILE-002] completed (commit: def456), update Notion to "Done"
↓ (Produces: API spec)

@pm: Start [FILE-003], assign @frontend
@frontend: Create file upload component with preview
@frontend: Commit with message "feat: add file upload UI component"
@pm: [FILE-003] completed (commit: ghi789), update Notion to "Done"
↓ (Produces: UI implementation)

@pm: Notify @ui-ux for review
@ui-ux: Review upload UX and provide feedback
@pm: Record UX feedback in [FILE-003] task notes

@pm: Start [FILE-004], assign @qa
@qa: Test file upload flow E2E
@pm: [FILE-004] completed, all tests passing

@pm: Start [FILE-005], assign @doc
@doc: Document file upload API and usage
@pm: [FILE-005] completed

@pm: Generate completion report
→ Feature: File Upload with Preview
→ Total commits: 3 (abc123, def456, ghi789)
→ Tasks completed: 5/5 (100%)
→ Total time: 18 hours (vs estimated 20 hours)
→ Status: ✅ All acceptance criteria met
```

### Agent Context Switching

When switching between agent roles during a conversation:

1. **Declare Role**: Explicitly state which agent role you're adopting
2. **Scope of Work**: Clearly define what you'll work on in this role
3. **Handoff**: When passing to another agent, summarize what was completed

**Example**:
```
[As @backend] I've completed the RPC service implementation.
The Proto contract defines: CreateDepartment, UpdateDepartment, DeleteDepartment, GetDepartment, ListDepartments.

[Switching to @api] Now implementing REST API endpoints based on the RPC contract...

[Handoff to @frontend] API endpoints are ready at /api/v1/departments.
Please implement the UI using these endpoints.
```

## Important Notes

- **File Naming**: Uses `go_zero` style (snake_case) as defined in Makefile
- **I18n Enabled**: Project supports internationalization (`PROJECT_I18N=true`)
- **Ent Features**: `sql/execquery`, `intercept`, `sql/modifier` are enabled
- **Custom Forks**: Project uses forked versions of go-zero (`simple-admin-tools`) and base64Captcha (see `go.mod` replace directives)
- **Service Dependencies**: API service depends on RPC service - start RPC first
- **Configuration**: Both services require database and Redis to be running
- **Spec-Driven Development**: For complex features, always create spec.md before coding
- **AI Agent Roles**: Use specialized agent roles (@backend, @api, @frontend, @ui-ux, @devops, @qa, @doc) to improve task efficiency

## Module System

Simple Admin is modular. This is the Core module. Other official modules:
- **FMS**: File management
- **Job**: Scheduled tasks (partially integrated)
- **MMS**: Member management
- **MCMS**: Message center (partially integrated)

The API layer includes handlers for Job and MCMS modules, but they require separate RPC services to be deployed.

## Deployment

- **Docker Compose**: See `deploy/docker-compose/` for various deployment configurations
- **Kubernetes**: See `deploy/k8s/` for k8s manifests
- Service discovery works natively with k8s endpoints (`k8s://default/<service>:<port>`)

## Commit Message Convention

Follow Angular/Vue style:
- `feat`: New feature
- `fix`: Bug fix
- `refactor`: Code refactoring
- `perf`: Performance improvement
- `docs`: Documentation
- `style`: Code style changes
- `test`: Tests
- `chore`: Dependency updates/configuration
