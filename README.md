# go-ls: A Modern Go-based Alternative to ls

go-ls is a modern alternative to the traditional `ls` command, written in Go. It provides enhanced features like color-coding, file metadata display, tree view, and Git integration, similar to Rust's `eza` tool, but implemented in Go.

## Features

- Enhanced file display with color coding based on file types
- Multiple output formats: grid, list, tree view
- Git integration shows the status of files in Git repositories
- Comprehensive file metadata display (permissions, size, timestamps)
- Configurable sorting options
- Cross-platform compatibility (Linux, macOS, Windows)

## Installation

To install go-ls, run:

```bash
go install github.com/logando-al/go-ls
```

## Usage

```bash
# Basic usage (similar to ls)
go-ls

# Long view with file details
go-ls -l

# Include hidden files
go-ls -a

# Recursive view in tree format
go-ls -T

# Show Git status alongside file listings
go-ls --git

# Sort by modification time
go-ls --sort=time
```

## Status

This is currently a work in progress. For more details about the implementation plan, see [plan.md](plan.md).

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

MIT License