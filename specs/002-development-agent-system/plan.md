# Technical Plan: Development Agent System

**Related Spec**: [spec.md](./spec.md)
**Created**: 2025-10-07
**Status**: Draft
**Estimated Effort**: 120-150 hours (6-8 weeks)

## Architecture Overview

### System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    Developer Interface                           │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐   │
│  │  CLI Tool      │  │ GitHub Actions │  │  VSCode Ext    │   │
│  │  agent run     │  │  on: pr_opened │  │  @agent cmd    │   │
│  └───────┬────────┘  └───────┬────────┘  └───────┬────────┘   │
└──────────┼────────────────────┼────────────────────┼────────────┘
           │                    │                    │
           └────────────────────┼────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Agent Orchestrator                            │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Job Queue (Redis)                                        │  │
│  │  - Priority queue                                         │  │
│  │  - Job persistence                                        │  │
│  │  - Distributed locks                                      │  │
│  └──────────────────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  Agent Registry                                           │  │
│  │  - Agent metadata                                         │  │
│  │  - Health checks                                          │  │
│  │  - Load balancing                                         │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
           │
           ├─────────────┬─────────────┬─────────────┬──────────┐
           │             │             │             │          │
           ▼             ▼             ▼             ▼          ▼
┌──────────────┐  ┌──────────┐  ┌──────────┐  ┌────────┐  ┌────────┐
│  @generator  │  │ @reviewer│  │ @deployer│  │ @doc   │  │ @spec  │
│              │  │          │  │          │  │ -sync  │  │ -writer│
│  Code Gen    │  │ Code     │  │ Deploy   │  │ Docs   │  │ Spec   │
│  Templates   │  │ Analysis │  │ Docker   │  │ Swagger│  │ Tasks  │
└──────┬───────┘  └─────┬────┘  └─────┬────┘  └───┬────┘  └───┬────┘
       │                │              │           │           │
       └────────────────┴──────────────┴───────────┴───────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Shared Services Layer                         │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐       │
│  │ LLM API  │  │ Git Ops  │  │ Linter   │  │ Notion   │       │
│  │ Claude   │  │ GitHub   │  │ golangci │  │ API      │       │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘       │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Data & Logs Layer                             │
│  ┌──────────────────┐  ┌──────────────────┐                    │
│  │ PostgreSQL       │  │ Elasticsearch    │                    │
│  │ - Agent jobs     │  │ - Agent logs     │                    │
│  │ - Actions        │  │ - Metrics        │                    │
│  └──────────────────┘  └──────────────────┘                    │
└─────────────────────────────────────────────────────────────────┘
```

### Agent Communication Protocol

```
┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│   Trigger   │  →HTTP  │ Orchestrator│  →Redis │    Agent    │
│  (CLI/GHA)  │         │   (Go API)  │  Queue  │  (Workers)  │
└─────────────┘         └─────────────┘         └─────────────┘
                                │
                                ▼
                        ┌───────────────┐
                        │  Job Message  │
                        ├───────────────┤
                        │ job_id        │
                        │ agent_name    │
                        │ input_params  │
                        │ priority      │
                        │ timeout       │
                        └───────────────┘
```

## Technology Stack

### Core Infrastructure
- **Language**: Go 1.25
- **Queue**: Redis 7.x (Bull/BullMQ style job queue)
- **Database**: PostgreSQL 15+ (agent job persistence)
- **Logs**: Elasticsearch + Kibana
- **Container**: Docker + Docker Compose

### Agent Runtime
- **LLM**: Claude 3.5 Sonnet API (Anthropic)
- **Code Analysis**:
  - golangci-lint
  - staticcheck
  - go-vet
- **Git Operations**: go-git library
- **Template Engine**: Go templates + Sprig functions

### Integrations
- **GitHub**: GitHub API v3 (REST) + v4 (GraphQL)
- **Notion**: Notion API (official SDK)
- **Slack**: Slack Webhooks
- **CI/CD**: GitHub Actions

## Implementation Details

### Phase 1: Core Infrastructure (Week 1-2, 30-40 hours)

#### Task 1.1: Agent Orchestrator Service
**File**: `tools/agent-system/cmd/orchestrator/main.go`

**Architecture**:
```go
package main

type Orchestrator struct {
    jobQueue    *redis.Client
    registry    *AgentRegistry
    db          *sql.DB
    logger      *zap.Logger
}

