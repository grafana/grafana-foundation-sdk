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
	labels?: Record<string, string>;
	// New type for datasource reference
	// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
	datasource?: {
		name?: string;
	};
	spec: any;
}

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
