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
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     */
    public ?\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource $datasource;

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

 * <span class="badge builder"></span> [athena.QueryV2Builder](../athena/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [azuremonitor.QueryV2Builder](../azuremonitor/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [bigquery.QueryV2Builder](../bigquery/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [cloudwatch.QueryV2Builder](../cloudwatch/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [datasource.QueryV2Builder](../datasource/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [elasticsearch.QueryV2Builder](../elasticsearch/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [expr.QueryV2Builder](../expr/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [googlecloudmonitoring.QueryV2Builder](../googlecloudmonitoring/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [grafanapyroscope.QueryV2Builder](../grafanapyroscope/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [loki.QueryV2Builder](../loki/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [parca.QueryV2Builder](../parca/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [prometheus.QueryV2Builder](../prometheus/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [tempo.QueryV2Builder](../tempo/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [testdata.QueryV2Builder](../testdata/builder-QueryV2Builder.md)
