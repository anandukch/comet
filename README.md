# 🌠 Comet: The Ultimate Code Comment Management CLI Tool 🌠

**Comet** is a powerful CLI tool designed to help developers scan, detect, manage, and document comments in code files. It supports multiple programming languages and allows you to remove comments with a single command, as well as generate comprehensive documentation of the comments. Perfect for cleaner code, streamlined collaboration, and effortless documentation!

---

## 🚀 Features

- **Multi-Language Support**: Detects comments in various programming languages (Go, Python, C, and more) based on file type.
- **Automatic Documentation**: Saves all comments in a structured format in `comments.txt`, including file paths and line numbers.
- **Optional Comment Removal**: Prompts you to remove comments during the scan process, with a simple "yes/no" option.
- **Directory Scanning**: Easily scan an entire project or specific directory for comments.
  
---

## 📂 Supported Languages & Comment Syntax

Comet automatically identifies comments based on file type, using the correct delimiter:

| Language | Extension | Comment Syntax    |
|----------|-----------|-------------------|
| Go       | `.go`     | `//`             |
| Python   | `.py`     | `#`              |
| C        | `.c`      | `//`, `/*...*/`  |
| Java     | `.java`   | `//`, `/*...*/`  |

*Additional languages can be easily configured as needed.*

---

## 🔧 Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/comet.git
   cd comet
2. **Install Dependencies** Ensure you have Go installed on your system. Run:

    ```bash
    go mod tidy
3. **Build the CLI**

    ```bash
    go build

## 🛠 Usage

### Basic Usage

Run Comet to scan for comments in the current directory:

```bash
./comet
```

### Scan a Specific Directory

Provide a project path to scan for comments in that specific folder:

```bash
./comet /path/to/your/project
```

### Generated Documentation

After scanning, Comet will generate a `comments.txt` file in the project root with a structured list of comments:

```bash
/path/to/your/file.go
Line 3: // This is a comment
Line 15: // Another comment

/path/to/your/file.py
Line 8: # Python comment example
```

### Comment Removal

When prompted, you can remove all detected comments by typing "y" (or keep them by typing "n"). This removes comments directly from the files scanned.

## 🤖 Example Output

```bash
Starting comment scan in directory: /path/to/your/project
Comment found at /path/to/your/file.go:3 - // Initial setup
Comment found at /path/to/your/file.go:15 - // Deprecated function

Would you like to remove the comments? (y/n): y
Comments removed and saved to comments.txt
```
