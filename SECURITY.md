# Security Policy

## Supported Versions

We release security updates for the latest tagged version of `gls`. Older versions are not maintained.

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |
| older   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability in `gls`, please report it privately.

**Do not open a public issue.** Public disclosure can put users at risk before a fix is available.

### How to report

- Open a **private vulnerability report** via GitHub: [Report a vulnerability](https://github.com/logando-al/gls/security/advisories/new)
- Alternatively, email the maintainer at `security@logando-al.dev` (replace with a valid address if different).

### What to include

- A clear description of the vulnerability
- Steps to reproduce, or a minimal proof of concept
- Affected versions
- Any suggested remediation

## Response Process

1. We will acknowledge receipt of your report within **5 business days**.
2. We will investigate and work on a fix.
3. We will keep you informed of our progress.
4. Once a fix is ready, we will coordinate disclosure and credit you unless you prefer to remain anonymous.

## Disclosure Policy

We follow a **coordinated disclosure** approach:

- A security advisory will be published on GitHub after a fix is released.
- We ask reporters to allow at least **90 days** before public disclosure, to give users time to update.

## Security Best Practices for Users

- Install `gls` only from the official GitHub releases page or via `go install github.com/logando-al/gls@latest`.
- Verify release artifacts when checksums/signatures are provided.
- Keep your installed version up to date.
