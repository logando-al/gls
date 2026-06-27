# gls: A Modern Go-based Alternative to ls

## Project Overview

gls is a modern alternative to the traditional `ls` command, written in Go. It aims to provide enhanced features like color-coded output, file metadata display, tree view, and Git integration, similar to Rust's `eza` tool, but implemented in Go.

## Goals

1. Provide a cross-platform, enhanced file listing tool
2. Integrate with Git to show file status
3. Support multiple output formats (grid, tree, details)
4. Add visual enhancements like colors and icons
5. Offer extensive customization options

## Project Structure

```
gls/
├── .gitignore
├── plan.md
├── README.md
├── go.mod
├── cmd/
│   └── gls/
│       └── main.go
├── pkg/
│   ├── colors/         # Terminal colors management
│   ├── git/            # Git status integration
│   ├── icons/          # File type icons
│   ├── output/         # Output formatting
│   ├── scanner/        # File system operations
│   └── sort/           # Sorting algorithms
└── test/               # Test files and directories
```

## Implementation Phases

### Phase 1: Basic Functionality
1. **File Scanning**: 
   - Use Go's `os` and `path/filepath` packages
   - Extract basic file information (name, size, mod time, permissions)
   
2. **CLI Interface**:
   - Use `spf13/cobra` for argument parsing
   - Support basic flags like `-l` (long listing), `-a` (all files), `-R` (recursive)

3. **Simple Output Formats**:
   - Grid view (default)
   - List view
   - One file per line

### Phase 2: Enhanced Features
1. **Color Coding**:
   - Implement file type detection
   - Use `fatih/color` for color output
   - Support different color themes

2. **File Metadata Display**:
   - File permissions
   - Owner and group
   - Timestamps with customizable formats
   - Size display with human-readable format

3. **Sorting Options**:
   - By name (default)
   - By size
   - By modification time
   - By extension

### Phase 3: Advanced Features
1. **Git Integration**:
   - Use `go-git/go-git/v5` to check Git status
   - Show modified, added, deleted, or untracked files
   - Display branch information

2. **Tree View**:
   - Implement tree structure display
   - Show hierarchical relationships between files and directories
   - Control depth with `-L` flag

3. **Icons Support**:
   - Add file type icons using Nerd fonts
   - Make icons configurable and optional

4. **Configuration System**:
   - Use `spf13/viper` for configuration
   - Support config files and environment variables

## Dependencies

- `spf13/cobra`: Powerful CLI applications library (for argument parsing)
- `spf13/viper`: Go configuration for flags and config files
- `fatih/color`: Color package for Go
- `go-git/go-git/v5`: Go implementation of Git

## Testing Strategy

1. **Unit Tests**: 
   - Test each package's functionality separately
   - Focus on core functionality like file scanning, sorting, and output formatting

2. **Integration Tests**:
   - Use sample test directories with various file types
   - Test different output formats and options
   - Verify Git integration with test repositories

3. **CLI Tests**:
  - Test command-line argument parsing
   - Verify output matches expected format

## Feature Roadmap

- [x] Setup basic project structure
- [x] Basic file scanning
- [x] Simple grid view
- [x] List view with basic metadata
- [x] Color coding for file types
- [ ] Git integration
- [ ] Tree view
- [ ] Icons support
- [ ] Configuration system
- [ ] Man pages and documentation
- [ ] Cross-platform builds and distribution

## Future Enhancements

1. **Advanced Filters**:
   - Filter files by type, size, or modification date
   - Support ignore patterns (similar to `.gitignore`)

2. **Extended Attributes**:
   - Display extended attributes on supported platforms
   - Show file permissions in octal format

3. **Performance Optimizations**:
   - Parallel scanning of directories
   - Caching for Git status in large repositories

4. **Terminal Integrations**:
   - Shell completions
   - Hyperlinks to files

## Build and Distribution

- Use standard Go build tools
- Create releases for multiple platforms (Linux, macOS, Windows)
- Package for popular package managers when available