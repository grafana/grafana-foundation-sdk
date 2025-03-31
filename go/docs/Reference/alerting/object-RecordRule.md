---
title: <span class="badge object-type-struct"></span> RecordRule
---
# <span class="badge object-type-struct"></span> RecordRule

## Definition

```go
type RecordRule struct {
    // Which expression node should be used as the input for the recorded metric.
    From string `json:"from"`
    // Name of the recorded metric.
    Metric string `json:"metric"`
    // Which data source should be used to write the output of the recording rule, specified by UID.
    TargetDatasourceUid *string `json:"target_datasource_uid,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RecordRule` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (recordRule *RecordRule) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RecordRule` objects.

```go
func (recordRule *RecordRule) Equals(other RecordRule) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RecordRule` fields for violations and returns them.

```go
func (recordRule *RecordRule) Validate() error
```

## See also

 * <span class="badge builder"></span> [RecordRuleBuilder](./builder-RecordRuleBuilder.md)
