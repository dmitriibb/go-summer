# go-summer
A Spring-like dependency injection framework for Go

## Overview

Go-Summer provides dependency injection capabilities similar to Java Spring Framework. It uses struct tags and reflection to automatically wire dependencies between components.

## Features

- **Dependency Injection**: Automatically inject dependencies using struct tags
- **Container-based**: Centralized bean container for managing component lifecycle
- **Type-safe**: Uses Go's reflection for type-safe dependency resolution
- **Thread-safe**: Container operations are protected with mutexes

## Quick Start

### Basic Usage

1. **Define your components** with struct tags:

```go
type UserRepository struct {
    // Repository implementation
}

type UserService struct {
    Repository *UserRepository `pebble:"autowired"`
}
```

2. **Register and initialize**:

```go
container := pebble.NewContainer()

// Register components
container.Register(NewUserRepository())
container.Register(NewUserService())

// Initialize (performs autowiring)
container.Initialize()

// Use your services
userService, _ := container.Get("userService")
```

### Struct Tags

- `pebble:"autowired"` - Marks a field for automatic dependency injection
- `pebble:"autowired,name:beanName"` - Specifies a specific bean name to inject

### Examples

See the `cmd/` directory for complete examples:
- `cmd/example/` - Basic usage example
- `cmd/example-advanced/` - Advanced usage with multiple dependencies

## Project Structure

- **internal/core/pebble/** - Core framework implementation
- **cmd/** - Example applications demonstrating framework usage

## How It Works

1. **Registration**: Components are registered with the container (by type or by name)
2. **Initialization**: The container scans all registered beans and performs autowiring
3. **Autowiring**: Fields marked with `pebble:"autowired"` are automatically populated with matching beans
4. **Resolution**: Dependencies are resolved by type or by explicit bean name

## Future Enhancements

- Package scanning to auto-discover components
- Lifecycle management (singleton, prototype scopes)
- Configuration injection
- Aspect-oriented programming support
- Web framework integration