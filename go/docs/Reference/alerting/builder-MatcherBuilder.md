---
title: <span class="badge builder"></span> MatcherBuilder
---
# <span class="badge builder"></span> MatcherBuilder

## Constructor

```go
func NewMatcherBuilder() *MatcherBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MatcherBuilder) Build() (Matcher, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *MatcherBuilder) Name(name string) *MatcherBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *MatcherBuilder) Type(typeArg alerting.MatchType) *MatcherBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *MatcherBuilder) Value(value string) *MatcherBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Matcher](./object-Matcher.md)
