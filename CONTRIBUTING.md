# Contributing to Ariaj

First off, thank you for considering contributing to Ariaj! This document provides guidelines and instructions for contributing.

## Development Process

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Testing

### Prerequisites
- Go 1.19 or higher
- Ollama installed and running
- goreleaser (for release testing)

### Running Tests

```bash
# Test Release Build
# For Windows:
.\scripts\test-release.bat

# For Linux/Unix/macOS:
./scripts/test-release.sh
```

### Test Cases

1. **CLI Tests**
```bash
# Test basic commands
ariaj --help
ariaj model

# Test interactive mode
ariaj

# Test prompt mode
ariaj "What is 2+2?"
```

1. **Installation Tests**
```bash
# Test install/uninstall
./ariaj install
./ariaj uninstall
```

## Code Style Guidelines

### Go Code
- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for public APIs
- Keep functions focused and small
- Use proper error handling

### Documentation
- Keep README.md up to date
- Document new features
- Include usage examples
- Update version numbers

## Pull Request Guidelines

1. **Before Submitting**
   - Run all tests
   - Update documentation
   - Format code
   - Add test cases

2. **PR Description**
   - Describe changes clearly
   - Reference related issues
   - List breaking changes
   - Include test results

3. **Review Process**
   - Respond to review comments
   - Make requested changes
   - Keep PR focused

## Release Process

1. **Version Bump**
   - Update version.txt
   - Update changelog
   - Tag release

2. **Testing**
   - Run test-release.sh/bat
   - Verify binaries
   - Check documentation

3. **Release**
   - Create GitHub release
   - Upload artifacts
   - Update documentation

## Bug Reports

Include:
- Version info
- OS/Environment
- Steps to reproduce
- Expected vs actual behavior
- Logs/error messages

## Feature Requests

Include:
- Use case description
- Problem being solved
- Proposed solution
- Alternative solutions

## Code of Conduct

- Be respectful
- Accept constructive feedback
- Focus on improvement
- Help others learn
- Follow project standards

## Getting Help

- Check existing issues
- Read documentation
- Ask clear questions
- Provide context

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
