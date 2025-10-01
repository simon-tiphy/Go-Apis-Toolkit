# 🚀 Go Hello API — Beginner Toolkit

> A minimal HTTP API built with Go that returns JSON responses. Perfect for learning Go basics!

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)](https://github.com/simon-tiphy/Go-Apis-Toolkit)

---

## 📋 What This Project Does

This is a **minimal Go HTTP server** that demonstrates:
- ✅ Setting up a basic HTTP API with Go's standard library
- ✅ Returning JSON responses
- ✅ Writing unit tests for your handlers
- ✅ Structuring a Go project with proper modules

**Sample Response:**
```json
{
  "message": "Hello, World!"
}
```

---

## 🎯 Prerequisites

Before you begin, make sure you have:

| Tool | Version | Check Installation |
|------|---------|-------------------|
| **Go** | 1.20+ | `go version` |
| **Git** | Latest | `git --version` |

---

## 📦 Installation & Setup

### Step 1: Install Go

Choose your platform and follow the instructions:

<details>
<summary><b>🍎 macOS</b></summary>

#### Using Homebrew (Recommended)
```bash
brew install go
```

#### Verify Installation
```bash
go version
```
✅ Expected output: `go version go1.20.x darwin/amd64`

</details>

<details>
<summary><b>🐧 Linux (Debian/Ubuntu)</b></summary>

#### Using APT
```bash
sudo apt update
sudo apt install -y golang
```

#### Verify Installation
```bash
go version
```
✅ Expected output: `go version go1.20.x linux/amd64`

</details>

<details>
<summary><b>🪟 Windows</b></summary>

1. Download the installer from [https://go.dev/doc/install](https://go.dev/doc/install)
2. Run the `.msi` file and follow the wizard
3. Restart your terminal

#### Verify Installation
```bash
go version
```
✅ Expected output: `go version go1.20.x windows/amd64`

</details>

---

### Step 2: Clone & Initialize the Project

```bash
# Create project directory
mkdir go-hello-api
cd go-hello-api

# Initialize Go module
go mod init github.com/simon-tiphy/Go-Apis-Toolkit
```

---

## 📁 Project Structure

```
go-hello-api/
├── 📄 main.go              # HTTP server with /hello endpoint
├── 📄 hello_test.go        # Unit tests for the handler
├── 📄 go.mod               # Module file (github.com/simon-tiphy/Go-Apis-Toolkit)
├── 📄 AI_PROMPTS.md        # AI prompts & learning reflections
├── 📄 README.md            # This file
└── 📄 .gitignore           # Git ignore rules
```

---

## 🚀 Quick Start

### Run the Server

```bash
# From project root directory
go run .
```

You should see:
```
Server starting on :8080
```

---

## 🧪 Test the API

### Option 1: Using Your Browser
Open [http://localhost:8080/hello](http://localhost:8080/hello)

### Option 2: Using cURL
```bash
curl http://localhost:8080/hello
```

### ✅ Expected Response:
```json
{
  "message": "Hello, World!"
}
```

---

## 🧪 Run Tests

Execute the unit tests to verify everything works:

```bash
go test ./...
```

### ✅ Expected Output:
```
PASS
ok      github.com/simon-tiphy/Go-Apis-Toolkit  0.XXXs
```

---

## 📚 Key Files Explained

### `main.go`
The main HTTP server that:
- Creates an HTTP handler for the `/hello` route
- Returns a JSON response with "Hello, World!"
- Listens on port 8080

### `hello_test.go`
Unit tests that verify:
- The handler returns the correct status code (200)
- The response is valid JSON
- The message content is correct

### `go.mod`
Defines your module path: `github.com/simon-tiphy/Go-Apis-Toolkit`

### `AI_PROMPTS.md`
Documents all the AI prompts used during development and personal learning reflections.

---

## 🐛 Troubleshooting

<details>
<summary><b>❌ Error: "go: cannot find main module"</b></summary>

**Solution:** Initialize the Go module first:
```bash
go mod init github.com/simon-tiphy/Go-Apis-Toolkit
```

</details>

<details>
<summary><b>❌ Error: "address already in use"</b></summary>

**Solution:** Port 8080 is occupied. Find and stop the process:

**macOS/Linux:**
```bash
lsof -i :8080
kill -9 <PID>
```

**Windows:**
```bash
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

Or change the port in `main.go` to 8081 or another available port.

</details>

<details>
<summary><b>❌ Error: "command not found: go"</b></summary>

**Solution:** Go isn't installed or not in your PATH.

1. Reinstall Go from [https://go.dev/doc/install](https://go.dev/doc/install)
2. Add Go to your PATH:

**macOS/Linux** (add to `~/.bashrc` or `~/.zshrc`):
```bash
export PATH=$PATH:/usr/local/go/bin
source ~/.bashrc
```

**Windows:** The installer should handle this automatically. Restart your terminal.

</details>

---

## 🎓 Learning Resources

### Official Documentation
- 📖 [Go Official Docs](https://go.dev/doc/) — Comprehensive Go documentation
- 🎯 [Go Tour](https://go.dev/tour/) — Interactive Go tutorial
- 📘 [Effective Go](https://go.dev/doc/effective_go) — Best practices guide

### Tutorials
- 🎬 [Go by Example](https://gobyexample.com/) — Hands-on examples
- 📝 [Go Web Examples](https://gowebexamples.com/) — Web development with Go
- 🎥 [FreeCodeCamp: Learn Go](https://www.youtube.com/watch?v=YS4e4q9oBaU) — Full course

### Community
- 💬 [Go Forum](https://forum.golangbridge.org/)
- 🐦 [r/golang](https://www.reddit.com/r/golang/)
- 💼 [Gophers Slack](https://gophers.slack.com/)

---

## 🤝 Contributing

Found a bug or want to improve this toolkit? Contributions are welcome!

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Simon Tiphy**

- 🔗 GitHub: [@simon-tiphy](https://github.com/simon-tiphy)
- 📦 Repository: [Go-Apis-Toolkit](https://github.com/simon-tiphy/Go-Apis-Toolkit)

---

## 🙏 Acknowledgments

- Built as part of the **Moringa AI Capstone Project**
- Thanks to the Go community for excellent documentation
- Special thanks to [ai.moringaschool.com](https://ai.moringaschool.com) for AI-powered learning support

---

<div align="center">

### ⭐ If this helped you learn Go, give it a star!

**Made with 💙 and Go**

[⬆ Back to Top](#-go-hello-api--beginner-toolkit)

</div>