# Community Go Collection for Ansible

[![Ansible Galaxy](https://img.shields.io/badge/ansible-galaxy-blue.svg)](https://galaxy.ansible.com/community/go_collection)
[![License](https://img.shields.io/badge/license-GPL--3.0-brightgreen.svg)](LICENSE)

A collection of Ansible modules written in Go, providing enhanced performance and capabilities for infrastructure automation tasks.

## Overview

This collection contains Ansible modules implemented in Go, offering better performance characteristics than traditional Python-based modules. The Go implementation allows for faster execution, lower memory footprint, and better concurrency handling for resource-intensive operations.

## Included Modules

### hello_world

A simple module demonstrating Go integration with Ansible:

```yaml
- name: Say hello
  community.go_collection.hello_world:
    name: 'Ansible'
    greeting: 'Hello'
    uppercase: true
```

**Parameters:**

- `name`: The name to greet (default: "World")
- `greeting`: The greeting to use (default: "Hello")
- `uppercase`: Whether to convert the message to uppercase (default: false)

### file_manager

A file system management module for creating, checking, and removing files and directories:

```yaml
- name: Create a directory
  community.go_collection.file_manager:
    path: '/tmp/test_dir'
    state: directory
```

**Parameters:**

- `path`: Path to the file or directory
- `state`: Desired state (choices: file, directory, touch, absent)

## Installation

### From Ansible Galaxy

```bash
ansible-galaxy collection install community.go_collection
```

### From Source

```bash
git clone https://github.com/your-username/ansible-go-collection.git
cd ansible-go-collection
make install
```

## Requirements

- Ansible 2.13 or later
- Go 1.21 or later (for development only)

## Development

### Building the Modules

```bash
make build
```

### Testing

```bash
make test
```

### Clean Build Artifacts

```bash
make clean
```

### Cross-Compile for Different Platforms

```bash
make build-all
```

## Project Structure

```
community/go_collection/
├── galaxy.yml             # Collection metadata
├── Makefile               # Build automation
├── plugins/
│   └── modules/           # Compiled Go binaries
├── playbooks/             # Example playbooks
│   └── test_modules.yaml  # Test playbook
└── src/                   # Go source code
    ├── common/            # Shared Go code
    ├── file_manager/      # File manager module
    ├── hello_world/       # Hello world module
    └── go.mod             # Go module definition
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the GNU General Public License v3.0 - see the LICENSE file for details.

## Author

- Sagar Paul
