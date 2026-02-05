---
title: <span class="badge builder"></span> MatcherConfigBuilder
---
# <span class="badge builder"></span> MatcherConfigBuilder

## Constructor

```go
func NewMatcherConfigBuilder() *MatcherConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MatcherConfigBuilder) Build() (MatcherConfig, error)
```

### <span class="badge object-method"></span> Id

The matcher id. This is used to find the matcher implementation from registry.

```go
func (builder *MatcherConfigBuilder) Id(id string) *MatcherConfigBuilder
```

### <span class="badge object-method"></span> Options

The matcher options. This is specific to the matcher implementation.

```go
func (builder *MatcherConfigBuilder) Options(options any) *MatcherConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MatcherConfig](./object-MatcherConfig.md)
