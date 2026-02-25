---
title: <span class="badge object-type-interface"></span> DataQueryKind
---
# <span class="badge object-type-interface"></span> DataQueryKind

## Definition

```typescript
export interface DataQueryKind {
	kind: "DataQuery";
	group: string;
	version: string;
	// New type for datasource reference
	// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
	datasource?: {
		name?: string;
	};
	spec: any;
}

```
## See also

 * <span class="badge builder"></span> [athena.QueryBuilder](../athena/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [azuremonitor.QueryBuilder](../azuremonitor/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [bigquery.QueryBuilder](../bigquery/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [cloudwatch.QueryBuilder](../cloudwatch/builder-QueryBuilder.md)
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
