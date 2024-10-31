---
title: <span class="badge builder"></span> BaseBucketAggregationBuilder
---
# <span class="badge builder"></span> BaseBucketAggregationBuilder

## Constructor

```go
func NewBaseBucketAggregationBuilder() *BaseBucketAggregationBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BaseBucketAggregationBuilder) Build() (BaseBucketAggregation, error)
```

### <span class="badge object-method"></span> Id

```go
func (builder *BaseBucketAggregationBuilder) Id(id string) *BaseBucketAggregationBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *BaseBucketAggregationBuilder) Settings(settings any) *BaseBucketAggregationBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BaseBucketAggregationBuilder) Type(typeArg elasticsearch.BucketAggregationType) *BaseBucketAggregationBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BaseBucketAggregation](./object-BaseBucketAggregation.md)
