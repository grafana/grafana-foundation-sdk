---
title: <span class="badge object-type-struct"></span> MovingAverageOrDerivativeOrCumulativeSumOrBucketScript
---
# <span class="badge object-type-struct"></span> MovingAverageOrDerivativeOrCumulativeSumOrBucketScript

## Definition

```go
type MovingAverageOrDerivativeOrCumulativeSumOrBucketScript struct {
    MovingAverage *elasticsearch.MovingAverage `json:"MovingAverage,omitempty"`
    Derivative *elasticsearch.Derivative `json:"Derivative,omitempty"`
    CumulativeSum *elasticsearch.CumulativeSum `json:"CumulativeSum,omitempty"`
    BucketScript *elasticsearch.BucketScript `json:"BucketScript,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` as JSON.

```go
func (movingAverageOrDerivativeOrCumulativeSumOrBucketScript *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` from JSON.

```go
func (movingAverageOrDerivativeOrCumulativeSumOrBucketScript *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageOrDerivativeOrCumulativeSumOrBucketScript *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` objects.

```go
func (movingAverageOrDerivativeOrCumulativeSumOrBucketScript *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) Equals(other MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` fields for violations and returns them.

```go
func (movingAverageOrDerivativeOrCumulativeSumOrBucketScript *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) Validate() error
```

