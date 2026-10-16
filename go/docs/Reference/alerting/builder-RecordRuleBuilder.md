---
title: <span class="badge builder"></span> RecordRuleBuilder
---
# <span class="badge builder"></span> RecordRuleBuilder

## Constructor

```go
func NewRecordRuleBuilder() *RecordRuleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RecordRuleBuilder) Build() (RecordRule, error)
```

### <span class="badge object-method"></span> From

Which expression node should be used as the input for the recorded metric.

```go
func (builder *RecordRuleBuilder) From(from string) *RecordRuleBuilder
```

### <span class="badge object-method"></span> Metric

Name of the recorded metric.

```go
func (builder *RecordRuleBuilder) Metric(metric string) *RecordRuleBuilder
```

### <span class="badge object-method"></span> TargetDatasourceUid

Which data source should be used to write the output of the recording rule, specified by UID.

```go
func (builder *RecordRuleBuilder) TargetDatasourceUid(targetDatasourceUid string) *RecordRuleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RecordRule](./object-RecordRule.md)
