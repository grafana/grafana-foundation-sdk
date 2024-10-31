---
title: <span class="badge builder"></span> BucketAggregationWithFieldBuilder
---
# <span class="badge builder"></span> BucketAggregationWithFieldBuilder

## Constructor

```go
func NewBucketAggregationWithFieldBuilder() *BucketAggregationWithFieldBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BucketAggregationWithFieldBuilder) Build() (BucketAggregationWithField, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *BucketAggregationWithFieldBuilder) Field(field string) *BucketAggregationWithFieldBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *BucketAggregationWithFieldBuilder) Id(id string) *BucketAggregationWithFieldBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *BucketAggregationWithFieldBuilder) Settings(settings any) *BucketAggregationWithFieldBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BucketAggregationWithFieldBuilder) Type(typeArg elasticsearch.BucketAggregationType) *BucketAggregationWithFieldBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BucketAggregationWithField](./object-BucketAggregationWithField.md)
