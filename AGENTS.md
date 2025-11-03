# AGENTS.md — AI Orchestration Protocol (Per-Task Execution)

This repository uses **one markdown file per task** (e.g., `TASK-01.md`, `TASK-02.md`).

**Each task file includes the FULL 4-step prompt** that AI agents must follow:

1. **Analyze & Confirm**
2. **Generate the Tests**
3. **Generate the Code**
4. **QA & Mark Complete**

---

## ✅ **How to run any task**

1. Open the specific `TASK-XXX.md` file from the index.
2. Copy the section **"PROMPT — FULL 4-STEP FLOW (execute sequentially)"** into your AI tool (Claude, Codex, GPT, etc.).
3. Provide any requested inputs during **Step 1**.
4. After you confirm Step 1, the AI will proceed with Steps 2-4 for that task.

---

## 📋 Task Index

### Phase 1: Core Implementation

- [TASK-01: Project Setup & Structure](tasks/TASK-01.md)
- [TASK-02: Transforms Module (Hex/Bin/Case)](tasks/TASK-02.md)
- [TASK-03: Formatters Module (Punctuation/Quotes)](tasks/TASK-03.md)
- [TASK-04: FSM Processor Core](tasks/TASK-04.md)
- [TASK-05: Main Entry Point](tasks/TASK-05.md)

### Phase 2: Testing & Validation

- [TASK-06: Unit Tests for Transforms](tasks/TASK-06.md)
- [TASK-07: Integration Tests](tasks/TASK-07.md)
- [TASK-08: Golden Test Suite](tasks/TASK-08.md)

### Phase 3: Documentation

- [TASK-09: README.md](tasks/TASK-09.md)
- [TASK-10: ANALYSIS.md](tasks/TASK-10.md)

---

## 🏗️ Repository Structure

```
text-editor/
├── AGENTS.md              # This file (orchestration protocol)
├── tasks/
│   ├── TASK-01.md         # Project setup instructions
│   ├── TASK-02.md         # Transforms module
│   ├── TASK-03.md         # Formatters module
│   └── ...
├── README.md              # Project overview
├── docs/
│   └── ANALYSIS.md        # Architecture analysis
├── main.go                # Entry point
├── fsm/                   # FSM implementation
├── transforms/            # Transformation logic
├── formatters/            # Formatting logic
└── tests/                 # Test suite
```

---

## 🤖 Execution Policy (Agent Focus)

- Treat `docs/ANALYSIS.md` as the canonical context for tasks.
- Each task is independent but follows the architecture defined in ANALYSIS.md.
- Tests must pass before marking task complete.
- Follow Go best practices and the FSM architecture pattern.

---

## 🔄 CI/CD Integration

### Local Testing
```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./transforms
go test ./formatters
go test ./fsm

# Run with coverage
go test -cover ./...
```

### Build
```bash
# Build executable
go build -o text-editor .

# Run
./text-editor input.txt output.txt
```

---

## 📊 Progress Tracking

| Task | Status | Tests | Description |
|------|--------|-------|-------------|
| TASK-01 | ⬜ TODO | N/A | Project setup |
| TASK-02 | ⬜ TODO | ⬜ | Transforms module |
| TASK-03 | ⬜ TODO | ⬜ | Formatters module |
| TASK-04 | ⬜ TODO | ⬜ | FSM processor |
| TASK-05 | ⬜ TODO | ⬜ | Main entry point |
| TASK-06 | ⬜ TODO | ⬜ | Unit tests |
| TASK-07 | ⬜ TODO | ⬜ | Integration tests |
| TASK-08 | ⬜ TODO | ⬜ | Golden tests |
| TASK-09 | ⬜ TODO | N/A | README |
| TASK-10 | ⬜ TODO | N/A | ANALYSIS |

**Legend:** ⬜ TODO | 🟨 IN PROGRESS | ✅ COMPLETE

---

## 🎯 For AI Agents

When executing tasks:
1. Read `docs/ANALYSIS.md` for architecture context
2. Follow the 4-step prompt in each TASK file
3. Run tests after code generation
4. Update progress table above
5. Commit with descriptive message: `feat: complete TASK-XX - [description]`

---

## 📝 Notes

- Each task should be completed independently
- Follow the order for dependencies (e.g., TASK-02 before TASK-04)
- All code must have tests
- Documentation tasks (TASK-09, TASK-10) can be done in parallel with implementation

---

*Last Updated: October 26, 2025*