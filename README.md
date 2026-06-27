# gls: A Modern Go-based Alternative to ls

gls is a modern alternative to the traditional `ls` command, written in Go. It provides enhanced features like color-coding, file metadata display, tree view, and Git integration, similar to Rust's `eza` tool, but implemented in Go.

## Features

- Enhanced file display with color coding based on file types
- Multiple output formats: grid, list, tree view
- Git integration shows the status of files in Git repositories
- Comprehensive file metadata display (permissions, size, timestamps)
- Configurable sorting options
- Cross-platform compatibility (Linux, macOS, Windows)

## Installation

To install gls, run:

```bash
go install github.com/logando-al/gls
```

## Usage

```bash
# Basic usage (similar to ls)
gls

# Long view with file details
gls -l

# Include hidden files
gls -a

# Recursive view in tree format
gls -T

# Show Git status alongside file listings
gls --git

# Sort by modification time
gls --sort=time
```

## Linux Alias Setup

If you're using Linux or a Unix-like system, you can create an alias to use `gls` in place of the standard `ls` command:

```bash
# Add this to your ~/.bashrc or ~/.zshrc file
alias ls="gls"

# Then reload your shell configuration
source ~/.bashrc  # or source ~/.zshrc
```

## Status

This is currently a work in progress. For more details about the implementation plan, see [plan.md](plan.md).

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

MIT License

---

*Note: This project has been created with assistance from ilmu-v3.1 (ILMU AI)*