func (o *Orchestrator) SubmitJob(ctx context.Context, req *JobRequest) (*JobResponse, error) {
    // 1. Validate request
    // 2. Create job record in DB
    // 3. Enqueue to Redis
    // 4. Return job ID
}

func (o *Orchestrator) MonitorJobs(ctx context.Context) {
    // Background goroutine monitoring job status
    // Update Notion, send notifications
}
```

**Database Schema**:
```sql
CREATE TABLE agent_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    agent_name VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL, -- pending, running, success, failed, rolled_back
    input_params JSONB NOT NULL,
    output_result JSONB,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_by VARCHAR(100) -- GitHub username or CLI user
);

CREATE TABLE agent_actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id UUID REFERENCES agent_jobs(id),
    action_type VARCHAR(50) NOT NULL, -- file_create, file_modify, git_commit, api_call
    target VARCHAR(500) NOT NULL, -- file path, API endpoint
    changes JSONB, -- diff or request payload
    requires_approval BOOLEAN DEFAULT FALSE,
    approved_by VARCHAR(100),
    approved_at TIMESTAMP,
    executed_at TIMESTAMP
);

CREATE INDEX idx_jobs_status ON agent_jobs(status, created_at);
CREATE INDEX idx_actions_job ON agent_actions(job_id);
```

**Testing**:
```go
func TestOrchestrator_SubmitJob(t *testing.T) {
    orch := setupTestOrchestrator()

    req := &JobRequest{
        AgentName: "generator",
        Params: map[string]interface{}{
            "entity": "Department",
            "fields": []string{"name", "parent_id"},
        },
    }

    resp, err := orch.SubmitJob(context.Background(), req)
    assert.NoError(t, err)
    assert.NotEmpty(t, resp.JobID)
}
```

#### Task 1.2: CLI Tool
**File**: `tools/agent-system/cmd/agent/main.go`

```go
package main

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "agent",
    Short: "Simple Admin Agent System CLI",
}

var runCmd = &cobra.Command{
    Use:   "run [agent] [args...]",
    Short: "Run an agent",
    Example: `
  agent run generator create-module Department
  agent run reviewer --pr 123
  agent run deployer --env staging --version v1.7.2
    `,
    Run: func(cmd *cobra.Command, args []string) {
        // Parse agent name and params
        // Call orchestrator API
        // Stream logs to stdout
    },
}

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all agents",
    Run: func(cmd *cobra.Command, args []string) {
        // Query orchestrator for agent registry
        // Display in table format
    },
}

func init() {
    rootCmd.AddCommand(runCmd)
    rootCmd.AddCommand(listCmd)
}

func main() {
    rootCmd.Execute()
}
```

**Installation**:
```bash
# Build CLI tool
cd tools/agent-system
go build -o agent cmd/agent/main.go

# Install globally
sudo mv agent /usr/local/bin/

# Usage
agent list
agent run generator create-module Department
agent run reviewer --pr 123
```

#### Task 1.3: Job Queue (Redis)
**File**: `tools/agent-system/pkg/queue/redis_queue.go`

```go
package queue

type RedisQueue struct {
    client *redis.Client
}

func NewRedisQueue(addr string) *RedisQueue {
    return &RedisQueue{
        client: redis.NewClient(&redis.Options{Addr: addr}),
    }
}

func (q *RedisQueue) Enqueue(ctx context.Context, job *Job) error {
    data, _ := json.Marshal(job)

    // Use sorted set for priority queue
    score := float64(job.Priority)
    return q.client.ZAdd(ctx, "agent:jobs:pending", &redis.Z{
        Score:  score,
        Member: data,
    }).Err()
}

func (q *RedisQueue) Dequeue(ctx context.Context) (*Job, error) {
    // Pop highest priority job
    result, err := q.client.ZPopMax(ctx, "agent:jobs:pending").Result()
    if err != nil || len(result) == 0 {
        return nil, ErrQueueEmpty
    }

    var job Job
    json.Unmarshal([]byte(result[0].Member.(string)), &job)
    return &job, nil
}

