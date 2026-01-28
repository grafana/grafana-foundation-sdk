---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> Alertmanager

Name of the alertmanager used as a source for alerts

```go
func (builder *OptionsBuilder) Alertmanager(alertmanager string) *OptionsBuilder
```

### <span class="badge object-method"></span> ExpandAll

Expand all alert groups by default

```go
func (builder *OptionsBuilder) ExpandAll(expandAll bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Labels

Comma-separated list of values used to filter alert results

```go
func (builder *OptionsBuilder) Labels(labels string) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
