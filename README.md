# ModelSync

ModelSync is an open-source tool designed to simplify model synchronization across different programming languages. It watches your backend model files (e.g., Go structs) and automatically generates equivalent models in multiple target languages (e.g., TypeScript, C#, C++) based on predefined templates. This tool is ideal for developers working in polyglot environments where consistent model definitions are critical.

---

## Features
- **File Watching**: Automatically detects changes in your source file and regenerates target models in real-time.
- **Multi-Language Support**: Generates code for multiple target languages, including TypeScript, C#, C++, Python, and Java.
- **Automatic Configuration**: Creates a default configuration file if none exists, making setup seamless.

---

## Installation
You can install ModelSync using Go:

```bash
go install github.com/RyanFloresTT/ModelSync/cmd/modelsync@latest
```

This will install the `modelsync` binary to your `$GOPATH/bin` directory, making it accessible system-wide.

---

## Usage

### 1. **Initialize the Project**
Run the program to automatically generate a default configuration file if none exists:

```bash
modelsync
```
Output:
```
Configuration file syncConfig.json does not exist. Creating a default one.
Initialized project with config file: syncConfig.json
```
Edit `syncConfig.json` to specify your source file and target languages:

```json
{
  "WatchFile": "models/sample.go",
  "Targets": [
    {
      "Language": "typescript",
      "Output": "frontend/src/models/BookResponse.tsx"
    },
    {
      "Language": "csharp",
      "Output": "shared/Models/BookResponse.cs"
    }
  ]
}
```

### 2. **Run the Program**
After editing the configuration, run the tool to start watching the specified file:

```bash
modelsync
```
When the source file changes, the tool automatically generates updated target files based on the configuration.

#### Running in Background
The program runs in the foreground by default, blocking the terminal while watching for changes. If you prefer to run it in the background:

- **Linux/macOS**: Use the `&` operator to run the process in the background:
  ```bash
  modelsync &
  ```
- **Windows**: Use `start` to run the program in a new background window:
  ```cmd
  start modelsync
  ```
Alternatively, you can use a process manager (e.g., `systemd`, `launchctl`, or `Task Scheduler`) to manage it as a daemon. This approach is left up to the user.

---

## Configuration
The `syncConfig.json` file contains the following fields:

- **WatchFile**: The path to the source file (e.g., a Go file containing structs).
- **Targets**: An array of target languages and output specifications.
  - **Language**: The target language (e.g., `typescript`, `csharp`, `cpp`).
  - **Output**: The path where the generated file should be written.

Example configuration:
```json
{
  "WatchFile": "models/sample.go",
  "Targets": [
    {
      "Language": "typescript",
      "Output": "frontend/src/models/BookResponse.tsx"
    },
    {
      "Language": "csharp",
      "Output": "shared/Models/BookResponse.cs"
    }
  ]
}
```

---

## Supported Languages
| Language   | Configuration Value |
|------------|----------------------|
| TypeScript | `typescript`         |
| C#         | `csharp`             |
| C++        | `cpp`                |
| Python     | `python`             |
| Java       | `java`               |

---

## Contributing
We welcome contributions! To get started:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with clear messages.
4. Submit a pull request.

---

## License
ModelSync is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments
Special thanks to the open-source community for their inspiration and contributions to cross-language tools.
