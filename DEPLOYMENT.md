# Deployment Guide

## Option 1: GitHub Pages (Static Demo)

### Setup:
1. Push your code to GitHub
2. Go to repository Settings → Pages
3. Source: Deploy from a branch
4. Branch: `main` → `/docs/gh-pages`
5. Save

Your site will be live at: `https://YOUR_USERNAME.github.io/go-reloaded/`

**Note:** The demo is client-side JavaScript simulation, not actual Go code.

---

## Option 2: Go Playground Links

Share executable code snippets:
1. Go to https://go.dev/play/
2. Paste your code
3. Click "Share" to get a permanent link
4. Add link to README.md

Example: https://go.dev/play/p/YOUR_CODE_ID

---

## Option 3: Replit (Interactive Demo)

### Setup:
1. Go to https://replit.com
2. Create new Repl → Import from GitHub
3. Enter your repository URL
4. Replit will auto-detect Go
5. Click "Run" to start

Users can test your code interactively!

---

## Option 4: Docker + Cloud Run (Full Backend)

### Create Dockerfile:
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o go-reloaded .
CMD ["./go-reloaded"]
```

### Deploy to Google Cloud Run:
```bash
gcloud run deploy go-reloaded --source .
```

---

## Option 5: GitHub Actions + Releases

Automatically build binaries for multiple platforms:

### Create `.github/workflows/release.yml`:
```yaml
name: Release
on:
  push:
    tags:
      - 'v*'
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - name: Build
        run: |
          GOOS=linux GOARCH=amd64 go build -o go-reloaded-linux
          GOOS=darwin GOARCH=amd64 go build -o go-reloaded-mac
          GOOS=windows GOARCH=amd64 go build -o go-reloaded.exe
      - uses: softprops/action-gh-release@v1
        with:
          files: |
            go-reloaded-linux
            go-reloaded-mac
            go-reloaded.exe
```

---

## Recommended: GitHub Pages + Replit

**Best combination for showcasing:**
1. **GitHub Pages** - Professional landing page with documentation
2. **Replit** - Interactive demo where users can test the actual Go code
3. **GitHub Releases** - Downloadable binaries for all platforms

Add badges to README.md:
```markdown
[![Demo](https://img.shields.io/badge/demo-live-success)](https://YOUR_USERNAME.github.io/go-reloaded/)
[![Try on Replit](https://img.shields.io/badge/try-replit-orange)](https://replit.com/@YOUR_USERNAME/go-reloaded)
```
