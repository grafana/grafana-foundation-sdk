---
title: <span class="badge object-type-class"></span> DataQueryKind
---
# <span class="badge object-type-class"></span> DataQueryKind

## Definition

```php
class DataQueryKind implements \JsonSerializable
{
    public string $kind;

    public string $group;

    public string $version;

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource $datasource;

    /**
     * @var mixed|null
     */
    public $spec;

}
```
## Methods

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [athena.QueryBuilder](../athena/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [azuremonitor.QueryBuilder](../azuremonitor/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [bigquery.QueryBuilder](../bigquery/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [cloudwatch.QueryBuilder](../cloudwatch/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [DataQueryKindBuilder](./builder-DataQueryKindBuilder.md)
 * <span class="badge builder"></span> [datasource.QueryBuilder](../datasource/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [elasticsearch.QueryBuilder](../elasticsearch/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [expr.QueryBuilder](../expr/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [googlecloudmonitoring.QueryBuilder](../googlecloudmonitoring/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [grafanapyroscope.QueryBuilder](../grafanapyroscope/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [loki.QueryBuilder](../loki/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [parca.QueryBuilder](../parca/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [prometheus.QueryBuilder](../prometheus/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [tempo.QueryBuilder](../tempo/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [testdata.QueryBuilder](../testdata/builder-QueryBuilder.md)
