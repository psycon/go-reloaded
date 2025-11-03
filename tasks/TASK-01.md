# TASK-01: Project Setup & Structure

**Status:** ⬜ TODO  
**Dependencies:** None  
**Estimated Time:** 15 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Initialize a new Go project for a text editor using FSM architecture.

**Requirements:**
1. Create project directory structure
2. Initialize Go module
3. Create placeholder directories for all packages
4. Create basic `.gitignore` for Go projects
5. Set up folder structure as per architecture design

**Directory Structure:**
```
go-reloaded/
├── go.mod                 # Go module file
├── go.sum                 # Dependencies (will be auto-generated)
├── .gitignore             # Git ignore file
├── README.md              # Project overview (placeholder)
├── AGENTS.md              # AI orchestration protocol
├── AUTHORS.md             # Project authors
├── docs/
│   └── ANALYSIS.md        # Architecture analysis (placeholder)
├── assets/
│   └── .gitkeep           # Keep empty folder in git
├── tasks/
│   └── .gitkeep           # Task files location
├── main.go                # Entry point (placeholder)
├── fsm/
│   └── .gitkeep           # FSM implementation
├── transforms/
│   └── .gitkeep           # Transformation logic
├── formatters/
│   └── .gitkeep           # Formatting logic
└── tests/
    └── .gitkeep           # Test suite
```

**Acceptance Criteria:**
- [ ] Go module initialized with `go mod init go-reloaded`
- [ ] All directories created
- [ ] `.gitignore` includes Go-specific patterns
- [ ] Git repository initialized (optional)
- [ ] Project compiles without errors (empty main.go)

**Questions for Human:**
1. Project name for go module? (Default: `go-reloaded`)
2. Initialize git repository? (Default: yes)
3. Go version? (Default: 1.25)

**AI Response:** 
"I understand the requirements. I will create the project structure with Go module and directory layout. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create validation tests to ensure project structure is correct.

**Test File:** `tests/structure_test.go`

**Test Coverage:**

1. **Directory Existence Tests:**
   - Verify all required directories exist
   - Verify go.mod file exists
   - Verify main.go exists

2. **Go Module Test:**
   - Verify go.mod contains correct module name
   - Verify go.mod specifies Go version

**Test Structure:**
```go
package tests

import (
    "os"
    "testing"
)

func TestProjectStructure(t *testing.T) {
    requiredDirs := []string{
        "fsm",
        "transforms",
        "formatters",
        "tests",
        "docs",
        "assets",
        "tasks",
    }
    
    for _, dir := range requiredDirs {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            t.Errorf("Required directory does not exist: %s", dir)
        }
    }
}

func TestGoModExists(t *testing.T) {
    if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
        t.Error("go.mod file does not exist")
    }
}

func TestMainGoExists(t *testing.T) {
    if _, err := os.Stat("main.go"); os.IsNotExist(err) {
        t.Error("main.go file does not exist")
    }
}
```

**AI Deliverable:** Complete `tests/structure_test.go` file

---

### **STEP 3: Generate the Code**

**Task:** Create all required files and directories.

**Files to Create:**

#### **1. go.mod**
```go
module go-reloaded

go 1.21
```

#### **2. .gitignore**
```
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
go-reloaded

# Test binary
*.test

# Output coverage
*.out

# Go workspace file
go.work

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Output files
result.txt
output.txt
```

#### **3. main.go (placeholder)**
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: Implement in TASK-05
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}
	
	fmt.Println("Text Editor - Coming soon!")
}
```

#### **4. Create all directories**
```bash
mkdir -p fsm transforms formatters tests docs assets tasks
touch fsm/.gitkeep transforms/.gitkeep formatters/.gitkeep
touch tests/.gitkeep docs/.gitkeep assets/.gitkeep tasks/.gitkeep
```

#### **5. README.md (placeholder)**
```markdown
# Text Editor Project

Text processing tool using FSM architecture.

## Status

🚧 Under construction - See `AGENTS.md` for build progress

## Quick Start

Coming soon!

## Documentation

- See `docs/ANALYSIS.md` for architecture details
- See `AGENTS.md` for AI-assisted build instructions
```

**Implementation Steps:**
1. Create directory structure
2. Initialize go module: `go mod init go-reloaded`
3. Create all placeholder files
4. Initialize git (optional): `git init`

**AI Deliverable:** Complete project structure with all files

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify project setup is correct.

**QA Checklist:**

1. **Structure Validation:**
   ```bash
   # Verify all directories exist
   ls -la
   
   # Should see: fsm/, transforms/, formatters/, tests/, docs/, assets/, tasks/
   ```

2. **Go Module Test:**
   ```bash
   # Verify go module
   go mod verify
   
   # Should output: all modules verified
   ```

3. **Compilation Test:**
   ```bash
   # Verify project compiles
   go build -o go-reloaded .
   
   # Should compile without errors
   ```

4. **Run Structure Tests:**
   ```bash
   go test ./tests/structure_test.go -v
   
   # Should pass all structure validation tests
   ```

5. **Git Check (if initialized):**
   ```bash
   git status
   
   # Should show untracked files
   ```

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-01 → ✅ COMPLETE
- Git commit:
  ```bash
  git add .
  git commit -m "feat: complete TASK-01 - project structure and setup"
  ```

**AI Final Response:**
"TASK-01 completed successfully. Project structure created and verified. Ready to proceed to TASK-02."

---

## 📊 Success Metrics

- ✅ All directories created
- ✅ Go module initialized
- ✅ .gitignore configured
- ✅ Project compiles
- ✅ Structure tests pass

---

## 🔗 Related Tasks

- **Previous:** None (first task)
- **Next:** TASK-02 (Transforms Module)
- **Blocks:** All other tasks depend on this

---

*Task created: October 26, 2025*