func (q *RedisQueue) AcquireLock(ctx context.Context, resource string, ttl time.Duration) (bool, error) {
    // Distributed lock for file-level coordination
    return q.client.SetNX(ctx, "agent:lock:"+resource, "locked", ttl).Result()
}
```

### Phase 2: Code Generation Agent (@generator) (Week 3, 20-25 hours)

#### Task 2.1: Template System
**Directory**: `tools/agent-system/templates/`

```
templates/
├── ent/
│   └── entity.go.tmpl        # Ent schema template
├── proto/
│   └── message.proto.tmpl    # Proto message template
├── api/
│   └── endpoint.api.tmpl     # API definition template
└── rpc/
    └── logic.go.tmpl         # RPC logic template
```

**Ent Template Example**:
```go
// templates/ent/entity.go.tmpl
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/mixin"
)

// {{ .EntityName }} holds the schema definition for the {{ .EntityName }} entity.
type {{ .EntityName }} struct {
    ent.Schema
}

func ({{ .EntityName }}) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixins.BaseMixin{},
        mixins.StatusMixin{},
    }
}

func ({{ .EntityName }}) Fields() []ent.Field {
    return []ent.Field{
        {{ range .Fields }}
        field.{{ .Type }}("{{ .Name }}").
            Comment("{{ .Comment }}").
            {{ if .Optional }}Optional().{{ end }}
            {{ if .Unique }}Unique().{{ end }}
            {{ if .Default }}Default({{ .Default }}).{{ end }},
        {{ end }}
    }
}

func ({{ .EntityName }}) Edges() []ent.Edge {
    return []ent.Edge{
        {{ range .Edges }}
        edge.{{ .Type }}("{{ .Name }}", {{ .Target }}.Type).
            {{ if .Required }}Required().{{ end }},
        {{ end }}
    }
}
```

#### Task 2.2: Generator Agent Implementation
**File**: `tools/agent-system/agents/generator/generator.go`

```go
package generator

type GeneratorAgent struct {
    llm         *claude.Client
    templates   *template.Template
    fileWriter  *FileWriter
}

func (a *GeneratorAgent) Execute(ctx context.Context, input *GeneratorInput) (*GeneratorOutput, error) {
    // Step 1: Analyze input with LLM
    entitySpec, err := a.analyzeEntityRequirements(ctx, input.Description)
    if err != nil {
        return nil, err
    }

    // Step 2: Generate Ent schema
    entFile, err := a.generateEntSchema(entitySpec)
    if err != nil {
        return nil, err
    }

    // Step 3: Generate Proto message
    protoFile, err := a.generateProtoMessage(entitySpec)
    if err != nil {
        return nil, err
    }

    // Step 4: Generate API definition
    apiFile, err := a.generateAPIDefinition(entitySpec)
    if err != nil {
        return nil, err
    }

    // Step 5: Generate RPC logic
    rpcFiles, err := a.generateRPCLogic(entitySpec)
    if err != nil {
        return nil, err
    }

    // Step 6: Run code generation
    if err := a.runCodeGeneration(); err != nil {
        return nil, err
    }

    // Step 7: Run tests
    if err := a.runTests(); err != nil {
        return nil, err
    }

    return &GeneratorOutput{
        Files: append([]string{entFile, protoFile, apiFile}, rpcFiles...),
        Commands: []string{
            "make gen-ent",
            "make gen-rpc",
            "make gen-api",
        },
    }, nil
}

func (a *GeneratorAgent) analyzeEntityRequirements(ctx context.Context, description string) (*EntitySpec, error) {
    prompt := fmt.Sprintf(`
Analyze this entity description and extract structured information:

Description: %s

Output JSON with:
- entity_name: CamelCase name
- table_name: snake_case name
- fields: array of {name, type, optional, unique, default, comment}
- relationships: array of {name, type (one-to-one, one-to-many), target_entity, required}
- i18n_keys: array of translation keys needed

Follow Simple Admin Core conventions:
- Use mixins.BaseMixin (id, created_at, updated_at, created_by, updated_by)
- Use mixins.StatusMixin (status uint8, NORMAL=1, BAN=2)
- Field names use snake_case
- Entity names use CamelCase
`, description)

    resp, err := a.llm.Complete(ctx, prompt)
    if err != nil {
        return nil, err
    }

    var spec EntitySpec
    json.Unmarshal([]byte(resp), &spec)
    return &spec, nil
}
```

**Usage Example**:
```bash
# CLI
agent run generator create-module Department \
    --description "Department entity with name, parent_id, sort, remark fields. Tree structure with parent-child relationship."

