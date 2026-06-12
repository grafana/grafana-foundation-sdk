---
title: <span class="badge builder"></span> MatcherConfigBuilder
---
# <span class="badge builder"></span> MatcherConfigBuilder

NOTE: (copied from dashboard_kind.cue, since not exported)

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

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
