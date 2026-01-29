---
title: <span class="badge object-type-class"></span> DataQueryKind
---
# <span class="badge object-type-class"></span> DataQueryKind

## Definition

```python
class DataQueryKind:
    kind: typing.Literal["DataQuery"]
    group: str
    version: str
    # New type for datasource reference
    # Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource: typing.Optional[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]
    spec: typing.Optional[object]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [athena.Query](../athena/builder-Query.md)
 * <span class="badge builder"></span> [azuremonitor.Query](../azuremonitor/builder-Query.md)
 * <span class="badge builder"></span> [bigquery.Query](../bigquery/builder-Query.md)
 * <span class="badge builder"></span> [cloudwatch.Query](../cloudwatch/builder-Query.md)
 * <span class="badge builder"></span> [DataQueryKind](./builder-DataQueryKind.md)
 * <span class="badge builder"></span> [datasource.Query](../datasource/builder-Query.md)
 * <span class="badge builder"></span> [elasticsearch.Query](../elasticsearch/builder-Query.md)
 * <span class="badge builder"></span> [expr.Query](../expr/builder-Query.md)
 * <span class="badge builder"></span> [googlecloudmonitoring.Query](../googlecloudmonitoring/builder-Query.md)
 * <span class="badge builder"></span> [grafanapyroscope.Query](../grafanapyroscope/builder-Query.md)
 * <span class="badge builder"></span> [loki.Query](../loki/builder-Query.md)
 * <span class="badge builder"></span> [parca.Query](../parca/builder-Query.md)
 * <span class="badge builder"></span> [prometheus.Query](../prometheus/builder-Query.md)
 * <span class="badge builder"></span> [tempo.Query](../tempo/builder-Query.md)
 * <span class="badge builder"></span> [testdata.Query](../testdata/builder-Query.md)