# Output
✓ Generated rpc/ent/schema/department.go
✓ Generated rpc/desc/department.proto
✓ Generated api/desc/core/department.api
✓ Generated rpc/internal/logic/department/create_department_logic.go
✓ Generated rpc/internal/logic/department/update_department_logic.go
✓ Generated rpc/internal/logic/department/delete_department_logic.go
✓ Generated rpc/internal/logic/department/get_department_list_logic.go
✓ Ran make gen-ent (success)
✓ Ran make gen-rpc (success)
✓ Ran make gen-api (success)
✓ All tests passed

Next steps:
1. Review generated files
2. Add business logic to RPC handlers
3. Add i18n translations to zh-TW.json, en.json
4. Create pull request
```

### Phase 3: Code Review Agent (@reviewer) (Week 4, 25-30 hours)

#### Task 3.1: Review Engine
**File**: `tools/agent-system/agents/reviewer/reviewer.go`

```go
package reviewer

type ReviewerAgent struct {
    llm        *claude.Client
    linter     *Linter
    github     *github.Client
    i18nChecker *I18nChecker
}

type ReviewConfig struct {
    PRNumber      int
    Repo          string
    CheckLinting  bool
    CheckI18n     bool
    CheckTests    bool
    CheckSecurity bool
}

func (a *ReviewerAgent) Execute(ctx context.Context, config *ReviewConfig) (*ReviewResult, error) {
    // Step 1: Fetch PR diff
    pr, err := a.github.GetPullRequest(ctx, config.Repo, config.PRNumber)
    if err != nil {
        return nil, err
    }

    diff, err := a.github.GetPRDiff(ctx, config.Repo, config.PRNumber)
    if err != nil {
        return nil, err
    }

    // Step 2: Run linters
    var issues []Issue
    if config.CheckLinting {
        lintIssues, _ := a.linter.Check(diff)
        issues = append(issues, lintIssues...)
    }

    // Step 3: Check i18n keys
    if config.CheckI18n {
        i18nIssues, _ := a.i18nChecker.Check(diff)
        issues = append(issues, i18nIssues...)
    }

    // Step 4: Check test coverage
    if config.CheckTests {
        testIssues, _ := a.checkTestCoverage(diff)
        issues = append(issues, testIssues...)
    }

    // Step 5: LLM-based code review
    llmIssues, err := a.reviewWithLLM(ctx, diff)
    if err != nil {
        return nil, err
    }
    issues = append(issues, llmIssues...)

    // Step 6: Post review comments
    if len(issues) > 0 {
        for _, issue := range issues {
            a.github.CreateReviewComment(ctx, config.Repo, config.PRNumber, &issue)
        }
    }

    return &ReviewResult{
        TotalIssues:    len(issues),
        CriticalIssues: countCritical(issues),
        Suggestion:     a.generateSuggestion(issues),
    }, nil
}

func (a *ReviewerAgent) reviewWithLLM(ctx context.Context, diff string) ([]Issue, error) {
    prompt := fmt.Sprintf(`
Review this Go code changes for Simple Admin Core project:

%s

Check for:
1. Go best practices violations
2. Error handling issues (unchecked errors)
3. Potential bugs or race conditions
4. Security vulnerabilities
5. Missing i18n translations
6. API/RPC contract mismatches
7. Missing test coverage

Output JSON array of issues:
[
  {
    "file": "path/to/file.go",
    "line": 42,
    "severity": "error|warning|info",
    "category": "error_handling|security|i18n|testing|style",
    "message": "Description of issue",
    "suggestion": "How to fix"
  }
]
`, diff)

    resp, err := a.llm.Complete(ctx, prompt)
    if err != nil {
        return nil, err
    }

    var issues []Issue
    json.Unmarshal([]byte(resp), &issues)
    return issues, nil
}
```

**GitHub Action Integration**:
```yaml
# .github/workflows/agent-review.yml
name: Agent Code Review

on:
  pull_request:
    types: [opened, synchronize]

jobs:
  review:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Run Agent Reviewer
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
        run: |
          agent run reviewer \
            --pr ${{ github.event.pull_request.number }} \
            --repo ${{ github.repository }} \
            --check-linting \
            --check-i18n \
            --check-tests \
            --auto-comment
