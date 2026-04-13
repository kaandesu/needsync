# needsync

**needsync** is an internal [Kampai](https://trykampai.agentcompany.cloud/landing-page) tool that bridges AI agents and humans through Git-based workflows.

This binary will be used along with a coding agent that works on a repository derived from the [Fullstack4Agents](https://github.com/kaandesu/fullstack4agents.git) template.

It reads structured NEEDS from the repository, creates corresponding issues in Gitea, collects human input, and feeds it back into the agent’s context.

> For now Public source, not open source.

---

## What it does

needsync turns agent uncertainty into structured, trackable human tasks.

**Agent → NEED → Issue → Human Input → JSON → Agent resumes**

It is the coordination layer between:

- AI agents (producers of NEEDS)
- Humans (providers of input)
- Repositories (source of truth)

---

## Expected project structure

```bash
/needs/
  email_sending.json
  phone_calling.json
```

Each file represents a single NEED.

---

## NEED format (input)

Example:

```json
{
  "id": "email_sending",
  "title": "Provide SMTP credentials",
  "priority": "high",
  "status": "pending",
  "blocking": true,
  "description": "Email sending requires SMTP credentials.",
  "human_actions": [
    {
      "id": "choose_provider",
      "type": "decision",
      "title": "Choose email provider",
      "options": ["Gmail", "Resend"],
      "answer": null
    }
  ]
}
```

---

## Environment variables

needsync requires access to a Gitea instance:

```env
GITEA_BASE_URL=http://localhost:3000
GITEA_TOKEN=your_token
GITEA_OWNER=your_org
GITEA_REPO=your_repo
```

---

## CLI Usage

### Sync needs → create issues

```bash
needsync sync
```

What it does:

- scans `/needs/*.json`
- creates a Gitea issue for each NEED without an `issue.number`
- updates the file with:

```json
"issue": {
  "repo": "org/repo",
  "number": 42
}
```

- sets status → `waiting_human`

---

## API Usage

Run server:

```bash
go run cmd/api/main.go
```

### Endpoint

```
POST /sync
```

Triggers the same behavior as CLI `sync`.

---

## Lifecycle

```text
Agent creates NEED
        ↓
needsync creates Issue
        ↓
Human provides input
        ↓
needsync updates JSON
        ↓
Agent resumes work
        ↓
NEED marked as done
```

---

## State model

| Status        | Meaning                    |
| ------------- | -------------------------- |
| pending       | newly created              |
| waiting_human | issue created, needs input |
| ready         | all answers provided       |
| processing    | agent is working           |
| done          | completed                  |

---

## Human interaction

Humans interact through Gitea issues.

Typical flow:

1. Open issue
2. Read requested actions
3. Provide answers via comments or UI
4. needsync pulls and updates JSON

---

## Design principles

- JSON is the **source of truth**
- Issues are **UI only**
- Agents are **stateless**
- No hidden state, everything is file-based

---

## Non-goals

- Not a task manager
- Not a workflow engine
- Not a replacement for CI/CD

---

## Internal usage

This tool is designed for Kampai™ internal systems.

Source is public for transparency, but:

- no support guarantees
- no contribution model
- no stability guarantees

---

## 🧪 Minimal example

```bash
# 1. agent creates file
echo '{ "id": "test", ... }' > needs/test.json

# 2. run sync
needsync sync

# 3. issue appears in Gitea
# 4. human responds

# 5. (future) sync pull updates JSON
```

---
