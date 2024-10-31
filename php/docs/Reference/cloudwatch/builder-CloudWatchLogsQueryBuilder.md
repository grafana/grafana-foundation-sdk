---
title: <span class="badge builder"></span> CloudWatchLogsQueryBuilder
---
# <span class="badge builder"></span> CloudWatchLogsQueryBuilder

## Constructor

```php
new CloudWatchLogsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> expression

The CloudWatch Logs Insights query to execute

```php
expression(string $expression)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> id

```php
id(string $id)
```

### <span class="badge object-method"></span> logGroupNames

@deprecated use logGroups

@param array<string> $logGroupNames

```php
logGroupNames(array $logGroupNames)
```

### <span class="badge object-method"></span> logGroups

Log groups to query

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\LogGroup>> $logGroups

```php
logGroups(array $logGroups)
```

### <span class="badge object-method"></span> queryMode

Whether a query is a Metrics, Logs, or Annotations query

```php
queryMode(\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode)
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

### <span class="badge object-method"></span> region

AWS region to query for the logs

```php
region(string $region)
```

### <span class="badge object-method"></span> statsGroups

Fields to group the results by, this field is automatically populated whenever the query is updated

@param array<string> $statsGroups

```php
statsGroups(array $statsGroups)
```

## See also

 * <span class="badge object-type-class"></span> [CloudWatchLogsQuery](./object-CloudWatchLogsQuery.md)