```

### Phase 4: Other Agents (Week 5-8, 60-70 hours)

#### @deployer Agent (15-20 hours)
**Key Features**:
- Build Docker images with tags (version, git SHA, latest)
- Push to Docker registry
- Update Kubernetes manifests (image tags)
- Apply to cluster with kubectl
- Run smoke tests
- Monitor rollout status
- Auto-rollback on failure
- Send deployment notification to Notion + Slack

**Code Snippet**:
```go
func (a *DeployerAgent) Execute(ctx context.Context, input *DeployInput) error {
    // 1. Build images
    for _, service := range []string{"api", "rpc"} {
        if err := a.buildDockerImage(service, input.Version); err != nil {
            return err
        }
    }

    // 2. Push to registry
    if err := a.pushImages(input.Version); err != nil {
        return err
    }

    // 3. Update k8s manifests
    if err := a.updateManifests(input.Environment, input.Version); err != nil {
        return err
    }

    // 4. Apply to cluster
    if err := a.applyManifests(input.Environment); err != nil {
        return err
    }

    // 5. Monitor rollout
    if err := a.monitorRollout(ctx, input.Environment, 5*time.Minute); err != nil {
        a.rollback(input.Environment)
        return err
    }

    // 6. Run smoke tests
    if err := a.runSmokeTests(input.Environment); err != nil {
        a.rollback(input.Environment)
        return err
    }

    // 7. Notify
    a.notifySuccess(input.Environment, input.Version)
    return nil
}
```

#### @doc-sync Agent (10-15 hours)
**Key Features**:
- Detect API changes (new endpoints, modified schemas)
- Run `make gen-swagger`
- Update CLAUDE.md (Common Commands, Architecture sections)
- Generate changelog entry
- Create/update Notion documentation pages
- Commit changes to docs branch

#### @test-gen Agent (15-20 hours)
**Key Features**:
- Parse Go function signatures
- Identify test scenarios (happy path, error cases, edge cases)
- Generate table-driven tests
- Create mock objects for dependencies
- Generate test fixtures (data files)
- Ensure >70% code coverage

#### @migrator Agent (10-15 hours)
**Key Features**:
- Analyze Ent schema changes
- Generate forward SQL migration
- Generate rollback SQL migration
- Include transaction handling
- Validate with `psql --dry-run`
- Create migration version file

#### @spec-writer Agent (15-20 hours)
**Key Features**:
- Interactive CLI prompts for feature requirements
- Generate spec.md following spec-kit format
- Analyze technical feasibility
- Generate plan.md with implementation phases
- Break down into Notion tasks with dependencies
- Link all artifacts together

## Performance Considerations

### LLM API Optimization
- **Prompt Caching**: Cache common prompts (templates, project conventions)
- **Response Caching**: Cache responses for identical inputs (1 hour TTL)
- **Batch Processing**: Group multiple small requests into one
- **Streaming**: Use streaming for long responses to provide progress feedback

### Resource Usage
- **CPU**: Agent workers run in separate goroutines, limit to 4 concurrent agents
- **Memory**: Each agent worker uses ~100MB, set max 1GB per agent process
- **Disk**: Store logs in Elasticsearch, rotate after 30 days
- **Network**: LLM API calls average 2KB request, 10KB response

### Scalability
- **Horizontal Scaling**: Deploy multiple agent workers, coordinate via Redis locks
- **Queue Priority**: Critical jobs (PR reviews) have higher priority than batch jobs (doc sync)
- **Rate Limiting**: Respect LLM API rate limits (50 req/min), implement exponential backoff

## Deployment Strategy

### Development Environment
```bash
# Start all services
cd tools/agent-system
docker-compose up -d

# Services:
# - Redis: localhost:6379
# - PostgreSQL: localhost:5432
# - Elasticsearch: localhost:9200
# - Orchestrator API: localhost:8080
```

### Production Deployment
```yaml
# deploy/k8s/agent-system/
apiVersion: apps/v1
kind: Deployment
metadata:
  name: agent-orchestrator
spec:
  replicas: 2
  template:
    spec:
      containers:
      - name: orchestrator
        image: ryanpower/agent-orchestrator:v1.0.0
        env:
        - name: REDIS_URL
          value: redis://redis-service:6379
        - name: DB_URL
          valueFrom:
            secretKeyRef:
              name: agent-secrets
              key: db-url
        - name: ANTHROPIC_API_KEY
          valueFrom:
            secretKeyRef:
              name: agent-secrets
              key: anthropic-key
