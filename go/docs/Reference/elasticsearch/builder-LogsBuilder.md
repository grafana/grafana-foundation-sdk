---
title: <span class="badge builder"></span> LogsBuilder
---
# <span class="badge builder"></span> LogsBuilder

## Constructor

```go
func NewLogsBuilder() *LogsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LogsBuilder) Build() (Logs, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *LogsBuilder) Hide(hide bool) *LogsBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *LogsBuilder) Id(id string) *LogsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *LogsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchLogsSettings]) *LogsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Logs](./object-Logs.md)
