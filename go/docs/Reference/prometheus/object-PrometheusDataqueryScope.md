---
title: <span class="badge object-type-struct"></span> PrometheusDataqueryScope
---
# <span class="badge object-type-struct"></span> PrometheusDataqueryScope

## Definition

```go
type PrometheusDataqueryScope struct {
    Matchers string `json:"matchers"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PrometheusDataqueryScope` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (prometheusDataqueryScope *PrometheusDataqueryScope) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PrometheusDataqueryScope` objects.

```go
func (prometheusDataqueryScope *PrometheusDataqueryScope) Equals(other PrometheusDataqueryScope) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PrometheusDataqueryScope` fields for violations and returns them.

```go
func (prometheusDataqueryScope *PrometheusDataqueryScope) Validate() error
```

## See also

 * <span class="badge builder"></span> [PrometheusDataqueryScopeBuilder](./builder-PrometheusDataqueryScopeBuilder.md)
