---
title: <span class="badge builder"></span> LogGroupBuilder
---
# <span class="badge builder"></span> LogGroupBuilder

## Constructor

```go
func NewLogGroupBuilder() *LogGroupBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LogGroupBuilder) Build() (LogGroup, error)
```

### <span class="badge object-method"></span> AccountId

AccountId of the log group

```go
func (builder *LogGroupBuilder) AccountId(accountId string) *LogGroupBuilder
```

### <span class="badge object-method"></span> AccountLabel

Label of the log group

```go
func (builder *LogGroupBuilder) AccountLabel(accountLabel string) *LogGroupBuilder
```

### <span class="badge object-method"></span> Arn

ARN of the log group

```go
func (builder *LogGroupBuilder) Arn(arn string) *LogGroupBuilder
```

### <span class="badge object-method"></span> Name

Name of the log group

```go
func (builder *LogGroupBuilder) Name(name string) *LogGroupBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LogGroup](./object-LogGroup.md)
