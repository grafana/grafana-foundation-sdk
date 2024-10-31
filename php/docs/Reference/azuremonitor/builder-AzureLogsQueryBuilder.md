---
title: <span class="badge builder"></span> AzureLogsQueryBuilder
---
# <span class="badge builder"></span> AzureLogsQueryBuilder

## Constructor

```php
new AzureLogsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> dashboardTime

If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.

```php
dashboardTime(bool $dashboardTime)
```

### <span class="badge object-method"></span> intersectTime

@deprecated Use dashboardTime instead

```php
intersectTime(bool $intersectTime)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```php
query(string $query)
```

### <span class="badge object-method"></span> resource

@deprecated Use resources instead

```php
resource(string $resource)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

@param array<string> $resources

```php
resources(array $resources)
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```php
resultFormat(\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat)
```

### <span class="badge object-method"></span> timeColumn

If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated

```php
timeColumn(string $timeColumn)
```

### <span class="badge object-method"></span> workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat.

```php
workspace(string $workspace)
```

## See also

 * <span class="badge object-type-class"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
