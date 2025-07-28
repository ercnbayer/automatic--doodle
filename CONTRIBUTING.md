# Contributing to Automatic Doodle Backend

Thank you for your interest in contributing! This document outlines the process for contributing to the project.

## How to Contribute

1. **Fork the repository** and clone your fork locally.
2. **Create a new branch** for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes**. Please follow the existing code style and structure.
4. **Test your changes** locally:
   ```bash
   docker compose -f ./dev/docker-compose.yaml up -d
   go run test ./... -v
   ```
5. **Commit your changes** with clear, descriptive messages.
6. **Push to your fork** and open a Pull Request (PR) against the `main` branch.
7. **Describe your PR** clearly. Reference any related issues.

## Code Style
- Use idiomatic Go and follow [Effective Go](https://golang.org/doc/effective_go.html).
- Use `gofmt` before committing.
- Keep functions and files focused and well-documented.

## Branching
- Use feature branches: `feature/your-feature`, `bugfix/your-bug`, etc.
- Keep your branch up to date with `main`.

## Pull Requests
- Ensure your PR passes all tests and CI checks.
- Link to any relevant issues.
- Large changes? Please open an issue first to discuss.

## Issue Reporting
- Search for existing issues before opening a new one.
- Provide clear steps to reproduce bugs.
- Suggest enhancements with a clear use case.

## Questions?
Open an issue or start a discussion in the repository.

Thank you for helping improve Automatic Doodle! 