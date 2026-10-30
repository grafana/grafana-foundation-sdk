---
title: <span class="badge object-type-struct"></span> Expression
---
# <span class="badge object-type-struct"></span> Expression

## Definition

```go
type Expression struct {
    // The type of query if this is a query expression
    QueryType *string `json:"queryType,omitempty"`
    RelativeTimeRange *alertingv1beta1.RelativeTimeRange `json:"relativeTimeRange,omitempty"`
    // The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
    DatasourceUID *alertingv1beta1.DatasourceUID `json:"datasourceUID,omitempty"`
    Model any `json:"model"`
    // Used to mark the expression to be used as the final source for the rule evaluation
    // Only one expression in a rule can be marked as the source
    // For AlertRules, this is the expression that will be evaluated against the alerting condition
    // For RecordingRules, this is the expression that will be recorded
    Source *bool `json:"source,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Expression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (expression *Expression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Expression` objects.

```go
func (expression *Expression) Equals(other Expression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Expression` fields for violations and returns them.

```go
func (expression *Expression) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExpressionBuilder](./builder-ExpressionBuilder.md)
