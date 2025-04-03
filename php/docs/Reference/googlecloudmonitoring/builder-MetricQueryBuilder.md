---
title: <span class="badge builder"></span> MetricQueryBuilder
---
# <span class="badge builder"></span> MetricQueryBuilder

## Constructor

```php
new MetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```php
aliasBy(string $aliasBy)
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```php
alignmentPeriod(string $alignmentPeriod)
```

### <span class="badge object-method"></span> crossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```php
crossSeriesReducer(string $crossSeriesReducer)
```

### <span class="badge object-method"></span> editorMode

```php
editorMode(string $editorMode)
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

@param array<string> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> graphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```php
graphPeriod(string $graphPeriod)
```

### <span class="badge object-method"></span> groupBys

Array of labels to group data by.

@param array<string> $groupBys

```php
groupBys(array $groupBys)
```

### <span class="badge object-method"></span> metricKind

```php
metricKind(\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind)
```

### <span class="badge object-method"></span> metricType

```php
metricType(string $metricType)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```php
perSeriesAligner(string $perSeriesAligner)
```

### <span class="badge object-method"></span> preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```php
preprocessor(\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor)
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

### <span class="badge object-method"></span> valueType

```php
valueType(string $valueType)
```

### <span class="badge object-method"></span> view

```php
view(string $view)
```

## See also

 * <span class="badge object-type-class"></span> [MetricQuery](./object-MetricQuery.md)
