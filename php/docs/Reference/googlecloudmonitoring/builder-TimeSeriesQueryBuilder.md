---
title: <span class="badge builder"></span> TimeSeriesQueryBuilder
---
# <span class="badge builder"></span> TimeSeriesQueryBuilder

## Constructor

```php
new TimeSeriesQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> graphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```php
graphPeriod(string $graphPeriod)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```php
projectName(string $projectName)
```

### <span class="badge object-method"></span> query

MQL query to be executed.

```php
query(string $query)
```

## See also

 * <span class="badge object-type-class"></span> [TimeSeriesQuery](./object-TimeSeriesQuery.md)
