# Contributing to ModelSync

Thank you for considering contributing to ModelSync! We welcome contributions from everyone. Here are some guidelines to help you get started.

---

## How Can You Contribute?

### 1. Report Issues
- If you encounter a bug or have a suggestion, please [open an issue](https://github.com/RyanFloresTT/ModelSync/issues).
- Include as much detail as possible, such as:
  - Steps to reproduce the bug
  - Expected behavior
  - Your environment (OS, Go version, etc.)

### 2. Submit Pull Requests
- Fork the repository and clone it locally.
- Create a new branch for your changes:
  ```bash
  git checkout -b feature/your-feature-name
  ```
- Make your changes, and ensure they adhere to the coding style.
- Commit your changes with clear messages:
  ```bash
  git commit -m "Add feature: your-feature-name"
  ```
- Push your branch and [submit a pull request](https://github.com/RyanFloresTT/ModelSync/pulls).

---

## Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/YourUsername/ModelSync.git
   cd ModelSync
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o modelsync
   ```

4. Run the tests:
   ```bash
   go test ./...
   ```

---

## Coding Guidelines

- **Formatting**: Ensure your code adheres to Go's standard formatting:
  ```bash
  go fmt ./...
  ```

- **Comments**: Write clear and concise comments for functions and complex logic.

- **Testing**: Add tests for any new functionality or bug fixes:
  - Place tests in the `*_test.go` files in the appropriate package.
  - Use `go test` to verify your changes.

---

## Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

---

## Need Help?

If you have any questions, feel free to reach out by opening an issue. We’re happy to help!

