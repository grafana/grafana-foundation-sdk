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

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> editorMode

```php
editorMode(\Grafana\Foundation\Loki\QueryEditorMode $editorMode)
```

### <span class="badge object-method"></span> expr

The LogQL query.

```php
expr(string $expr)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> instant

@deprecated, now use queryType.

```php
instant(bool $instant)
```

### <span class="badge object-method"></span> legendFormat

Used to override the name of the series.

```php
legendFormat(string $legendFormat)
```

### <span class="badge object-method"></span> maxLines

Used to limit the number of log rows returned.

```php
maxLines(int $maxLines)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> range

@deprecated, now use queryType.

```php
range(bool $range)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> resolution

@deprecated, now use step.

```php
resolution(int $resolution)
```

### <span class="badge object-method"></span> step

Used to set step value for range queries.

```php
step(string $step)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
