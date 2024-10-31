---
title: <span class="badge object-type-struct"></span> BucketScript
---
# <span class="badge object-type-struct"></span> BucketScript

## Definition

```go
type BucketScript struct {
    Type string `json:"type"`
    PipelineVariables []elasticsearch.PipelineVariable `json:"pipelineVariables,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchBucketScriptSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BucketScript` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (bucketScript *BucketScript) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BucketScript` objects.

```go
func (bucketScript *BucketScript) Equals(other BucketScript) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BucketScript` fields for violations and returns them.

```go
func (bucketScript *BucketScript) Validate() error
```

## See also

 * <span class="badge builder"></span> [BucketScriptBuilder](./builder-BucketScriptBuilder.md)
