# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Bazel playground demonstrating multi-language builds with Java, Go, Python, and Protocol Buffers. The project uses Bazelisk for consistent Bazel version management.

## Build System

- Uses Bazel 8.2.1 (LTS) with Bzlmod for dependency management
- Bazelisk is the recommended way to invoke Bazel commands
- Maven dependencies are managed through rules_jvm_external
- Go dependencies use Gazelle for BUILD file generation
- Migrated from WORKSPACE to MODULE.bazel for Bazel 8 compatibility

## Essential Commands

### Building
```bash
bazelisk build //...                    # Build all targets
bazelisk build //src/java/...          # Build Java targets
bazelisk build //src/go/...            # Build Go targets
bazelisk build //src/python/...        # Build Python targets
```

### Running Applications
```bash
bazelisk run //src/java/org/scode/bazelplay:hello
bazelisk run //src/python/scode:hello
bazelisk run //src/go:main
```

### Testing
```bash
bazelisk test //...                     # Run all tests
bazelisk test //src/java/...           # Run Java tests
bazelisk test //src/go/...             # Run Go tests
bazelisk test //src/python/...         # Run Python tests
```

### Dependency Management
```bash
./scripts/gazelle.sh                   # Update Go BUILD files
./scripts/maven_install_pin.sh         # Pin Maven dependencies
```

## Architecture

- **src/java/**: Java packages with main application and library code
- **src/go/**: Go modules with main package and subpackages
- **src/python/**: Python modules
- **src/proto/**: Protocol Buffer definitions
- **scripts/**: Utility scripts for dependency management

The project demonstrates inter-language dependencies and protocol buffer usage across all supported languages.

## Environment Requirements

- Requires JDK (set JAVA_HOME if needed)
- On macOS with Homebrew OpenJDK: `JAVA_HOME=/usr/local/opt/openjdk/libexec/openjdk.jdk/Contents/Home`

## Bazel Configuration

The project includes a `.bazelrc` file with macOS-specific configuration for C++ compilation compatibility:
- Minimum macOS version set to 11.0 for modern C++ standard library support
- C++17 standard enabled for protobuf compatibility

## Version Information

- **Bazel**: 8.2.1 (latest LTS)
- **Go**: 1.21.13
- **Protobuf**: 24.4 (latest compatible version for macOS)
- **rules_go**: 0.51.0
- **rules_jvm_external**: 6.6