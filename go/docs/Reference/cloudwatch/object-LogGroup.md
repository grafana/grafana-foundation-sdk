---
title: <span class="badge object-type-struct"></span> LogGroup
---
# <span class="badge object-type-struct"></span> LogGroup

## Definition

```go
type LogGroup struct {
    // ARN of the log group
    Arn string `json:"arn"`
    // Name of the log group
    Name string `json:"name"`
    // AccountId of the log group
    AccountId *string `json:"accountId,omitempty"`
    // Label of the log group
    AccountLabel *string `json:"accountLabel,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LogGroup` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (logGroup *LogGroup) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LogGroup` objects.

```go
func (logGroup *LogGroup) Equals(other LogGroup) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LogGroup` fields for violations and returns them.

```go
func (logGroup *LogGroup) Validate() error
```

## See also

 * <span class="badge builder"></span> [LogGroupBuilder](./builder-LogGroupBuilder.md)
