# Contributing to gls

Thank you for your interest in contributing to `gls`! We welcome bug reports, feature requests, documentation improvements, and code contributions.

## How to Contribute

### Reporting Issues

- Use the [issue templates](https://github.com/logando-al/gls/issues/new/choose) when possible.
- Search existing issues first to avoid duplicates.
- Provide a clear title, description, and steps to reproduce.
- Include your operating system and Go version when relevant.

### Suggesting Features

- Open a feature request issue and describe the use case.
- Discuss the feature with maintainers before investing significant effort.

### Pull Requests

1. **Fork** the repository and create a branch from `main`.
2. **Follow the existing code style** (standard Go formatting with `gofmt`).
3. **Add or update tests** for new functionality.
4. **Run the test suite locally:** `go test ./...`
5. **Ensure the build passes:** `go build ./...`
6. **Update documentation** if your change affects usage or behavior.
7. **Open a pull request** using the provided PR template.

### Development Setup

```bash
git clone https://github.com/logando-al/gls.git
cd gls
go mod download
go build -o gls ./cmd/gls
./gls --help
```

### Running Tests

```bash
go test ./...
go vet ./...
```

### Coding Standards

- Use `gofmt` for formatting.
- Keep functions focused and well-named.
- Add comments for exported symbols.
- Avoid unnecessary dependencies.

### Commit Messages

- Use clear, descriptive commit messages.
- Prefer imperative mood (e.g., "Add color support for symlinks").
- Reference issues when applicable (e.g., "Fixes #12").

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Questions?

Feel free to open a [discussion](https://github.com/logando-al/gls/discussions) or ask in an issue.