```

### Rollout Plan
1. **Week 1**: Deploy infrastructure (orchestrator, queue, DB)
2. **Week 2**: Deploy @generator agent (internal testing)
3. **Week 3**: Deploy @reviewer agent (beta with 2-3 PRs)
4. **Week 4**: Deploy @deployer agent (staging environment only)
5. **Week 5-8**: Deploy remaining agents, production rollout

## Monitoring & Observability

### Metrics to Track
- **Job Metrics**:
  - Jobs submitted/completed/failed per hour
  - Average job duration by agent type
  - Queue depth over time
- **Agent Metrics**:
  - Agent uptime
  - Success rate by agent
  - LLM API latency (p50, p95, p99)
- **Business Metrics**:
  - Development velocity (story points/sprint)
  - Code review cycle time (PR open to merge)
  - Deployment frequency and success rate

### Alerting Rules
```yaml
# Prometheus alerting rules
groups:
- name: agent_system
  rules:
  - alert: HighJobFailureRate
    expr: rate(agent_jobs_failed[5m]) > 0.2
    annotations:
      summary: "Agent job failure rate >20%"

  - alert: QueueBacklog
    expr: agent_queue_depth > 100
    annotations:
      summary: "Job queue has >100 pending jobs"

  - alert: AgentDown
    expr: up{job="agent-worker"} == 0
    annotations:
      summary: "Agent worker is down"
```

## Security Considerations

### API Key Management
- Store LLM API keys in Kubernetes Secrets
- Rotate keys monthly
- Use separate keys for dev/staging/prod

### Code Safety
- All generated code goes through PR review
- Agents cannot push directly to main branch
- File-level permissions: agents can only modify specific directories

### Data Privacy
- Sanitize all logs (remove secrets, tokens, passwords)
- Encrypt agent job data at rest
- Limit log retention to 30 days

## Cost Analysis

### LLM API Costs (Claude 3.5 Sonnet)
- **Input**: $3 per million tokens
- **Output**: $15 per million tokens

**Estimated Monthly Usage**:
- @generator: 50 runs × 5K input + 10K output = 250K input + 500K output = $8.25
- @reviewer: 200 runs × 10K input + 5K output = 2M input + 1M output = $21
- @doc-sync: 100 runs × 3K input + 2K output = 300K input + 200K output = $3.90
- @test-gen: 80 runs × 4K input + 8K output = 320K input + 640K output = $10.56
- @spec-writer: 20 runs × 8K input + 15K output = 160K input + 300K output = $4.98

**Total Monthly Cost**: ~$50-60 (well within budget)

### Infrastructure Costs
- Redis: $10/month (managed service)
- PostgreSQL: $20/month (managed service)
- Elasticsearch: $30/month (managed service)
- Compute: $40/month (2 agent workers, 2 CPU, 4GB RAM each)

**Total Infrastructure**: ~$100/month

**Grand Total**: ~$150-170/month

## Success Metrics

### Technical Metrics
- ✅ @generator reduces module creation time by 75% (4h → 1h)
- ✅ @reviewer catches 80%+ issues before human review
- ✅ @deployer reduces deployment time by 83% (30min → 5min)
- ✅ @doc-sync maintains docs freshness >95%
- ✅ @test-gen achieves >70% coverage automatically

### Business Metrics
- ✅ Development velocity increases 40%
- ✅ Code review cycle time decreases 50%
- ✅ Zero production incidents from agent-generated code

## Timeline

| Phase | Duration | Deliverables |
|-------|----------|--------------|
| Phase 1: Infrastructure | 2 weeks | Orchestrator, CLI, Job Queue |
| Phase 2: @generator | 1 week | Code generation agent |
| Phase 3: @reviewer | 1 week | Code review agent |
| Phase 4: @deployer | 1 week | Deployment agent |
| Phase 5: @doc-sync | 1 week | Documentation agent |
| Phase 6: @test-gen | 1 week | Test generation agent |
| Phase 7: @migrator | 0.5 week | Migration agent |
| Phase 8: @spec-writer | 1 week | Spec workflow agent |
| **Total** | **8-9 weeks** | Complete agent system |

## Team Assignment

| Role | Responsibility | Hours |
|------|----------------|-------|
| Senior Backend Dev | Orchestrator, @generator, @reviewer | 60h |
| Backend Dev | @deployer, @doc-sync, @test-gen | 50h |
| DevOps Engineer | Infrastructure, deployment, monitoring | 30h |
| Frontend Dev (optional) | VSCode extension, web UI | 20h |

---

**Next Action**: Review with team, get approval, start Phase 1 prototype
