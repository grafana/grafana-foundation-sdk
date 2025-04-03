---
title: <span class="badge object-type-struct"></span> TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql
---
# <span class="badge object-type-struct"></span> TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql

## Definition

```go
type TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql struct {
    TypeMath *expr.TypeMath `json:"TypeMath,omitempty"`
    TypeReduce *expr.TypeReduce `json:"TypeReduce,omitempty"`
    TypeResample *expr.TypeResample `json:"TypeResample,omitempty"`
    TypeClassicConditions *expr.TypeClassicConditions `json:"TypeClassicConditions,omitempty"`
    TypeThreshold *expr.TypeThreshold `json:"TypeThreshold,omitempty"`
    TypeSql *expr.TypeSql `json:"TypeSql,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` as JSON.

```go
func (typeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` from JSON.

```go
func (typeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (typeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` objects.

```go
func (typeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) Equals(other TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` fields for violations and returns them.

```go
func (typeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) Validate() error
```

