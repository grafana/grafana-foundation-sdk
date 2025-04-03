---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```php
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> alias

Alias pattern

```php
alias(string $alias)
```

### <span class="badge object-method"></span> bucketAggs

List of bucket aggregations

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\DateHistogram>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Histogram>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Terms>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Filters>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\GeoHashGrid>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Nested>> $bucketAggs

```php
bucketAggs(array $bucketAggs)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> metrics

List of metric aggregations

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Count>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverage>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Derivative>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\CumulativeSum>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BucketScript>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\SerialDiff>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawData>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawDocument>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\UniqueCount>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Percentiles>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ExtendedStats>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Min>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Max>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Sum>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Average>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingFunction>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Logs>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Rate>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\TopMetrics>> $metrics

```php
metrics(array $metrics)
```

### <span class="badge object-method"></span> query

Lucene query

```php
query(string $query)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> timeField

Name of time field

```php
timeField(string $timeField)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
