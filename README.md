# Container

Central container for managing dependencies and performing dependency injection.

**Mote:** Container use [Caster](https://github.com/bopher/caster) for get dependencies by type.

## Create New Container

```go
import "github.com/bopher/container";
cnt := container.NewContainer();
```

## Usage

Container interface contains following methods:

### Register

Add new dependency to container

```go
// Signature:
Register(name string, dep any)

// Example:
cnt.Register("app-name", "My App")
```

### Resolve

Get dependency. this function return false as second return value if dependency not exists.

**Caution:** Resolve return data as `any` type. if you need to parse data as special type use helper parser methods described later.

```go
// Signature:
Resolve(name string) (any, bool)

// Example:
appName, exists := cnt.Resolve("app-name") // => "My App", true
```

### Exists

Check if dependency exists.

```go
// Signature:
Exists(name string) bool

// Example:
exists := cnt.Exists("app-name") // => true
```

### Cast

Parse dependency as caster.

```go
// Signature:
Cast(name string) caster.Caster

// Example:
v, err := cnt.Cast("timeout").Int()
```
