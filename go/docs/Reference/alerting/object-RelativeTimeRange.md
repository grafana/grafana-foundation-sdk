---
title: <span class="badge object-type-struct"></span> RelativeTimeRange
---
# <span class="badge object-type-struct"></span> RelativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

## Definition

```go
type RelativeTimeRange struct {
    // A Duration represents the elapsed time between two instants
    // as an int64 nanosecond count. The representation limits the
    // largest representable duration to approximately 290 years.
    From *alerting.Duration `json:"from,omitempty"`
    // A Duration represents the elapsed time between two instants
    // as an int64 nanosecond count. The representation limits the
    // largest representable duration to approximately 290 years.
    To *alerting.Duration `json:"to,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RelativeTimeRange` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (relativeTimeRange *RelativeTimeRange) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RelativeTimeRange` objects.

```go
func (relativeTimeRange *RelativeTimeRange) Equals(other RelativeTimeRange) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RelativeTimeRange` fields for violations and returns them.

```go
func (relativeTimeRange *RelativeTimeRange) Validate() error
```

