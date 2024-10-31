---
title: <span class="badge object-type-enum"></span> BucketAggregationType
---
# <span class="badge object-type-enum"></span> BucketAggregationType

## Definition

```go
type BucketAggregationType string
const (
	BucketAggregationTypeTerms BucketAggregationType = "terms"
	BucketAggregationTypeFilters BucketAggregationType = "filters"
	BucketAggregationTypeGeohashGrid BucketAggregationType = "geohash_grid"
	BucketAggregationTypeDateHistogram BucketAggregationType = "date_histogram"
	BucketAggregationTypeHistogram BucketAggregationType = "histogram"
	BucketAggregationTypeNested BucketAggregationType = "nested"
)

```
