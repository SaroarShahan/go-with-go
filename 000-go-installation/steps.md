# ✅ How to Install Go (Golang) on macOS

## 📦 Step 1: Download the Go Installer

- Go to the official Go download page:  
  👉 [https://go.dev/dl/](https://go.dev/dl/)

- Download the **macOS installer** (`.pkg` file) based on your system:
  - `go1.xx.x.darwin-amd64.pkg` (for Intel Macs)
  - `go1.xx.x.darwin-arm64.pkg` or `universal.pkg` (for Apple Silicon Macs)


## 🛠 Step 2: Run the Installer

- Open the downloaded `.pkg` file.
- Follow the installation steps (just like installing any regular Mac app).
- It will install Go in `/usr/local/go`.

## ✅ Step 3: Verify Installation

Open your Terminal and run:

```bash
go version
```
✅ You should see something like:
```bash
go version go1.x.x darwin/arm64
```

⚙️ (Optional) Step 4: Set Go Environment Variables

- If Go is installed in /usr/local/go, add the following to your shell config (~/.zshrc, ~/.bash_profile, or ~/.bashrc):

```bash
export PATH=$PATH:/usr/local/go/bin
```
- Then apply the changes:
```bash
source ~/.zshrc  # or source ~/.bash_profile or source ~/.bashrc    
```

[🚀 Run Your First "Hello, World!" Program](../001-hello-world/explanation.md)
