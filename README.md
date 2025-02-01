# NyxCore

**NyxCore** is a powerful CLI tool designed to streamline API testing, test generation, and automation. With **NyxCore**, you can easily generate tests from your code, automate API testing, and manage your test suite with minimal effort.

## Features
- **Test Generation**: Automatically generate tests from your API code.
- **Test Automation**: Run tests with ease, automate your testing workflows.
- **Load Testing**: Built-in tools for API performance testing and monitoring.
- **Customizable**: Easily configure tests and testing behavior to match your workflow.

## Installation

To install **NyxCore**, follow these steps:

### Prerequisites
- **Go** (version 1.XX or higher) is required for building **NyxCore** from source.
- Ensure you have a Go environment set up (e.g., `$GOPATH` configured).

### Installing from Source

```bash
git clone https://github.com/your-username/nyx-core.git
cd nyx-core
make build  # (optional, if you have a Makefile with build commands)
```

This will compile the tool into a binary that can be run locally.

Alternatively, you can install it directly using Go/

```bash
go install github.com/TravisBubb/nyx-core@latest
```

### Pre-built Binaries
You can download pre-built binaries from the **Releases** section on the GitHub repo.

## Usage
Once installed, you can start using NyxCore directly from the command line. Here are some examples:

### Generate Tests from an OpenAPI Specification
```bash
nyx generate myapi.yaml
```

### Running a Test Suite:
```bash
nyx run tests/myapi_test.json
```

### Load Testing:
```bash
nyx loadtest tests/myapi_loadtest.json
```

### Configuration:
You can configure NyxCore by modifiying the ```nyxcore.yaml``` config file in your project root. Example configuration:
```yaml
api_base_url: https://yourapi.com
timeout: 30s
```

## Contributing
If you'd like to contribute to NyxCore, feel free to fork the repo, create a branch, and submit a pull request. We welcome bug fixes, new features, and improvements.
1. Fork the repo
2. Create a branch (```git checkout -b feature/your-feature```)
3. Make your changes
4. Commit your changes (```git commit -am 'Add new feature'```)
5. Push to the remote repository (```git push origin feature/your-feature```)
6. Open a pull request

Please see **CONTRIBUTING.md** for further details.

## License
NyxCore is open-source and available under the (MIT License)[https://tlo.mit.edu/understand-ip/exploring-mit-open-source-license-comprehensive-guide]
