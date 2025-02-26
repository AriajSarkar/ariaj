# 🤖 Ariaj - CLI Interface for Ollama

<div align="center">

![Ariaj Logo](https://raw.githubusercontent.com/AriajSarkar/ariaj/main/assets/logo.png)

[![Go Version](https://img.shields.io/github/go-mod/go-version/AriajSarkar/ariaj)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/AriajSarkar/ariaj)](https://github.com/AriajSarkar/ariaj/releases/latest)
[![Release](https://github.com/AriajSarkar/ariaj/actions/workflows/release.yml/badge.svg)](https://github.com/AriajSarkar/ariaj/actions/workflows/release.yml)

*A powerful CLI tool for interacting with Large Language Models through Ollama*

</div>

## ✨ Features

- 🚀 **Easy Installation** - Simple global installation process
- 💬 **Interactive Mode** - Chat with your LLM in real-time
- 🔄 **Model Switching** - Easily switch between different LLM models
- 🖥️ **Process Management** - Automatic Ollama server management
- ⚡ **Streaming Responses** - Get real-time streaming responses
- 🎯 **Single Prompt Mode** - Quick one-off queries

## 🚀 Installation

### Prerequisites

1. Install [Go](https://golang.org/doc/install) (version 1.19 or higher)
2. Install [Ollama](https://ollama.ai)

### Quick Install

```bash
go install github.com/AriajSarkar/ariaj@latest
```

### Download Binary

Choose the appropriate version for your system:

| Platform | Architecture | Download |
|----------|-------------|----------|
| Windows | x64 | [Download](https://github.com/AriajSarkar/ariaj/releases/download/v0.1.5/ariaj_0.1.5_Windows_x86_64.zip) |
| Windows | ARM64 | [Download](https://github.com/AriajSarkar/ariaj/releases/download/v0.1.5/ariaj_0.1.5_Windows_arm64.zip) |
| Linux | x64 | [Download](https://github.com/AriajSarkar/ariaj/releases/download/v0.1.5/ariaj_0.1.5_Linux_x86_64.tar.gz) |
| Linux | ARM64 | [Download](https://github.com/AriajSarkar/ariaj/releases/download/v0.1.5/ariaj_0.1.5_Linux_arm64.tar.gz) |
| macOS | x64 | [Download](https://github.com/AriajSarkar/ariaj/releases/download/v0.1.5/ariaj_0.1.5_Darwin_x86_64.tar.gz) |
| macOS | ARM64 | [Download](https://github.com/AriajSarkar/ariaj/releases/download/v0.1.5/ariaj_0.1.5_Darwin_arm64.tar.gz) |

[View all releases](https://github.com/AriajSarkar/ariaj/releases)

### Manual Installation

```bash
# Clone the repository
git clone https://github.com/AriajSarkar/ariaj.git

# Build and install globally
cd ariaj
go build
./ariaj install
```

## 📚 Usage

### Basic Commands

```bash
# Interactive mode
ariaj

# Single prompt
ariaj "What is the capital of France?"

# Change model
ariaj model

# Uninstall
ariaj uninstall
```

### Interactive Mode

Start an interactive session:
```bash
$ ariaj
Using model: llama2
Enter your prompt (or 'exit' to quit): 
```

### Switching Models

```bash
$ ariaj model
? Select Model: 
  ▸ llama2
    codellama
    mistral
    phi
```

## 🔧 Configuration

Configuration is automatically managed in:
- Windows: `%APPDATA%/ariaj/config.json`
- Linux/Mac: `$HOME/.config/ariaj/config.json`

## 🤝 Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, development process, and how to contribute to the project.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Ollama](https://ollama.ai) for the amazing LLM server
- All the contributors who help improve this project

---

<div align="center">
Made with ❤️ by Raj
</div>
