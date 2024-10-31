---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```php
new AnnotationQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
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

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

@param array<string> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> groupBys

Array of labels to group data by.

@param array<string> $groupBys

```php
groupBys(array $groupBys)
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

### <span class="badge object-method"></span> secondaryAlignmentPeriod

Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```php
secondaryAlignmentPeriod(string $secondaryAlignmentPeriod)
```

### <span class="badge object-method"></span> secondaryCrossSeriesReducer

Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```php
secondaryCrossSeriesReducer(string $secondaryCrossSeriesReducer)
```

### <span class="badge object-method"></span> secondaryGroupBys

Only present if a preprocessor is selected. Array of labels to group data by.

@param array<string> $secondaryGroupBys

```php
secondaryGroupBys(array $secondaryGroupBys)
```

### <span class="badge object-method"></span> secondaryPerSeriesAligner

Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.

```php
secondaryPerSeriesAligner(string $secondaryPerSeriesAligner)
```

### <span class="badge object-method"></span> text

Annotation text.

```php
text(string $text)
```

### <span class="badge object-method"></span> title

Annotation title.

```php
title(string $title)
```

### <span class="badge object-method"></span> view

Data view, defaults to FULL.

```php
view(string $view)
```

## See also

 * <span class="badge object-type-class"></span> [AnnotationQuery](./object-AnnotationQuery.md)
