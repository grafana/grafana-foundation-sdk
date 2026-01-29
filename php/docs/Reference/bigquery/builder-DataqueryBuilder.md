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

### <span class="badge object-method"></span> convertToUTC

```php
convertToUTC(bool $convertToUTC)
```

### <span class="badge object-method"></span> dataset

```php
dataset(string $dataset)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Common\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> editorMode

```php
editorMode(\Grafana\Foundation\Bigquery\EditorMode $editorMode)
```

### <span class="badge object-method"></span> format

```php
format(\Grafana\Foundation\Bigquery\QueryFormat $format)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> location

```php
location(string $location)
```

### <span class="badge object-method"></span> partitioned

```php
partitioned(bool $partitioned)
```

### <span class="badge object-method"></span> partitionedField

```php
partitionedField(string $partitionedField)
```

### <span class="badge object-method"></span> project

```php
project(string $project)
```

### <span class="badge object-method"></span> queryPriority

```php
queryPriority(\Grafana\Foundation\Bigquery\QueryPriority $queryPriority)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> rawQuery

```php
rawQuery(bool $rawQuery)
```

### <span class="badge object-method"></span> rawSql

```php
rawSql(string $rawSql)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> sharded

```php
sharded(bool $sharded)
```

### <span class="badge object-method"></span> sql

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\SQLExpression> $sql

```php
sql(\Grafana\Foundation\Cog\Builder $sql)
```

### <span class="badge object-method"></span> table

```php
table(string $table)
```

### <span class="badge object-method"></span> timeShift

```php
timeShift(string $timeShift)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
