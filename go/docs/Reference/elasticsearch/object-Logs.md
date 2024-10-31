---
title: <span class="badge object-type-struct"></span> Logs
---
# <span class="badge object-type-struct"></span> Logs

## Definition

```go
type Logs struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchLogsSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Logs` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (logs *Logs) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Logs` objects.

```go
func (logs *Logs) Equals(other Logs) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Logs` fields for violations and returns them.

```go
func (logs *Logs) Validate() error
```

## See also

 * <span class="badge builder"></span> [LogsBuilder](./builder-LogsBuilder.md)
