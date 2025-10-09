# Notion Auto-Update Script - User Guide

## Overview

This automated script updates all Spec-004 tasks in your Notion Tasks database to "Done" status with completion metadata.

## Prerequisites

### 1. Install jq (JSON processor)

```bash
# macOS
brew install jq

# Ubuntu/Debian
sudo apt-get install jq

# Verify installation
jq --version
```

### 2. Get Notion API Key

1. Go to [Notion Integrations](https://www.notion.so/my-integrations)
2. Click "**+ New integration**"
3. Name it: "**Simple Admin Tasks Updater**"
4. Select your workspace
5. Copy the **Internal Integration Token** (starts with `secret_...`)

### 3. Share Database with Integration

1. Open your **Tasks** database in Notion
2. Click "**...**" (top right) → **Connections** → **Connect to**
3. Select "**Simple Admin Tasks Updater**"
4. Grant write access

### 4. Get Database ID

**Method 1: From URL**
```
https://www.notion.so/<workspace>/abcd1234efgh5678?v=...
                                   ^^^^^^^^^^^^^^^^
                                   This is your Database ID
```

**Method 2: Using Notion API**
```bash
curl -X GET 'https://api.notion.com/v1/search' \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -H 'Notion-Version: 2022-06-28' \
  | jq '.results[] | select(.object == "database") | {id, title}'
```

## Usage

### Method 1: Command Line Argument

```bash
cd specs/004-user-module-proto-completion

# Replace YOUR_API_KEY with actual token
./notion-auto-update.sh "secret_abc123..."
```

### Method 2: Environment Variables

```bash
# Set environment variables
export NOTION_API_KEY="secret_abc123..."
export NOTION_DATABASE_ID="abcd1234efgh5678"

# Run script
./notion-auto-update.sh
```

### Method 3: Create .env File

```bash
# Create .env file (DO NOT commit to git!)
cat > .env << 'EOF'
export NOTION_API_KEY="secret_abc123..."
export NOTION_DATABASE_ID="abcd1234efgh5678"
EOF

# Source and run
source .env
./notion-auto-update.sh
```

## Expected Output

```
[INFO] Starting Spec-004 Notion Tasks auto-update...
[INFO] Database ID: abcd1234efgh5678
[INFO] Tasks to update: 8

[INFO] Processing: ZH-TW-007 - Extend core.proto with User RPC methods
[INFO] Querying task: ZH-TW-007
[INFO] Updating task ZH-TW-007 to status: Done
[SUCCESS] Updated ZH-TW-007

[INFO] Processing: ZH-TW-008 - Update user.proto for Proto-First generation
[INFO] Querying task: ZH-TW-008
[INFO] Updating task ZH-TW-008 to status: Done
[SUCCESS] Updated ZH-TW-008

... (6 more tasks)

========================================
[INFO] Update Summary
========================================
[SUCCESS] Successfully updated: 8 tasks
========================================
[SUCCESS] All tasks updated successfully! ✅
```

## Tasks Updated

The script automatically updates the following 8 tasks:

| Task ID | Description | Status | Estimated | Actual | Commit |
|---------|-------------|--------|-----------|--------|--------|
| ZH-TW-007 | Extend core.proto with User RPC methods | Done | 6h | 6h | eac6379d |
| ZH-TW-008 | Update user.proto for Proto-First generation | Done | 4h | 4h | eac6379d |
| USER-001 | Implement authentication RPC logic | Done | 6h | 6h | eac6379d |
| USER-002 | Implement registration RPC logic | Done | 4h | 4h | eac6379d |
| USER-003 | Implement password management RPC logic | Done | 4h | 4h | eac6379d |
| USER-004 | Implement user info retrieval RPC logic | Done | 4h | 4h | eac6379d |
| USER-005 | Implement token management RPC logic | Done | 4h | 4h | eac6379d |
| USER-006 | Generate API file from user.proto | Done | 2h | 2h | eac6379d |

## Fields Updated

For each task, the script updates:

1. **Status**: "In progress" → "Done"
2. **Estimated Hours**: Set from task definition
3. **Actual Hours**: Set from task definition
4. **Completed At**: Current timestamp (UTC)
5. **Commit Hash**: `eac6379d` (Spec-004 completion commit)
6. **Progress**: Set to 100%

## Troubleshooting

### Error: "jq is not installed"

**Solution**: Install jq using package manager (see Prerequisites)

### Error: "NOTION_API_KEY not set"

**Solution**: Provide API key via command line or environment variable

### Error: "Please set NOTION_DATABASE_ID"

**Solution**: Update script or set environment variable with your database ID

### Error: "401 Unauthorized"

**Possible Causes**:
1. Invalid API key
2. Integration not shared with database
3. API key expired

**Solution**:
1. Verify API key at https://www.notion.so/my-integrations
2. Share database with integration (see Prerequisites)

### Error: "Task not found"

**Possible Causes**:
1. Task ID doesn't match exactly (case-sensitive)
2. Task doesn't exist in database
3. Wrong database ID

**Solution**:
1. Check task IDs in Notion match script definitions
2. Verify database ID is correct

### Error: "Rate limit exceeded"

**Solution**: Script includes 0.5s delay between requests, but if you hit rate limit:
1. Wait 1 minute
2. Increase `sleep` value in script (line 214)

## Manual Verification

After running the script, verify in Notion:

```bash
# Check Tasks page
1. Open Notion Tasks database
2. Filter by "Status = Done"
3. Verify all 8 tasks are marked complete
4. Check "Completed At" timestamps match script run time
5. Verify "Commit Hash" = eac6379d for all tasks
```

## Security Notes

⚠️ **IMPORTANT**:

1. **Never commit** `.env` file or API keys to git
2. **Rotate API keys** regularly
3. **Limit integration permissions** to only required databases
4. **Use environment variables** in CI/CD pipelines

Add to `.gitignore`:
```
.env
**/notion-api-key.txt
```

## Alternative: Manual CSV Import

If you prefer manual update, use the CSV import method from `notion-task-updates.md`:

1. Export CSV from Notion
2. Update rows manually
3. Re-import CSV
4. Merge updates

## FAQ

**Q: Can I update specific tasks only?**
A: Yes, comment out unwanted tasks in the `TASKS` array (lines 71-80)

**Q: Can I change the status to something other than "Done"?**
A: Yes, modify the status value in `TASKS` array (e.g., "In progress", "Blocked")

**Q: Will this overwrite manual changes in Notion?**
A: Yes, the script will overwrite Status, Hours, and Commit Hash fields. Other fields are preserved.

**Q: Can I dry-run first?**
A: Add `--dry-run` flag support by commenting out the `curl -X PATCH` line (line 157)

## Support

For issues with:
- **Script**: Contact @pm or check script logs
- **Notion API**: See [Notion API Docs](https://developers.notion.com/)
- **Spec-004**: See `completion-report.md` and `acceptance-checklist.md`

---

**Last Updated**: 2025-10-10
**Version**: 1.0
**Related**: Spec-004 User Module Proto Completion
