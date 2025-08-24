# Contributing to PlayHarbor

Thank you for your interest in contributing to **PlayHarbor**!  
We welcome contributions from everyone, whether it's bug fixes, new features, documentation improvements, or testing.

## How to Contribute

1. **Fork the repository** and clone it locally.
2. Create a **feature branch** for your changes:
   ```bash
   git checkout -b feature/my-feature
   ```
3. Make your changes in the branch.
4. Make sure to **run and test** your code before committing.
5. Commit your changes using **Conventional Commits** (see below).
6. Push your branch to your fork:
   ```bash
   git push origin feature/my-feature
   ```
7. Open a **Pull Request** against the `main` branch of this repository.

---

## Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification to make our commit history clean and readable.

### Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

Common types we use:

- **feat**: a new feature
- **fix**: a bug fix
- **docs**: documentation only changes
- **style**: changes that do not affect meaning of code (white-space, formatting, etc.)
- **refactor**: code changes that neither fixes a bug nor adds a feature
- **perf**: code changes that improve performance
- **test**: adding or updating tests
- **chore**: changes to build process, auxiliary tools, libraries, etc.

### Examples

```bash
git commit -m "feat(go-launcher): add URL validation before launch"
git commit -m "fix(rust-launcher): handle process not found error"
git commit -m "docs: update README with usage instructions"
git commit -m "chore: update sysinfo dependency to 0.30"
```

---

## Pull Request Guidelines

- Base your PR on the latest `main`.
- Include a clear description of what the PR does.
- Reference any related issues (if applicable).
- Ensure that all tests pass.

---

## Code Style

- Keep the code clean and readable.
- Follow idiomatic Go and Rust styles in their respective folders.
- Document new functions and modules with meaningful comments.

---

## Reporting Issues

If you find a bug or have a suggestion, please open an **issue** with a clear description and steps to reproduce.

---

Thank you for contributing to PlayHarbor!

---

## License
All contributions are subject to the [Apache License 2.0](LICENSE) of this project.
