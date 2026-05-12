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
    labels: typing.Optional[dict[str, str]]
    # New type for datasource reference
    # Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource: typing.Optional[dashboardv2.Dashboardv2DataQueryKindDatasource]
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

 * <span class="badge builder"></span> [athena.QueryV2](../athena/builder-QueryV2.md)
 * <span class="badge builder"></span> [azuremonitor.QueryV2](../azuremonitor/builder-QueryV2.md)
 * <span class="badge builder"></span> [bigquery.QueryV2](../bigquery/builder-QueryV2.md)
 * <span class="badge builder"></span> [cloudwatch.QueryV2](../cloudwatch/builder-QueryV2.md)
 * <span class="badge builder"></span> [datasource.QueryV2](../datasource/builder-QueryV2.md)
 * <span class="badge builder"></span> [elasticsearch.QueryV2](../elasticsearch/builder-QueryV2.md)
 * <span class="badge builder"></span> [expr.QueryV2](../expr/builder-QueryV2.md)
 * <span class="badge builder"></span> [googlecloudmonitoring.QueryV2](../googlecloudmonitoring/builder-QueryV2.md)
 * <span class="badge builder"></span> [grafanapyroscope.QueryV2](../grafanapyroscope/builder-QueryV2.md)
 * <span class="badge builder"></span> [loki.QueryV2](../loki/builder-QueryV2.md)
 * <span class="badge builder"></span> [parca.QueryV2](../parca/builder-QueryV2.md)
 * <span class="badge builder"></span> [prometheus.QueryV2](../prometheus/builder-QueryV2.md)
 * <span class="badge builder"></span> [tempo.QueryV2](../tempo/builder-QueryV2.md)
 * <span class="badge builder"></span> [testdata.QueryV2](../testdata/builder-QueryV2.md)
